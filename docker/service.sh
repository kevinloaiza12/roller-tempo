#!/bin/bash

ROOT=$(pwd)
DOCKER_FILE="$ROOT/docker-compose.yaml"
export PSQL_DATA_DIR="$ROOT/data"

if [ "$1" = "up" ]; then
    docker-compose --file "$DOCKER_FILE" up --detach
fi

if [ "$1" = "down" ]; then
    docker-compose --file "$DOCKER_FILE" down
fi