GO ?= go
GOBIN ?= $$($(GO) env GOPATH)/bin
GOLANGCI_LINT ?= $(GOBIN)/golangci-lint
GOLANGCI_LINT_VERSION ?= v0.1.0

VERSION ?= "$(shell git describe --tags --abbrev=0 | cut -c2-)"
COMMIT_HASH ?= "$(shell git describe --long --dirty --always --match "" || true)"
CLEAN_COMMIT ?= "$(shell git describe --long --always --match "" || true)"
COMMIT_TIME ?= "$(shell git show -s --format=%ct $(CLEAN_COMMIT) || true)"
LDFLAGS ?= -s -w \
-X github.com/omni-network/omni.version="$(VERSION)" \
-X github.com/omni-network/omni.commitHash="$(COMMIT_HASH)" \
-X github.com/omni-network/omni.commitTime="$(COMMIT_TIME)"


.PHONY: all
all: clean binary # lint test-race

.PHONY: build
build: export CGO_ENABLED=0
build:
	$(GO) build -trimpath -ldflags "$(LDFLAGS)" ./...

.PHONY: binary
binary: export CGO_ENABLED=0
binary: dist
	$(GO) version
	$(GO) build -trimpath -ldflags "$(LDFLAGS)" -o build/omni ./cmd/omni

#.PHONY: lint
#lint: linter
#	$(GOLANGCI_LINT) run ./...
#
#.PHONY: linter
#linter:
#	test -f $(GOLANGCI_LINT) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$($(GO) env GOPATH)/bin $(GOLANGCI_LINT_VERSION)


.PHONY: clean
clean:
	$(GO) clean
	rm -rf dist/

dist:
	mkdir $@


