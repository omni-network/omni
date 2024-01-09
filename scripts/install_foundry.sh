#!/usr/bin/env bash
# Install foundryup and foundry toolkit (forge, cast, anvil)

set -e

# foundryup installs versions by tag (https://github.com/foundry-rs/foundry/tags)
TAG="nightly-02292f2d2caa547968bd039c06dc53d98b72bf39"

# output of forge --version, to check $TAG installed
VERSION="forge 0.2.0 (02292f2 2024-01-08T00:26:30.856271000Z)"


# If not running interactively (like in Github Actions), specify FOUNDRY_DIR
# This tells https://foundry.paradigm.xyz where to install foundryup -
# $FOUNDRY_DIR/bin. This dir is added to $GITHUB_PATH in
# .github/workflows/pre-commit.yml
case $- in
  *i*) ;;
  *) export FOUNDRY_DIR="$HOME/.foundry";
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
