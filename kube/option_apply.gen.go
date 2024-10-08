// Code generated by 'option-gen'. DO NOT EDIT.

package kube

import (
	prompt "github.com/c-bata/go-prompt"
)

var applyOptions = []prompt.Suggest{
	{Text: "--all", Description: "选择指定资源类型所在命名空间中的所有资源。 "},
	{Text: "--allow-missing-template-keys", Description: "如果为真，在模板中缺少字段或映射键时，忽略模板中的任何错误。仅适用于 golang 和 jsonpath 输出格式。 "},
	{Text: "--cascade", Description: "如果为真，则级联删除此资源所管理的资源（例如由复制控制器创建的 Pod）。默认值为真。 "},
	{Text: "--dry-run", Description: "如果为真，仅打印将会被发送的对象，而不实际发送。警告 :  --dry-run cannot accurately output the result of merging the local manifest and the server-side data. Use --server-dry-run to get the merged result instead."},
	{Text: "--field-manager", Description: "用于跟踪字段所有权的管理器的名称。 "},
	{Text: "-f", Description: "其中包含要应用的配置 "},
	{Text: "--filename", Description: "其中包含要应用的配置 "},
	{Text: "--force", Description: "仅在宽限期为 0 时使用。如果为真，会立即从 API 中移除资源并绕过优雅删除。请注意，立即删除某些资源可能会导致不一致或数据丢失，需要进行确认。 "},
	{Text: "--force-conflicts", Description: "如果为真，服务器端应用将强制针对冲突进行更改。 "},
	{Text: "--grace-period", Description: "以秒为单位的时间段，给予资源以优雅地终止。如果为负数则被忽略。设置为 1 表示立即关闭。只有在 --force 为 true（强制删除）时才能设置为 0 。 "},
	{Text: "-k", Description: "处理一个 kustomization 目录。此标志不能与 -f 或 -R 一起使用。 "},
	{Text: "--kustomize", Description: "处理一个 kustomization 目录。此标志不能与 -f 或 -R 一起使用。 "},
	{Text: "--openapi-patch", Description: "如果为 true，当 OpenAPI 存在且资源能在 OpenAPI 规范中找到时，使用 OpenAPI 来计算差异。否则，回退使用内置类型。 "},
	{Text: "-o", Description: "输出格式。其中之一 :  json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-file."},
	{Text: "--output", Description: "输出格式。其中之一 :  json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-file."},
	{Text: "--overwrite", Description: "通过使用修改后的配置中的值，自动解决修改后的配置与实时配置之间的冲突 "},
	{Text: "--prune", Description: "自动删除资源对象，包括未初始化的对象，这些对象未出现在配置中，并且是通过 `apply` 或 `create --save-config` 创建的。应与 `-l` 或 `--all` 一起使用。 "},
	{Text: "--prune-whitelist", Description: "覆盖用于 --prune 的默认白名单为 <组/版本/种类> "},
	{Text: "--record", Description: "在资源注解中记录当前的 kubectl 命令。若设置为 false，则不记录命令。若设置为 true，则记录命令。若未设置，默认仅在已有注解值时更新现有值。 "},
	{Text: "-R", Description: "对在 -f、--filename 中使用的目录进行递归处理。在您想要管理组织在同一目录内的相关清单时很有用。 "},
	{Text: "--recursive", Description: "对在 -f、--filename 中使用的目录进行递归处理。在您想要管理组织在同一目录内的相关清单时很有用。 "},
	{Text: "-l", Description: "选择器（标签查询）用于筛选，支持 '='、'==' 和 '!=' 。（例如 -l key1=value1,key2=value2） "},
	{Text: "--selector", Description: "选择器（标签查询）用于筛选，支持 '='、'==' 和 '!=' 。（例如 -l key1=value1,key2=value2） "},
	{Text: "--server-dry-run", Description: "如果为真，请求将带着试运行标志发送到服务器，这意味着所做的修改不会被持久保存。这是一个 alpha 特性和标志。 "},
	{Text: "--server-side", Description: "如果为真，`apply` 在服务器而非客户端运行。 "},
	{Text: "--template", Description: "当 `-o=go-template` 或 `-o=go-template-file` 时使用的模板字符串或模板文件的路径。该模板格式为 Go 语言模板 [http] : //golang.org/pkg/text/template/#pkg-overview]."},
	{Text: "--timeout", Description: "在放弃删除操作之前等待的时长，零表示根据对象的大小确定超时时间。 "},
	{Text: "--validate", Description: "如果为真，则在发送输入之前使用模式对其进行验证 "},
	{Text: "--wait", Description: "如果为真，在返回之前等待资源消失。这会等待终结器。 "},
}
