COMMIT_SHORT_FETCH:= $(shell git rev-parse --short HEAD)
export COMMIT_SHORT=$(COMMIT_SHORT_FETCH)

help:  ## Display this help message.
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-30s\033[0m %s\n", $$1, $$2}'

gen:  ## generates our graphql codegen
	pnpm run codegen -o -v

format-check:  ## Check if the code is formatted correctly
	npx prettier --check .

format: ## Format the code
	npx prettier --write .

.PHONY: build-docker
build-docker: ## Build the UI Docker Image
	@docker buildx build . -t omniops/explorer-ui:$(COMMIT_SHORT) -t omniops/explorer-ui:main --no-cache
