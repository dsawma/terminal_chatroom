#!/bin/bash

# 1. Load the .env file
if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)
fi

# 2. Check if DB_URL is set
if [ -z "$DB_URL" ]; then
    echo "Error: DB_URL is not set in .env"
    exit 1
fi

echo "Initializing schema..."
# This ensures the 'public' schema exists and is owned by your user
psql "$DB_URL" -c "CREATE SCHEMA IF NOT EXISTS public;"

# 3. Run Goose Migrations
echo "Running Goose migrations up..."

goose -dir sql/schema postgres "$DB_URL" up

# 4. Check if Goose succeeded
if [ $? -eq 0 ]; then
    echo "Goose migrations completed successfully!"
else
    echo "Goose migrations failed!"
    exit 1
fi