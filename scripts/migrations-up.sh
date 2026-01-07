#!/bin/bash

echo "Validating if the .env file exists"
if [ -f .env ]; then
  export $(grep -v '^#' .env | xargs)
fi

echo "Validating if the DATABASE_URL environment variable is declared"
if [ -z "${DATABASE_URL:-}" ]; then
  echo "❌ DATABASE_URL não está definida"
  exit 1
fi

echo "Runing migrations"
migrate -path migrations -database "$DATABASE_URL" up

echo "Migrations successfully implemented"