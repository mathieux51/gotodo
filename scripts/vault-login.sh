#!/usr/bin/env bash

set -o errexit
set -o nounset

echo Enter token:
read TOKEN

VAULT_ADDR="http://127.0.0.1:8200"
vault login -method=github token=$TOKEN

