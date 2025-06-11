#!/usr/bin/env bash

# Runs anchor build inside omniops/anchor:v0.31.0 docker container
# which has correct versions of rust, solana and anchor.
# See everclear for source: from https://github.com/everclearorg/monorepo/tree/dev/packages/contracts/solana-spoke/docker
# Note this container is huge, 2GB 🥲#
#
# It also has prebuilt artefacts located at /app/target which is copy to mounted to /anchor/target to speed up builds.

set -euo pipefail

anchor --version
solana --version
cargo --version
rustc --version

# link /app/target to /anchor/target
cd /anchor || exit 1
mv ./target ./target_bak || true
ln -s /app/target /anchor/target

# Ensure the correct keypair is used
cp ./target_bak/deploy/solver_inbox-keypair.json ./target/deploy/solver_inbox-keypair.json

anchor keys sync
anchor build

mkdir -p ./target_bak/idl
mkdir -p ./target_bak/deploy
mkdir -p ./target_bak/types

cp ./target/idl/solver_inbox.json ./target_bak/idl/solver_inbox.json
cp ./target/deploy/solver_inbox.so ./target_bak/deploy/solver_inbox.so
cp ./target/types/solver_inbox.ts ./target_bak/types/solver_inbox.ts

# delete the symlink
rm ./target
mv ./target_bak ./target
