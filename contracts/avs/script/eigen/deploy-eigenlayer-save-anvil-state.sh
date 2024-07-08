#!/usr/bin/env bash

project_root="$(git rev-parse --show-toplevel)"

# we dump anvil state to a tmp file, and copy to e2e/app/static
e2e_anvil_state_file="${project_root}/e2e/app/static/el-anvil-state.json"
tmp_anvil_state_file=$(mktemp)

# forge script outputs to ./output/script/deployments.json, we copy this to e2e/app/static
e2e_deployments_file="${project_root}/e2e/app/static/el-deployments.json"
script_output_file="${project_root}/contracts/avs/script/eigen/output/deployments.json"

# anvil private key 9
DEPLOYER_KEY="0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6"

anvil --dump-state $tmp_anvil_state_file &

# wait for anvil to start
sleep 1

forge script script/eigen/DeployEigenLayer.s.sol \
  --rpc-url http://localhost:8545 \
  --private-key $DEPLOYER_KEY \
  --broadcast

# kill anvil to save its state
pkill anvil

# set block number to 0 (e2e tests break otherwise), and copy to e2e static folder
jq '.block.number = "0x0"' $tmp_anvil_state_file > $e2e_anvil_state_file

# copy deployments to e2e static folder
jq '.' $script_output_file > $e2e_deployments_file
