ifneq ("$(wildcard .env)","")
	include .env
	export $(shell sed 's/=.*//' .env)
endif


help:  ## Display this help message.
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-30s\033[0m %s\n", $$1, $$2}'

version: ## Print tool versions.
	@forge --version
	@abigen --version

build: version ## Build contracts.
	forge build

bindings: build ## Generate golang contract bindings.
	./bindings/scripts/gen.sh OmniPortal FeeOracleV1 OmniAVS TransparentUpgradeableProxy ProxyAdmin DelegationManager StrategyManager StrategyBase MockERC20 AVSDirectory
	go run ./bindings/scripts/commenttypes.go -- bindings/strategymanager.go IStrategyManagerDeprecatedStructQueuedWithdrawal IStrategyManagerDeprecatedStructWithdrawerAndNonce
	go run ./bindings/scripts/commenttypes.go -- bindings/avsdirectory.go ISignatureUtilsSignatureWithSaltAndExpiry

examples-bindings: build ## Generate golang contract bindings for example contracts
	PKG=examples DIR=./bindings/examples ./bindings/scripts/gen.sh PingPong

deploy-eigenlayer-save-anvil-state: ## Deploy eigenlayer contracts and dump anvil state.
	./script/eigen/deploy-eigenlayer-save-anvil-state.sh

avs-fork-test-goerli: ## Run OmniAVS tests on a goerli fork.
	@if [ -z "${INFURA_KEY}" ]; then echo "INFURA_KEY is not set"; exit 1; fi
	@FOUNDRY_FUZZ_RUNS=10 forge test --fork-url https://goerli.infura.io/v3/${INFURA_KEY} --fork-block-number 10580259 --match-contract OmniAVS_Test
