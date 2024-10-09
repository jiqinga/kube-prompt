package kube

import (
	"fmt"
	"io/ioutil"
	"os"

	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

func createTempKubeconfig(namespaceName string) (string, error) {
	// 加载当前的 kubeconfig
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	config, err := loadingRules.Load()
	if err != nil {
		return "", fmt.Errorf("failed to load kubeconfig: %v", err)
	}

	// 获取当前上下文
	currentContext := config.CurrentContext
	context, exists := config.Contexts[currentContext]
	if !exists {
		return "", fmt.Errorf("current context does not exist")
	}

	// 创建新的上下文,使用指定的命名空间
	newContext := *context // 复制当前上下文
	newContext.Namespace = namespaceName
	// 创建新的 kubeconfig，只包含指定的上下文
	newConfig := clientcmdapi.NewConfig()
	newConfig.Contexts = map[string]*clientcmdapi.Context{
		currentContext: &newContext,
	}
	newConfig.CurrentContext = currentContext

	// 添加必要的集群和认证信息
	newConfig.Clusters = map[string]*clientcmdapi.Cluster{
		context.Cluster: config.Clusters[context.Cluster],
	}
	newConfig.AuthInfos = map[string]*clientcmdapi.AuthInfo{
		context.AuthInfo: config.AuthInfos[context.AuthInfo],
	}

	// 创建临时文件
	tempFile, err := ioutil.TempFile(os.TempDir(), ".kubeconfig-*.yaml")
	if err != nil {
		return "", fmt.Errorf("failed to create temp file: %v", err)
	}
	tempFile.Close()

	// 将新的 kubeconfig 写入临时文件
	err = clientcmd.WriteToFile(*newConfig, tempFile.Name())
	if err != nil {
		os.Remove(tempFile.Name())
		return "", fmt.Errorf("failed to write kubeconfig: %v", err)
	}

	return tempFile.Name(), nil
}
