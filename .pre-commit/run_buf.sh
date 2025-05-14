#!/usr/bin/env bash

if ! which buf 1>/dev/null; then
  echo "Installing buf"
  go generate scripts/tools.go
fi

EXPECT=$(go list -f "{{.Module.Version}}" github.com/bufbuild/buf/cmd/buf)
ACTUAL="v$(buf --version)"
if [[ "${EXPECT}" != "${ACTUAL}" ]]; then
  echo "Updating buf"
  go generate scripts/tools.go
fi

./scripts/buf_generate.sh
buf lint
