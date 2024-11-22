package kube

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/jiqinga/kube-prompt/internal/debug"
)

func handleExit() {
	rawModeOff := exec.Command("/bin/stty", "-raw", "echo")
	rawModeOff.Stdin = os.Stdin
	_ = rawModeOff.Run()
	err := rawModeOff.Wait()
	if err != nil {
		return
	}
}

var ExitChan = make(chan os.Signal, 1)

func init() {
	signal.Notify(ExitChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ExitChan
		// 恢复终端设置
		// 恢复终端设置到初始状态
		handleExit()
		fmt.Println("Bye!")
		if err := os.Remove(Getns("KUBECONFIG")); err != nil {
			debug.Log("清理临时配置文件失败: " + err.Error())
		}
		os.Exit(0)
	}()
}
