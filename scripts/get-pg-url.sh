#!/bin/bash

CURRENT_DIR=$(dirname $0)
. $CURRENT_DIR/read-env.sh

read_env ${1:-.env}

echo "postgres://${DB_USER}:${DB_PASS}@localhost:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSLMODE}"
