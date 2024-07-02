# This is a TOML config file.
# For more information, see https://github.com/toml-lang/toml

# The version of the Halo binary that created or
# last modified the config file. Do not modify this.
version = "{{ .Version}}"

# Omni network to participate in: mainnet, testnet, or devnet.
network = "{{ .Network }}"

#######################################################################
###                         Monitor Options                         ###
#######################################################################

# Path to the ethereum private key used to sign avs omni sync transactions.
private-key = "{{ .PrivateKey }}"

# The address that the monitor listens for metric scrape requests.
monitoring-addr = "{{ .MonitoringAddr }}"

# The URL of the halo node to connect to.
halo-url = "{{ .HaloURL }}"

#######################################################################
###                             X-Chain                             ###
#######################################################################

[xchain]

# Cross-chain EVM RPC endpoints to use for voting; only required for validators. One per supported EVM is required.
# It is strongly advised to operate fullnodes for each chain and NOT to use free public RPCs.
[xchain.evm-rpc-endpoints]
{{- if not .RPCEndpoints }}
# ethereum = "http://my-ethreum-node:8545"
# optimism = "https://my-op-node.com"
{{ end -}}
{{- range $key, $value := .RPCEndpoints }}
{{ $key }} = "{{ $value }}"
{{ end }}

#######################################################################
###                         Logging Options                         ###
#######################################################################

[log]
# Logging level. Note cometBFT internal logs are configured in config.yaml.
# Options are: debug, info, warn, error.
level = "{{ .Log.Level }}"

# Logging format. Options are: console, json.
format = "{{ .Log.Format }}"

# Logging color if console format is chosen. Options are: auto, force, disable.
color = "{{ .Log.Color }}"

#######################################################################
###                         Load Generation                         ###
#######################################################################

# Note that load generation is only used for testing purposes; ie on devent or staging.
[loadgen]
# Validator keys glob defines the validator keys to use for self-delegation.
validator-keys-glob = "{{ .LoadGen.ValidatorKeysGlob }}"
