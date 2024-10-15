package cmd

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"net/url"
	"strings"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/umath"

	"github.com/cometbft/cometbft/rpc/client/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"

	"github.com/spf13/cobra"
)

const minSelfDelegation = uint64(100)
const minDelegation = uint64(1)

func newCreateValCmd() *cobra.Command {
	var cfg createValConfig

	cmd := &cobra.Command{
		Use:   "create-validator",
		Short: "Create new validator initialized with a self-delegation",
		Long:  `Sign and broadcast a create-validator transaction that registers a new validator on the omni consensus chain initialized with a self-delegation`,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := cfg.validate(); err != nil {
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

// eoaConfig defines the required data to sign and submit evm transactions.
type eoaConfig struct {
	network        netconf.ID
	executionRPC   string
	consensusRPC   string
	privateKeyFile string
}

func (v eoaConfig) privateKey() (*ecdsa.PrivateKey, error) {
	if v.privateKeyFile == "" {
		return nil, errors.New("required flag --private-key-file not set")
	}

	opPrivKey, err := crypto.LoadECDSA(v.privateKeyFile)
	if err != nil {
		return nil, errors.Wrap(err, "load private key")
	}

	return opPrivKey, nil
}

func (v eoaConfig) validate() error {
	if _, err := v.privateKey(); err != nil {
		return errors.Wrap(err, "verify --private-key-file flag")
	}

	if err := v.network.Verify(); err != nil {
		return errors.Wrap(err, "verify --network flag")
	}

	if v.executionRPC != "" {
		if _, err := url.Parse(v.executionRPC); err != nil {
			return errors.Wrap(err, "verify --execution-rpc flag")
		}
	}

	if v.consensusRPC != "" {
		if _, err := url.Parse(v.consensusRPC); err != nil {
			return errors.Wrap(err, "verify --consensus-rpc flag")
		}
	}

	return nil
}

type createValConfig struct {
	eoaConfig
	consensusPubKeyHex string
	selfDelegation     uint64
}

func (c createValConfig) consensusPublicKey() (*ecdsa.PublicKey, error) {
	if strings.HasPrefix(c.consensusPubKeyHex, "0x") {
		return nil, errors.New("consensus pubkey hex should not have 0x prefix")
	}

	bz, err := hex.DecodeString(c.consensusPubKeyHex)
	if err != nil {
		return nil, errors.Wrap(err, "decode consensus pubkey hex")
	}

	resp, err := crypto.DecompressPubkey(bz)
	if err != nil {
		return nil, errors.Wrap(err, "decompress consensus pubkey")
	}

	return resp, nil
}

func (c createValConfig) validate() error {
	if err := c.eoaConfig.validate(); err != nil {
		return err
	}

	if _, err := c.consensusPublicKey(); err != nil {
		return errors.Wrap(err, "verify --consensus-pubkey-hex flag")
	}

	if c.selfDelegation < minSelfDelegation {
		return errors.New("insufficient --self-delegation", "minimum", minSelfDelegation, "self_delegation", c.selfDelegation)
	}

	if c.selfDelegation > 1e3*minSelfDelegation {
		return errors.New("excessive --self-delegation", "maximum", 1e3*minSelfDelegation, "self_delegation", c.selfDelegation)
	}

	return nil
}

func createValidator(ctx context.Context, cfg createValConfig) error {
	operatorPriv, err := cfg.privateKey()
	if err != nil {
		return err
	}
	opAddr := crypto.PubkeyToAddress(operatorPriv.PublicKey)

	eth, cprov, backend, err := setupClients(cfg.eoaConfig, operatorPriv)
	if err != nil {
		return err
	}

	// check if we already have an existing validator
	if _, ok, err := cprov.SDKValidator(ctx, opAddr); err != nil {
		return err
	} else if ok {
		return &CliError{
			Msg:     "Operator address already a validator: " + opAddr.Hex(),
			Suggest: "Ensure correct operator address",
		}
	}

	contract, err := bindings.NewStaking(common.HexToAddress(predeploys.Staking), backend)
	if err != nil {
		return err
	}

	// only check if validator is on allow list if the allow list is enabled
	bindOpts := &bind.CallOpts{Context: ctx}
	if enabled, err := contract.IsAllowlistEnabled(bindOpts); err != nil {
		return err
	} else if enabled {
		if ok, err := contract.IsAllowedValidator(bindOpts, opAddr); err != nil {
			return err
		} else if !ok {
			return &CliError{
				Msg:     "Operator address not allowed to create validator: " + opAddr.Hex(),
				Suggest: "Contact Omni team to be included in validator allow list",
			}
		}
	}

	bal, err := eth.EtherBalanceAt(ctx, opAddr)
	if err != nil {
		return err
	} else if bal <= float64(cfg.selfDelegation) {
		return &CliError{
			Msg:     fmt.Sprintf("Operator address has insufficient balance=%.2f OMNI, address=%s", bal, opAddr),
			Suggest: "Fund the operator address with sufficient OMNI for self-delegation and gas",
		}
	}

	txOpts, err := backend.BindOpts(ctx, opAddr)
	if err != nil {
		return err
	}
	txOpts.Value = new(big.Int).Mul(umath.NewBigInt(cfg.selfDelegation), big.NewInt(params.Ether)) // Send self-delegation
	consPubkey, err := cfg.consensusPublicKey()
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

	link := fmt.Sprintf("https://%s.omniscan.network/tx/%s", cfg.network, tx.Hash().Hex())
	log.Info(ctx, "🎉 Create-validator transaction sent and included on-chain", "link", link, "block", rec.BlockNumber.Uint64())

	return nil
}

type delegateConfig struct {
	eoaConfig
	amount uint64
	self   bool
}

func (d delegateConfig) validate() error {
	if d.amount < minDelegation {
		return errors.New("insufficient --amount", "minimum", minDelegation, "amount", d.amount)
	}

	return d.eoaConfig.validate()
}

func newDelegateCmd() *cobra.Command {
	var cfg delegateConfig

	cmd := &cobra.Command{
		Use:   "delegate",
		Short: "Delegate Omni tokens to a validator",
		Long:  `Delegate an amount of Omni tokens to a validator from your wallet. Only self-delegation by validators supported at the moment.`,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := cfg.validate(); err != nil {
				return errors.Wrap(err, "verify flags")
			}

			err := delegate(cmd.Context(), cfg)
			if err != nil {
				return errors.Wrap(err, "create-validator")
			}

			return nil
		},
	}

	bindDelegateConfig(cmd, &cfg)

	return cmd
}

func delegate(ctx context.Context, cfg delegateConfig) error {
	delegatorPriv, err := cfg.privateKey()
	if err != nil {
		return err
	}
	delegatorAddr := crypto.PubkeyToAddress(delegatorPriv.PublicKey)

	eth, cprov, backend, err := setupClients(cfg.eoaConfig, delegatorPriv)
	if err != nil {
		return err
	}

	// check if we already have an existing validator
	if _, ok, err := cprov.SDKValidator(ctx, delegatorAddr); err != nil {
		return err
	} else if !ok {
		return &CliError{
			Msg:     "Operator address is not a validator: " + delegatorAddr.Hex(),
			Suggest: "Ensure operator is already created as validator, see create-validator command",
		}
	}

	contract, err := bindings.NewStaking(common.HexToAddress(predeploys.Staking), backend)
	if err != nil {
		return err
	}

	bal, err := eth.EtherBalanceAt(ctx, delegatorAddr)
	if err != nil {
		return err
	} else if bal <= float64(cfg.amount) {
		return &CliError{
			Msg:     fmt.Sprintf("Delegator address has insufficient balance=%.2f OMNI, address=%s", bal, delegatorAddr),
			Suggest: "Fund the operator address with sufficient OMNI for self-delegation and gas",
		}
	}

	txOpts, err := backend.BindOpts(ctx, delegatorAddr)
	if err != nil {
		return err
	}
	txOpts.Value = new(big.Int).Mul(umath.NewBigInt(cfg.amount), big.NewInt(params.Ether)) // Send self-delegation

	tx, err := contract.Delegate(txOpts, delegatorAddr)
	if err != nil {
		return errors.Wrap(err, "create validator")
	}

	rec, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	}

	link := fmt.Sprintf("https://%s.omniscan.network/tx/%s", cfg.network, tx.Hash().Hex())
	log.Info(ctx, "🎉 Delegate transaction sent and included on-chain", "link", link, "block", rec.BlockNumber.Uint64())

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

	// todo reuse setupClient

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

	if val, ok, err := cprov.SDKValidator(ctx, opAddr); err != nil {
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

// setupClients is a helper that creates the omni evm client,
// omni consensus client and a backend set with the operator private key.
func setupClients(
	conf eoaConfig,
	operatorPriv *ecdsa.PrivateKey,
) (ethclient.Client, cchain.Provider, *ethbackend.Backend, error) {
	static := conf.network.Static()
	chainID := static.OmniExecutionChainID

	chainMeta, ok := evmchain.MetadataByID(chainID)
	if !ok {
		return nil, nil, nil, errors.New("chain metadata not found")
	}

	if conf.executionRPC == "" {
		conf.executionRPC = static.ExecutionRPC()
	}

	if conf.consensusRPC == "" {
		conf.consensusRPC = static.ConsensusRPC()
	}

	cl, err := http.New(conf.consensusRPC, "/websocket")
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "new tendermint client")
	}

	cprov := provider.NewABCIProvider(cl, conf.network, netconf.ChainVersionNamer(conf.network))

	eth, err := ethclient.Dial(chainMeta.Name, conf.executionRPC)
	if err != nil {
		return nil, nil, nil, err
	}

	backend, err := ethbackend.NewBackend(
		chainMeta.Name,
		chainID,
		chainMeta.BlockPeriod,
		eth,
		operatorPriv,
	)
	if err != nil {
		return nil, nil, nil, err
	}

	return eth, cprov, backend, err
}
