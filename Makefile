help:  ## Display this help message
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-30s\033[0m %s\n", $$1, $$2}'

###############################################################################
###                                Docker                                 	###
###############################################################################

build-docker: ensure-go-releaser ## Builds the docker images.
	@goreleaser release --snapshot --clean

###############################################################################
###                                Contracts                                 ###
###############################################################################

contract-bindings: ## Generate golang contract bindings.
	make -C ./contracts bindings

###############################################################################
###                                Utils                                 	###
###############################################################################

gen: ## Runs our gen code
	cd lib/db/ent && go generate ./

ensure-go-releaser: ## Installs the go-releaser tool.
	@which goreleaser > /dev/null || echo "go-releaser not installed, see https://goreleaser.com/install/"

install-pre-commit: ## Installs the pre-commit tool as the git pre-commit hook for this repo.
	@which pre-commit > /dev/null || echo "pre-commit not installed, see https://pre-commit.com/#install"
	@pre-commit install --install-hooks

lint: ## Runs linters via pre-commit.
	@pre-commit run -v --all-files

.PHONY: install-pre-commit lint
