#!/usr/bin/env bash

# get dir
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# nav to sdk dir
SDK_DIR="$SCRIPT_DIR/../sdk"

# when running pre-commit, files are passed as args
# when running git hook, find ourselves
if [ $# -eq 0 ]; then
    TS_FILES=$(git diff --cached --name-only --diff-filter=ACMR | grep "^sdk.*\.\(ts\|tsx\)$" || true)
else
    TS_FILES="$@"
fi

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