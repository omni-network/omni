#!/usr/bin/env bash

set -euo pipefail

CHANGED=$(git diff --name-only --cached)
FILES=$(echo "$CHANGED" | grep '^anchor/' || true)

if [ -z "$FILES" ]; then
  echo "no /anchor/* changes, skipping rust check..."
  exit 0
fi

# nav to sdk
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
ANCHOR_DIR="$SCRIPT_DIR/../anchor"
cd "${ANCHOR_DIR}"

cargo fmt
cargo check
cargo clippy -- -D warning
