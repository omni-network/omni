#!/usr/bin/env bash

# get dir
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# nav to sdk dir
SDK_DIR="$SCRIPT_DIR/../sdk"

# get changed TS files
TS_FILES=$(git diff --cached --name-only --diff-filter=ACMR | grep "^sdk.*\.\(ts\|tsx\)$" || true)

if [ -n "$TS_FILES" ]; then
    echo "Running SDK checks..."
    cd "$SDK_DIR"

    # clean + build
    echo "Building SDK..."
    pnpm build:clean

    # precommit checks (Biome)
    echo "Running SDK checks..."
    pnpm precommit

    echo "SDK checks completed..."
fi
