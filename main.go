package main

import (
	"fmt"
	"os"

	prompt "github.com/c-bata/go-prompt"
	"github.com/c-bata/go-prompt/completer"
	"github.com/jiqinga/kube-prompt/internal/debug"
	"github.com/jiqinga/kube-prompt/kube"

	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	//_ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

var (
	version  string
	revision string
)

func main() {
	c, err := kube.NewCompleter()
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}

	defer debug.Teardown()
	fmt.Printf("kube-prompt %s (rev-%s)\n", version, revision)
	fmt.Println("请使用 `exit` or `Ctrl-D` 退出此程序.")
	defer fmt.Println("Bye!")
	p := prompt.New(
		kube.Executor,
		c.Complete,
		prompt.OptionTitle("kube-prompt: interactive kubernetes client"),
		prompt.OptionPrefix(">>> "),
		prompt.OptionInputTextColor(prompt.Yellow),
		prompt.OptionCompletionWordSeparator(completer.FilePathCompletionSeparator),
	)
	p.Run()
}
