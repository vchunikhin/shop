#!/usr/bin/env bash
. ./env/vars.sh

echo "start the local environment..."
docker-compose -f "$CURRENT_DIR"/local/docker-compose.yml up -d