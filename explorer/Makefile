help:  ## Display this help message.
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-30s\033[0m %s\n", $$1, $$2}'

gen-db: ## generates our ent orm schemas
	make -C ./db gen

open-dev: ## opens all of the local host endpoints
	python -m webbrowser "http://localhost:3000"
	python -m webbrowser "http://localhost:8080"

build: ## builds our docker images
	@docker-compose build
	@echo "Building containers"
	@make -C ../ build-docker

run: stop ## Deploys e2e devnet1, copies network.json locally, and starts explorer docker compose
	@echo "Starting devnet"
	@make -C ../ devnet-deploy
	@cp ../test/e2e/runs/devnet1/relayer/network.json ./network.json
	@docker-compose up

stop: ## stops our docker compose
	@echo "Stopping devnet"
	@docker-compose down
	@make -C ../ devnet-clean

run-clean: build run ## docker build and run

gen-graphql: ## generates our graphql bindings for the ui
	cp ./graphql/app/schema.graphql ./ui/app/graphql/schema.graphql
	make -C ./ui gen-graphql
