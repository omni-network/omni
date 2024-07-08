ifneq ("$(wildcard .env)","")
	include .env
	export $(shell sed 's/=.*//' .env)
endif


.PHONY: help
help:  ## Display this help message.
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-30s\033[0m %s\n", $$1, $$2}'

.PHONY: version
version: ## Print tool versions.
	@forge --version

.PHONY: build
build: version ## Build contracts.
	forge build

.PHONY: deploy-eigenlayer-save-anvil-state
deploy-eigenlayer-save-anvil-state: ## Deploy eigenlayer contracts and dump anvil state.
	./script/eigen/deploy-eigenlayer-save-anvil-state.sh

.PHONY: avs-fork-test
avs-fork-test: ## Run AVS tests againts a local fork.
	@FOUNDRY_FUZZ_RUNS=5 forge test --fork-url http://localhost:8545 --match-path "test/avs/**/*.sol"

.PHONY: fork-holesky
fork-holesky: ## Run an anvil holesky fork.
	anvil --fork-url https://holesky.infura.io/v3/$(INFURA_KEY)

.PHONY: fork-mainnet
fork-mainnet: ## Run an anvil mainnet fork.
	anvil --fork-url https://mainnet.infura.io/v3/$(INFURA_KEY)

.PHONY: deploy-mainnet-avs
deploy-mainnet-avs: ## Deploy AVS contracts to mainnet.
	forge script DeployAVS --rpc-url https://mainnet.infura.io/v3/$(INFURA_KEY) --broadcast

.PHONY: deploy-local-avs
deploy-local-avs: ## Deploy AVS contracts to a local fork.
	forge script DeployAVS --rpc-url http://localhost:8545 --broadcast

.PHONY: deploy-mainnet-proxy-admin
deploy-mainnet-proxy-admin: ## Deploy ProxyAdmin contract to mainnet.
	forge script DeployProxyAdmin --rpc-url https://mainnet.infura.io/v3/$(INFURA_KEY) --broadcast

.PHONY: deploy-local-proxy-admin
deploy-local-proxy-admin: ## Deploy ProxyAdmin contract to a local fork.
	forge script DeployProxyAdmin --rpc-url http://localhost:8545 --broadcast
