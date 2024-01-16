.PHONY: gen

gen:
	go run github.com/ogen-go/ogen/cmd/ogen@v0.81.1 --target openapi --clean ./openapi/openapi.yaml
	cp ./openapi/openapi.yaml ./static/openapi.yaml
