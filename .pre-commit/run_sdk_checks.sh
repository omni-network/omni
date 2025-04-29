#!/usr/bin/env bash

set -euo pipefail

CHANGED=$(git diff --name-only --cached)
FILES=$(echo "$CHANGED" | grep '^sdk/' || true)

if [ -z "$FILES" ]; then
  echo "no sdk changes, skipping SDK check..."
  exit 0
fi

echo "Running Biome on SDK files..."

# nav to sdk
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
SDK_DIR="$SCRIPT_DIR/../sdk"
cd "$SDK_DIR"

# install deps
echo "Installing deps via pnpm..."
pnpm install

# clean + build
pnpm build

pnpm run check
