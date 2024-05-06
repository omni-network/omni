#!/usr/bin/env bash

set -e

OS=$(uname -s)
ARCH=$(uname -m)

# Check if running on Windows
if [[ "$OS" == CYGWIN* ]] || [[ "$OS" == MINGW* ]] || [[ "$OS" == MSYS* ]]; then
    echo "😮 It looks like you're running this script on Windows."
    echo "✅ Please use Ubuntu or macOS to run this script, or use WSL or another VM service if you are on Windows."
    exit 0
fi

case $ARCH in
    arm64) ARCH="arm64" ;;
    aarch64) ARCH="arm64" ;;
    x86_64) ARCH="x86_64" ;;
    *) echo "Unsupported architecture"; exit 1 ;;
esac

URL="https://github.com/omni-network/omni/releases/latest/download/omni_${OS}_${ARCH}.tar.gz"
TARGET="$HOME/bin/omni"

echo "ℹ️ Downloading omni from $URL"

# Create target directory
echo "ℹ️ Installing omni to $TARGET"
mkdir -p "$(dirname "${TARGET}")"

# Download and extract omni
curl -L -s "$URL" -o omni.tar.gz
tar -xz -C "$(dirname "${TARGET}")" -f omni.tar.gz
rm omni.tar.gz
chmod +x "$TARGET"

# Add to PATH if not already there
if ! echo $PATH | grep -q "$HOME/bin"; then
    SHELL_PROFILE=""
    if [[ $SHELL == *"zsh"* ]]; then
        FILE=".zshrc"
        SHELL_PROFILE="$HOME/$FILE"
    elif [[ $SHELL == *"bash"* ]]; then
        FILE=".bashrc"
        SHELL_PROFILE="$HOME/$FILE"
    else
        echo "Unknown shell: $SHELL. Please add $HOME/bin to your PATH manually."
        exit 1
    fi

    echo "ℹ️ Adding '\$HOME/bin' to your PATH in $SHELL_PROFILE"
    echo 'export PATH="$PATH:$HOME/bin"' >> "$SHELL_PROFILE"
    export PATH=\$PATH:$HOME/bin
fi

# Check if omni is correctly installed
if command -v omni &> /dev/null; then
    echo "✅ omni is now installed. Run 'omni' to get started."
else
    echo "❌ Error: omni executable not found in PATH. You may need to add '\$HOME/bin' to your PATH manually."
fi
