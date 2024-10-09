package kube

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"

	"k8s.io/client-go/tools/remotecommand"

	"github.com/c-bata/go-prompt"
	"github.com/jiqinga/kube-prompt/internal/debug"
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	policyv1beta1 "k8s.io/api/policy/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const thresholdFetchInterval = 10 * time.Second

func init() {
	lastFetchedAt = new(sync.Map)
	podList = new(sync.Map)
	endpointList = new(sync.Map)
	deploymentList = new(sync.Map)
	daemonSetList = new(sync.Map)
	eventList = new(sync.Map)
	secretList = new(sync.Map)
	ingressList = new(sync.Map)
	limitRangeList = new(sync.Map)
	persistentVolumeClaimsList = new(sync.Map)
	podTemplateList = new(sync.Map)
	replicaSetList = new(sync.Map)
	replicationControllerList = new(sync.Map)
	resourceQuotaList = new(sync.Map)
	serviceAccountList = new(sync.Map)
	serviceList = new(sync.Map)
	jobList = new(sync.Map)
}

/* LastFetchedAt */

var lastFetchedAt *sync.Map

func shouldFetch(key string) bool {
	v, ok := lastFetchedAt.Load(key)
	if !ok {
		return true
	}
	t, ok := v.(time.Time)
	if !ok {
		return true
	}
	return time.Since(t) > thresholdFetchInterval
}

func updateLastFetchedAt(key string) {
	lastFetchedAt.Store(key, time.Now())
}

/* Component Status */

var componentStatusList atomic.Value

func fetchComponentStatusList(client *kubernetes.Clientset) {
	key := "component_status"
	if !shouldFetch(key) {
		return
	}
	l, _ := client.CoreV1().ComponentStatuses().List(context.TODO(), metav1.ListOptions{})
	componentStatusList.Store(l)
	updateLastFetchedAt(key)
}

func getComponentStatusCompletions(client *kubernetes.Clientset) []prompt.Suggest {
	go fetchComponentStatusList(client)
	l, ok := componentStatusList.Load().(*corev1.ComponentStatusList)
	if !ok || len(l.Items) == 0 {
		return []prompt.Suggest{}
	}
	s := make([]prompt.Suggest, len(l.Items))
	for i := range l.Items {
		s[i] = prompt.Suggest{
			Text: l.Items[i].Name,
		}
	}
	return s
}

/* Config Maps */

var configMapsList atomic.Value

func fetchConfigMapList(client *kubernetes.Clientset, namespace string) {
	key := "config_map_" + namespace
	if !shouldFetch(key) {
		return
	}
	updateLastFetchedAt(key)
	l, _ := client.CoreV1().ConfigMaps(namespace).List(context.TODO(), metav1.ListOptions{})
	configMapsList.Store(l)
}

func getConfigMapSuggestions(client *kubernetes.Clientset, namespace string) []prompt.Suggest {
	go fetchConfigMapList(client, namespace)
	l, ok := configMapsList.Load().(*corev1.ConfigMapList)
	if !ok || len(l.Items) == 0 {
		return []prompt.Suggest{}
	}
	s := make([]prompt.Suggest, len(l.Items))
	for i := range l.Items {
		s[i] = prompt.Suggest{
			Text: l.Items[i].Name,
		}
	}
	return s
}

/* Contexts */

var contextList atomic.Value

func fetchContextList() {
	key := "context"
	if !shouldFetch(key) {
		return
	}
	updateLastFetchedAt(key)
	r := ExecuteAndGetResults("config get-contexts --no-headers -o name")
	r = strings.TrimRight(r, "\n")
	contextList.Store(strings.Split(r, "\n"))
}

func getContextSuggestions() []prompt.Suggest {
	go fetchContextList()
	l, ok := contextList.Load().([]string)
	if !ok || len(l) == 0 {
		return []prompt.Suggest{}
	}
	s := make([]prompt.Suggest, len(l))
	for i := range l {
		s[i] = prompt.Suggest{
			Text: l[i],
		}
	}
	return s
}

/* Pod */

var podList *sync.Map

func fetchPods(client *kubernetes.Clientset, namespace string) {
	key := "pod_" + namespace
	if !shouldFetch(key) {
		return
	}
	updateLastFetchedAt(key)

	l, _ := client.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	podList.Store(namespace, l)
}

// 路径补全函数
func completerfile(path string, dirOnly bool) []prompt.Suggest {
	var dir string
	var isRelative bool

	if path == "" || path == "/" {
		dir = "/"
	} else if strings.HasPrefix(path, "./") {
		dir = filepath.Clean(path)
		isRelative = true
	} else if filepath.IsAbs(path) {
		if strings.HasSuffix(path, "/.") {
			dir = path
		} else {
			dir = filepath.Clean(path)
		}
	} else {
		dir = filepath.Join(".", path)
	}

	// 如果路径以 / 结尾或者是 "." 或 ".."，我们应该列出该目录的内容
	if strings.HasSuffix(path, "/") || path == "." || path == ".." {
		// if strings.HasSuffix(path, "/") {
		dir = path
	} else {
		// 否则，我们列出父目录的内容
		dir = filepath.Dir(dir)
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return []prompt.Suggest{}
	}

	var suggestions []prompt.Suggest

	for _, file := range files {
		// 如果 dirOnly 为 true 且当前项不是目录，则跳过
		if dirOnly && !file.IsDir() {
			continue
		}
		name := file.Name()
		var fullPath string

		if isRelative {
			relPath, _ := filepath.Rel(".", filepath.Join(dir, name))
			fullPath = "./" + relPath
		} else if dir == "/" {
			fullPath = "/" + name
		} else {
			fullPath = filepath.Join(dir, name)
		}

		if file.IsDir() {
			fullPath += string(os.PathSeparator)
		}

		// 如果路径以 / 结尾或者是 "." 或 ".."，我们显示所有文件和目录
		// 否则，我们只显示与最后一个路径组件匹配的建议
		if strings.HasSuffix(path, "/") || path == "." || path == ".." ||
			strings.HasPrefix(strings.ToLower(name), strings.ToLower(filepath.Base(path))) {
			suggestions = append(suggestions, prompt.Suggest{
				Text: fullPath,
				Description: func() string {
					if file.IsDir() {
						return "目录"
					}
					return "文件"
				}(),
			})
		}
	}

	return suggestions
}

func lsfiles(argfs string, client *kubernetes.Clientset, namespace string, config *rest.Config, d prompt.Document, dirmode bool) []prompt.Suggest {
	if strings.HasPrefix(argfs, ".") || strings.HasPrefix(argfs, "/") {
		return prompt.FilterContains(completerfile(argfs, dirmode), argfs, true)
		// return completerfile(argfs)
	} else if strings.Contains(argfs, ":") {
		return getpodfiles(client, namespace, config, argfs, dirmode) // return
	} else {
		return prompt.FilterContains(getPodSuggestions(client, namespace), argfs, true)
		// return getPodSuggestions(client, namespace)
	}
}

func getpodfiles(client *kubernetes.Clientset, namespace string, config *rest.Config, podname string, dirOnly bool) []prompt.Suggest {
	var podnames []string
	var cmd []string
	if strings.Contains(podname, ":") {
		podnames = strings.SplitN(podname, ":", 2)
	}
	//if len(podnames) >= 2 {
	//	fmt.Println("冒号前的内容:", podnames[0])
	//} else {
	//	fmt.Println("字符串中没有找到冒号", podnames, podname)
	//}
	if podnames[1] == "" || podnames[1] == "/" {
		cmd = []string{"sh", "-c", "ls -pd /* /.[!.]* 2>/dev/null || true"}
	} else {
		if strings.HasPrefix(podnames[1], "/") {
			// 移除开头的/
			// ls -pd /ro*
			// podnamespath := strings.TrimPrefix(podnames[1], "/")
			// fmt.Println(podnames[1])
			cmd = []string{"sh", "-c", fmt.Sprintf("ls -pd %s* %s.[!.]* 2>/dev/null || true", podnames[1], podnames[1])}
		} else {
			return []prompt.Suggest{}
		}
	}
	req := client.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(podnames[0]).
		Namespace(namespace).
		SubResource("exec")

	req.VersionedParams(&corev1.PodExecOptions{
		// Container: containerName,
		Stdin:   false,
		Stdout:  true,
		Stderr:  true,
		TTY:     false,
		Command: cmd,
	}, scheme.ParameterCodec)
	// 5. 执行命令并获取输出
	executor, err := remotecommand.NewSPDYExecutor(config, "POST", req.URL())
	if err != nil {
		panic(err.Error())
	}
	var stdoutBuffer, stderrBuffer bytes.Buffer
	streamOpts := remotecommand.StreamOptions{
		Stdin:  nil,
		Stdout: &stdoutBuffer,
		Stderr: &stderrBuffer,
		Tty:    false,
	}
	err = executor.Stream(streamOpts)
	if err != nil {
		// fmt.Println("执行命令失败:", err)
		// todo:可以输出debug日志
		return []prompt.Suggest{}
	}
	stdout := stdoutBuffer.String()
	stderr := stderrBuffer.String()
	if stderr != "" {
		// fmt.Println("执行命令失败:", stderr)
		// todo:可以输出debug日志
		return []prompt.Suggest{}
	}
	lines := strings.Split(stdout, "\n")
	// fmt.Println(lines)
	var suggestions []prompt.Suggest
	for _, line := range lines {
		if strings.Contains(line, "[") || line == "" {
			continue
		}
		if strings.HasSuffix(line, "/") {
			suggestions = append(suggestions, prompt.Suggest{
				Text:        line,
				Description: "目录",
			})
		} else if !dirOnly {
			suggestions = append(suggestions, prompt.Suggest{
				Text:        line,
				Description: "文件",
			})
		}
	}
	return suggestions
}

func getPodSuggestions(client *kubernetes.Clientset, namespace string) []prompt.Suggest {
	// go fetchPods(client, namespace)
	fetchPods(client, namespace)
	x, ok := podList.Load(namespace)
	if !ok {
		return []prompt.Suggest{}
	}
	l, ok := x.(*corev1.PodList)
	if !ok || len(l.Items) == 0 {
		fmt.Println("8888888")
		return []prompt.Suggest{}
	}
	s := make([]prompt.Suggest, len(l.Items))
	for i := range l.Items {
		s[i] = prompt.Suggest{
			Text:        l.Items[i].Name,
			Description: string(l.Items[i].Status.Phase),
		}
	}
	return s
}

func getPod(namespace, podName string) (corev1.Pod, bool) {
	x, ok := podList.Load(namespace)
	if !ok {
		return corev1.Pod{}, false
	}
	l, ok := x.(*corev1.PodList)
	if !ok || len(l.Items) == 0 {
		return corev1.Pod{}, false
	}
	for i := range l.Items {
		if podName == l.Items[i].Name {
			return l.Items[i], true
		}
	}
	return corev1.Pod{}, false
}

func getPortsFromPodName(namespace string, podName string) []prompt.Suggest {
	pod, found := getPod(namespace, podName)
	if !found {
		return []prompt.Suggest{}
	}

	// Extract unique ports
	portSet := make(map[int32]struct{})
	for i := range pod.Spec.Containers {
		ports := pod.Spec.Containers[i].Ports
		for j := range ports {
			portSet[ports[j].ContainerPort] = struct{}{}
		}
	}

	// Sort
	var ports []int
	for k := range portSet {
		ports = append(ports, int(k))
	}
	sort.Ints(ports)

	// Prepare suggestions
	suggests := make([]prompt.Suggest, 0, len(ports))
	for i := range ports {
		suggests = append(suggests, prompt.Suggest{
			Text: fmt.Sprintf("%d:%d", ports[i], ports[i]),
		})
	}
	return suggests
}

func getContainerNamesFromCachedPods(client *kubernetes.Clientset, namespace string) []prompt.Suggest {
	go fetchPods(client, namespace)

	x, ok := podList.Load(namespace)
	if !ok {
		return []prompt.Suggest{}
	}
	l, ok := x.(*corev1.PodList)
	if !ok || len(l.Items) == 0 {
		return []prompt.Suggest{}
	}
	// container name -> pod name
	set := make(map[string]string, len(l.Items))
	for i := range l.Items {
		for j := range l.Items[i].Spec.Containers {
			set[l.Items[i].Spec.Containers[j].Name] = l.Items[i].Name
		}
	}
	s := make([]prompt.Suggest, 0, len(set))
	for key := range set {
		s = append(s, prompt.Suggest{
			Text:        key,
			Description: "Pod Name: " + set[key],
		})
	}
	return s
}

func getContainerName(client *kubernetes.Clientset, namespace string, podName string) []prompt.Suggest {
	go fetchPods(client, namespace)

	pod, found := getPod(namespace, podName)
	if !found {
		return []prompt.Suggest{}
	}
	s := make([]prompt.Suggest, len(pod.Spec.Containers))
	for i := range pod.Spec.Containers {
		s[i] = prompt.Suggest{
			Text:        pod.Spec.Containers[i].Name,
			Description: "",
		}
	}
	return s
}

/* Daemon Sets */

var daemonSetList *sync.Map

func fetchDaemonSetList(client *kubernetes.Clientset, namespace string) {
	key := "daemon_" + namespace
	if !shouldFetch(key) {
		return
	}
	updateLastFetchedAt(key)

	l, _ := client.AppsV1().DaemonSets(namespace).List(context.TODO(), metav1.ListOptions{})
	daemonSetList.Store(namespace, l)
	return
}

func getDaemonSetSuggestions(client *kubernetes.Clientset, namespace string) []prompt.Suggest {
	go fetchDaemonSetList(client, namespace)
	x, ok := daemonSetList.Load(namespace)
	if !ok {
		return []prompt.Suggest{}
	}
	l, ok := x.(appsv1.DaemonSetList)
	if !ok || len(l.Items) == 0 {
		return []prompt.Suggest{}
	}
	s := make([]prompt.Suggest, len(l.Items))
	for i := range l.Items {
		s[i] = prompt.Suggest{
			Text: l.Items[i].Name,
		}
	}
	return s
}

/* Deployment */

var deploymentList *sync.Map

func fetchDeployments(client *kubernetes.Clientset, namespace string) {
	key := "deployment_" + namespace
	if !shouldFetch(key) {
		return
	}
	updateLastFetchedAt(key)

	l, _ := client.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	deploymentList.Store(namespace, l)
	return
}

func getDeploymentSuggestions(client *kubernetes.Clientset, namespace string) []prompt.Suggest {
	go fetchDeployments(client, namespace)
	x, ok := deploymentList.Load(namespace)
	if !ok {
		return []prompt.Suggest{}
	}
	l, ok := x.(*appsv1.DeploymentList)
	if !ok || len(l.Items) == 0 {
		return []prompt.Suggest{}
	}
	s := make([]prompt.Suggest, len(l.Items))
	for i := range l.Items {
		s[i] = prompt.Suggest{
			Text: l.Items[i].Name,
		}
	}
	return s
}

/* Endpoint */

var endpointList *sync.Map

func fetchEndpoints(client *kubernetes.Clientset, namespace string) {
	key := "endpoint_" + namespace
	if !shouldFetch(key) {
		return
	}
	updateLastFetchedAt(key)

	l, _ := client.CoreV1().Endpoints(namespace).List(context.TODO(), metav1.ListOptions{})
	endpointList.Store(key, l)
	return
}

func getEndpointsSuggestions(client *kubernetes.Clientset, namespace string) []prompt.Suggest {
	go fetchEndpoints(client, namespace)
	x, ok := endpointList.Load(namespace)
	if !ok {
		return []prompt.Suggest{}
	}
	l, ok := x.(*corev1.EndpointsList)
	if !ok || len(l.Items) == 0 {
		return []prompt.Suggest{}
	}
	s := make([]prompt.Suggest, len(l.Items))
	for i := range l.Items {
		s[i] = prompt.Suggest{
			Text: l.Items[i].Name,
		}
	}
	return s
}

/* Events */

var eventList *sync.Map

func fetchEvents(client *kubernetes.Clientset, namespace string) {
	key := "event_" + namespace
	if !shouldFetch(key) {
		return
	}
	updateLastFetchedAt(key)

	l, _ := client.CoreV1().Events(namespace).List(context.TODO(), metav1.ListOptions{})
	eventList.Store(namespace, l)
	return
}

func getEventsSuggestions(client *kubernetes.Clientset, namespace string) []prompt.Suggest {
	go fetchEvents(client, namespace)
	x, ok := eventList.Load(namespace)
	if !ok {
		return []prompt.Suggest{}
	}
	l, ok := x.(*corev1.EventList)
	if !ok || len(l.Items) == 0 {
		return []prompt.Suggest{}
	}
	s := make([]prompt.Suggest, len(l.Items))
	for i := range l.Items {
		s[i] = prompt.Suggest{
			Text: l.Items[i].Name,
		}
	}
	return s
}

/* Node */

var nodeList atomic.Value

func fetchNodeList(client *kubernetes.Clientset) {
	key := "node"
	if !shouldFetch(key) {
		return
	}
	updateLastFetchedAt(key)

	l, _ := client.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	nodeList.Store(l)
	return
}

func getNodeSuggestions(client *kubernetes.Clientset) []prompt.Suggest {
	go fetchNodeList(client)
	l, ok := nodeList.Load().(*corev1.NodeList)
	if !ok || len(l.Items) == 0 {
		return []prompt.Suggest{}
	}
	s := make([]prompt.Suggest, len(l.Items))
	for i := range l.Items {
		s[i] = prompt.Suggest{
			Text: l.Items[i].Name,
		}
	}
	return s
}

/* Secret */

var secretList *sync.Map

func fetchSecretList(client *kubernetes.Clientset, namespace string) {
	key := "secret_" + namespace
	if !shouldFetch(key) {
		return
	}
	updateLastFetchedAt(key)

	l, _ := client.CoreV1().Secrets(namespace).List(context.TODO(), metav1.ListOptions{})
	secretList.Store(namespace, l)
	return
}

func getSecretSuggestions(client *kubernetes.Clientset, namespace string) []prompt.Suggest {
	go fetchSecretList(client, namespace)
	x, ok := secretList.Load(namespace)
	if !ok {
		return []prompt.Suggest{}
	}
	l, ok := x.(*corev1.SecretList)
	if !ok || len(l.Items) == 0 {
		return []prompt.Suggest{}
	}
	s := make([]prompt.Suggest, len(l.Items))
	for i := range l.Items {
		s[i] = prompt.Suggest{
			Text: l.Items[i].Name,
		}
	}
	return s
}

/* Ingress */

var ingressList *sync.Map

func fetchIngresses(client *kubernetes.Clientset, namespace string) {
	key := "ingress_" + namespace
	if !shouldFetch(key) {
		return
	}
	updateLastFetchedAt(key)

	l, _ := client.ExtensionsV1beta1().Ingresses(namespace).List(context.TODO(), metav1.ListOptions{})
	ingressList.Store(namespace, l)
}

func getIngressSuggestions(client *kubernetes.Clientset, namespace string) []prompt.Suggest {
	go fetchIngresses(client, namespace)

	x, ok := ingressList.Load(namespace)
	if !ok {
		return []prompt.Suggest{}
	}
	l, ok := x.(*extensionsv1beta1.IngressList)
	if !ok {
		debug.Log("must not reach here")
		return []prompt.Suggest{}
	}
	if len(l.Items) == 0 {
		return []prompt.Suggest{}
	}
	s := make([]prompt.Suggest, len(l.Items))
	for i := range l.Items {
		s[i] = prompt.Suggest{
			Text: l.Items[i].Name,
		}
	}
	return s
}

/* LimitRange */

var limitRangeList *sync.Map

func fetchLimitRangeList(client *kubernetes.Clientset, namespace string) {
	key := "limit_range_" + namespace
	if !shouldFetch(key) {
		return
	}
	updateLastFetchedAt(key)

	l, _ := client.CoreV1().LimitRanges(namespace).List(context.TODO(), metav1.ListOptions{})
	limitRangeList.Store(namespace, l)
	return
}

func getLimitRangeSuggestions(client *kubernetes.Clientset, namespace string) []prompt.Suggest {
	go fetchLimitRangeList(client, namespace)
	x, ok := limitRangeList.Load(namespace)
	if !ok {
		return []prompt.Suggest{}
	}
	l, ok := x.(*corev1.NamespaceList)
	if !ok || len(l.Items) == 0 {
		return []prompt.Suggest{}
	}
	s := make([]prompt.Suggest, len(l.Items))
	for i := range l.Items {
		s[i] = prompt.Suggest{
			Text: l.Items[i].Name,
		}
	}
	return s
}

/* NameSpaces */

func getNameSpaceSuggestions(namespaceList *corev1.NamespaceList) []prompt.Suggest {
	if namespaceList == nil || len(namespaceList.Items) == 0 {
		return []prompt.Suggest{}
	}
	s := make([]prompt.Suggest, len(namespaceList.Items))
	for i := range namespaceList.Items {
		s[i] = prompt.Suggest{
			Text: namespaceList.Items[i].Name,
		}
	}
	return s
}

/* Persistent Volume Claims */

var persistentVolumeClaimsList *sync.Map

func fetchPersistentVolumeClaimsList(client *kubernetes.Clientset, namespace string) {
	key := "persistent_volume_claims" + namespace
	if !shouldFetch(key) {
		return
	}
	updateLastFetchedAt(key)

	l, _ := client.CoreV1().PersistentVolumeClaims(namespace).List(context.TODO(), metav1.ListOptions{})
	persistentVolumeClaimsList.Store(namespace, l)
	return
}

func getPersistentVolumeClaimSuggestions(client *kubernetes.Clientset, namespace string) []prompt.Suggest {
	go fetchPersistentVolumeClaimsList(client, namespace)
	x, ok := persistentVolumeClaimsList.Load(namespace)
	if !ok {
		return []prompt.Suggest{}
	}
	l, ok := x.(*corev1.PersistentVolumeClaimList)
	if !ok || len(l.Items) == 0 {
		return []prompt.Suggest{}
	}
	s := make([]prompt.Suggest, len(l.Items))
	for i := range l.Items {
		s[i] = prompt.Suggest{
			Text: l.Items[i].Name,
		}
	}
	return s
}

/* Persistent Volumes */

var persistentVolumesList atomic.Value

func fetchPersistentVolumeList(client *kubernetes.Clientset) {
	key := "persistent_volume"
	if !shouldFetch(key) {
		return
	}
	updateLastFetchedAt(key)

	l, _ := client.CoreV1().PersistentVolumes().List(context.TODO(), metav1.ListOptions{})
	persistentVolumesList.Store(l)
	return
}

func getPersistentVolumeSuggestions(client *kubernetes.Clientset) []prompt.Suggest {
	go fetchPersistentVolumeList(client)
	l, ok := persistentVolumesList.Load().(*corev1.PersistentVolumeList)
	if !ok || len(l.Items) == 0 {
		return []prompt.Suggest{}
	}
	s := make([]prompt.Suggest, len(l.Items))
	for i := range l.Items {
		s[i] = prompt.Suggest{
			Text: l.Items[i].Name,
		}
	}
	return s
}

/* Pod Security Policies */

var podSecurityPolicyList atomic.Value

func fetchPodSecurityPolicyList(client *kubernetes.Clientset) {
	key := "pod_security_policy"
	if !shouldFetch(key) {
		return
	}
	updateLastFetchedAt(key)

	l, _ := client.ExtensionsV1beta1().PodSecurityPolicies().List(context.TODO(), metav1.ListOptions{})
	podSecurityPolicyList.Store(l)
	return
}

func getPodSecurityPolicySuggestions(client *kubernetes.Clientset) []prompt.Suggest {
	go fetchPodSecurityPolicyList(client)
	l, ok := podSecurityPolicyList.Load().(policyv1beta1.PodSecurityPolicyList)
	if !ok || len(l.Items) == 0 {
		return []prompt.Suggest{}
	}
	s := make([]prompt.Suggest, len(l.Items))
	for i := range l.Items {
		s[i] = prompt.Suggest{
			Text: l.Items[i].Name,
		}
	}
	return s
}

/* Pod Templates */

var podTemplateList *sync.Map

func fetchPodTemplateList(client *kubernetes.Clientset, namespace string) {
	key := "pod_template_" + namespace
	if !shouldFetch(key) {
		return
	}
	updateLastFetchedAt(key)

	l, _ := client.CoreV1().PodTemplates(namespace).List(context.TODO(), metav1.ListOptions{})
	podTemplateList.Store(namespace, l)
	return
}

func getPodTemplateSuggestions(client *kubernetes.Clientset, namespace string) []prompt.Suggest {
	go fetchPodTemplateList(client, namespace)
	x, ok := podTemplateList.Load(namespace)
	if !ok {
		return []prompt.Suggest{}
	}
	l, ok := x.(*corev1.PodTemplateList)
	if !ok || len(l.Items) == 0 {
		return []prompt.Suggest{}
	}
	s := make([]prompt.Suggest, len(l.Items))
	for i := range l.Items {
		s[i] = prompt.Suggest{
			Text: l.Items[i].Name,
		}
	}
	return s
}

/* Replica Sets */

var replicaSetList *sync.Map

func fetchReplicaSetList(client *kubernetes.Clientset, namespace string) {
	key := "replica_set_" + namespace
	if !shouldFetch(key) {
		return
	}
	updateLastFetchedAt(key)

	l, _ := client.AppsV1beta2().ReplicaSets(namespace).List(context.TODO(), metav1.ListOptions{})
	replicaSetList.Store(namespace, l)
	return
}

func getReplicaSetSuggestions(client *kubernetes.Clientset, namespace string) []prompt.Suggest {
	go fetchReplicaSetList(client, namespace)
	x, ok := replicaSetList.Load(namespace)
	if !ok {
		return []prompt.Suggest{}
	}
	l, ok := x.(appsv1.ReplicaSetList)
	if !ok || len(l.Items) == 0 {
		return []prompt.Suggest{}
	}
	s := make([]prompt.Suggest, len(l.Items))
	for i := range l.Items {
		s[i] = prompt.Suggest{
			Text: l.Items[i].Name,
		}
	}
	return s
}

/* Replication Controller */

var replicationControllerList *sync.Map

func fetchReplicationControllerList(client *kubernetes.Clientset, namespace string) {
	key := "replication_controller" + namespace
	if !shouldFetch(key) {
		return
	}
	updateLastFetchedAt(key)

	l, _ := client.CoreV1().ReplicationControllers(namespace).List(context.TODO(), metav1.ListOptions{})
	replicationControllerList.Store(namespace, l)
	return
}

func getReplicationControllerSuggestions(client *kubernetes.Clientset, namespace string) []prompt.Suggest {
	go fetchReplicationControllerList(client, namespace)
	x, ok := replicationControllerList.Load(namespace)
	if !ok {
		return []prompt.Suggest{}
	}
	l, ok := x.(*corev1.ReplicationControllerList)
	if !ok || len(l.Items) == 0 {
		return []prompt.Suggest{}
	}
	s := make([]prompt.Suggest, len(l.Items))
	for i := range l.Items {
		s[i] = prompt.Suggest{
			Text: l.Items[i].Name,
		}
	}
	return s
}

/* Resource quotas */

var resourceQuotaList *sync.Map

func fetchResourceQuotaList(client *kubernetes.Clientset, namespace string) {
	key := "resource_quota" + namespace
	if !shouldFetch(key) {
		return
	}
	updateLastFetchedAt(key)

	l, _ := client.CoreV1().ResourceQuotas(namespace).List(context.TODO(), metav1.ListOptions{})
	resourceQuotaList.Store(namespace, l)
	return
}

func getResourceQuotasSuggestions(client *kubernetes.Clientset, namespace string) []prompt.Suggest {
	go fetchResourceQuotaList(client, namespace)
	x, ok := resourceQuotaList.Load(namespace)
	if !ok {
		return []prompt.Suggest{}
	}
	l, ok := x.(*corev1.ResourceQuotaList)
	if !ok || len(l.Items) == 0 {
		return []prompt.Suggest{}
	}
	s := make([]prompt.Suggest, len(l.Items))
	for i := range l.Items {
		s[i] = prompt.Suggest{
			Text: l.Items[i].Name,
		}
	}
	return s
}

/* Service Account */

var serviceAccountList *sync.Map

func fetchServiceAccountList(client *kubernetes.Clientset, namespace string) {
	key := "service_account_" + namespace
	if !shouldFetch(key) {
		return
	}
	updateLastFetchedAt(key)

	l, _ := client.CoreV1().ServiceAccounts(namespace).List(context.TODO(), metav1.ListOptions{})
	serviceAccountList.Store(namespace, l)
	return
}

func getServiceAccountSuggestions(client *kubernetes.Clientset, namespace string) []prompt.Suggest {
	go fetchServiceAccountList(client, namespace)
	x, ok := serviceAccountList.Load(namespace)
	if !ok {
		return []prompt.Suggest{}
	}
	l, ok := x.(*corev1.ServiceAccountList)
	if !ok || len(l.Items) == 0 {
		return []prompt.Suggest{}
	}
	s := make([]prompt.Suggest, len(l.Items))
	for i := range l.Items {
		s[i] = prompt.Suggest{
			Text: l.Items[i].Name,
		}
	}
	return s
}

/* Service */

var serviceList *sync.Map

func fetchServiceList(client *kubernetes.Clientset, namespace string) {
	key := "service_" + namespace
	if !shouldFetch(key) {
		return
	}
	updateLastFetchedAt(key)

	l, _ := client.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	serviceList.Store(namespace, l)
	return
}

func getServiceSuggestions(client *kubernetes.Clientset, namespace string) []prompt.Suggest {
	go fetchServiceList(client, namespace)
	x, ok := serviceList.Load(namespace)
	if !ok {
		return []prompt.Suggest{}
	}
	l, ok := x.(*corev1.ServiceList)
	if !ok || len(l.Items) == 0 {
		return []prompt.Suggest{}
	}
	s := make([]prompt.Suggest, len(l.Items))
	for i := range l.Items {
		s[i] = prompt.Suggest{
			Text: l.Items[i].Name,
		}
	}
	return s
}

/* Job */

var jobList *sync.Map

func fetchJobs(client *kubernetes.Clientset, namespace string) {
	key := "job_" + namespace
	if !shouldFetch(key) {
		return
	}
	updateLastFetchedAt(key)

	l, _ := client.BatchV1().Jobs(namespace).List(context.TODO(), metav1.ListOptions{})
	jobList.Store(namespace, l)
}

func getJobSuggestions(client *kubernetes.Clientset, namespace string) []prompt.Suggest {
	go fetchJobs(client, namespace)
	x, ok := jobList.Load(namespace)
	if !ok {
		return []prompt.Suggest{}
	}
	l, ok := x.(*batchv1.JobList)
	if !ok || len(l.Items) == 0 {
		return []prompt.Suggest{}
	}
	s := make([]prompt.Suggest, len(l.Items))
	for i := range l.Items {
		s[i] = prompt.Suggest{
			Text:        l.Items[i].Name,
			Description: l.Items[i].Status.StartTime.String(),
		}
	}
	return s
}
