.PHONY: help
help:  ## Display this help message.
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-30s\033[0m %s\n", $$1, $$2}'

.PHONY: version
version: ## Print tool versions.
	@forge --version
	@abigen --version

.PHONY: build
build: version ## Build contracts.
	(cd core && forge build)
	(cd avs && forge build)


.PHONY: all
all: build bindings allocs ## Build contracts, generate bindings and predeploy allocations.


CORE_CONTRACTS := OmniPortal FeeOracleV1 Create3 TransparentUpgradeableProxy \
			Staking Slashing OmniBridgeL1 OmniBridgeNative Omni WOmni \
			PortalRegistry AllocPredeploys PingPong ProxyAdmin Admin

AVS_CONTRACTS := OmniAVS DelegationManager StrategyManager StrategyBase AVSDirectory \
			avs/test/common/MockERC20.sol:MockERC20

.PHONY: bindings
bindings: check-abigen-version build ## Generate golang contract bindings.
	ROOT=./core ./bindings/scripts/gen.sh $(CORE_CONTRACTS)
	ROOT=./avs ./bindings/scripts/gen.sh $(CONTRACTS)
	go run ./bindings/scripts/commenttypes.go -- bindings/strategymanager.go IStrategyManagerDeprecatedStructQueuedWithdrawal IStrategyManagerDeprecatedStructWithdrawerAndNonce
	go run ./bindings/scripts/commenttypes.go -- bindings/avsdirectory.go ISignatureUtilsSignatureWithSaltAndExpiry

.PHONY: allocs
allocs: bindings ## Generate predeploy allocations.
	go run ./allocs/scripts/genallocs.go

.PHONY: check-abigen-version
check-abigen-version: ## Check abigen version, exit(1) if not 1.13.14-stable.
	@version=$$(abigen --version); \
	if [ "$$version" != "abigen version 1.13.14-stable" ]; then \
		echo "abigen version is not 1.13.14-stable"; \
		echo "Install with go install github.com/ethereum/go-ethereum/cmd/abigen@v1.13.14"; \
		exit 1; \
	fi
