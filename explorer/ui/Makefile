help:  ## Display this help message.
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-30s\033[0m %s\n", $$1, $$2}'

gen-graphql:
	pnpm run codegen

format-check:  ## Check if the code is formatted correctly
	npx prettier --check .

format: ## Format the code
	npx prettier --write .
