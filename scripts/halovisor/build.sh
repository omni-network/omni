#!/usr/bin/env bash

# ./build.sh
# This scripts builds the halovisor docker image
# Halovisor wraps cosmovisor and multiple halo versions into a single docker image.
# It allows for docker based deployments that support halo network upgrades.

set -e
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

HALO_VERSION_N_LATEST=$(git rev-parse --short=7 HEAD)
echo "Using head as HALO_VERSION_N_LATEST: ${HALO_VERSION_N_LATEST}"

IMAGEREF="omniops/halovisor:${HALO_VERSION_N_LATEST}"
IMAGEMAIN="omniops/halovisor:main"

docker build \
  --build-arg HALO_VERSION_N_LATEST="${HALO_VERSION_N_LATEST}" \
  -t "${IMAGEREF}" \
  -t "${IMAGEMAIN}" \
  "${SCRIPT_DIR}"

echo "Built docker image: ${IMAGEREF}"
echo "Built docker image: ${IMAGEMAIN}"
