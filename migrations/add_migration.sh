#!/bin/bash

export ROOT=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

MIGRATION_NAME=$1

cd ${ROOT}/..
migrate create -ext sql -dir ${ROOT}/scripts -seq $MIGRATION_NAME
