help:  ## Display this help message.
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-30s\033[0m %s\n", $$1, $$2}'

gen: ## generates our openapi.yaml bindings for our ui
	npx openapi-typescript ./app/openapi/openapi.yaml -o ./app/openapi/openapi.d.ts

gen-graphql: # generates our graphql codegen
	pnpm run codegen -o -v

format-check:  ## Check if the code is formatted correctly
	npx prettier --check .

format: ## Format the code
	npx prettier --write .
