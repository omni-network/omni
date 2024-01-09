#!/usr/bin/env bash
# Install foundryup and foundry toolkit (forge, cast, anvil)

set -e

# foundryup installs versions by tag (https://github.com/foundry-rs/foundry/tags)
TAG="nightly-02292f2d2caa547968bd039c06dc53d98b72bf39"

# output of forge --version, to check $TAG installed
VERSION="forge 0.2.0 (6fc7463 2024-01-05T00:17:41.668342000Z)"


# If not running interactively (like in Github Actions), specify FOUNDRY_BIN_DIR
# This tells https://foundry.paradigm.xyz where to install foundryup
case $- in
  *i*) ;;
  *) export FOUNDRY_BIN_DIR="$HOME/.config/.foundry/bin";
esac


# If foundryup not installed, install it
if ! which foundryup 1>/dev/null; then
  echo "Installing foundryup"
  curl -L https://foundry.paradigm.xyz | bash
fi


# If correct version not installed, install it
if ! which forge 1>/dev/null || [[ $(forge --version) != $VERSION ]]; then
  echo "Installing $VERSION"
  foundryup --version $TAG
fi
