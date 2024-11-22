package main

import (
	"fmt"
	"os"

	"github.com/c-bata/go-prompt"
	"github.com/jiqinga/kube-prompt/internal/debug"
	"github.com/jiqinga/kube-prompt/kube"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	//_ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

var version string

func main() {
	c, err := kube.NewCompleter("")
	if err != nil {
		fmt.Println("error", err)
		// 恢复终端状态
		// 通知主程序优雅退出
		close(kube.ExitChan)
	}
	defer debug.Teardown()
	fmt.Printf("kube-prompt %s \n", version)
	fmt.Println("请使用 `exit` or `Ctrl-D` 退出此程序.")
	defer fmt.Println("Bye!")
	// 添加一个函数来获取当前命名空间
	getCurrentNamespace := func() string {
		ns, ok := kube.GetGlobalState("namespace")
		if !ok || ns == "" {
			return "default"
		}
		return ns.(string)
	}
	p := prompt.New(
		kube.Executor,
		c.Complete,
		prompt.OptionTitle("kube-prompt: interactive kubernetes client"),
		// prompt.OptionPrefix(">>> "),
		prompt.OptionPrefix(getCurrentNamespace()+">>> "), // 初始提示符
		prompt.OptionLivePrefix(func() (string, bool) {
			return getCurrentNamespace() + ">>> ", true
		}),
		prompt.OptionCompletionWordSeparator(string([]byte{' ', ':'})),
		// prompt.OptionCompletionWordSeparator(completer.FilePathCompletionSeparator),
		prompt.OptionInputTextColor(prompt.Yellow),
	)
	p.Run()
	defer func() {
		// 删除临时文件
		os.Remove(kube.Getns("KUBECONFIG"))
	}()
}
