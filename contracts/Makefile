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



CONTRACTS := OmniPortal FeeOracleV1 OmniAVS OmniStake Create3 TransparentUpgradeableProxy \
			DelegationManager StrategyManager StrategyBase AVSDirectory \
			src/deploy/ProxyAdmin.sol:ProxyAdmin \
			test/avs/common/MockERC20.sol:MockERC20

.PHONY: bindings
bindings: check-abigen-version build ## Generate golang contract bindings.
	./bindings/scripts/gen.sh $(CONTRACTS)
	./bindings/scripts/genmore.sh OmniStake Create3 TransparentUpgradeableProxy src/deploy/ProxyAdmin.sol:ProxyAdmin
	go run ./bindings/scripts/commenttypes.go -- bindings/strategymanager.go IStrategyManagerDeprecatedStructQueuedWithdrawal IStrategyManagerDeprecatedStructWithdrawerAndNonce
	go run ./bindings/scripts/commenttypes.go -- bindings/avsdirectory.go ISignatureUtilsSignatureWithSaltAndExpiry

.PHONY: examples-bindings
examples-bindings: check-abigen-version build ## Generate golang contract bindings for example contracts
	PKG=examples DIR=./bindings/examples ./bindings/scripts/gen.sh PingPong

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


.PHONY: check-abigen-version
check-abigen-version: ## Check abigen version, exit(1) if not 1.13.14-stable.
	@version=$$(abigen --version); \
	if [ "$$version" != "abigen version 1.13.14-stable" ]; then \
		echo "abigen version is not 1.13.14-stable"; \
		echo "Install with go install github.com/ethereum/go-ethereum/cmd/abigen@v1.13.14"; \
		exit 1; \
	fi
