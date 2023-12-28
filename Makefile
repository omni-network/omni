GO ?= go
GOFMT ?= gofmt
GOBIN ?= $$($(GO) env GOPATH)/bin
GOLANGCI_LINT ?= $(GOBIN)/golangci-lint
GOLANGCI_LINT_VERSION ?= v1.55.0

VERSION ?= "$(shell git describe --tags --abbrev=0 | cut -c2-)"
COMMIT_HASH ?= "$(shell git describe --long --dirty --always --match "" || true)"
CLEAN_COMMIT ?= "$(shell git describe --long --always --match "" || true)"
LDFLAGS ?= -s -w \
-X github.com/omni-network/omni.version="$(VERSION)" \
-X github.com/omni-network/omni.commitHash="$(COMMIT_HASH)" \
-X github.com/omni-network/omni.commitTime="$(COMMIT_TIME)"
DEST ?= "$(shell (go list ./... ))"


.PHONY: all
all: clean lint vet test-race binary

.PHONY: binary
binary: export CGO_ENABLED=0
binary: dist
	$(GO) version
	$(GO) build -trimpath -ldflags "$(LDFLAGS)" -o dist/omni ./cmd/omni

.PHONY: build
build: export CGO_ENABLED=0
build:
	$(GO) build -trimpath -ldflags "$(LDFLAGS)" ./...

dist:
	mkdir $@

FOLDER=$(shell pwd)

.PHONY: install-formatters
install-formatters:
	$(GO) get github.com/daixiang0/gci
	$(GO) install mvdan.cc/gofumpt@latest

.PHONY: format
format:
	$(GOFMT) -s -w .
	$(GOBIN)/gofumpt -l -w $(FOLDER)
	$(GOBIN)/gci -w -local $(go list -m) `find $(FOLDER) -type f \! -name "*.go" \! -path \*/\.git/\* -exec echo {} \;`

.PHONY: lint
lint: linter
	$(GOLANGCI_LINT) run ./...

.PHONY: linter
linter:
	test -f $(GOLANGCI_LINT) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$($(GO) env GOPATH)/bin $(GOLANGCI_LINT_VERSION)

.PHONY: vet
vet:
	$(GO) vet "$(DEST)"

.PHONY: test-race
test-race:
	$(GO) test -race -timeout 300000ms -v "$(DEST)"

.PHONY: test
test:
	$(GO) test -v "$(DEST)"


.PHONY: clean
clean:
	$(GO) clean
	rm -rf dist/
	rm -rf build/


