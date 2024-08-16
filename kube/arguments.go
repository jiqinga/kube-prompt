package kube

import prompt "github.com/c-bata/go-prompt"

var commands = []prompt.Suggest{
	{Text: "get", Description: "显示一个或多个资源"},
	{Text: "describe", Description: "显示特定资源或资源组的详细信息"},
	{Text: "create", Description: "按文件名或标准输入创建资源"},
	{Text: "replace", Description: "通过文件名或标准输入来替换资源"},
	{Text: "patch", Description: "使用策略合并补丁更新资源的字段"},
	{Text: "delete", Description: "通过文件名、标准输入、资源和名称，或者通过资源和标签选择器删除资源 "},
	{Text: "edit", Description: "在服务器上编辑资源"},
	{Text: "apply", Description: "通过文件名或标准输入将配置应用于资源 "},
	{Text: "namespace", Description: "已过时：设置并查看当前的 Kubernetes 命名空间"},
	{Text: "logs", Description: "打印一个 Pod 中容器的日志"},
	{Text: "rolling-update", Description: "对给定的ReplicationController执行滚动更新"},
	{Text: "scale", Description: "为Deployment, ReplicaSet, Replication Controller,Job设置新的实例数"},
	{Text: "cordon", Description: "将节点标记为不可调度"},
	{Text: "drain", Description: "为维护做准备而排空节点,驱逐节点上所有Pod"},
	{Text: "uncordon", Description: "将节点标记为可调度"},
	{Text: "attach", Description: "连接到正在运行的容器"},
	{Text: "exec", Description: "进入容器中执行命令"},
	{Text: "port-forward", Description: "将一个或多个本地端口转发到一个Pod"},
	{Text: "proxy", Description: "运行到 Kubernetes API 服务器的代理 "},
	{Text: "run", Description: "在集群上运行特定的镜像"},
	{Text: "expose", Description: "获取一个controller, service, pod 并将其作为新的 Kubernetes 服务进行公开"},
	{Text: "expose", Description: "获取一个controller, service, pod 并将其作为新的 Kubernetes 服务进行公开"},
	{Text: "autoscale", Description: "自动缩放Deployment, ReplicaSet, ReplicationController"},
	{Text: "rollout", Description: "管理资源"},
	{Text: "label", Description: "更新资源上的标签"},
	{Text: "annotate", Description: "更新资源上的注释"},
	{Text: "config", Description: "配置修改 kubeconfig 文件"},
	{Text: "cluster-info", Description: "显示集群信息"},
	{Text: "api-versions", Description: "在服务器上以'group/version'的形式打印受支持的 API 版本."},
	{Text: "version", Description: "打印客户端和服务端版本信息"},
	{Text: "explain", Description: "资源文档"},
	{Text: "convert", Description: "在不同的 API 版本之间转换配置文件 "},
	{Text: "top", Description: "显示资源 (CPU/Memory/Storage) 使用情况"},
	{Text: "kustomize", Description: "从目录或远程 URL 构建一个自定义（kustomization）目标 "},

	// Custom command.
	{Text: "exit", Description: "退出此程序"},
}

var resourceTypes = []prompt.Suggest{
	{Text: "clusters"}, // valid only for federation apiservers
	{Text: "componentstatuses"},
	{Text: "configmaps"},
	{Text: "daemonsets"},
	{Text: "deployments"},
	{Text: "endpoints"},
	{Text: "events"},
	{Text: "horizontalpodautoscalers"},
	{Text: "ingresses"},
	{Text: "jobs"},
	{Text: "cronjobs"},
	{Text: "limitranges"},
	{Text: "namespaces"},
	{Text: "networkpolicies"},
	{Text: "nodes"},
	{Text: "persistentvolumeclaims"},
	{Text: "persistentvolumes"},
	{Text: "pod"},
	{Text: "podsecuritypolicies"},
	{Text: "podtemplates"},
	{Text: "replicasets"},
	{Text: "replicationcontrollers"},
	{Text: "resourcequotas"},
	{Text: "secrets"},
	{Text: "serviceaccounts"},
	{Text: "services"},
	{Text: "statefulsets"},
	{Text: "storageclasses"},
	{Text: "thirdpartyresources"},

	// aliases
	{Text: "cs"},
	{Text: "cm"},
	{Text: "ds"},
	{Text: "deploy"},
	{Text: "ep"},
	{Text: "hpa"},
	{Text: "ing"},
	{Text: "limits"},
	{Text: "ns"},
	{Text: "no"},
	{Text: "pvc"},
	{Text: "pv"},
	{Text: "po"},
	{Text: "psp"},
	{Text: "rs"},
	{Text: "rc"},
	{Text: "quota"},
	{Text: "sa"},
	{Text: "svc"},
}

func (c *Completer) argumentsCompleter(namespace string, args []string) []prompt.Suggest {
	if len(args) <= 1 {
		return prompt.FilterHasPrefix(commands, args[0], true)
	}

	first := args[0]
	switch first {
	case "get":
		second := args[1]
		if len(args) == 2 {
			subcommands := []prompt.Suggest{
				{Text: "componentstatuses"},
				{Text: "configmaps"},
				{Text: "daemonsets"},
				{Text: "deployments"},
				{Text: "endpoints"},
				{Text: "events"},
				{Text: "horizontalpodautoscalers"},
				{Text: "ingresses"},
				{Text: "jobs"},
				{Text: "cronjobs"},
				{Text: "limitranges"},
				{Text: "namespaces"},
				{Text: "networkpolicies"},
				{Text: "nodes"},
				{Text: "persistentvolumeclaims"},
				{Text: "persistentvolumes"},
				{Text: "pod"},
				{Text: "podsecuritypolicies"},
				{Text: "podtemplates"},
				{Text: "replicasets"},
				{Text: "replicationcontrollers"},
				{Text: "resourcequotas"},
				{Text: "secrets"},
				{Text: "serviceaccounts"},
				{Text: "services"},
				{Text: "statefulsets"},
				{Text: "storageclasses"},
				{Text: "thirdpartyresources"},
				// aliases
				{Text: "cs"},
				{Text: "cm"},
				{Text: "ds"},
				{Text: "deploy"},
				{Text: "ep"},
				{Text: "hpa"},
				{Text: "ing"},
				{Text: "limits"},
				{Text: "ns"},
				{Text: "no"},
				{Text: "pvc"},
				{Text: "pv"},
				{Text: "po"},
				{Text: "psp"},
				{Text: "rs"},
				{Text: "rc"},
				{Text: "quota"},
				{Text: "sa"},
				{Text: "svc"},
			}
			return prompt.FilterHasPrefix(subcommands, second, true)
		}

		third := args[2]
		if len(args) == 3 {
			switch second {
			case "componentstatuses", "cs":
				return prompt.FilterContains(getComponentStatusCompletions(c.client), third, true)
			case "configmaps", "cm":
				return prompt.FilterContains(getConfigMapSuggestions(c.client, namespace), third, true)
			case "daemonsets", "ds":
				return prompt.FilterContains(getDaemonSetSuggestions(c.client, namespace), third, true)
			case "deploy", "deployments":
				return prompt.FilterContains(getDeploymentSuggestions(c.client, namespace), third, true)
			case "endpoints", "ep":
				return prompt.FilterContains(getEndpointsSuggestions(c.client, namespace), third, true)
			case "ingresses", "ing":
				return prompt.FilterContains(getIngressSuggestions(c.client, namespace), third, true)
			case "limitranges", "limits":
				return prompt.FilterContains(getLimitRangeSuggestions(c.client, namespace), third, true)
			case "namespaces", "ns":
				return prompt.FilterContains(getNameSpaceSuggestions(c.namespaceList), third, true)
			case "no", "nodes":
				return prompt.FilterContains(getNodeSuggestions(c.client), third, true)
			case "po", "pod", "pods":
				return prompt.FilterContains(getPodSuggestions(c.client, namespace), third, true)
			case "persistentvolumeclaims", "pvc":
				return prompt.FilterContains(getPersistentVolumeClaimSuggestions(c.client, namespace), third, true)
			case "persistentvolumes", "pv":
				return prompt.FilterContains(getPersistentVolumeSuggestions(c.client), third, true)
			case "podsecuritypolicies", "psp":
				return prompt.FilterContains(getPodSecurityPolicySuggestions(c.client), third, true)
			case "podtemplates":
				return prompt.FilterContains(getPodTemplateSuggestions(c.client, namespace), third, true)
			case "replicasets", "rs":
				return prompt.FilterContains(getReplicaSetSuggestions(c.client, namespace), third, true)
			case "replicationcontrollers", "rc":
				return prompt.FilterContains(getReplicationControllerSuggestions(c.client, namespace), third, true)
			case "resourcequotas", "quota":
				return prompt.FilterContains(getResourceQuotasSuggestions(c.client, namespace), third, true)
			case "secrets":
				return prompt.FilterContains(getSecretSuggestions(c.client, namespace), third, true)
			case "sa", "serviceaccounts":
				return prompt.FilterContains(getServiceAccountSuggestions(c.client, namespace), third, true)
			case "svc", "services":
				return prompt.FilterContains(getServiceSuggestions(c.client, namespace), third, true)
			case "job", "jobs":
				return prompt.FilterContains(getJobSuggestions(c.client, namespace), third, true)
			}
		}
	case "describe":
		second := args[1]
		if len(args) == 2 {
			return prompt.FilterHasPrefix(resourceTypes, second, true)
		}

		third := args[2]
		if len(args) == 3 {
			switch second {
			case "componentstatuses", "cs":
				return prompt.FilterContains(getComponentStatusCompletions(c.client), third, true)
			case "configmaps", "cm":
				return prompt.FilterContains(getConfigMapSuggestions(c.client, namespace), third, true)
			case "daemonsets", "ds":
				return prompt.FilterContains(getDaemonSetSuggestions(c.client, namespace), third, true)
			case "deploy", "deployments":
				return prompt.FilterContains(getDeploymentSuggestions(c.client, namespace), third, true)
			case "endpoints", "ep":
				return prompt.FilterContains(getEndpointsSuggestions(c.client, namespace), third, true)
			case "ingresses", "ing":
				return prompt.FilterContains(getIngressSuggestions(c.client, namespace), third, true)
			case "limitranges", "limits":
				return prompt.FilterContains(getLimitRangeSuggestions(c.client, namespace), third, true)
			case "namespaces", "ns":
				return prompt.FilterContains(getNameSpaceSuggestions(c.namespaceList), third, true)
			case "no", "nodes":
				return prompt.FilterContains(getNodeSuggestions(c.client), third, true)
			case "po", "pod", "pods":
				return prompt.FilterContains(getPodSuggestions(c.client, namespace), third, true)
			case "persistentvolumeclaims", "pvc":
				return prompt.FilterContains(getPersistentVolumeClaimSuggestions(c.client, namespace), third, true)
			case "persistentvolumes", "pv":
				return prompt.FilterContains(getPersistentVolumeSuggestions(c.client), third, true)
			case "podsecuritypolicies", "psp":
				return prompt.FilterContains(getPodSecurityPolicySuggestions(c.client), third, true)
			case "podtemplates":
				return prompt.FilterContains(getPodTemplateSuggestions(c.client, namespace), third, true)
			case "replicasets", "rs":
				return prompt.FilterContains(getReplicaSetSuggestions(c.client, namespace), third, true)
			case "replicationcontrollers", "rc":
				return prompt.FilterContains(getReplicationControllerSuggestions(c.client, namespace), third, true)
			case "resourcequotas", "quota":
				return prompt.FilterContains(getResourceQuotasSuggestions(c.client, namespace), third, true)
			case "secrets":
				return prompt.FilterContains(getSecretSuggestions(c.client, namespace), third, true)
			case "sa", "serviceaccounts":
				return prompt.FilterContains(getServiceAccountSuggestions(c.client, namespace), third, true)
			case "svc", "services":
				return prompt.FilterContains(getServiceSuggestions(c.client, namespace), third, true)
			case "job", "jobs":
				return prompt.FilterContains(getJobSuggestions(c.client, namespace), third, true)
			}
		}
	case "create":
		subcommands := []prompt.Suggest{
			{Text: "configmap", Description: "Create a configmap from a local file, directory or literal value"},
			{Text: "deployment", Description: "Create a deployment with the specified name."},
			{Text: "namespace", Description: "Create a namespace with the specified name"},
			{Text: "quota", Description: "Create a quota with the specified name."},
			{Text: "secret", Description: "Create a secret using specified subcommand"},
			{Text: "service", Description: "Create a service using specified subcommand."},
			{Text: "serviceaccount", Description: "Create a service account with the specified name"},
		}
		if len(args) == 2 {
			return prompt.FilterHasPrefix(subcommands, args[1], true)
		}
	case "delete":
		second := args[1]
		if len(args) == 2 {
			return prompt.FilterHasPrefix(resourceTypes, second, true)
		}

		third := args[2]
		if len(args) == 3 {
			switch second {
			case "componentstatuses", "cs":
				return prompt.FilterContains(getComponentStatusCompletions(c.client), third, true)
			case "configmaps", "cm":
				return prompt.FilterContains(getConfigMapSuggestions(c.client, namespace), third, true)
			case "daemonsets", "ds":
				return prompt.FilterContains(getDaemonSetSuggestions(c.client, namespace), third, true)
			case "deploy", "deployments":
				return prompt.FilterContains(getDeploymentSuggestions(c.client, namespace), third, true)
			case "endpoints", "ep":
				return prompt.FilterContains(getEndpointsSuggestions(c.client, namespace), third, true)
			case "ingresses", "ing":
				return prompt.FilterContains(getIngressSuggestions(c.client, namespace), third, true)
			case "limitranges", "limits":
				return prompt.FilterContains(getLimitRangeSuggestions(c.client, namespace), third, true)
			case "namespaces", "ns":
				return prompt.FilterContains(getNameSpaceSuggestions(c.namespaceList), third, true)
			case "no", "nodes":
				return prompt.FilterContains(getNodeSuggestions(c.client), third, true)
			case "po", "pod", "pods":
				return prompt.FilterContains(getPodSuggestions(c.client, namespace), third, true)
			case "persistentvolumeclaims", "pvc":
				return prompt.FilterContains(getPersistentVolumeClaimSuggestions(c.client, namespace), third, true)
			case "persistentvolumes", "pv":
				return prompt.FilterContains(getPersistentVolumeSuggestions(c.client), third, true)
			case "podsecuritypolicies", "psp":
				return prompt.FilterContains(getPodSecurityPolicySuggestions(c.client), third, true)
			case "podtemplates":
				return prompt.FilterContains(getPodTemplateSuggestions(c.client, namespace), third, true)
			case "replicasets", "rs":
				return prompt.FilterContains(getReplicaSetSuggestions(c.client, namespace), third, true)
			case "replicationcontrollers", "rc":
				return prompt.FilterContains(getReplicationControllerSuggestions(c.client, namespace), third, true)
			case "resourcequotas", "quota":
				return prompt.FilterContains(getResourceQuotasSuggestions(c.client, namespace), third, true)
			case "secrets":
				return prompt.FilterContains(getSecretSuggestions(c.client, namespace), third, true)
			case "sa", "serviceaccounts":
				return prompt.FilterContains(getServiceAccountSuggestions(c.client, namespace), third, true)
			case "svc", "services":
				return prompt.FilterContains(getServiceSuggestions(c.client, namespace), third, true)
			case "job", "jobs":
				return prompt.FilterContains(getJobSuggestions(c.client, namespace), third, true)
			}
		}
	case "edit":
		if len(args) == 2 {
			return prompt.FilterHasPrefix(resourceTypes, args[1], true)
		}

		if len(args) == 3 {
			third := args[2]
			switch args[1] {
			case "componentstatuses", "cs":
				return prompt.FilterContains(getComponentStatusCompletions(c.client), third, true)
			case "configmaps", "cm":
				return prompt.FilterContains(getConfigMapSuggestions(c.client, namespace), third, true)
			case "daemonsets", "ds":
				return prompt.FilterContains(getDaemonSetSuggestions(c.client, namespace), third, true)
			case "deploy", "deployments":
				return prompt.FilterContains(getDeploymentSuggestions(c.client, namespace), third, true)
			case "endpoints", "ep":
				return prompt.FilterContains(getEndpointsSuggestions(c.client, namespace), third, true)
			case "ingresses", "ing":
				return prompt.FilterContains(getIngressSuggestions(c.client, namespace), third, true)
			case "limitranges", "limits":
				return prompt.FilterContains(getLimitRangeSuggestions(c.client, namespace), third, true)
			case "namespaces", "ns":
				return prompt.FilterContains(getNameSpaceSuggestions(c.namespaceList), third, true)
			case "no", "nodes":
				return prompt.FilterContains(getNodeSuggestions(c.client), third, true)
			case "po", "pod", "pods":
				return prompt.FilterContains(getPodSuggestions(c.client, namespace), third, true)
			case "persistentvolumeclaims", "pvc":
				return prompt.FilterContains(getPersistentVolumeClaimSuggestions(c.client, namespace), third, true)
			case "persistentvolumes", "pv":
				return prompt.FilterContains(getPersistentVolumeSuggestions(c.client), third, true)
			case "podsecuritypolicies", "psp":
				return prompt.FilterContains(getPodSecurityPolicySuggestions(c.client), third, true)
			case "podtemplates":
				return prompt.FilterContains(getPodTemplateSuggestions(c.client, namespace), third, true)
			case "replicasets", "rs":
				return prompt.FilterContains(getReplicaSetSuggestions(c.client, namespace), third, true)
			case "replicationcontrollers", "rc":
				return prompt.FilterContains(getReplicationControllerSuggestions(c.client, namespace), third, true)
			case "resourcequotas", "quota":
				return prompt.FilterContains(getResourceQuotasSuggestions(c.client, namespace), third, true)
			case "secrets":
				return prompt.FilterContains(getSecretSuggestions(c.client, namespace), third, true)
			case "sa", "serviceaccounts":
				return prompt.FilterContains(getServiceAccountSuggestions(c.client, namespace), third, true)
			case "svc", "services":
				return prompt.FilterContains(getServiceSuggestions(c.client, namespace), third, true)
			case "job", "jobs":
				return prompt.FilterContains(getJobSuggestions(c.client, namespace), third, true)
			}
		}

	case "namespace":
		if len(args) == 2 {
			return prompt.FilterContains(getNameSpaceSuggestions(c.namespaceList), args[1], true)
		}
	case "logs":
		if len(args) == 2 {
			return prompt.FilterContains(getPodSuggestions(c.client, namespace), args[1], true)
		}
	case "rolling-update", "rollingupdate":
		if len(args) == 2 {
			return prompt.FilterContains(getReplicationControllerSuggestions(c.client, namespace), args[1], true)
		} else if len(args) == 3 {
			return prompt.FilterContains(getReplicationControllerSuggestions(c.client, namespace), args[2], true)
		}
	case "scale", "resize":
		if len(args) == 2 {
			// Deployment, ReplicaSet, Replication Controller, or Job.
			r := getDeploymentSuggestions(c.client, namespace)
			r = append(r, getReplicaSetSuggestions(c.client, namespace)...)
			r = append(r, getReplicationControllerSuggestions(c.client, namespace)...)
			return prompt.FilterContains(r, args[1], true)
		}
	case "cordon":
		fallthrough
	case "drain":
		fallthrough
	case "uncordon":
		if len(args) == 2 {
			return prompt.FilterHasPrefix(getNodeSuggestions(c.client), args[1], true)
		}
	case "attach":
		if len(args) == 2 {
			return prompt.FilterContains(getPodSuggestions(c.client, namespace), args[1], true)
		}
	case "exec":
		if len(args) == 2 {
			return prompt.FilterContains(getPodSuggestions(c.client, namespace), args[1], true)
		}
	case "port-forward":
		if len(args) == 2 {
			return prompt.FilterContains(getPodSuggestions(c.client, namespace), args[1], true)
		}
		if len(args) == 3 {
			return prompt.FilterHasPrefix(getPortsFromPodName(namespace, args[1]), args[2], true)
		}
	case "rollout":
		subCommands := []prompt.Suggest{
			{Text: "history", Description: "view rollout history"},
			{Text: "pause", Description: "Mark the provided resource as paused"},
			{Text: "resume", Description: "Resume a paused resource"},
			{Text: "undo", Description: "undoes a previous rollout"},
		}
		if len(args) == 2 {
			return prompt.FilterHasPrefix(subCommands, args[1], true)
		}
	case "annotate":
	case "config":
		subCommands := []prompt.Suggest{
			{Text: "current-context", Description: "Displays the current-context"},
			{Text: "delete-cluster", Description: "Delete the specified cluster from the kubeconfig"},
			{Text: "delete-context", Description: "Delete the specified context from the kubeconfig"},
			{Text: "get-clusters", Description: "Display clusters defined in the kubeconfig"},
			{Text: "get-contexts", Description: "Describe one or many contexts"},
			{Text: "set", Description: "Sets an individual value in a kubeconfig file"},
			{Text: "set-cluster", Description: "Sets a cluster entry in kubeconfig"},
			{Text: "set-context", Description: "Sets a context entry in kubeconfig"},
			{Text: "set-credentials", Description: "Sets a user entry in kubeconfig"},
			{Text: "unset", Description: "Unsets an individual value in a kubeconfig file"},
			{Text: "use-context", Description: "Sets the current-context in a kubeconfig file"},
			{Text: "view", Description: "Display merged kubeconfig settings or a specified kubeconfig file"},
		}
		if len(args) == 2 {
			return prompt.FilterHasPrefix(subCommands, args[1], true)
		}
		if len(args) == 3 {
			third := args[2]
			switch args[1] {
			case "use-context":
				return prompt.FilterContains(getContextSuggestions(), third, true)
			}
		}
	case "cluster-info":
		subCommands := []prompt.Suggest{
			{Text: "dump", Description: "Dump lots of relevant info for debugging and diagnosis"},
		}
		if len(args) == 2 {
			return prompt.FilterHasPrefix(subCommands, args[1], true)
		}
	case "explain":
		return prompt.FilterHasPrefix(resourceTypes, args[1], true)
	case "top":
		second := args[1]
		if len(args) == 2 {
			subcommands := []prompt.Suggest{
				{Text: "nodes"},
				{Text: "pod"},
				// aliases
				{Text: "no"},
				{Text: "po"},
			}
			return prompt.FilterHasPrefix(subcommands, second, true)
		}

		third := args[2]
		if len(args) == 3 {
			switch second {
			case "no", "node", "nodes":
				return prompt.FilterContains(getNodeSuggestions(c.client), third, true)
			case "po", "pod", "pods":
				return prompt.FilterContains(getPodSuggestions(c.client, namespace), third, true)
			}
		}
	default:
		return []prompt.Suggest{}
	}
	return []prompt.Suggest{}
}
