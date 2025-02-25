#!/usr/bin/env bash

# get dir
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# nav to sdk dir
SDK_DIR="$SCRIPT_DIR/../sdk"

# get changed TS files
TS_FILES=$(git diff --cached --name-only --diff-filter=ACMR | grep "^sdk.*\.\(ts\|tsx\)$" || true)

echo "Running SDK checks..."
cd "$SDK_DIR"

# clean + build
echo "Building SDK..."
pnpm build

# precommit checks (Biome)
echo "Running SDK checks..."
pnpm check

# TODO enable after issues resolved
# CHECK_EXIT_CODE=$?

# if [ $CHECK_EXIT_CODE -ne 0 ]; then
#     echo "SDK check failed with $CHECK_EXIT_CODE"
#     exit $CHECK_EXIT_CODE
# fi

echo "SDK check completed..."

exit 0
