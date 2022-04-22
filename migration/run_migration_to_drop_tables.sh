#!/usr/bin/env bash
. ./migration/vars.sh

if [ ! -d "$MIGRATION_PATH" ]
then
  echo "Migration scripts are missing"
else
  echo "running migration scripts..."
  echo y | migrate -path "$MIGRATION_PATH" -database "$DATABASE" down
fi