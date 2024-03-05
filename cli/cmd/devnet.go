package cmd

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/txmgr"

	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"

	"github.com/spf13/cobra"
)

const (
	// privKeyHex0 of pre-funded anvil account 0.
	privKeyHex0 = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
)

func newDevnetCmds() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "devnet",
		Short: "Local devnet commands",
		Args:  cobra.NoArgs,
	}

	cmd.AddCommand(
		newDevnetFundCmd(),
		newDevnetAVSAllow(),
	)

	return cmd
}

func newDevnetFundCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "fund",
		Short:   "Fund a local devnet account with 1 ETH",
		Example: "  omni devnet fund <address> <rpc_url>",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return fund(cmd.Context(), args[0], args[1])
		},
	}

	return cmd
}

func newDevnetAVSAllow() *cobra.Command {
	var omniAVSAddress string

	cmd := &cobra.Command{
		Use:     "avs-allow",
		Short:   "Add an operator to the omni AVS allow list",
		Example: "  omni devnet avs-allow <operator-address> <rpc-url> [avs-address]",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return devnetAllow(cmd.Context(), args[0], args[1], omniAVSAddress)
		},
	}

	cmd.Flags().StringVar(&omniAVSAddress, "omni-avs-address", omniAVSAddress, "Optional address of the Omni AVS contract.")

	return cmd
}

func devnetAllow(ctx context.Context, operatorAddr string, rpcURL string, avsAddr string) error {
	if !common.IsHexAddress(operatorAddr) {
		return errors.New("invalid operator address", "address", operatorAddr)
	}

	avsOwner, backend, err := devnetBackend(ctx, rpcURL)
	if err != nil {
		return err
	}

	chainID, err := backend.ChainID(ctx)
	if err != nil {
		return errors.Wrap(err, "get chain id")
	}

	avsAddress, err := avsAddressOrDefault(avsAddr, chainID)
	if err != nil {
		return err
	}

	omniAVS, err := bindings.NewOmniAVS(avsAddress, backend)
	if err != nil {
		return errors.Wrap(err, "omni avs")
	}

	txOpts, err := backend.BindOpts(ctx, avsOwner)
	if err != nil {
		return err
	}

	tx, err := omniAVS.AddToAllowlist(txOpts, common.HexToAddress(operatorAddr))
	if err != nil {
		return errors.Wrap(err, "add to allowlist")
	}

	if _, err := backend.WaitMined(ctx, tx); err != nil {
		return errors.Wrap(err, "wait mined")
	}

	return nil
}

func fund(ctx context.Context, address string, rpcURL string) error {
	if !common.IsHexAddress(address) {
		return errors.New("invalid ETH address", "address", address)
	}

	funder, backend, err := devnetBackend(ctx, rpcURL)
	if err != nil {
		return err
	}

	addr := common.HexToAddress(address)
	tx, _, err := backend.Send(ctx, funder, txmgr.TxCandidate{
		To:       &addr,
		GasLimit: 100_000,
		Value:    big.NewInt(params.Ether),
	})
	if err != nil {
		return errors.Wrap(err, "send tx")
	} else if _, err := backend.WaitMined(ctx, tx); err != nil {
		return errors.Wrap(err, "wait mined")
	}

	b, err := backend.BalanceAt(ctx, addr, nil)
	if err != nil {
		return errors.Wrap(err, "get balance")
	}

	bf, _ := b.Float64()
	bf /= params.Ether

	log.Info(ctx, "Account funded", "address", address, "balance", fmt.Sprintf("%.4f ETH", bf))

	return nil
}

// devnetBackend returns a backend populated with the default anvil account 0 private key.
func devnetBackend(ctx context.Context, rpcURL string) (common.Address, *ethbackend.Backend, error) {
	ethCl, err := ethclient.Dial("", rpcURL)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "dial eth client", "url", rpcURL)
	}

	chainID, err := ethCl.ChainID(ctx)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "get chain id")
	}

	funderPrivKey, err := ethcrypto.HexToECDSA(strings.TrimPrefix(privKeyHex0, "0x"))
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "parse private key")
	}

	backend, err := ethbackend.NewBackend("", chainID.Uint64(), time.Second, ethCl)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "create backend")
	}

	funderAddr, err := backend.AddAccount(funderPrivKey)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "add account")
	}

	return funderAddr, backend, nil
}
