#!/usr/bin/env bash
# Install foundryup and foundry toolkit (forge, cast, anvil)

set -e

BASE_DIR=${XDG_CONFIG_HOME:-$HOME}

# tells https://foundry.paradigm.xyz where to install foundryup
# this way we can explicity add FOUNDRY_BIN_DIR to PATH
export FOUNDRY_DIR=${FOUNDRY_DIR-"$BASE_DIR/.foundry"}
export FOUNDRY_BIN_DIR="$FOUNDRY_DIR/bin"

# foundryup installs versions by tag (https://github.com/foundry-rs/foundry/tags)
TAG="nightly-6fc74638b797b8e109452d3df8e26758f86f31fe"

# output of forge --version, to check $TAG installed
VERSION="forge 0.2.0 (6fc7463 2024-01-05T00:17:41.668342000Z)"

# If not running interactively (like in GithubActions), add FOUNDRY_BIN_DIR to PATH
case $- in
  *i*) ;;
  *) [[ ":$PATH:" != *":${FOUNDRY_BIN_DIR}:"* ]] && PATH="${PATH}:${FOUNDRY_BIN_DIR}";;
esac


# If foundryup not installed, install it
if ! which foundryup 1>/dev/null; then
  echo "Installing foundryup"
  curl -L https://foundry.paradigm.xyz | bash

  # again if not running interactively, add FOUNDRY_BIN_DIR to PATH
  case $- in
    *i*)
      ;;
    *) [[ ":$PATH:" != *":${FOUNDRY_BIN_DIR}:"* ]] && PATH="${PATH}:${FOUNDRY_BIN_DIR}";;
  esac
fi


# If correct version not installed, install it
if ! which forge 1>/dev/null || [[ $(forge --version) != $VERSION ]]; then
  foundryup --version $TAG
fi
