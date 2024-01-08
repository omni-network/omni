#!/usr/bin/env bash

# Runs `forge test` for every unique foundry project derived from the list
# of files provided as arguments by pre-commit.

source scripts/install_foundry.sh

# searches upwards from a filepath for a directory containing foundry.toml
foundryroot() {
  if [ -z "$1" ]; then
    echo "$1"
    echo "Usage: foundryroot <filepath>"
    return
  fi

  dir=$(dirname $1)

  while [[ "$dir" != "." && "$dir" != "/" ]]; do
    if [ -f "$dir/foundry.toml" ]; then
      echo "$dir"
      return
    fi

    dir=$(dirname $dir)
  done
}

# get foundryroot for every path provided
ROOTS=()
for file in $@; do
  root=$(foundryroot $file)
  if [ -n "$root" ]; then
    ROOTS+=("$root")
  fi
done

# remove duplicates
ROOTS=($(echo "${ROOTS[@]}" | tr ' ' '\n' | sort -u))

# run tests
for dir in ${ROOTS[@]}; do
  echo "Running forge tests in ./$dir"
  (cd $dir && forge test)
done
