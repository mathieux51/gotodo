#!/usr/bin/env bash

set -o errexit
set -o nounset

kubectl -n kube-system describe secret default | tail -n2 | awk '{print $2 }' | pbcopy
echo Token copied to clipboard

