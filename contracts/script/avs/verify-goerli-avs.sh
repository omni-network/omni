#!/usr/bin/env bash


COMPILER_VERSION=0.8.12
DEPLOYMENT_FILE=script/avs/output/deploy-goerli-avs.json


read() {
  echo $(cat $DEPLOYMENT_FILE | jq -r "$1")
}

contracts=("OmniAVS" "TransparentUpgradeableProxy")

for contract in "${contracts[@]}"; do
  addr=$(read ".contracts.$contract.address")
  args=$(read ".contracts.$contract.constructorArgs")

  echo "Verifying $contract at $addr with args $args"

  forge verify-contract $addr $contract \
      --constructor-args $args \
      --chain goerli \
      --compiler-version $COMPILER_VERSION \
      --watch
done
