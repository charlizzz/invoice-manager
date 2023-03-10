#!/bin/sh

set -e

echo "run the migration for the app"
/app/migrate -path /app/migration -database "$DATABASE_SOURCE" -v up

echo "start the app"
exec "$0"