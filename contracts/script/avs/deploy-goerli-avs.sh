#!/usr/bin/env bash

confirm() {
    local message=$1
    echo "$message (y/n)"
    read input

    if [ "$input" = "y" ]; then
      return 0
    else
      return 1
    fi
}

RPC_URL=$1

echo "Deploying to Goerli with RPC URL: $RPC_URL"
confirm "Do you want to continue?" || exit 1

if confirm "Deploy Create3 ?"; then
  echo "Deploying..."
  forge script DeployCreate3Factory --rpc-url $RPC_URL --broadcast
  echo "Done."
fi


if confirm "Deploy ProxyAdmin?"; then
  echo "Deploying..."
  forge script DeployProxyAdmin --rpc-url $RPC_URL --broadcast
  echo "Done."
fi

if confirm "Do you want to deploy the OmniAVS?"; then
  echo "Deploying..."
  forge script DeployGoerliAVS --rpc-url $RPC_URL --broadcast
  echo "Done."
else
  exit 1
fi
