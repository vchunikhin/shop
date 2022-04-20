#!/usr/bin/env bash

# shellcheck disable=SC2034
MIGRATION_PATH=./migration/schema
# shellcheck disable=SC2125
DATABASE=postgres://gotest:gotest@localhost:5432/shop?sslmode=disable