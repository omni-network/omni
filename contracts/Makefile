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
	@abigen --version

.PHONY: build
build: version ## Build contracts.
	forge build

.PHONY: bindings
bindings: check-abigen-version build ## Generate golang contract bindings.
	./bindings/scripts/gen.sh OmniPortal FeeOracleV1 OmniAVS TransparentUpgradeableProxy ProxyAdmin DelegationManager StrategyManager StrategyBase MockERC20 AVSDirectory
	go run ./bindings/scripts/commenttypes.go -- bindings/strategymanager.go IStrategyManagerDeprecatedStructQueuedWithdrawal IStrategyManagerDeprecatedStructWithdrawerAndNonce
	go run ./bindings/scripts/commenttypes.go -- bindings/avsdirectory.go ISignatureUtilsSignatureWithSaltAndExpiry

.PHONY: examples-bindings
examples-bindings: check-abigen-version build ## Generate golang contract bindings for example contracts
	PKG=examples DIR=./bindings/examples ./bindings/scripts/gen.sh PingPong

.PHONY: deploy-eigenlayer-save-anvil-state
deploy-eigenlayer-save-anvil-state: ## Deploy eigenlayer contracts and dump anvil state.
	./script/eigen/deploy-eigenlayer-save-anvil-state.sh

GOERLI_FORK_BLOCK_NUMBER ?= 10641114

.PHONY: avs-fork-test-goerli
avs-fork-test-goerli: ## Run OmniAVS tests on a goerli fork.
	@if [ -z "${INFURA_KEY}" ]; then echo "INFURA_KEY is not set"; exit 1; fi
	@FOUNDRY_FUZZ_RUNS=10 forge test --fork-url https://goerli.infura.io/v3/${INFURA_KEY} --fork-block-number ${GOERLI_FORK_BLOCK_NUMBER} --match-path "test/avs/**/*.sol"

.PHONY: avs-fork-test-local
avs-fork-test-local: ## Run OmniAVS tests on a local fork.
	@if [ -z "${INFURA_KEY}" ]; then echo "INFURA_KEY is not set"; exit 1; fi
	@FOUNDRY_FUZZ_RUNS=10 forge test --fork-url http://localhost:8545 --match-path "test/avs/**/*.sol"

.PHONY: avs-fork-goerli
anvil-fork-goerli: ## Run an anvil goerli fork
	@if [ -z "${INFURA_KEY}" ]; then echo "INFURA_KEY is not set"; exit 1; fi
	@anvil --fork-url https://goerli.infura.io/v3/${INFURA_KEY} --fork-block-number ${GOERLI_FORK_BLOCK_NUMBER}


.PHONY: check-abigen-version
check-abigen-version: ## Check abigen version, exit(1) if not 1.13.14-stable.
	@version=$$(abigen --version); \
	if [ "$$version" != "abigen version 1.13.14-stable" ]; then \
		echo "abigen version is not 1.13.14-stable"; \
		echo "Install with go install github.com/ethereum/go-ethereum/cmd/abigen@v1.13.14"; \
		exit 1; \
	fi
