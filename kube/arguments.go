package kube

import (
	"github.com/c-bata/go-prompt"
)

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
	{Text: "cp", Description: "在容器与外部之间复制文件和目录。"},
	{Text: "kustomize", Description: "从目录或远程 URL 构建一个自定义（kustomization）目标 "},
	{Text: "set", Description: "设置各种 Kubernetes 资源"},
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
	{Text: "pod", Description: "显示一个或多个pod"},
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

func (c *Completer) argumentsCompleter(namespace string, args []string, d prompt.Document) []prompt.Suggest {
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
	case "set":
		second := args[1]
		if len(args) == 2 {
			subcommands := []prompt.Suggest{
				{Text: "ns", Description: "设置默认命名空间"},
				{Text: "namespace", Description: "设置默认命名空间"},
				{Text: "env", Description: "更新 Pod 模板上的环境变量。"},
				{Text: "image", Description: "更新 Pod 模板的镜像。"},
				{Text: "resources", Description: "更新带有 Pod 模板的对象上的资源请求和限制。"},
				{Text: "selector", Description: "设置资源上的选择器"},
				{Text: "serviceaccount", Description: "更新资源的服务账号"},
				{Text: "subject", Description: "更新角色绑定（RoleBinding）或集群角色绑定（ClusterRoleBinding）中的用户、组或服务账号"},
			}
			return prompt.FilterHasPrefix(subcommands, second, true)
		}
		third := args[2]
		if len(args) == 3 {
			switch second {
			case "ns", "namespace":

				return prompt.FilterContains(getNameSpaceSuggestions(c.namespaceList), third, true)
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
			{Text: "configmap", Description: "从本地文件、目录或者直接给定的字面量值创建一个配置映射（ConfigMap）。"},
			{Text: "deployment", Description: "创建一个具有指定名称的部署（Deployment）。"},
			{Text: "namespace", Description: "创建一个具有指定名称的命名空间。"},
			{Text: "quota", Description: "创建一个具有指定名称的资源配额。"},
			{Text: "secret", Description: "使用指定的子命令创建一个密钥（secret）。"},
			{Text: "service", Description: "使用指定的子命令创建一个服务。"},
			{Text: "serviceaccount", Description: "创建一个具有指定名称的服务账号。"},
			{Text: "clusterrole", Description: "创建一个集群角色。"},
			{Text: "clusterrolebinding", Description: "为特定的集群角色（ClusterRole）创建一个集群角色绑定（ClusterRoleBinding）。"},
			{Text: "cronjob", Description: "创建一个具有指定名称的定时任务（CronJob）。"},
			{Text: "job", Description: "创建一个具有指定名称的任务作业（Job）。"},
			{Text: "poddisruptionbudget", Description: "创建一个具有指定名称的 Pod 中断预算。"},
			{Text: "priorityclass", Description: "创建一个具有指定名称的优先级类别。"},
			{Text: "role", Description: "创建一个包含单个规则的角色。"},
			{Text: "rolebinding", Description: "为特定的角色（Role）或者集群角色（ClusterRole）创建一个角色绑定（RoleBinding）。"},
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

	// case "namespace":
	//	if len(args) == 2 {
	//		return prompt.FilterContains(getNameSpaceSuggestions(c.namespaceList), args[1], true)
	//	}
	case "logs":
		if len(args) == 2 {
			return prompt.FilterContains(getPodSuggestions(c.client, namespace), args[1], true)
		}
	// case "rolling-update", "rollingupdate":
	//	if len(args) == 2 {
	//		return prompt.FilterContains(getReplicationControllerSuggestions(c.client, namespace), args[1], true)
	//	} else if len(args) == 3 {
	//		return prompt.FilterContains(getReplicationControllerSuggestions(c.client, namespace), args[2], true)
	//	}
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
	case "cp":
		// fmt.Println("*********", args, len(args), "*************")
		// args = compressEmptyStrings(args)
		// fmt.Println("########", args, len(args), "#############")
		if len(args) == 2 {
			// return prompt.FilterContains(lsfiles(args[1], c.client, namespace, c.config, d), args[1], true)

			pathCompletion := lsfiles(args[1], c.client, namespace, c.config, d, false)
			// 如果补全项只剩下一个就继续补全
			if len(pathCompletion) == 1 {
				return lsfiles(pathCompletion[0].Text, c.client, namespace, c.config, d, false)
			}
			// 多个补项全直接返回结果
			return pathCompletion
			// return prompt.FilterContains(completerfile(args[1], d), args[1], true)
		}
		if len(args) == 3 {
			pathCompletion := lsfiles(args[2], c.client, namespace, c.config, d, true)
			// 如果补全项只剩下一个就继续补全
			if len(pathCompletion) == 1 {
				return lsfiles(pathCompletion[0].Text, c.client, namespace, c.config, d, true)
			}
			// 多个补项全直接返回结果
			return pathCompletion
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
			{Text: "history", Description: "查看部署的滚动更新历史记录。"},
			{Text: "pause", Description: "将所提供的资源标记为暂停状态。"},
			{Text: "resume", Description: "恢复一个已暂停的资源。"},
			{Text: "undo", Description: "撤销上一次的部署（回滚到上一次部署之前的状态）。"},
			{Text: "restart", Description: "重启一个资源。"},
			{Text: "status", Description: "展示部署（rollout）的状态。"},
		}
		if len(args) == 2 {
			return prompt.FilterHasPrefix(subCommands, args[1], true)
		}
	case "annotate":
	case "config":
		subCommands := []prompt.Suggest{
			{Text: "current-context", Description: "显示当前上下文。"},
			{Text: "delete-cluster", Description: "从kubeconfig 中删除指定的集群。"},
			{Text: "delete-context", Description: "从kubeconfig配置文件中删除指定的上下文环境信息。"},
			{Text: "get-clusters", Description: "展示在 kubeconfig 中定义的集群。"},
			{Text: "rename-context", Description: "重命名 kubeconfig 文件中的一个上下文（context）。"},
			{Text: "get-contexts", Description: "描述一个或者多个上下文环境。"},
			{Text: "set", Description: "在 kubeconfig 文件中设置单个值。"},
			{Text: "set-cluster", Description: "在 kubeconfig 中设置一个集群条目。"},
			{Text: "set-context", Description: "在 kubeconfig 中设置一个上下文项。"},
			{Text: "set-credentials", Description: "在 kubeconfig 中设置一个用户条目。"},
			{Text: "unset", Description: "在 kubeconfig（Kubernetes 配置）文件中取消设置单个值。"},
			{Text: "use-context", Description: "在一个 Kubernetes 配置文件（kubeconfig）中设置当前上下文。"},
			{Text: "view", Description: "显示合并后的 Kubernetes 配置（kubeconfig）设置或者指定的一个 kubeconfig 文件的内容。"},
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
			{Text: "dump", Description: "输出大量用于调试和诊断的相关信息。"},
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
