help:  ## Display this help message.
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-30s\033[0m %s\n", $$1, $$2}'

gen-db: ## generates our ent orm schemas
	make -C ./db gen

build-ui: ## builds our ui docker image
	@make -C ./ui build-docker

build: ## builds our docker images
	@echo "building ui container"
	@make -C ./ui build-docker
	@echo "Building containers"
	@make -C ../ build-docker

run-devnet: stop ## starts our devnet with the main tag for our docker containers
	@echo "Starting devnet"
	@cd .. && go run github.com/omni-network/omni/e2e deploy -f e2e/manifests/devnet2.toml \
	--explorer-db-conn="postgres://omni:password@explorer_db:5432/omni_db" \
	--omni-image-tag="main" \
	--graphql-url="http://localhost:21335/query"

run-explorer: ## starts our explorer in production mode
	@echo "Starting explorer"
	@cd ./ui && pnpm i && pnpm run build && pnpm run start

run: stop run-devnet ## Deploys e2e devnet1, copies network.json locally, and starts explorer docker compose

stop: ## stops our docker compose
	@echo "Stopping devnet"
	@make -C ../ devnet-clean

run-clean: build run-devnet
