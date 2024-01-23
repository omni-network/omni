help:  ## Display this help message.
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-30s\033[0m %s\n", $$1, $$2}'

gen-api: ## generates our openapi.yaml bindings for the api and ui
	make -C ./api gen
	cp ./api/openapi/openapi.yaml ./ui/app/openapi/openapi.yaml
	make -C ./ui gen

gen-db: ## generates our ent orm schemas
	make -C ./db gen
