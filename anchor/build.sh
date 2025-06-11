#!/usr/bin/env bash

# Build.sh builds the anchor inbox program for the specified network and generate go bindings if the network is localnet.
# It requires inbox program private key (from which the program ID is derived) to be present in ./$NET/solver_inbox-keypair.json.
# It must be run from the omni/anchor directory.
# It copies the compiled artifacts to ./$NET/ (IDL, share object, and TypeScript bindings).
# It uses backpackapp/build:v0.31.0 container to avoid local dependencies, see anchorbuild.sh for details.
# Go bindings do however require solana-anchor-go to be installed locally.

set -euo pipefail

NET=${1}
echo "Building network: ${NET}"

# If target_back exists, remove target and restore target_back
if [ -d "./target_bak" ]; then
    rm -rf ./target
    mv ./target_bak ./target
fi

# Ensure the "${NET}" keypair is used
mkdir -p ./target/deploy
cp ./"${NET}"/solver_inbox-keypair.json ./target/deploy/solver_inbox-keypair.json

# Mount . in backpackapp/build:v0.31.0 and run ./build.sh
docker run --rm -v "${PWD}":/anchor -w /anchor omniops/anchor:v0.31.0 ./anchorbuild.sh

# Copy the compiled artefacts to "${NET}"
cp ./target/idl/solver_inbox.json ./"${NET}"/solver_inbox.json
cp ./target/deploy/solver_inbox.so ./"${NET}"/solver_inbox.so
cp ./target/types/solver_inbox.ts ./"${NET}"/solver_inbox.ts

# Generate go bindings (only if localnet)
if [ "${NET}" != "localnet" ]; then exit 0; fi

solana-anchor-go -src=./target/idl/solver_inbox.json -pkg=anchorinbox -dst=./anchorinbox -remove-account-suffix
# Fix go bindings (replace .ParamsOrderId with .Params.OrderId)
sed -i.bak 's/\.ParamsOrderId/\.Params.OrderId/g' anchorinbox/open.go
rm anchorinbox/open.go.bak
