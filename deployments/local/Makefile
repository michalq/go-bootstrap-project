.PHONY: run migration

PROJECT_NAME=go-project
DOCKER_COMPOSE=docker-compose -p ${PROJECT_NAME}

run: 
	${DOCKER_COMPOSE} up

migration: 
	${DOCKER_COMPOSE} run postgres-migration
	${DOCKER_COMPOSE} run mongo-migration
	