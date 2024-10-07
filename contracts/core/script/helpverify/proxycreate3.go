package main

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/spf13/cobra"
)

func newParseProxyCreate3TxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "parse-proxy-create3-tx",
		Short: "Parse proxy create3 tx",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			chain := ChainName(args[0])
			if err := chain.Validate(); err != nil {
				return err
			}

			txHash := common.HexToHash(args[1])

			res, err := parseProxyCreate3Tx(ctx, chain, txHash)
			if err != nil {
				return err
			}

			log.Info(ctx, "✅", "implementation", res.Implementation.Hex(), "constructor_args", res.ConstructorArgs)

			return nil
		},
	}

	return cmd
}

type ProxyCreate3Tx struct {
	Implementation  common.Address
	ConstructorArgs string
}

func parseProxyCreate3Tx(ctx context.Context, chain ChainName, txHash common.Hash) (ProxyCreate3Tx, error) {
	client, err := ethclient.Dial(string(chain), chain.RPCURL())
	if err != nil {
		return ProxyCreate3Tx{}, errors.Wrap(err, "dial", "chain", chain)
	}

	tx, _, err := client.TransactionByHash(ctx, txHash)
	if err != nil {
		return ProxyCreate3Tx{}, errors.Wrap(err, "get transaction", "hash", txHash.Hex())
	}

	txData := tx.Data()

	// first 4 bytes are signature
	deployArgsI, err := create3ABI.Methods["deploy"].Inputs.Unpack(txData[4:])
	if err != nil {
		return ProxyCreate3Tx{}, errors.Wrap(err, "unpack deploy args", "chain", chain, "tx", txHash.Hex())
	}

	creationCode, ok := deployArgsI[1].([]byte)
	if !ok {
		return ProxyCreate3Tx{}, errors.New("cast creation code")
	}

	constructorArgs := creationCode[len(mustDecodeHex(bindings.TransparentUpgradeableProxyBin)):]

	// implementation is first 20 byte word of constructor args
	// TODO(kevin): EthAddress length != 32
	impl, err := cast.EthAddress(constructorArgs[:20])
	if err != nil {
		return ProxyCreate3Tx{}, err
	}

	return ProxyCreate3Tx{
		Implementation:  impl,
		ConstructorArgs: hexutil.Encode(constructorArgs),
	}, nil
}

func mustDecodeHex(hex string) []byte {
	data, err := hexutil.Decode(hex)
	if err != nil {
		panic(err)
	}

	return data
}
