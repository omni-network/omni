# This is a TOML config file.
# For more information, see https://github.com/toml-lang/toml

# The version of the Halo binary that created or
# last modified the config file. Do not modify this.
version = "{{ .Version}}"

# Omni network to participate in: mainnet, testnet, or devnet.
network = "{{ .Network }}"

#######################################################################
###                           Trade Options                         ###
#######################################################################

# The address to listen on for RPC requests.
rpc-listen = "{{ .RPCListen }}"

# Postgres connection string to use for the database.
db-conn = "{{ .DBConn }}"

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
