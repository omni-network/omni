help:  ## Display this help message
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-30s\033[0m %s\n", $$1, $$2}'

###############################################################################
###                                Docker                                 	###
###############################################################################

.PHONY: build-docker
build-docker: ensure-go-releaser ## Builds the docker images.
	@goreleaser release --snapshot --clean

###############################################################################
###                                Contracts                                 ###
###############################################################################

.PHONY: contracts-gen
contract-bindings: ## Generate golang contract bindings.
	make -C ./contracts bindings

###############################################################################
###                               Explorer                                 	###
###############################################################################

.PHONY: explorer-gen
explorer-gen: ## Generates code for our explorer
	make -C ./explorer/db gen
	make -C ./explorer/api gen

###############################################################################
###                                Utils                                 	###
###############################################################################

.PHONY: ensure-go-releaser
ensure-go-releaser: ## Installs the go-releaser tool.
	@which goreleaser > /dev/null || echo "go-releaser not installed, see https://goreleaser.com/install/"

.PHONY: install-pre-commit
install-pre-commit: ## Installs the pre-commit tool as the git pre-commit hook for this repo.
	@which pre-commit > /dev/null || echo "pre-commit not installed, see https://pre-commit.com/#install"
	@pre-commit install --install-hooks

.PHONY: lint
lint: ## Runs linters via pre-commit.
	@pre-commit run -v --all-files

###############################################################################
###                                Testing                                 	###
###############################################################################

.PHONY: halo-simnet
halo-simnet: ## Runs halo in simnet mode.
	@go install github.com/omni-network/omni/halo
	@halo init --home=/tmp/halo --network=simnet --clean
	@halo run --home=/tmp/halo
