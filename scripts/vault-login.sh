#!/usr/bin/env bash

set -o errexit
set -o nounset

TOKEN=$(pbpaste)
# if TOKEN not set ask user input
if -z $TOKEN
then
  echo Enter token:
  read TOKEN
else
  echo Using token from clipboard
fi

echo $TOKEN
# VAULT_ADDR="http://127.0.0.1:8200"
# vault login -method=github token=$TOKEN

