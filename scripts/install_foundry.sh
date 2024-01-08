#!/usr/bin/env bash
# Install foundryup and foundry toolkit (forge, cast, anvil)

set -e

# foundryup installs versions by tag (https://github.com/foundry-rs/foundry/tags)
TAG="nightly-6fc74638b797b8e109452d3df8e26758f86f31fe"

# output of forge --version, to check $TAG installed
VERSION="forge 0.2.0 (6fc7463 2024-01-05T00:17:41.668342000Z)"

if ! which foundryup 1>/dev/null; then
  echo "Installing foundryup"
  curl -L https://foundry.paradigm.xyz | bash
fi

if ! which forge 1>/dev/null || [[ $(forge --version) != $VERSION ]]; then
  echo "Installing $VERSION"
  foundryup --version $TAG
fi
