package kube

// import "C"
import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/jiqinga/kube-prompt/internal/debug"

	"github.com/jiqinga/kubecolor/command"
)

func ExecuteAndGetResults(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		debug.Log("you need to pass the something arguments")
		return ""
	}

	out := &bytes.Buffer{}
	cmd := exec.Command("/bin/sh", "-c", "kubectl "+s)
	cmd.Stdin = os.Stdin
	cmd.Stdout = out
	if err := cmd.Run(); err != nil {
		debug.Log(err.Error())
		return ""
	}
	r := string(out.Bytes())
	return r
}

func color(kargs, env []string) string {
	_ = command.Run(kargs[1:], Version, env)
	//if err != nil {
	//	var ke *command.KubectlError
	//	if errors.As(err, &ke) {
	//		os.Exit(ke.ExitCode)
	//	}
	//	os.Exit(1)
	//}
	return ""
}

func Executor(s string) {
	s = strings.TrimSpace(s)

	if s == "" {
		return
	} else if s == "quit" || s == "exit" {
		// 修复退出时导致终端字符不显示
		cmd := exec.Command("reset")
		_ = cmd.Run()
		fmt.Println("Bye!")

		os.Remove(Getns("KUBECONFIG"))

		os.Exit(0)
		return
	}
	if o := strings.Fields(s); o[0] == "set" {
		if c := strings.Fields(s); c[1] == "ns" || c[1] == "namespace" {
			tempfile, err := createTempKubeconfig(c[2])
			if err != nil {
				fmt.Println(err)
				return
			}
			if err != nil {
				fmt.Println("error", err)
				os.Exit(1)
			}
			// completer := &Completer{} // 假设你已经有了这个实例
			// completer.namespace = "s[:len(s)-1]"
			globalState.Store("KUBECONFIG", tempfile)
			globalState.Store("namespace", c[2])
			fmt.Printf("成功切换到命名空间: %s\n", c[2])

			return
		}
	}
	env := os.Environ()
	value, ok := globalState.Load("KUBECONFIG")
	if ok {
		// 键 "KUBECONFIG" 存在
		env = append(env, fmt.Sprintf("KUBECONFIG=%s", value))
	}
	color(strings.Fields("kubectl "+s), env)
	//
	//shell := os.Getenv("SHELL")
	//if shell == "" {
	//	shell = "/bin/sh"
	//}
	//
	//// 获取当前工作目录
	//pwd, err := os.Getwd()
	//if err != nil {
	//	pwd = "/"
	//}
	//cmd := exec.Command(shell, "-c", "kubectl "+s)
	//cmd.Env = env
	//cmd.Dir = pwd
	//cmd.Stdin = os.Stdin
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	//if err := cmd.Run(); err != nil {
	//	fmt.Printf("Got error: %s\n", err.Error())
	//}

	return
}

// 添加这个新函数
func GetGlobalState(key string) (interface{}, bool) {
	return globalState.Load(key)
}
