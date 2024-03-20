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
	./bindings/scripts/gen.sh OmniPortal FeeOracleV1 OmniAVS OmniStake Create3 TransparentUpgradeableProxy ProxyAdmin DelegationManager StrategyManager StrategyBase AVSDirectory test/avs/common/MockERC20.sol:MockERC20
	./bindings/scripts/genmore.sh OmniStake
	go run ./bindings/scripts/commenttypes.go -- bindings/strategymanager.go IStrategyManagerDeprecatedStructQueuedWithdrawal IStrategyManagerDeprecatedStructWithdrawerAndNonce
	go run ./bindings/scripts/commenttypes.go -- bindings/avsdirectory.go ISignatureUtilsSignatureWithSaltAndExpiry

.PHONY: examples-bindings
examples-bindings: check-abigen-version build ## Generate golang contract bindings for example contracts
	PKG=examples DIR=./bindings/examples ./bindings/scripts/gen.sh PingPong

.PHONY: deploy-eigenlayer-save-anvil-state
deploy-eigenlayer-save-anvil-state: ## Deploy eigenlayer contracts and dump anvil state.
	./script/eigen/deploy-eigenlayer-save-anvil-state.sh

GOERLI_FORK_BLOCK_NUMBER ?= 10641114
GOERLI_RPC = https://goerli.infura.io/v3/${INFURA_KEY}
LOCAL_RPC = http://localhost:8545

.PHONY: avs-fork-test-goerli
avs-fork-test-goerli: check-infura-key ## Run OmniAVS tests on a goerli fork.
	@FOUNDRY_FUZZ_RUNS=10 forge test --fork-url ${GOERLI_RPC} --fork-block-number ${GOERLI_FORK_BLOCK_NUMBER} --match-path "test/avs/**/*.sol"

.PHONY: avs-fork-test-local
avs-fork-test-local: check-infura-key ## Run OmniAVS tests on a local fork.
	@FOUNDRY_FUZZ_RUNS=10 forge test --fork-url ${LOCAL_RPC} --match-path "test/avs/**/*.sol"

.PHONY: avs-fork-goerli
anvil-fork-goerli: check-infura-key ## Run an anvil goerli fork
	@anvil --fork-url ${GOERLI_RPC} --fork-block-number ${GOERLI_FORK_BLOCK_NUMBER}

.PHONY: deploy-goerli-avs
deploy-goerli-avs: check-infura-key ## Deploy OmniAVS to goerli.
	@./script/avs/deploy-goerli-avs.sh ${GOERLI_RPC}

.PHONY: deploy-local-goerli-avs
deploy-local-goerli-avs: ## Deploy OmniAVS to local goerli fork.
	@./script/avs/deploy-goerli-avs.sh ${LOCAL_RPC}

.PHONY: verify-goerli-avs
verify-goerli-avs: ## Verify OmniAVS on goerli etherscan.
	@./script/avs/verify-goerli-avs.sh

.PHONY: check-abigen-version
check-abigen-version: ## Check abigen version, exit(1) if not 1.13.14-stable.
	@version=$$(abigen --version); \
	if [ "$$version" != "abigen version 1.13.14-stable" ]; then \
		echo "abigen version is not 1.13.14-stable"; \
		echo "Install with go install github.com/ethereum/go-ethereum/cmd/abigen@v1.13.14"; \
		exit 1; \
	fi

.PHONY: check-infura-key
check-infura-key: ## Check if INFURA_KEY is set.
	@if [ -z "${INFURA_KEY}" ]; then echo "INFURA_KEY is not set"; exit 1; fi
