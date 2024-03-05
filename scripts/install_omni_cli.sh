#!/usr/bin/env bash

set -e

OS=$(uname -s)
ARCH=$(uname -m)
URL="https://github.com/omni-network/omni/releases/latest/download/omni_${OS}_${ARCH}.tar.gz"
TARGET="$HOME/bin/omni"
SHELL_PROFILE=""

echo "ℹ️ Downloading omni from $URL"
echo "ℹ️ Installing omni to $TARGET"

mkdir -p "$(dirname "${TARGET}")"

curl -L -s "$URL" | tar -xz -C "$(dirname "${TARGET}")" omni
chmod +x "$TARGET"

if ! echo $PATH | grep -q "$HOME/bin"; then
    if [[ $SHELL == *"zsh"* ]]; then
        SHELL_PROFILE="$HOME/.zshrc"
    elif [[ $SHELL == *"bash"* ]]; then
        SHELL_PROFILE="$HOME/.bashrc"
    else
        echo "Unknown shell: $SHELL, you may need to add $HOME/bin to your PATH manually"
        exit 1
    fi

    echo "ℹ️ Adding $HOME/bin to your PATH in $SHELL_PROFILE"
    echo "export PATH="$PATH:$HOME/bin"" >> "$SHELL_PROFILE"
    export PATH=\$PATH:$HOME/bin
fi

which omni &> /dev/null && echo "✅ omni is now installed: try running 'omni --help' to get started" || echo "Error: omni executable not found in PATH, you may need to add $HOME/bin to your PATH"
