// Code generated by 'option-gen'. DO NOT EDIT.

package kube

import (
	prompt "github.com/c-bata/go-prompt"
)

var runOptions = []prompt.Suggest{
	{Text: "--allow-missing-template-keys", Description: "如果为真，在模板中缺少字段或映射键时忽略模板中的任何错误。仅适用于 golang 和 jsonpath 输出格式。 "},
	{Text: "--attach", Description: "如果为 true，等待 Pod 开始运行，然后附加到 Pod，就像调用 'kubectl attach...' 一样。默认值为 false，除非设置了 '-i/--stdin'，在这种情况下默认值为 true。使用 '--restart=Never' 时，会返回容器进程的退出代码。 "},
	{Text: "--cascade", Description: "如果为真，则级联删除此资源所管理的资源（例如由复制控制器创建的 Pod）。默认值为真。 "},
	{Text: "--command", Description: "如果为真且存在额外的参数，则将它们用作容器中的“command”字段，而非默认的“args”字段。 "},
	{Text: "--dry-run", Description: "如果为真，仅打印将会发送的对象，而不实际发送。 "},
	{Text: "--env", Description: "在容器中设置的环境变量 "},
	{Text: "--expose", Description: "如果为 true，则为运行的容器创建一个公开的外部服务 "},
	{Text: "-f", Description: "用于替换资源所使用的。 "},
	{Text: "--filename", Description: "用于替换资源所使用的。 "},
	{Text: "--force", Description: "仅在宽限期为 0 时使用。如果为真，会立即从 API 中删除资源并绕过优雅删除。请注意，某些资源的立即删除可能会导致不一致或数据丢失，并且需要确认。 "},
	{Text: "--generator", Description: "要使用的 API 生成器的名称，请参见 http : //kubernetes.io/docs/user-guide/kubectl-conventions/#generators for a list."},
	{Text: "--grace-period", Description: "以秒为单位的时间段，给予资源以优雅地终止。如果为负数则被忽略。设置为 1 则立即关闭。只有当 --force 为 true（强制删除）时才能设置为 0 。 "},
	{Text: "--hostport", Description: "容器端口的主机端口映射。用于演示单主机容器。 "},
	{Text: "--image", Description: "要运行的容器的镜像。 "},
	{Text: "--image-pull-policy", Description: "容器的镜像拉取策略。如果留空，此值将不会由客户端指定，并由服务器设置默认值。 "},
	{Text: "-k", Description: "处理一个 kustomization 目录。此标志不能与 -f 或 -R 一起使用。 "},
	{Text: "--kustomize", Description: "处理一个 kustomization 目录。此标志不能与 -f 或 -R 一起使用。 "},
	{Text: "-l", Description: "以逗号分隔的标签，应用于 Pod。将覆盖先前的值。 "},
	{Text: "--labels", Description: "以逗号分隔的标签，应用于 Pod。将覆盖先前的值。 "},
	{Text: "--leave-stdin-open", Description: "如果 Pod 是以交互模式启动或带有标准输入，在首次连接完成后保持标准输入处于打开状态。默认情况下，在首次连接完成后标准输入将被关闭。 "},
	{Text: "--limits", Description: "此容器的资源需求限制。例如，'cpu=200m,memory=512Mi'。请注意，服务器端组件可能会根据服务器配置（例如限制范围）分配限制。 "},
	{Text: "-o", Description: "输出格式。其中之一 :  json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-file."},
	{Text: "--output", Description: "输出格式。其中之一 :  json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-file."},
	{Text: "--overrides", Description: "一个用于生成对象的内联 JSON 覆盖。如果此内容不为空，则用于覆盖生成的对象。要求对象提供有效的 apiVersion 字段。 "},
	{Text: "--pod-running-timeout", Description: "等待至少一个 Pod 处于运行状态的时长（例如 5 秒、2 分钟或 3 小时，大于零） "},
	{Text: "--port", Description: "此容器暴露的端口。如果 --expose 为 true，这也是所创建服务使用的端口。 "},
	{Text: "--quiet", Description: "如果为真，则抑制提示消息。 "},
	{Text: "--record", Description: "在资源注释中记录当前的 kubectl 命令。如果设置为 false ，则不记录该命令。如果设置为 true ，则记录该命令。如果未设置，默认情况下仅在已存在注释值时更新现有注释值。 "},
	{Text: "-R", Description: "对在 -f、--filename 中使用的目录进行递归处理。在您想要管理在同一目录中组织的相关清单时很有用。 "},
	{Text: "--recursive", Description: "对在 -f、--filename 中使用的目录进行递归处理。在您想要管理在同一目录中组织的相关清单时很有用。 "},
	{Text: "-r", Description: "要为此容器创建的副本数量。默认值为 1 。 "},
	{Text: "--replicas", Description: "要为此容器创建的副本数量。默认值为 1 。 "},
	{Text: "--requests", Description: "此容器的资源需求请求。例如，'cpu=100m，memory=256Mi'。请注意，服务器端组件可能会根据服务器配置（例如限制范围）来分配请求。 "},
	{Text: "--restart", Description: "此 Pod 的重启策略。合法值为 [Always（总是）、OnFailure（失败时）、Never（从不）] 。如果设置为“Always”，则创建部署；如果设置为“OnFailure”，则创建任务；如果设置为“Never”，则创建常规 Pod。对于后两种情况，--replicas 必须为 1。默认值为“Always”，对于 CronJobs 为“Never”。 "},
	{Text: "--rm", Description: "如果为真，则删除此命令为附加容器创建的资源。 "},
	{Text: "--save-config", Description: "如果为真，当前对象的配置将保存在其注解中。否则，注解将保持不变。此标志在您未来想要对此对象执行 kubectl apply 操作时很有用。 "},
	{Text: "--schedule", Description: "以 Cron 格式表示的作业应运行的调度计划。 "},
	{Text: "--service-generator", Description: "用于创建服务的生成器的名称。仅在 --expose 为 true 时使用 "},
	{Text: "--service-overrides", Description: "对于生成的服务对象的内联 JSON 覆盖。如果此内容不为空，则用于覆盖生成的对象。要求对象提供有效的 apiVersion 字段。仅在 --expose 为 true 时使用。 "},
	{Text: "--serviceaccount", Description: "在 Pod 规范中要设置的服务账号 "},
	{Text: "-i", Description: "即使没有任何连接，也使 Pod 中的容器保持标准输入处于打开状态。 "},
	{Text: "--stdin", Description: "即使没有任何连接，也使 Pod 中的容器保持标准输入处于打开状态。 "},
	{Text: "--template", Description: "在 -o=go-template 或 -o=go-template-file 时使用的模板字符串或模板文件的路径。该模板格式为 golang 模板 [http] : //golang.org/pkg/text/template/#pkg-overview]."},
	{Text: "--timeout", Description: "在放弃删除操作前等待的时长，零表示根据对象的大小确定超时时间 "},
	{Text: "-t", Description: "为 Pod 中的每个容器分配一个 TTY 。 "},
	{Text: "--tty", Description: "为 Pod 中的每个容器分配一个 TTY 。 "},
	{Text: "--wait", Description: "如果为真，则等待资源消失后再返回。此操作会等待终结器。 "},
}
