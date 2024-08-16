// Code generated by 'option-gen'. DO NOT EDIT.

package kube

import (
	prompt "github.com/c-bata/go-prompt"
)

var configSetCredentialsOptions = []prompt.Suggest{
	{Text: "--auth-provider", Description: "Auth provider for the user entry in kubeconfig"},
	{Text: "--auth-provider-arg", Description: "'key=value' arguments for the auth provider"},
	{Text: "--embed-certs", Description: "Embed client cert/key for the user entry in kubeconfig"},
}