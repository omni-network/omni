help:  ## Display this help message.
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-30s\033[0m %s\n", $$1, $$2}'

gen: ## generates our openapi.yaml bindings
	go run github.com/ogen-go/ogen/cmd/ogen@v0.81.1 --target openapi --clean ./openapi/openapi.yaml
	cp ./openapi/openapi.yaml ./static/openapi.yaml
