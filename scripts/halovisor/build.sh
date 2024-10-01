#!/usr/bin/env bash

# ./build.sh <HALO_VERSION_0_GENESIS> <HALO_VERSION_1_ULUWATU>
# This scripts builds the halovisor docker image
# Halovisor wraps cosmovisor and multiple halo versions into a single docker image.
# It allows for docker based deployments that support halo network upgrades.

set -e
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

HALO_VERSION_0_GENESIS="${1}"
if [ -z "$HALO_VERSION_0_GENESIS" ]; then
  # TODO(corver): Replace with v0.8.1 when released. Below is v0.8.1-rc1
  HALO_VERSION_0_GENESIS=v0.8.0-rc1
  echo "Using HALO_VERSION_GENESIS: ${HALO_VERSION_0_GENESIS}"
fi

HALO_VERSION_1_ULUWATU="${2}"
if [ -z "$HALO_VERSION_1_ULUWATU" ]; then
  HALO_VERSION_1_ULUWATU=$(git rev-parse --short=7 HEAD)
  echo "Using head as HALO_VERSION_ULUWATU: ${HALO_VERSION_1_ULUWATU}"
fi

IMAGEREF="omniops/halovisor:${HALO_VERSION_1_ULUWATU}"
IMAGEMAIN="omniops/halovisor:main"

docker build \
  --build-arg HALO_VERSION_0_GENESIS="${HALO_VERSION_0_GENESIS}" \
  --build-arg HALO_VERSION_1_ULUWATU="${HALO_VERSION_1_ULUWATU}" \
  -t "${IMAGEREF}" \
  -t "${IMAGEMAIN}" \
  "${SCRIPT_DIR}"

echo "Built docker image: ${IMAGEREF}"
echo "Built docker image: ${IMAGEMAIN}"
