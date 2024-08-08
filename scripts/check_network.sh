#!/usr/bin/env bash

# ./check_network.sh (staging|omega|mainnet)
# This script checks whether network connections (TCP/HTTP) can be established to
# the wellknown bootstrap/seed nodes of the provided network. This ensures that that
# the nodes are up and running and that networking from the local machine is working.
# Note this script must be located in `scripts` folder of the omni mono repo to work.

set -e
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

NETWORK="${1}"
if [ -z "$NETWORK" ]; then
  echo "Please provide the network name: ./check_network.sh (staging|omega|mainnet)"
  exit 1
fi

# Iterate over each consensus seed in ../lib/netconf/<network>/consensus-seeds.txt
for SEED in $(cat "${SCRIPT_DIR}/../lib/netconf/${NETWORK}/consensus-seeds.txt"); do
  # Get hostname and port from: 623ab30714ecfc8a0c0da0227a1b5bd9cf5b3d8b@seed01.omega.omni.network:26656
  HOSTPORT=$(echo "${SEED}" | cut -d@ -f2)
  HOST=$(echo "${HOSTPORT}" | cut -d: -f1)
  PORT=$(echo "${HOSTPORT}" | cut -d: -f2)
  echo "Checking connection to consensus seed: ${HOST}:${PORT}"
  # Check TCP connection
  nc -z -v -w 1 "${HOST}" "${PORT}"
done

# Iterate over each execution seed in ../lib/netconf/<network>/execution-seeds.txt
for SEED in $(cat "${SCRIPT_DIR}/../lib/netconf/${NETWORK}/execution-seeds.txt"); do
  # Get hostname and port from: enode://dad1ad8680fc41c4bac3ef@seed01.omega.omni.network:30303
  HOSTPORT=$(echo "${SEED}" | cut -d@ -f2)
  HOST=$(echo "${HOSTPORT}" | cut -d: -f1)
  PORT=$(echo "${HOSTPORT}" | cut -d: -f2)
  echo "Checking connection to execution seed: ${HOST}:${PORT}"
  # Check TCP connection
  nc -z -v -w 1 "${HOST}" "${PORT}"
  # Check UDP connection
  nc -z -v -u -w 1 "${HOST}" "${PORT}"
done
