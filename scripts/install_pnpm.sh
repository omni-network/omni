#!/usr/bin/env bash

VERSION="9.12.1"

if ! which pnpm 1>/dev/null || [[ $(pnpm --version) != "$VERSION" ]]; then
  echo "Installing pnpm@$VERSION"
  npm install -g pnpm@$VERSION
fi
