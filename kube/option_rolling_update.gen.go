// Code generated by 'option-gen'. DO NOT EDIT.

package kube

import (
	prompt "github.com/c-bata/go-prompt"
)

var rollingUpdateOptions = []prompt.Suggest{
	{Text: "--allow-missing-template-keys", Description: "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats."},
	{Text: "--container", Description: "Container name which will have its image upgraded. Only relevant when --image is specified, ignored otherwise. Required when using --image on a multi-container pod"},
	{Text: "--deployment-label-key", Description: "The key to use to differentiate between two different controllers, default 'deployment'.  Only relevant when --image is specified, ignored otherwise"},
	{Text: "--dry-run", Description: "If true, only print the object that would be sent, without sending it."},
	{Text: "-f", Description: "Filename or URL to file to use to create the new replication controller."},
	{Text: "--filename", Description: "Filename or URL to file to use to create the new replication controller."},
	{Text: "--image", Description: "Image to use for upgrading the replication controller. Must be distinct from the existing image (either new image or new image tag).  Can not be used with --filename/-f"},
	{Text: "--image-pull-policy", Description: "Explicit policy for when to pull container images. Required when --image is same as existing image, ignored otherwise."},
	{Text: "--include-extended-apis", Description: "If true, include definitions of new APIs via calls to the API server. [default true]"},
	{Text: "--no-headers", Description: "When using the default or custom-column output format, don't print headers (default print headers)."},
	{Text: "-o", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See custom columns [http://kubernetes.io/docs/user-guide/kubectl-overview/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://kubernetes.io/docs/user-guide/jsonpath]."},
	{Text: "--output", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See custom columns [http://kubernetes.io/docs/user-guide/kubectl-overview/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://kubernetes.io/docs/user-guide/jsonpath]."},
	{Text: "--poll-interval", Description: "Time delay between polling for replication controller status after the update. Valid time units are \"ns\", \"us\" (or \"µs\"), \"ms\", \"s\", \"m\", \"h\"."},
	{Text: "--rollback", Description: "If true, this is a request to abort an existing rollout that is partially rolled out. It effectively reverses current and next and runs a rollout"},
	{Text: "-a", Description: "When printing, show all resources (default show all pods including terminated one.)"},
	{Text: "--show-all", Description: "When printing, show all resources (default show all pods including terminated one.)"},
	{Text: "--show-labels", Description: "When printing, show all labels as the last column (default hide labels column)"},
	{Text: "--sort-by", Description: "If non-empty, sort list types using this field specification.  The field specification is expressed as a JSONPath expression (e.g. '{.metadata.name}'). The field in the API resource specified by this JSONPath expression must be an integer or a string."},
	{Text: "--template", Description: "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."},
	{Text: "--timeout", Description: "Max time to wait for a replication controller to update before giving up. Valid time units are \"ns\", \"us\" (or \"µs\"), \"ms\", \"s\", \"m\", \"h\"."},
	{Text: "--update-period", Description: "Time to wait between updating pods. Valid time units are \"ns\", \"us\" (or \"µs\"), \"ms\", \"s\", \"m\", \"h\"."},
	{Text: "--validate", Description: "If true, use a schema to validate the input before sending it"},
}