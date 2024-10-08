// Code generated by 'option-gen'. DO NOT EDIT.

package kube

import (
	prompt "github.com/c-bata/go-prompt"
)

var setImageOptions = []prompt.Suggest{
	{Text: "--all", Description: "选择指定资源类型所在命名空间中的所有资源，包括未初始化的资源 "},
	{Text: "--allow-missing-template-keys", Description: "如果为真，在模板中缺少字段或映射键时忽略模板中的任何错误。仅适用于 golang 和 jsonpath 输出格式。 "},
	{Text: "--dry-run", Description: "如果为真，仅打印将会发送的对象，而不实际发送。 "},
	{Text: "-f", Description: "文件名、目录或指向文件的 URL，用于标识要从服务器获取的资源。 "},
	{Text: "--filename", Description: "文件名、目录或指向文件的 URL，用于标识要从服务器获取的资源。 "},
	{Text: "-k", Description: "处理 kustomization 目录。此标志不能与 -f 或 -R 一起使用。 "},
	{Text: "--kustomize", Description: "处理 kustomization 目录。此标志不能与 -f 或 -R 一起使用。 "},
	{Text: "--local", Description: "如果为真，设置镜像将不会联系 API 服务器，而是在本地运行。 "},
	{Text: "-o", Description: "输出格式。其中之一 :  json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-file."},
	{Text: "--output", Description: "输出格式。其中之一 :  json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-file."},
	{Text: "--record", Description: "在资源注释中记录当前的 kubectl 命令。如果设置为 false，则不记录该命令。如果设置为 true，则记录该命令。如果未设置，默认情况下仅在已存在注释值时更新现有注释值。 "},
	{Text: "-R", Description: "对在 -f、--filename 中使用的目录进行递归处理。当您想要管理在同一目录中组织的相关清单时很有用。 "},
	{Text: "--recursive", Description: "对在 -f、--filename 中使用的目录进行递归处理。当您想要管理在同一目录中组织的相关清单时很有用。 "},
	{Text: "-l", Description: "选择器（标签查询）用于筛选，不包括未初始化的，支持 '='、'==' 和 '!=' 。（例如 -l key1=value1,key2=value2） "},
	{Text: "--selector", Description: "选择器（标签查询）用于筛选，不包括未初始化的，支持 '='、'==' 和 '!=' 。（例如 -l key1=value1,key2=value2） "},
	{Text: "--template", Description: "当 -o=go-template 或 -o=go-template-file 时使用的模板字符串或模板文件的路径。该模板格式为 Go 语言模板 [http] : //golang.org/pkg/text/template/#pkg-overview]."},
}
