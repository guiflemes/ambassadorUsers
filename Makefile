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

dump_schema:
	pg_dump -s $(POSTGRES_URI) | psql $(POSTGRES_TEST_URI)

drop_shema:
	./entrypoint.sh .env ./script.sh


# test_integration:
#	dump_schema test drop_shema

# test:
# 	dump_schema run_test