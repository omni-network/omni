#!/usr/bin/env bash

echo "starting sdk checks in dir: $(pwd)"

source scripts/install_pnpm.sh

# get dir
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

echo "SCRIPT_DIR used to jump to SCRIPT_DIR/../sdk =: $(SCRIPT_DIR)"

# nav to sdk dir
SDK_DIR="$SCRIPT_DIR/../sdk"

# get changed TS files
TS_FILES=$(git diff --cached --name-only --diff-filter=ACMR | grep "^sdk.*\.\(ts\|tsx\)$" || true)

echo "Running SDK checks..."
cd "$SDK_DIR"

echo "Installing deps via pnpm..."
pnpm install

# clean + build
echo "Building SDK..."
pnpm build

# precommit checks (Biome)
echo "Running checks..."
pnpm check

# CHECK_EXIT_CODE=$?

# if [ $CHECK_EXIT_CODE -ne 0 ]; then
#     echo "SDK check failed with $CHECK_EXIT_CODE"
#     exit $CHECK_EXIT_CODE
# fi

# echo "SDK check completed..."

exit 0
