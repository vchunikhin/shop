#!/usr/bin/env bash
. ./env/vars.sh

echo "drop the local environment..."

docker container stop "$DOCKER_DB_CONTAINER_NAME" && docker rm "$DOCKER_DB_CONTAINER_NAME"
echo "$SUDO_PASSWORD" | sudo -S rm -rf "$CURRENT_DIR"/local/pg-data-volume