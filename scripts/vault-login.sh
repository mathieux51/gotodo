#!/usr/bin/env bash

set -o errexit
set -o nounset

echo Enter token:
read TOKEN

VAULT_ADDR="http://127.0.0.1:8200"
echo Trying to log to Vault with github token
vault login -method=github token=$TOKEN

