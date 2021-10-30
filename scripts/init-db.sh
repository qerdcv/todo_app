#!/bin/bash

set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
  CREATE TABLE IF NOT EXISTS todos (
    id SERIAL,
    text VARCHAR(255)
  );
  CREATE UNIQUE INDEX IF NOT EXISTS todos_index ON todos (id);
EOSQL
