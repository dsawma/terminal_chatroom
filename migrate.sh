#!/bin/bash

if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)
fi

if [ -z "$DB_URL" ]; then
    echo "Error: DB_URL is not set in .env"
    exit 1
fi

echo "Initializing schema..."
psql "$DB_URL" -c "CREATE SCHEMA IF NOT EXISTS public;"

echo "Running Goose migrations up..."

goose -dir sql/schema postgres "$DB_URL" up

if [ $? -eq 0 ]; then
    echo "Goose migrations completed successfully!"
else
    echo "Goose migrations failed!"
    exit 1
fi