package xchain

import (
	"strconv"

	"github.com/omni-network/omni/lib/errors"

	"github.com/spf13/pflag"
)

type RPCEndpoints map[string]string

func (e RPCEndpoints) ByNameOrID(name string, chainID uint64) (string, error) {
	if val, ok := e[name]; ok {
		return val, nil
	} else if val, ok := e[strconv.FormatUint(chainID, 10)]; ok {
		return val, nil
	}

	return "", errors.New("no rpc endpoint for chain", "chain_name", name, "chain_id", chainID)
}

func (e RPCEndpoints) Keys() []string {
	keys := make([]string, 0, len(e))
	for k := range e {
		keys = append(keys, k)
	}

	return keys
}

// BindFlags binds the xchain evm rpc flag.
func BindFlags(flags *pflag.FlagSet, endpoints *RPCEndpoints) {
	flags.StringToStringVar((*map[string]string)(endpoints), "xchain-evm-rpc-endpoints", *endpoints, "Cross-chain EVM RPC endpoints. e.g. \"ethereum=http://geth:8545,optimism=https://optimism.io\"")
}
