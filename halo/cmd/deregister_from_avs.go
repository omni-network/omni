package cmd

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/spf13/cobra"
)

func DeRegisterOperatorFromOmniAVS(cfg *OperatorConfig) *cobra.Command {
	deregisterFromAVSCmd := &cobra.Command{
		Use:   "deregister",
		Short: "deregister validator from omni avs",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg.HaloConfig.HomeDir = cfg.HomeDir

			// read the comet config based on the home directory
			cometCfg, err := parseCometConfig(cmd.Context(), cfg.HaloConfig.HomeDir)
			if err != nil {
				return err
			}
			cfg.CometConfig = cometCfg

			return deregister(cmd.Context(), cfg)
		},
	}

	bindOperatorFlags(deregisterFromAVSCmd.Flags(), cfg)

	return deregisterFromAVSCmd
}

func deregister(ctx context.Context, cfg *OperatorConfig) error {
	privVal, client, chain, err := loadKeysAndChain(ctx, cfg)
	if err != nil {
		return err
	}

	err = validateContractAddresses(ctx, cfg, client)
	if err != nil {
		return err
	}

	omniAvs, err := bindings.NewOmniAVS(common.HexToAddress(cfg.OmniAVSAddr), client)
	if err != nil {
		return err
	}

	operPK, err := crypto.ToECDSA(privVal.Key.PrivKey.Bytes())
	if err != nil {
		return errors.Wrap(err, "could not convert pk to ecdsa")
	}
	pubKey, err := privVal.GetPubKey()
	if err != nil {
		return errors.Wrap(err, "get pubkey")
	}
	operAddr, err := k1util.PubKeyToAddress(pubKey)
	if err != nil {
		return errors.Wrap(err, "could not convert to ethereum address")
	}

	txOpts, err := bind.NewKeyedTransactorWithChainID(operPK, big.NewInt(int64(chain.ID)))
	if err != nil {
		return errors.Wrap(err, "error getting txopts")
	}
	txOpts.Context = ctx
	tx, err := omniAvs.DeregisterOperatorFromAVS(txOpts, operAddr)
	if err != nil {
		return err
	}
	log.Info(ctx, "Submitted de-registration to AVS")

	rcpt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		return errors.Wrap(err, "error waiting for mining tx")
	}
	log.Info(ctx, "Operator de-registered with AVS", "address", operAddr.String(), "txHash", rcpt.TxHash)

	return nil
}
