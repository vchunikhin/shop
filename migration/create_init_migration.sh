#!/usr/bin/env bash
. ./migration/vars.sh

if [ ! -d "$MIGRATION_PATH" ]
then
  echo "start initializing the migration script..."
  migrate create -ext sql -dir "$MIGRATION_PATH" -seq init
else
  echo "schema folder already exists."
fi