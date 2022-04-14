#!/bin/bash

echo "DESTROYING TESTDB"

URI=$POSTGRES_TEST_URI
URI=${URI//testdb}

set -e
psql "${URI%?}" <<-EOSQL
  DROP DATABASE IF EXISTS testdb;
  CREATE DATABASE testdb;
EOSQL

echo "TESTDB CREATED"