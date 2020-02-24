#!/usr/bin/env bash

set -o errexit
set -o nounset

export POD_NAME=$(kubectl get pods -n default -l "app=kubernetes-dashboard,release=kubernetes-dashboard" -o jsonpath="{.items[0].metadata.name}")
  echo https://127.0.0.1:8443/
kubectl -n default port-forward $POD_NAME 8443:8443

