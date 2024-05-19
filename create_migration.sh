#!/bin/bash

help()
{
   echo "syntax: ./create_migration.sh MIGRATION_NAME"
}

MIGRATION_NAME="$1"
if [ -z "$MIGRATION_NAME" ]; then help; exit; fi

goose -dir=src/migrations create "${MIGRATION_NAME}" sql