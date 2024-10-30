package cmd

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"net/url"
	"strings"
	"time"

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
	Network        netconf.ID
	ExecutionRPC   string
	ConsensusRPC   string
	PrivateKeyFile string
}

func (v eoaConfig) privateKey() (*ecdsa.PrivateKey, error) {
	if v.PrivateKeyFile == "" {
		return nil, errors.New("required flag --private-key-file not set")
	}

	opPrivKey, err := crypto.LoadECDSA(v.PrivateKeyFile)
	if err != nil {
		return nil, errors.Wrap(err, "load private key")
	}

	return opPrivKey, nil
}

func (v eoaConfig) validate() error {
	if _, err := v.privateKey(); err != nil {
		return errors.Wrap(err, "verify --private-key-file flag")
	}

	if err := v.Network.Verify(); err != nil {
		return errors.Wrap(err, "verify --network flag")
	}

	if v.ExecutionRPC != "" {
		if _, err := url.Parse(v.ExecutionRPC); err != nil {
			return errors.Wrap(err, "verify --execution-rpc flag")
		}
	}

	if v.ConsensusRPC != "" {
		if _, err := url.Parse(v.ConsensusRPC); err != nil {
			return errors.Wrap(err, "verify --consensus-rpc flag")
		}
	}

	return nil
}

type createValConfig struct {
	eoaConfig
	ConsensusPubKeyHex string
	SelfDelegation     uint64
}

func (c createValConfig) consensusPublicKey() (*ecdsa.PublicKey, error) {
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

func (c createValConfig) validate() error {
	if err := c.eoaConfig.validate(); err != nil {
		return err
	}

	if _, err := c.consensusPublicKey(); err != nil {
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

	link := fmt.Sprintf("https://%s.omniscan.network/tx/%s", cfg.Network, tx.Hash().Hex())
	log.Info(ctx, "🎉 Create-validator transaction sent and included on-chain", "link", link, "block", rec.BlockNumber.Uint64())

	return nil
}

type delegateConfig struct {
	eoaConfig
	Amount uint64
	Self   bool
}

func (d delegateConfig) validate() error {
	if !d.Self {
		return errors.New("required --self", "required_value", "true")
	}

	if d.Amount < minDelegation {
		return errors.New("insufficient --amount", "minimum", minDelegation, "amount", d.Amount)
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
				return errors.Wrap(err, "delegate")
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
	} else if bal <= float64(cfg.Amount) {
		return &CliError{
			Msg:     fmt.Sprintf("Delegator address has insufficient balance=%.2f OMNI, address=%s", bal, delegatorAddr),
			Suggest: "Fund the delegator address with sufficient OMNI for self-delegation and gas",
		}
	}

	txOpts, err := backend.BindOpts(ctx, delegatorAddr)
	if err != nil {
		return err
	}
	txOpts.Value = new(big.Int).Mul(umath.NewBigInt(cfg.Amount), big.NewInt(params.Ether)) // Send self-delegation

	tx, err := contract.Delegate(txOpts, delegatorAddr)
	if err != nil {
		return errors.Wrap(err, "create validator")
	}

	rec, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	}

	link := fmt.Sprintf("https://%s.omniscan.network/tx/%s", cfg.Network, tx.Hash().Hex())
	log.Info(ctx, "🎉 Delegate transaction sent and included on-chain", "link", link, "block", rec.BlockNumber.Uint64())

	return nil
}

func newUnjailCmd() *cobra.Command {
	var cfg eoaConfig

	cmd := &cobra.Command{
		Use:   "unjail",
		Short: "Unjail a validator",
		Long: "Sign and broadcast a unjail transaction that unjails a jailed validator. " +
			"This transaction must be sent by the operator address and costs 0.1 OMNI.",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := cfg.validate(); err != nil {
				return errors.Wrap(err, "verify flags")
			}

			err := unjailValidator(cmd.Context(), cfg)
			if err != nil {
				return errors.Wrap(err, "unjail")
			}

			return nil
		},
	}

	bindEOAConfig(cmd, &cfg)

	return cmd
}

func unjailValidator(ctx context.Context, cfg eoaConfig) error {
	opPrivKey, err := crypto.LoadECDSA(cfg.PrivateKeyFile)
	if err != nil {
		return errors.Wrap(err, "load private key")
	}
	opAddr := crypto.PubkeyToAddress(opPrivKey.PublicKey)

	_, cprov, backend, err := setupClients(cfg, opPrivKey)
	if err != nil {
		return err
	}

	validator, ok, err := cprov.SDKValidator(ctx, opAddr)
	if err != nil {
		return err
	} else if !ok {
		return &CliError{
			Msg:     "Operator address not a validator: " + opAddr.Hex(),
			Suggest: "Ensure operator address is a validator",
		}
	} else if !validator.IsJailed() {
		return &CliError{
			Msg:     "Validator not jailed: " + opAddr.Hex(),
			Suggest: "Ensure validator is jailed before unjailing",
		}
	}

	if err := checkUnjailPeriod(ctx, cprov, validator); err != nil {
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

// checkUnjailPeriod returns an error if the validator is still in the unjail period.
func checkUnjailPeriod(ctx context.Context, cprov cchain.Provider, val cchain.SDKValidator) error {
	consCmtAddr, err := val.ConsensusCmtAddr()
	if err != nil {
		return err
	}
	// Ensure the validator can be unjailed
	infos, err := cprov.SDKSigningInfos(ctx)
	if err != nil {
		return err
	}
	var found bool
	for _, info := range infos {
		if addr, err := info.ConsensusCmtAddr(); err != nil {
			return err
		} else if !bytes.Equal(addr, consCmtAddr) {
			continue
		}
		found = true
		if info.JailedUntil.After(time.Now()) {
			return &CliError{
				Msg:     "Validator cannot be unjailed yet",
				Suggest: "Retry after unjail period ends at " + info.JailedUntil.String(),
			}
		}

		break
	}
	if !found {
		return errors.New("missing signing info for validator [BUG]")
	}

	return nil
}

// setupClients is a helper that creates the omni evm client,
// omni consensus client and a backend set with the operator private key.
func setupClients(
	conf eoaConfig,
	operatorPriv *ecdsa.PrivateKey,
) (ethclient.Client, cchain.Provider, *ethbackend.Backend, error) {
	static := conf.Network.Static()
	chainID := static.OmniExecutionChainID

	chainMeta, ok := evmchain.MetadataByID(chainID)
	if !ok {
		return nil, nil, nil, errors.New("chain metadata not found")
	}

	if conf.ExecutionRPC == "" {
		conf.ExecutionRPC = static.ExecutionRPC()
	}

	if conf.ConsensusRPC == "" {
		conf.ConsensusRPC = static.ConsensusRPC()
	}

	cl, err := http.New(conf.ConsensusRPC, "/websocket")
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "new tendermint client")
	}

	cprov := provider.NewABCI(cl, conf.Network)

	eth, err := ethclient.Dial(chainMeta.Name, conf.ExecutionRPC)
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
