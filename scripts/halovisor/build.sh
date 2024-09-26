#!/usr/bin/env bash

# ./build.sh <HALO_VERSION_0_GENESIS>
# This scripts builds the halovisor docker image
# Halovisor wraps cosmovisor and multiple halo versions into a single docker image.
# It allows for docker based deployments that support halo network upgrades.

set -e
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

HALO_VERSION_0_GENESIS="${1}"
if [ -z "${HALO_VERSION_0_GENESIS}" ]; then
  HALO_VERSION_0_GENESIS=$(git rev-parse --short=7 HEAD)
  echo "Using head as HALO_VERSION_GENESIS: ${HALO_VERSION_0_GENESIS}"
fi

IMAGEREF="omniops/halovisor:${HALO_VERSION_0_GENESIS}"
IMAGEMAIN="omniops/halovisor:main"

docker build \
  --build-arg HALO_VERSION_0_GENESIS="${HALO_VERSION_0_GENESIS}" \
  -t "${IMAGEREF}" \
  -t "${IMAGEMAIN}" \
  "${SCRIPT_DIR}"

echo "Built docker image: ${IMAGEREF}"
echo "Built docker image: ${IMAGEMAIN}"

# TODOs:
# - Add support for multiple halo versions/upgrades
# - Add support for official releases
# - Support multiple architectures
