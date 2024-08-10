#!/usr/bin/env bash

# ./build.sh <HALO_VERSION_GENESIS>
# This scripts builds the halovisor docker image
# Halovisor wraps cosmovisor and multiple halo versions into a single docker image.
# It allows for docker based deployments that support halo network upgrades.

set -e
set -x
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

HALO_VERSION_GENESIS="${1}"
if [ -z "$HALO_VERSION_GENESIS" ]; then
  HALO_VERSION_GENESIS=$(git rev-parse --short=7 HEAD~2)
  echo "Using head as HALO_VERSION_GENESIS: ${HALO_VERSION_GENESIS}"
fi
echo "HALO_VERSION_GENESIS: ${HALO_VERSION_GENESIS}"

HALO_VERSION_V2="${2}"
if [ -z "$HALO_VERSION_V2" ]; then
  HALO_VERSION_V2=$(git rev-parse --short=7 HEAD^)
  echo "Using head as HALO_VERSION_V2: ${HALO_VERSION_V2}"
fi
echo "HALO_VERSION_V2: ${HALO_VERSION_V2}"

IMAGEREF="omniops/halovisor:${HALO_VERSION_V2}"
IMAGEMAIN="omniops/halovisor:main"

HALO_VERSION_V3="${3}"
if [ -z "$HALO_VERSION_V3" ]; then
  HALO_VERSION_V3=$(git rev-parse --short=7 HEAD)
  echo "Using head as HALO_VERSION_V3: ${HALO_VERSION_V3}"
fi
echo "HALO_VERSION_V3: ${HALO_VERSION_V3}"

IMAGEREF="omniops/halovisor:${HALO_VERSION_V3}"
IMAGEMAIN="omniops/halovisor:main"

docker build \
  --build-arg HALO_VERSION_GENESIS="${HALO_VERSION_GENESIS}" \
  --build-arg HALO_VERSION_V2="${HALO_VERSION_V2}" \
  --build-arg HALO_VERSION_V3="${HALO_VERSION_V3}" \
  -t "${IMAGEREF}" \
  -t "${IMAGEMAIN}" \
  "${SCRIPT_DIR}"

echo "Built docker image: ${IMAGEREF}"
echo "Built docker image: ${IMAGEMAIN}"

# TODOs:
# - Add support for multiple halo versions/upgrades
# - Add support for official releases
# - Support multiple architectures
