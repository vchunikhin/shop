#!/usr/bin/env bash

CURRENT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
echo "CURRENT_DIR: $CURRENT_DIR"

DOCKER_DB_CONTAINER_NAME=shop-db
echo "DOCKER_DB_CONTAINER_NAME: $DOCKER_DB_CONTAINER_NAME"

# shellcheck disable=SC2034
SUDO_PASSWORD=1234