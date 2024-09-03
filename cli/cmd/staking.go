package cmd

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/umath"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"

	"github.com/spf13/cobra"
)

const minSelfDelegation = uint64(100)

func newCreateValCmd() *cobra.Command {
	var cfg createValConfig

	cmd := &cobra.Command{
		Use:   "create-validator",
		Short: "Create new validator initialized with a self-delegation",
		Long:  `Sign and broadcast a create-validator transaction that registers a new validator on the omni consensus chain initialized with a self-delegation`,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := cfg.Verify(); err != nil {
				return errors.Wrap(err, "verify flags")
			}

			err := createValidator(cmd.Context(), cfg)
			if err != nil {
				return errors.Wrap(err, "create-validator")
			}

			return nil
		},
	}

	bindCreateValConfig(cmd, &cfg)

	return cmd
}

type createValConfig struct {
	Network            netconf.ID
	PrivateKeyFile     string
	ConsensusPubKeyHex string
	SelfDelegation     uint64
}

func (c createValConfig) ConsensusPubKey() (*ecdsa.PublicKey, error) {
	if strings.HasPrefix(c.ConsensusPubKeyHex, "0x") {
		return nil, errors.New("consensus pubkey hex should not have 0x prefix")
	}

	bz, err := hex.DecodeString(c.ConsensusPubKeyHex)
	if err != nil {
		return nil, errors.Wrap(err, "decode consensus pubkey hex")
	}

	resp, err := crypto.DecompressPubkey(bz)
	if err != nil {
		return nil, errors.Wrap(err, "decompress consensus pubkey")
	}

	return resp, nil
}

func (c createValConfig) Verify() error {
	if c.PrivateKeyFile == "" {
		return errors.New("required flag --private-key-file not set")
	}

	if err := c.Network.Verify(); err != nil {
		return errors.Wrap(err, "verify --network flag")
	}

	if _, err := c.ConsensusPubKey(); err != nil {
		return errors.Wrap(err, "verify --consensus-pubkey-hex flag")
	}

	if c.SelfDelegation < minSelfDelegation {
		return errors.New("insufficient --self-delegation", "minimum", minSelfDelegation, "self_delegation", c.SelfDelegation)
	}

	if c.SelfDelegation > 1e3*minSelfDelegation {
		return errors.New("excessive --self-delegation", "maximum", 1e3*minSelfDelegation, "self_delegation", c.SelfDelegation)
	}

	return nil
}

func createValidator(ctx context.Context, cfg createValConfig) error {
	opPrivKey, err := crypto.LoadECDSA(cfg.PrivateKeyFile)
	if err != nil {
		return errors.Wrap(err, "load private key")
	}
	opAddr := crypto.PubkeyToAddress(opPrivKey.PublicKey)

	chainID := cfg.Network.Static().OmniExecutionChainID
	chainMeta, ok := evmchain.MetadataByID(chainID)
	if !ok {
		return errors.New("chain metadata not found")
	}

	ethCl, err := ethclient.Dial(chainMeta.Name, cfg.Network.Static().ExecutionRPC())
	if err != nil {
		return err
	}

	cprov, err := provider.Dial(cfg.Network)
	if err != nil {
		return err
	}

	if _, ok, err = cprov.Validator(ctx, opAddr); err != nil {
		return err
	} else if ok {
		return &CliError{
			Msg:     "Operator address already a validator: " + opAddr.Hex(),
			Suggest: "Ensure correct operator address",
		}
	}

	backend, err := ethbackend.NewBackend(chainMeta.Name, chainID, chainMeta.BlockPeriod, ethCl, opPrivKey)
	if err != nil {
		return err
	}

	contract, err := bindings.NewStaking(common.HexToAddress(predeploys.Staking), backend)
	if err != nil {
		return err
	}

	if ok, err := contract.IsAllowedValidator(nil, opAddr); err != nil {
		return err
	} else if !ok {
		return &CliError{
			Msg:     "Operator address not allowed to create validator: " + opAddr.Hex(),
			Suggest: "Contact Omni team to be included in validator allow list",
		}
	}

	bal, err := ethCl.EtherBalanceAt(ctx, opAddr)
	if err != nil {
		return err
	} else if bal <= float64(cfg.SelfDelegation) {
		return &CliError{
			Msg:     fmt.Sprintf("Operator address has insufficient balance=%.2f OMNI, address=%s", bal, opAddr),
			Suggest: "Fund the operator address with sufficient OMNI for self-delegation and gas",
		}
	}

	txOpts, err := backend.BindOpts(ctx, opAddr)
	if err != nil {
		return err
	}
	txOpts.Value = new(big.Int).Mul(umath.NewBigInt(cfg.SelfDelegation), big.NewInt(params.Ether)) // Send self-delegation

	consPubkey, err := cfg.ConsensusPubKey()
	if err != nil {
		return err
	}

	tx, err := contract.CreateValidator(txOpts, crypto.CompressPubkey(consPubkey))
	if err != nil {
		return errors.Wrap(err, "create validator")
	}

	rec, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	}

	link := fmt.Sprintf("https://%s.omniscan.network/tx/%s", cfg.Network, tx.Hash().Hex())
	log.Info(ctx, "🎉 Create-validator transaction sent and included on-chain", "link", link, "block", rec.BlockNumber.Uint64())

	return nil
}

func newUnjailCmd() *cobra.Command {
	var cfg unjailConfig

	cmd := &cobra.Command{
		Use:   "unjail",
		Short: "Unjail a validator",
		Long: "Sign and broadcast a unjail transaction that unjails a jailed validator. " +
			"This transaction must be sent by the operator address and costs 0.1 OMNI.",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := cfg.Verify(); err != nil {
				return errors.Wrap(err, "verify flags")
			}

			err := unjailValidator(cmd.Context(), cfg)
			if err != nil {
				return errors.Wrap(err, "unjail")
			}

			return nil
		},
	}

	bindUnjailConfig(cmd, &cfg)

	return cmd
}

type unjailConfig struct {
	Network        netconf.ID
	PrivateKeyFile string
}

func (c unjailConfig) Verify() error {
	if c.PrivateKeyFile == "" {
		return errors.New("required flag --private-key-file not set")
	}

	if err := c.Network.Verify(); err != nil {
		return errors.Wrap(err, "verify --network flag")
	}

	return nil
}

func unjailValidator(ctx context.Context, cfg unjailConfig) error {
	opPrivKey, err := crypto.LoadECDSA(cfg.PrivateKeyFile)
	if err != nil {
		return errors.Wrap(err, "load private key")
	}
	opAddr := crypto.PubkeyToAddress(opPrivKey.PublicKey)

	chainID := cfg.Network.Static().OmniExecutionChainID
	chainMeta, ok := evmchain.MetadataByID(chainID)
	if !ok {
		return errors.New("chain metadata not found")
	}

	ethCl, err := ethclient.Dial(chainMeta.Name, cfg.Network.Static().ExecutionRPC())
	if err != nil {
		return err
	}

	cprov, err := provider.Dial(cfg.Network)
	if err != nil {
		return err
	}

	if val, ok, err := cprov.Validator(ctx, opAddr); err != nil {
		return err
	} else if !ok {
		return &CliError{
			Msg:     "Operator address not a validator: " + opAddr.Hex(),
			Suggest: "Ensure operator address is a validator",
		}
	} else if !val.IsJailed() {
		return &CliError{
			Msg:     "Validator not jailed: " + opAddr.Hex(),
			Suggest: "Ensure validator is jailed before unjailing",
		}
	}

	backend, err := ethbackend.NewBackend(chainMeta.Name, chainID, chainMeta.BlockPeriod, ethCl, opPrivKey)
	if err != nil {
		return err
	}

	contract, err := bindings.NewSlashing(common.HexToAddress(predeploys.Slashing), backend)
	if err != nil {
		return err
	}

	fee, err := contract.Fee(&bind.CallOpts{Context: ctx})
	if err != nil {
		return err
	}

	txOpts, err := backend.BindOpts(ctx, opAddr)
	if err != nil {
		return err
	}

	txOpts.Value = fee
	tx, err := contract.Unjail(txOpts)
	if err != nil {
		return errors.Wrap(err, "unjail validator")
	}

	rec, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	}

	link := fmt.Sprintf("https://%s.omniscan.network/tx/%s", cfg.Network, rec.TxHash.Hex())
	log.Info(ctx, "🎉 Unjail transaction sent and included on-chain", "link", link, "block", rec.BlockNumber.Uint64())

	return nil
}
