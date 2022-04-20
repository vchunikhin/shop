#!/usr/bin/env bash
. ./migration/vars.sh

if [ ! -d "$MIGRATION_PATH" ]
then
  echo "Migration scripts are missing"
else
  echo "running migration scripts..."
  migrate -path "$MIGRATION_PATH" -database "$DATABASE" up
fi