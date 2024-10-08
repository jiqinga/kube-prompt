# Kube-Prompt 🚀

![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)
![Go Report Card](https://goreportcard.com/badge/github.com/jiqinga/kube-prompt)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/jiqinga/kube-prompt)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/jiqinga/kube-prompt)

> 一个具有自动完成功能的交互式Kubernetes客户端,基于[kube-prompt](https://github.com/c-bata/kube-prompt)构建。

![demo](https://github.com/jiqinga/assets/raw/master/kube-prompt/docs/kube-prompt.gif)

## ✨ 特性

- 🖥️ **交互式命令提示**: 实时命令补全和智能建议
- 🔍 **资源管理**: 轻松查看、创建、编辑和删除 Kubernetes 资源
- 🎨 **彩色输出**: 使用彩色文本增强可读性,基于[kubecolor](https://github.com/hidetatz/kubecolor)
- 🔍 智能自动完成
- 🔗 支持管道操作
- 🚀 与 kubectl 命令兼容
- 🔐 无需额外权限,直接使用现有 kubeconfig
- 📁 本地文件路径和 Pod 内路径的智能补全
- 🔄 默认命名空间快速切换


## 🚀 快速开始

## 🛠️ 安装

### 下载独立二进制文件

从 [GitHub Releases](https://github.com/jiqinga/kube-prompt/releases) 下载适合您系统的二进制文件。

<details>
<summary>macOS (darwin) - amd64</summary>

```bash
wget https://github.com/jiqinga/kube-prompt/releases/download/v1.0.11/kube-prompt_v1.0.11_darwin_amd64.zip
unzip kube-prompt_v1.0.11_darwin_amd64.zip
chmod +x kube-prompt
sudo mv ./kube-prompt /usr/local/bin/kube-prompt
```

</details>

<details>
<summary>Linux - amd64</summary>

```bash
wget https://github.com/jiqinga/kube-prompt/releases/download/v1.0.11/kube-prompt_v1.0.11_linux_amd64.zip
unzip kube-prompt_v1.0.11_linux_amd64.zip
chmod +x kube-prompt
sudo mv ./kube-prompt /usr/local/bin/kube-prompt
```

</details>
### 从源码构建

```bash
GO111MODULE=on go build .
```

## 🚀 快速开始

安装完成后,直接在终端中运行:

```bash
kube-prompt
```

然后,您就可以开始使用 Kubernetes 命令,无需 `kubectl` 前缀:

```
>>> get pod | grep web
web-1144924021-2spbr        1/1     Running     4       25d
web-1144924021-5r1fg        1/1     Running     4       25d
web-1144924021-pqmfq        1/1     Running     4       25d
>>> describe deployment nginx
Name:                   nginx
Namespace:              default
CreationTimestamp:      Thu, 15 Nov 2018 12:42:23 +0100
Labels:                 app=nginx
Annotations:            deployment.kubernetes.io/revision=1
Selector:               app=nginx
Replicas:               1 desired | 1 updated | 1 total | 1 available | 0 unavailable
StrategyType:           RollingUpdate
MinReadySeconds:        0
RollingUpdateStrategy:  25% max unavailable, 25% max surge
...

>>> logs -f nginx-65899c769f-wv2gp
172.17.0.1 - - [15/Nov/2018:11:45:06 +0000] "GET / HTTP/1.1" 200 612 "-" "Mozilla/5.0 ..."
172.17.0.1 - - [15/Nov/2018:11:45:06 +0000] "GET /favicon.ico HTTP/1.1" 404 571 "http://..."
...
```

## 📚 支持的命令

kube-prompt 支持大多数常用的 Kubernetes 命令,包括但不限于:

- get, describe, create, replace, patch, delete
- edit, apply, namespace, logs
- scale, cordon, drain, uncordon
- exec, port-forward, proxy, run, expose

完整列表请参考:
```
get            Display one or many resources
describe       Show details of a specific resource or group of resources
create         Create a resource by filename or stdin
replace        Replace a resource by filename or stdin.
patch          Update field(s) of a resource using strategic merge patch.
delete         Delete resources by filenames, stdin, resources and names, or by resources and label selector.
edit           Edit a resource on the server
apply          Apply a configuration to a resource by filename or stdin
namespace      SUPERSEDED: Set and view the current Kubernetes namespace
logs           Print the logs for a container in a pod.
rolling-update Perform a rolling update of the given ReplicationController.
scale          Set a new size for a Deployment, ReplicaSet, Replication Controller, or Job.
cordon         Mark node as unschedulable
drain          Drain node in preparation for maintenance
uncordon       Mark node as schedulable
attach         Attach to a running container.
exec           Execute a command in a container.
port-forward   Forward one or more local ports to a pod.
proxy          Run a proxy to the Kubernetes API server
run            Run a particular image on the cluster.
expose         Take a replication controller, service, or pod and expose it as a new Kubernetes Service
autoscale      Auto-scale a Deployment, ReplicaSet, or ReplicationController
rollout        rollout manages a deployment
label          Update the labels on a resource
annotate       Update the annotations on a resource
config         config modifies kubeconfig files
cluster-info   Display cluster info
api-versions   Print the supported API versions on the server, in the form of "group/version".
version        Print the client and server version information.
explain        Documentation of resources.
convert        Convert config files between different API versions
top            Display Resource (CPU/Memory/Storage) usage
```


## 📄 许可证

本项目采用 MIT 许可证。详情请见 [LICENSE](./LICENSE) 文件。


