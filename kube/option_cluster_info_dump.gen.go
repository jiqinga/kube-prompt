// Code generated by 'option-gen'. DO NOT EDIT.

package kube

import (
	prompt "github.com/c-bata/go-prompt"
)

var clusterInfoDumpOptions = []prompt.Suggest{
	{Text: "-A", Description: "如果为真，则转储所有命名空间。如果为真，则忽略 --namespaces 。 "},
	{Text: "--all-namespaces", Description: "如果为真，则转储所有命名空间。如果为真，则忽略 --namespaces 。 "},
	{Text: "--allow-missing-template-keys", Description: "如果为真，当模板中缺少字段或映射键时，忽略模板中的任何错误。仅适用于 golang 和 jsonpath 输出格式。 "},
	{Text: "--namespaces", Description: "一个以逗号分隔的要转储的命名空间列表。 "},
	{Text: "-o", Description: "输出格式。其中之一 :  json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-file."},
	{Text: "--output", Description: "输出格式。其中之一 :  json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-file."},
	{Text: "--output-directory", Description: "输出文件的位置。如果为空或‘-’则使用标准输出，否则在该目录中创建目录层级。 "},
	{Text: "--pod-running-timeout", Description: "等待至少一个 Pod 处于运行状态的时长（例如 5 秒、2 分钟或 3 小时，需大于零） "},
	{Text: "--template", Description: "在 -o=go-template 或 -o=go-template-file 时使用的模板字符串或模板文件的路径。该模板格式为 golang 模板 [http] : //golang.org/pkg/text/template/#pkg-overview]."},
}
