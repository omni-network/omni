.PHONY: help
help:  ## Display this help message.
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-30s\033[0m %s\n", $$1, $$2}'

.PHONY: version
version: ## Print tool versions.
	@forge --version
	@abigen --version

PNPM_VERSION := 9.12.1
ABIGEN_VERSION := 1.14.8

.PHONY: install-pnpm
install-pnpm:
	@which npm > /dev/null || echo "npm not installed, see https://nodejs.org/en/download/package-manager"
	@which pnpm > /dev/null || (echo "pnpm not installed, installing..."; npm install -g pnpm@$(PNPM_VERSION))

.PHONY: install-deps
install-deps: check-pnpm-version ## Install dependencies.
	(cd avs && pnpm install --frozen-lockfile)
	(cd core && pnpm install --frozen-lockfile)
	(cd solve && pnpm install --frozen-lockfile)
	(cd xapps && pnpm install --frozen-lockfile)
	(cd nomina && pnpm install --frozen-lockfile)
	go install github.com/ethereum/go-ethereum/cmd/abigen@v1.14.8

.PHONY: build
build: version ## Build contracts.
	forge build --force --root core
	forge build --force --root avs
	forge build --force --root solve
	forge build --force --root xapps
	forge build --force --root nomina

.PHONY: all
all: install-deps build bindings allocs ## Build contracts, generate bindings and predeploy allocations.


CORE_CONTRACTS := OmniPortal FeeOracleV1 Create3 TransparentUpgradeableProxy \
			Staking Slashing Distribution Redenom OmniBridgeL1 OmniBridgeNative Omni WOmni \
			PortalRegistry AllocPredeploys PingPong ProxyAdmin Admin \
			OmniGasPump OmniGasStation FeeOracleV2 \
			node_modules/@openzeppelin/contracts/token/ERC20/IERC20.sol:IERC20 \
			test/StakingProxy.sol:StakingProxy

SOLVE_CONTRACTS := ISolverNetBindings SolverNetInbox SolverNetOutbox SolverNetExecutor \
			test/utils/MockERC20.sol:MockERC20 test/utils/MockVault.sol:MockVault SolverNetAdmin

AVS_CONTRACTS := OmniAVS DelegationManager StrategyManager StrategyBase AVSDirectory

XAPPS_CONTRACTS := Bridge Lockbox StablecoinUpgradeable StablecoinProxy

NOMINA_CONTRACTS := Nomina NominaBridgeL1 NominaBridgeNative

.PHONY: bindings
bindings: check-abigen-version build ## Generate golang contract bindings.
	./bindings/scripts/enumgen.sh
	ROOT=./core ./bindings/scripts/gen.sh $(CORE_CONTRACTS)
	ROOT=./avs ./bindings/scripts/gen.sh $(AVS_CONTRACTS)
	ROOT=./solve ./bindings/scripts/gen.sh $(SOLVE_CONTRACTS)
	ROOT=./xapps ./bindings/scripts/gen.sh $(XAPPS_CONTRACTS)
	ROOT=./nomina ./bindings/scripts/gen.sh $(NOMINA_CONTRACTS)
	go run ./bindings/scripts/commenttypes.go -- bindings/strategymanager.go IStrategyManagerDeprecatedStructQueuedWithdrawal IStrategyManagerDeprecatedStructWithdrawerAndNonce
	go run ./bindings/scripts/commenttypes.go -- bindings/avsdirectory.go ISignatureUtilsSignatureWithSaltAndExpiry
	go run ./bindings/scripts/commenttypes.go -- bindings/solvernetinbox.go SolverNetCall SolverNetFillOriginData SolverNetTokenExpense
	go run ./bindings/scripts/commenttypes.go -- bindings/solvernetoutbox.go IERC7683FillInstruction IERC7683Output IERC7683ResolvedCrossChainOrder ISolverNetOutboxInboxConfig

.PHONY: allocs
allocs: bindings ## Generate predeploy allocations.
	go run ./allocs/scripts/genallocs.go

.PHONY: check-abigen-version
check-abigen-version: ## Check abigen version, exit(1) if incorrect.
	@version=$$(abigen --version); \
	if [ "$$version" != "abigen version $(ABIGEN_VERSION)-stable" ]; then \
		echo "abigen version is not $(ABIGEN_VERSION)"; \
		echo "Install with go install github.com/ethereum/go-ethereum/cmd/abigen@v$(ABIGEN_VERSION)"; \
		exit 1; \
	fi

.PHONY: check-pnpm-version
check-pnpm-version: ## Check pnpm version, exit(1) if incorrect.
	@version=$$(pnpm --version); \
	if [ "$$version" != "$(PNPM_VERSION)" ]; then \
		echo "pnpm version is not $(PNPM_VERSION)"; \
		echo "Install with npm install -g pnpm@$(PNPM_VERSION)"; \
		exit 1; \
	fi
