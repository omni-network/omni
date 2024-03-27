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

copy-network-file:
	if [ -z "$(NETWORK)" ]; then echo "⚠️  Please specify a network: NETWORK=devnet-1" && exit 1; fi
	@echo "Using NETWORK=$(NETWORK)"
	@cp ../e2e/runs/$(NETWORK)/relayer/network.json ./network.json

run-devnet: ## starts our devnet
	@echo "Starting devnet"
	@cp ../e2e/runs/devnet1/relayer/network.json ./network.json
	@make -C ../ devnet-deploy

run-explorer: ## starts our explorer
	@echo "Starting explorer"
	@docker-compose up

run: stop run-devnet run-explorer ## Deploys e2e devnet1, copies network.json locally, and starts explorer docker compose

stop: ## stops our docker compose
	@echo "Stopping devnet"
	@docker-compose down
	@make -C ../ devnet-clean

run-clean: build run ## docker build and run

gen-graphql: ## generates our graphql bindings for the ui
	cp ./graphql/app/schema.graphql ./ui/app/graphql/schema.graphql
	make -C ./ui gen-graphql
