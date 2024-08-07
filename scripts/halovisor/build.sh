#!/usr/bin/env bash

# ./build.sh <halo_genesis_version>
# This scripts builds the halovisor docker image
# Halovisor wraps cosmovisor and multiple halo versions into a single docker image.
# It allows for docker based deployments that support halo network upgrades.

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

HALO_GENESIS_VERSION="${1}"
if [ -z "$HALO_GENESIS_VERSION" ]; then
  HALO_GENESIS_VERSION=$(git rev-parse --short=7 HEAD)
  echo "Using head as HALO_GENESIS_VERSION: ${HALO_GENESIS_VERSION}"
fi

IMAGEREF="omniops/halovisor:${HALO_GENESIS_VERSION}"
IMAGEMAIN="omniops/halovisor:main"

docker build \
  --build-arg HALO_GENESIS_VERSION="${HALO_GENESIS_VERSION}" \
  -t "${IMAGEREF}" \
  -t "${IMAGEMAIN}" \
  "${SCRIPT_DIR}"

echo "Built docker image: ${IMAGEREF}"
echo "Built docker image: ${IMAGEMAIN}"

# TODOs:
# - Add support for multiple halo versions/upgrades
# - Add support for official releases
# - Support multiple architectures
