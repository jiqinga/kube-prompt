package kube

import (
	"context"
	"os"
	"strings"
	"sync"

	"k8s.io/client-go/rest"

	"github.com/c-bata/go-prompt"
	"github.com/c-bata/go-prompt/completer"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	globalState sync.Map
	Version     = "unset"
)

func Getns(key string) string {
	value, ok := globalState.Load(key)
	if !ok {
		// 键 "KUBECONFIG" 不存在
		return ""
	}
	return value.(string)
}

func NewCompleter(kubeconfigPath string) (*Completer, error) {
	var loadingRules *clientcmd.ClientConfigLoadingRules

	if kubeconfigPath != "" {
		loadingRules = &clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfigPath}
	} else {
		loadingRules = clientcmd.NewDefaultClientConfigLoadingRules()
	}

	loader := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		loadingRules,
		&clientcmd.ConfigOverrides{},
	)

	config, err := loader.ClientConfig()
	if err != nil {
		return nil, err
	}

	namespace, _, err := loader.Namespace()
	globalState.Store("namespace", namespace)
	if err != nil {
		return nil, err
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	namespaces, err := client.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		if statusError, ok := err.(*errors.StatusError); ok && statusError.Status().Code == 403 {
			namespaces = nil
		} else {
			return nil, err
		}
	}

	return &Completer{
		// namespace:     namespace,
		namespaceList: namespaces,
		client:        client,
		config:        config,
	}, nil
}

type Completer struct {
	// namespace     string
	namespaceList *corev1.NamespaceList
	client        *kubernetes.Clientset
	config        *rest.Config
}

func removeSpaces(slice []string) []string {
	var s []string
	for i, str := range slice {
		slice[i] = strings.ReplaceAll(str, " ", "")
		if strings.TrimSpace(str) == "" {
			continue
		}
		s = append(s, str)
	}
	return s
}

func (c *Completer) Complete(d prompt.Document) []prompt.Suggest {
	// out := os.Stdout
	// fmt.Fprintln(out, d.CurrentLine())
	if d.TextBeforeCursor() == "" {
		return []prompt.Suggest{}
	}
	// 如果多空格个,空格也会当作参数
	args := strings.Split(d.TextBeforeCursor(), " ")
	// args = removeEmptyStrings(args)
	w := d.GetWordBeforeCursor()

	// If PIPE is in text before the cursor, returns empty suggestions.
	for i := range args {
		if args[i] == "|" {
			return []prompt.Suggest{}
		}
	}

	// If word before the cursor starts with "-", returns CLI flag options.
	if strings.HasPrefix(w, "-") {
		return optionCompleter(args, strings.HasPrefix(w, "--"))
	}

	// Return suggestions for option
	if suggests, found := c.completeOptionArguments(d); found {
		return suggests
	}

	namespace := checkNamespaceArg(d)
	if namespace == "" {
		namespace = Getns("namespace")
	}
	commandArgs, skipNext := excludeOptions(args)
	if skipNext {
		// when type 'get pod -o ', we don't want to complete pods. we want to type 'json' or other.
		// So we need to skip argumentCompleter.
		return []prompt.Suggest{}
	}
	//if len(w) > 0 && w[len(w)-1] == '/' {
	//	return c.argumentsCompleter(namespace, commandArgs, d)
	//}
	//fmt.Println("########", commandArgs, len(commandArgs), "#############")
	if strings.TrimSpace(commandArgs[len(commandArgs)-1]) == "" {
		commandArgs := removeSpaces(commandArgs)
		commandArgs = append(commandArgs, "")
		return c.argumentsCompleter(namespace, commandArgs, d)
	}
	commandArgs = removeSpaces(commandArgs)
	return c.argumentsCompleter(namespace, commandArgs, d)
}

func checkNamespaceArg(d prompt.Document) string {
	args := strings.Split(d.Text, " ")
	var found bool
	for i := 0; i < len(args); i++ {
		if found {
			return args[i]
		}
		if args[i] == "--namespace" || args[i] == "-n" {
			found = true
			continue
		}
	}
	return ""
}

/* Option arguments */

var yamlFileCompleter = completer.FilePathCompleter{
	IgnoreCase: true,
	Filter: func(fi os.FileInfo) bool {
		if fi.IsDir() {
			return true
		}
		if strings.HasSuffix(fi.Name(), ".yaml") || strings.HasSuffix(fi.Name(), ".yml") {
			return true
		}
		return false
	},
}

func getPreviousOption(d prompt.Document) (cmd, option string, found bool) {
	args := strings.Split(d.TextBeforeCursor(), " ")
	l := len(args)
	if l >= 2 {
		option = args[l-2]
	}
	if strings.HasPrefix(option, "-") {
		return args[0], option, true
	}
	return "", "", false
}

func (c *Completer) completeOptionArguments(d prompt.Document) ([]prompt.Suggest, bool) {
	cmd, option, found := getPreviousOption(d)
	if !found {
		return []prompt.Suggest{}, false
	}

	// namespace
	if option == "-n" || option == "--namespace" {
		return prompt.FilterHasPrefix(
			getNameSpaceSuggestions(c.namespaceList),
			d.GetWordBeforeCursor(),
			true,
		), true
	}

	// filename
	switch cmd {
	case "get", "describe", "create", "delete", "replace", "patch",
		"edit", "apply", "expose", "rolling-update", "rollout",
		"label", "annotate", "scale", "convert", "autoscale", "top":
		if option == "-f" || option == "--filename" {
			return yamlFileCompleter.Complete(d), true
		}
	}

	// container
	switch cmd {
	case "exec", "logs", "run", "attach", "port-forward", "cp":
		if option == "-c" || option == "--container" {
			cmdArgs := getCommandArgs(d)
			var suggestions []prompt.Suggest
			if cmdArgs == nil || len(cmdArgs) < 2 {
				suggestions = getContainerNamesFromCachedPods(c.client, Getns("namespace"))
			} else {
				suggestions = getContainerName(c.client, Getns("namespace"), cmdArgs[1])
			}
			return prompt.FilterHasPrefix(
				suggestions,
				d.GetWordBeforeCursor(),
				true,
			), true
		}
	}
	return []prompt.Suggest{}, false
}

func getCommandArgs(d prompt.Document) []string {
	args := strings.Split(d.TextBeforeCursor(), " ")

	// If PIPE is in text before the cursor, returns empty.
	for i := range args {
		if args[i] == "|" {
			return nil
		}
	}

	commandArgs, _ := excludeOptions(args)
	return commandArgs
}

func excludeOptions(args []string) ([]string, bool) {
	l := len(args)
	if l == 0 {
		return nil, false
	}
	cmd := args[0]
	filtered := make([]string, 0, l)

	var skipNextArg bool
	for i := 0; i < len(args); i++ {
		if skipNextArg {
			skipNextArg = false
			continue
		}

		if cmd == "logs" && args[i] == "-f" {
			continue
		}

		for _, s := range []string{
			"-f", "--filename",
			"-n", "--namespace",
			"-s", "--server",
			"--kubeconfig",
			"--cluster",
			"--user",
			"-o", "--output",
			"-c",
			"--container",
		} {
			if strings.HasPrefix(args[i], s) {
				if strings.Contains(args[i], "=") {
					// we can specify option value like '-o=json'
					skipNextArg = false
				} else {
					skipNextArg = true
				}
				continue
			}
		}
		if strings.HasPrefix(args[i], "-") {
			continue
		}

		filtered = append(filtered, args[i])
	}
	return filtered, skipNextArg
}
