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
	make -C ./explorer gen-db
	make -C ./explorer gen-api

###############################################################################
###                                Utils                                 	###
###############################################################################

.PHONY: ensure-go-releaser
ensure-go-releaser: ## Installs the go-releaser tool.
	@which goreleaser > /dev/null || echo "go-releaser not installed, see https://goreleaser.com/install/"

.PHONY: ensure-detect-secrets
ensure-detect-secrets: ## Checks if detect-secrets is installed.
	@which detect-secrets > /dev/null || echo "detect-secrets not installed, see https://github.com/Yelp/detect-secrets?tab=readme-ov-file#installation"

.PHONY: install-pre-commit
install-pre-commit: ## Installs the pre-commit tool as the git pre-commit hook for this repo.
	@which pre-commit > /dev/null || echo "pre-commit not installed, see https://pre-commit.com/#install"
	@pre-commit install --install-hooks

.PHONY: install-go-tools
install-go-tools: ## Installs the go-dev-tools, like buf.
	@go generate tools.go

.PHONY: lint
lint: ## Runs linters via pre-commit.
	@pre-commit run -v --all-files

.PHONY: bufgen
bufgen: ## Generates protobufs using buf generate.
	@./scripts/buf_generate.sh

.PHONY:
secrets-baseline: ensure-detect-secrets ## Update secrets baseline.
	@detect-secrets scan --exclude-file pnpm-lock.yaml > .secrets.baseline

###############################################################################
###                                Testing                                 	###
###############################################################################

PWD := $(shell pwd)

DEFAULT_E2E_FLAGS := --eigen-file $(PWD)/contracts/script/eigen/output/deployments.json \
					 --anvil-state mock_l1=$(PWD)/contracts/script/eigen/output/anvil-state.json

E2E_FLAGS ?= $(DEFAULT_E2E_FLAGS)

.PHONY: halo-simnet
halo-simnet: ## Runs halo in simnet mode.
	@go install github.com/omni-network/omni/halo
	@halo init --home=/tmp/halo --network=simnet --clean
	@halo run --home=/tmp/halo

.PHONY: devnet-deploy
devnet-deploy: ## Deploys devnet1
	@echo "Creating a docker-compose devnet in ./test/e2e/run/devnet1"
	@go run github.com/omni-network/omni/test/e2e -f test/e2e/manifests/devnet1.toml deploy $(E2E_FLAGS)

.PHONY: devnet-clean
devnet-clean: ## Deletes devnet1 containers
	@echo "Stopping the devnet in ./test/e2e/run/devnet1"
	@go run github.com/omni-network/omni/test/e2e -f test/e2e/manifests/devnet1.toml clean

.PHONY: e2e-run
e2e-run: ## Run specific e2e manifest (MANIFEST=single, MANIFEST=simple, etc). Note container remain running after the test.
	@if [ -z "$(MANIFEST)" ]; then echo "⚠️ Please specify a manifest: MANIFEST=simple make e2e-run" && exit 1; fi
	@echo "Using MANIFEST=$(MANIFEST)"
	@go run github.com/omni-network/omni/test/e2e -f test/e2e/manifests/$(MANIFEST).toml $(E2E_FLAGS)

.PHONY: e2e-logs
e2e-logs: ## Print the docker logs of previously ran e2e manifest (single, simple, etc).
	@if [ -z "$(MANIFEST)" ]; then echo "⚠️  Please specify a manifest: MANIFEST=simple make e2e-logs" && exit 1; fi
	@echo "Using MANIFEST=$(MANIFEST)"
	@go run github.com/omni-network/omni/test/e2e -f test/e2e/manifests/$(MANIFEST).toml logs

.PHONY: e2e-clean
e2e-clean: ## Deletes all running containers from previously ran e2e.
	@if [ -z "$(MANIFEST)" ]; then echo "⚠️  Please specify a manifest: MANIFEST=simple make e2e-clean" && exit 1; fi
	@echo "Using MANIFEST=$(MANIFEST)"
	@go run github.com/omni-network/omni/test/e2e -f test/e2e/manifests/$(MANIFEST).toml clean
