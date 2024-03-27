#!/usr/bin/env bash

# ./build_docker.sh <app_name>
# Builds a single docker image for the provided app using local git ref.

APP=$1

if [ -z "$APP" ]; then
  echo "Please provide the app name: ./single_docker.sh <app_name>"
  exit 1
fi

GOOS=linux goreleaser build --single-target --snapshot --clean --id="${APP}"

GITREF=$(git rev-parse --short=7 HEAD)

cd dist/${APP}* || exit 1

docker build -f "../../${APP}/Dockerfile" . -t "omniops/${APP}:${GITREF}"

echo "Built docker image: omniops/${APP}:${GITREF}"
