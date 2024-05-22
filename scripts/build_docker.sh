#!/usr/bin/env bash

# ./build_docker.sh <app_name>
# Builds a single docker image for the provided app using local git ref.

APP="${1}"
PARENT_DIR="${2}" # Defaults to repo root
GITREF="${3}"     # Defaults to git rev-parse HEAD
ARC="${4}"        # Defaults to go env GOARCH

if [ -z "$APP" ]; then
  echo "Please provide the app name: ./single_docker.sh <app_name>"
  exit 1
fi
if [ -z "${PARENT_DIR}" ]; then
    PARENT_DIR="."
fi
if [ -z "${GITREF}" ]; then
    GITREF=$(git rev-parse --short=7 HEAD)
fi
if [ -z "${ARC}" ]; then
    ARC=$(go env GOARCH)
fi

GOOS=linux GOARCH="${ARC}" goreleaser build --single-target --snapshot --clean --id="${APP}"

cd dist/${APP}* || exit 1

docker build -f "../../${PARENT_DIR}/${APP}/Dockerfile" . -t "omniops/${APP}:${GITREF}"

echo "Built docker image: omniops/${APP}:${GITREF}"
