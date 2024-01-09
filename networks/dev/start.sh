#!/usr/bin/env bash
# Start a local devnet with two chains - chain1 and chain2
# Portals deployed to both at 0x5FbDB2315678afecb367f032d93F642f64180aa3

docker-compose up -d

CHAIN_A_RPC="http://localhost:6545"
CHAIN_B_RPC="http://localhost:7545"

# Private key of pre-funded anvil account 0
# Public key: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
DEPLOYER_KEY="0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"

cd ../../contracts

for rpc in $CHAIN_A_RPC $CHAIN_B_RPC ; do
  echo "Deploying to $rpc"
  forge create --rpc-url $rpc --private-key $DEPLOYER_KEY OmniPortal
done
