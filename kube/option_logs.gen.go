// Code generated by 'option-gen'. DO NOT EDIT.

package kube

import (
	prompt "github.com/c-bata/go-prompt"
)

var logsOptions = []prompt.Suggest{
	{Text: "-c", Description: "Print the logs of this container"},
	{Text: "--container", Description: "Print the logs of this container"},
	{Text: "-f", Description: "Specify if the logs should be streamed."},
	{Text: "--follow", Description: "Specify if the logs should be streamed."},
	{Text: "--include-extended-apis", Description: "If true, include definitions of new APIs via calls to the API server. [default true]"},
	{Text: "--interactive", Description: "If true, prompt the user for input when required."},
	{Text: "--limit-bytes", Description: "Maximum bytes of logs to return. Defaults to no limit."},
	{Text: "--pod-running-timeout", Description: "The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running"},
	{Text: "-p", Description: "If true, print the logs for the previous instance of the container in a pod if it exists."},
	{Text: "--previous", Description: "If true, print the logs for the previous instance of the container in a pod if it exists."},
	{Text: "-l", Description: "Selector (label query) to filter on."},
	{Text: "--selector", Description: "Selector (label query) to filter on."},
	{Text: "--since", Description: "Only return logs newer than a relative duration like 5s, 2m, or 3h. Defaults to all logs. Only one of since-time / since may be used."},
	{Text: "--since-time", Description: "Only return logs after a specific date (RFC3339). Defaults to all logs. Only one of since-time / since may be used."},
	{Text: "--tail", Description: "Lines of recent log file to display. Defaults to -1 with no selector, showing all log lines otherwise 10, if a selector is provided."},
	{Text: "--timestamps", Description: "Include timestamps on each line in the log output"},
}