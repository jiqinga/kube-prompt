// Code generated by 'option-gen'. DO NOT EDIT.

package kube

import (
	prompt "github.com/c-bata/go-prompt"
)

var certificateDenyOptions = []prompt.Suggest{
	{Text: "--allow-missing-template-keys", Description: "如果为真，当模板中缺少字段或映射键时，忽略模板中的任何错误。仅适用于 golang 和 jsonpath 输出格式。 "},
	{Text: "-f", Description: "文件名、目录或指向标识要更新的资源的文件的 URL "},
	{Text: "--filename", Description: "文件名、目录或指向标识要更新的资源的文件的 URL "},
	{Text: "--force", Description: "即使 CSR 已被拒绝，仍对其进行更新。 "},
	{Text: "-k", Description: "处理 kustomization 目录。此标志不能与 -f 或 -R 一起使用。 "},
	{Text: "--kustomize", Description: "处理 kustomization 目录。此标志不能与 -f 或 -R 一起使用。 "},
	{Text: "-o", Description: "输出格式。其中之一 :  json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-file."},
	{Text: "--output", Description: "输出格式。其中之一 :  json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-file."},
	{Text: "-R", Description: "对在 -f、--filename 中使用的目录进行递归处理。在您想要管理在同一目录中组织的相关清单时很有用。 "},
	{Text: "--recursive", Description: "对在 -f、--filename 中使用的目录进行递归处理。在您想要管理在同一目录中组织的相关清单时很有用。 "},
	{Text: "--template", Description: "当 -o=go-template 或 -o=go-template-file 时使用的模板字符串或模板文件的路径。该模板格式为 Go 语言模板 [http] : //golang.org/pkg/text/template/#pkg-overview]."},
}