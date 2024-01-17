#!/usr/bin/env bash
# Set of foundry shell utils


# Get foundry root of a given filepath. Searches upwards from filepath for a
# directory containing foundry.toml
foundryroot() {
  dir=$(dirname $1)

  while [[ "$dir" != "." && "$dir" != "/" ]]; do
    if [ -f "$dir/foundry.toml" ]; then
      echo "$dir"
      return
    fi

    dir=$(dirname $dir)
  done
}

# Get a list of unique foundry roots for a list of files
foundryroots() {
  for file in $@; do
    foundryroot $file
  done | sort -u
}
