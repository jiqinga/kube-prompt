// Code generated by 'option-gen'. DO NOT EDIT.

package kube

import (
	prompt "github.com/c-bata/go-prompt"
)

var proxyOptions = []prompt.Suggest{
	{Text: "--accept-hosts", Description: "好的，请您提供 `kubectl --help` 的帮助说明内容，我来为您翻译。: 1\\]$':  Regular expression for hosts that the proxy should accept."},
	{Text: "--accept-paths", Description: "代理应该接受的路径的正则表达式。 "},
	{Text: "--address", Description: "服务所使用的 IP 地址。 "},
	{Text: "--api-prefix", Description: "用于服务代理 API 的前缀。 "},
	{Text: "--disable-filter", Description: "如果为真，则在代理中禁用请求过滤。这很危险，当与可访问的端口一起使用时，可能会使您容易受到 XSRF 攻击。 "},
	{Text: "--keepalive", Description: "keepalive 指定了活跃网络连接的保持活动周期。设置为 0 可禁用保持活动。 "},
	{Text: "-p", Description: "运行代理的端口。设置为 0 以选择随机端口。 "},
	{Text: "--port", Description: "运行代理的端口。设置为 0 以选择随机端口。 "},
	{Text: "--reject-methods", Description: "用于指定代理应拒绝的 HTTP 方法的正则表达式（例如 --reject-methods='POST,PUT,PATCH'） "},
	{Text: "--reject-paths", Description: "代理应拒绝的路径的正则表达式。此处指定的路径即使被 --accept-paths 接受也会被拒绝。 "},
	{Text: "-u", Description: "用于运行代理的 Unix 套接字。 "},
	{Text: "--unix-socket", Description: "用于运行代理的 Unix 套接字。 "},
	{Text: "-w", Description: "同时在指定的前缀下，从给定的目录提供静态文件服务。 "},
	{Text: "--www", Description: "同时在指定的前缀下，从给定的目录提供静态文件服务。 "},
	{Text: "-P", Description: "如果指定了静态文件目录，则为服务于其下的静态文件的前缀。 "},
	{Text: "--www-prefix", Description: "如果指定了静态文件目录，则为服务于其下的静态文件的前缀。 "},
}
