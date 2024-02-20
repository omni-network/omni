#!/usr/bin/env bash


export OUTPUT_DIR="${OUTPUT_DIR:-"script/eigen/output"}"
ANVIL_STATE_FILE="${OUTPUT_DIR}/anvil-state.json"


# anvil private key 9
DEPLOYER_KEY="${DEPLOYER_KEY:-0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6}"

anvil --dump-state $ANVIL_STATE_FILE &

forge script script/eigen/DeployEigenLayer.s.sol \
  --rpc-url http://localhost:8545 \
  --private-key $DEPLOYER_KEY \
  --broadcast

# # kill anvil to save its state
pkill anvil

mv $ANVIL_STATE_FILE $OUTPUT_DIR/anvil-state.json.tmp
jq '.block.number = "0x0"' $OUTPUT_DIR/anvil-state.json.tmp > $ANVIL_STATE_FILE
rm $OUTPUT_DIR/anvil-state.json.tmp
