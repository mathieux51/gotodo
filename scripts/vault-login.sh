#!/usr/bin/env bash

set -o errexit
set -o nounset

echo Enter token:
read TOKEN

VAULT_ADDR="http://127.0.0.1:8200"
echo Trying to log to Vault with github token
unset VAULT_TOKEN
vault login -method=github token=$TOKEN

# curl \
#   --request POST \
#   --data "{\"token\": \"$TOKEN\"}" 'http://127.0.0.1:8200/v1/auth/github/login'

