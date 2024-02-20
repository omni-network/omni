#!/usr/bin/env bash

project_root="$(git rev-parse --show-toplevel)"
anvil_state_file="${project_root}/test/e2e/app/static/el-anvil-state.json"
tmp_file=$(mktemp)

# anvil private key 9
DEPLOYER_KEY="${DEPLOYER_KEY:-0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6}"

anvil --dump-state $tmp_file &

forge script script/eigen/DeployEigenLayer.s.sol \
  --rpc-url http://localhost:8545 \
  --private-key $DEPLOYER_KEY \
  --broadcast

# # kill anvil to save its state
pkill anvil

jq '.block.number = "0x0"' $tmp_file > $anvil_state_file
