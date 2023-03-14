#!/bin/bash

export ROOT="$( readlink -f "$( dirname "${BASH_SOURCE[0]}" )" )"

MIGRATION_NAME=$1

cd ${ROOT}/..
migrate create -ext sql -dir ${ROOT}/scripts -seq $MIGRATION_NAME