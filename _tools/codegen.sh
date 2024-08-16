#!/bin/bash

DIR=$(cd $(dirname $0); pwd)
KUBE_DIR=$(cd $(dirname $DIR); pwd)/kube

# clean generated files
rm ${KUBE_DIR}/*.gen.go
mkdir -p bin

set -e

go build -o ./bin/option-gen ./main.go

subcmds=(
    "get"
    "set env"
    "set image"
    "set resources"
    "set selector"
    "set serviceaccount"
    "set subject"
    "describe"
    "create"
    "replace"
    "patch"
    "delete"
    "edit"
    "apply"
    "logs"
    "scale"
    "attach"
    "exec"
    "port-forward"
    "proxy"
    "run"
    "expose"
    "autoscale"
    "rollout history"
    "rollout pause"
    "rollout resume"
    "rollout status"
    "rollout undo"
    "rollout restart"
    "label"
    "explain"
    "cordon"
    "drain"
    "taint"
    "uncordon"
    "annotate"
    "convert"
    "top node"
    "top pod"
    "cluster-info dump"
    "config get-contexts"
    "config set"
    "config set-cluster"
    "config set-credentials"
    "config view"
    "certificate approve"
    "certificate deny"
    "cp"
    "auth can-i"
    "auth reconcile"
    "diff"
    "wait"
    "plugin list"
)

for cmd in "${subcmds[@]}"; do
  camelized=`echo ${cmd} | sed -r 's/[- ](.)/\U\1\E/g'`
  snaked=`echo ${cmd} | sed -r 's/[- ]/_/g'`
  kubectl ${cmd} --help | ./bin/option-gen -o ${KUBE_DIR}/option_${snaked}.gen.go -var ${camelized}Options
  gofumpt -l -w  ${KUBE_DIR}/option_${snaked}.gen.go
done