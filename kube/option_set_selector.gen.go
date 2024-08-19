// Code generated by 'option-gen'. DO NOT EDIT.

package kube

import (
	prompt "github.com/c-bata/go-prompt"
)

var setSelectorOptions = []prompt.Suggest{
	{Text: "--all", Description: "选择指定资源类型所在命名空间中的所有资源 "},
	{Text: "--allow-missing-template-keys", Description: "如果为真，当模板中缺少字段或映射键时，忽略模板中的任何错误。仅适用于 golang 和 jsonpath 输出格式。 "},
	{Text: "--dry-run", Description: "如果为真，仅打印将要发送的对象，而不发送它。 "},
	{Text: "-f", Description: "确定资源。 "},
	{Text: "--filename", Description: "确定资源。 "},
	{Text: "--local", Description: "如果为真，注解将不会联系 API 服务器，而是在本地运行。 "},
	{Text: "-o", Description: "输出格式。其中之一 :  json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-file."},
	{Text: "--output", Description: "输出格式。其中之一 :  json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-file."},
	{Text: "--record", Description: "在资源注解中记录当前的 kubectl 命令。若设为 false，则不记录命令。若设为 true，则记录命令。若未设置，默认仅在已存在注解值时更新现有值。 "},
	{Text: "-R", Description: "递归处理 -f 或 --filename 所使用的目录。在您想要管理在同一目录中组织的相关清单时很有用。 "},
	{Text: "--recursive", Description: "递归处理 -f 或 --filename 所使用的目录。在您想要管理在同一目录中组织的相关清单时很有用。 "},
	{Text: "--resource-version", Description: "如果不为空，选择器的更新只有在这是对象的当前资源版本时才会成功。仅在指定单个资源时有效。 "},
	{Text: "--template", Description: "在 -o=go-template 或 -o=go-template-file 时使用的模板字符串或模板文件的路径。该模板格式为 Go 语言模板 [http] : //golang.org/pkg/text/template/#pkg-overview]."},
}
