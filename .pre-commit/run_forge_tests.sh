#!/usr/bin/env bash

# Runs `forge test` for every unique foundry project derived from the list
# of files provided as arguments by pre-commit.

source scripts/install_foundry.sh
source scripts/install_pnpm.sh

# import foundryroots
source .pre-commit/foundry_utils.sh

for dir in $(foundryroots $@); do
  echo "Running 'forge test' in ./$dir"
  (cd $dir && pnpm install && forge test)
done
