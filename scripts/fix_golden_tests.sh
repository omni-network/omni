#!/usr/bin/env bash

# fix_golden_tests.sh fixes all golden test fixtures in the repo.

# Find all unique directories with _test.go files with tutil.RequireGolden* calls.
DIRS=$(grep -rl 'tutil.RequireGolden' . | xargs dirname | uniq)

# Run `go:generate` for each directory.
for DIR in $DIRS; do
  cd "${DIR}" || exit 1
  echo "Fixing golden tests in: ${DIR}"
  go generate
  cd - || exit 1
done
