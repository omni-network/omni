#!/usr/bin/env bash

set -e

OS=$(uname -s)
ARCH=$(uname -m)
URL="https://github.com/omni-network/omni/releases/latest/download/omni_${OS}_${ARCH}.tar.gz"
TARGET="$HOME/bin/omni"

echo "ℹ️ Downloading omni from $URL"
echo "ℹ️ Installing omni to $TARGET"

mkdir -p "$(dirname "${TARGET}")"

curl -L -s "$URL" | tar -xz -C "$(dirname "${TARGET}")" omni
chmod +x "$TARGET"


which -s omni || echo "$HOME/bin is not in your PATH. You can add it by running 'export PATH=\$PATH:$HOME/bin'"

echo "✅ omni is now installed: try running 'omni --help' to get started"
