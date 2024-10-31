#!/usr/bin/env bash
# Install foundryup and foundry toolkit (forge, cast, anvil)


# If already forge installed, exit. (we do not care about cast or anvil
if which forge 1>/dev/null; then
  echo "Foundry already installed: $(forge --version)."
  echo "Run 'foundryup' to update."
  return
fi

# This tells https://foundry.paradigm.xyz where to install foundryup -
# $FOUNDRY_DIR/bin. This dir is added to $GITHUB_PATH in
# .github/workflows/pre-commit.yaml
export FOUNDRY_DIR="$HOME/.foundry";


# If foundryup not installed, install it
if ! which foundryup 1>/dev/null; then
  echo "Installing foundryup"
  curl -L https://foundry.paradigm.xyz | bash
fi


# We use the nightly version, rather than pinning to a specific version.
# foundryup does allow pinning to a specific version, via github tag:
#   foundryup --version nightly-24abca6c9133618e0c355842d2be2dd4f36da46d
# Or via git commit:
#   foundryup --commit 24abca6
# But they delete github tags frequently, and installation via commit requires
# rust and lenghty builds. So we use nightly, until versioning becomes an
# issue.

foundryup --version nightly
