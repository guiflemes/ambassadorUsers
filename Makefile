# these will speed up builds, for docker-compose >= 1.25
export COMPOSE_DOCKER_CLI_BUILD=1
export DOCKER_BUILDKIT=1
include .env

all: down build up test

build:
	docker-compose build

up:
	docker-compose up

down:
	docker-compose down --remove-orphans

up_db_test:
	docker container run -d -p 5437:5432 --name testdb -v $(SQL_PATH):/docker-entrypoint-initdb.d/create_tables.sql\
		   -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -e POSTGRES_DB=testdb -e POSTGRES_HOST=testdb postgres:13-alpine


down_db_test:
	docker container rm -f testdb

run_test:
	- ./scripts/tests.sh


test: up_db_test run_test down_db_test

