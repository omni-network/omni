help:  ## Display this help message.
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-30s\033[0m %s\n", $$1, $$2}'

gen-db: ## generates our ent orm schemas
	make -C ./db gen

open-dev: ## opens all of the local host endpoints
	python -m webbrowser "http://localhost:3000"
	python -m webbrowser "http://localhost:8080"

build: ## builds our docker images
	make -C ../ build-docker

run: ## runs our docker compose up
	docker-compose up

stop: ## stops our docker compose
	docker-compose down

run-clean: ## docker build and run
	make -C ./ build
	make -C ./ run

gen-graphql: ## generates our graphql bindings for the ui
	cp ./graphql/app/schema.graphql ./ui/app/graphql/schema.graphql
	make -C ./ui gen-graphql
