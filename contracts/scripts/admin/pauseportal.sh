#!/usr/bin/env bash
#
# Pause a portal on the given chain
#
# Usage: RPC=... ADDRESS=... NETWORK=... FIREBLOCKS_API_KEY=... FIREBLOCKS_KEY_PATH=... ./pauseportal.sh
#
# Note: This script will only work for persistent networks, not devnet, as devnet uses
#   anvil accounts as admins.


gitroot() {
  git rev-parse --show-toplevel
}

usage() {
  echo "
  Require environment variables:
  RPC                 rpc endpoint of the chain
  ADDRESS             address of the portal contract
  NETWORK             network name
  FIREBLOCKS_API_KEY  fireblocks api key
  FIREBLOCKS_KEY_PATH fireblocks key path
"
}

cd $(gitroot)
source ./contracts/scripts/admin/util.sh

if ! require_env RPC ADDRESS NETWORK FIREBLOCKS_API_KEY FIREBLOCKS_KEY_PATH; then
  usage
  exit 1
fi

echo "Pausing ${NETWORK} portal ${PORTAL} on ${CHAIN}..."

fbproxy_addr="0.0.0.0:7070"

# run fbproxy
go run ./e2e/fbproxy \
  --network $NETWORK \
  --fireblocks-api-key $FIREBLOCKS_API_KEY \
  --fireblocks-key-path $FIREBLOCKS_KEY_PATH \
  --listen-addr $fbproxy_addr \
  --base-rpc $RPC \
  &

fbproxy_pid=$!

# kill fbproxy on exit
cleanup() {
  echo "Cleaning up..."
  echo "killing fbproxy pid: $fbproxy_pid"
  pkill -P $fbproxy_pid
}

trap cleanup EXIT

# wait for a bit, then check if fbproxy is running
sleep 2
if ! kill -0 $fbproxy_pid; then
  echo "fbproxy failed to start"
  exit 1
fi

# pause the portal
forge script PausePortal \
  --root ./contracts/core \
  --broadcast \
  --unlocked \
  --slow \
  --rpc-url $fbproxy_addr \
  --sig $(cast calldata "run(string,address)" $NETWORK $PORTAL)
