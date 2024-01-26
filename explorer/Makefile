help:  ## Display this help message.
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-30s\033[0m %s\n", $$1, $$2}'

gen-api: ## generates our openapi.yaml bindings for the api and ui
	make -C ./api gen
	cp ./api/openapi/openapi.yaml ./ui/app/openapi/openapi.yaml
	make -C ./ui gen

gen-db: ## generates our ent orm schemas
	make -C ./db gen

open-dev: ## opens all of the local host endpoints
	python -m webbrowser "http://localhost:8080"
	python -m webbrowser "http://localhost:3000"
	python -m webbrowser "http://localhost:8081/hello"

build: ## builds our docker images
	make -C ../ build-docker

run: ## runs our docker compose up
	docker-compose up

stop: ## stops our docker compose
	docker-compose down

run-clean: ## docker build and run
	make -C ./ build
	make -C ./ run
