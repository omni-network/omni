# shared admin script utils

# get network json
# usage: network_json <network>
netjson() {
  net=$(go run ./contracts/scripts/netjson $@)

  if [ -z "$net" ]; then
    echo "Network not found: $@"
    exit 1
  fi

  echo $net
}

# check if a network json has a chain
# usage: hasChain <network json> <chain name>
hasChain() {
  jq -e '.chains[] | select(.name == "'$2'")' <<< $1 > /dev/null
}

# check if a the chain id in network json matches the remote chain id
# usage: matchesRemoteChainID <network json> <chain name> <remote rpc>
matchesRemoteChainID() {
  remote=$(cast chain-id --rpc-url $3)
  local=$(chainID "$1" "$2")

  if [ "$remote" != "$local" ]; then
    return 1
  fi
}

# get chain id from an network json
# usage: chainID <network json> <chain name>
chainID() {
  jq -r '.chains[] | select(.name == "'$2'") | .id' <<< $1
}

# get portal address for a chain from network json
# usage: portalAddr <network json> <chain name>
portalAddr() {
  jq -r '.chains[] | select(.name == "'$2'") | .portal_address' <<< $1
}

# check if required environment variables are set
# usage: require_env <var1> <var2> ...
require_env() {
  for var in $@; do
    if [ -z "${!var}" ]; then
      echo "Missing required environment variable: $var"
      return 1
    fi
  done
}
