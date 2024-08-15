#!/bin/bash

DIR=$(cd $(dirname $0); pwd)
KUBE_DIR=$(cd $(dirname $(dirname $0)); pwd)/kube

# clean generated files
rm ${KUBE_DIR}/*.gen.go
mkdir -p bin

set -e

go build -o ./bin/option-gen ./_tools/option-gen/main.go

subcmds=(
    "get"
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
    "uncordon"
    "annotate"
    "convert"
    "top node"
    "top pod"
    "cluster-info"
    "config get-contexts"
    "config set"
    "config set-cluster"
    "config set-credentials"
    "config view"
)

for cmd in "${subcmds[@]}"; do
  camelized=`echo ${cmd} | sed -r 's/[- ](.)/\U\1\E/g'`
  snaked=`echo ${cmd} | sed -r 's/[- ]/_/g'`
  kubectl ${cmd} --help | ./bin/option-gen -o ${KUBE_DIR}/option_${snaked}.gen.go -var ${camelized}Options
  gofumpt -l -w  ${KUBE_DIR}/option_${snaked}.gen.go
done