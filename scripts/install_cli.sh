#!/bin/bash

# Set the desired version of omni
OMNI_VERSION="0.1.0"

# Determine OS and Architecture
OS=$(uname -s)
ARCH=$(uname -m)

case $OS in
    Darwin) OS="Darwin" ;;
    Linux) OS="Linux" ;;
    *) echo "Unsupported OS"; exit 1 ;;
esac

case $ARCH in
    arm64) ARCH="arm64" ;;
    x86_64) ARCH="x86_64" ;;
    *) echo "Unsupported architecture"; exit 1 ;;
esac

# Formulate download URL
DOWNLOAD_URL="https://github.com/omni-network/omni/releases/download/v${OMNI_VERSION}/omni_${OS}_${ARCH}.tar.gz"

# Download and extract omni
echo "Downloading omni for ${OS}/${ARCH}"
curl -L ${DOWNLOAD_URL} -o omni_${OS}_${ARCH}.tar.gz

echo "Extracting omni"
tar -xzf omni_${OS}_${ARCH}.tar.gz -C /tmp

# Assuming the omni binary is at the root of the tarball structure
# Adjust this path if it's nested in directories
OMNI_BINARY="/tmp/omni"

if [ -f "$OMNI_BINARY" ]; then
    echo "Installing omni to /usr/local/bin"
    sudo mv "$OMNI_BINARY" /usr/local/bin/omni
    chmod +x /usr/local/bin/omni
    echo "omni installed successfully"
else
    echo "omni binary not found after extraction"
    exit 1
fi

# Clean up downloaded tarball
rm omni_${OS}_${ARCH}.tar.gz

# Verify installation
if command -v omni &> /dev/null; then
    echo "omni is installed. Run 'omni version' to verify."
else
    echo "Failed to install omni."
fi
