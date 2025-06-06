package cmd

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log/slog"
	"net/url"
	"strings"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	k1 "github.com/cometbft/cometbft/crypto/secp256k1"
	"github.com/cometbft/cometbft/rpc/client/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"

	"cosmossdk.io/math"
	"github.com/spf13/cobra"
)

const minSelfDelegation = uint64(100)
const minDelegation = uint64(1)

func newCreateValCmd() *cobra.Command {
	var cfg CreateValConfig

	cmd := &cobra.Command{
		Use:   "create-validator",
		Short: "Create new validator initialized with a self-delegation",
		Long:  `Sign and broadcast a create-validator transaction that registers a new validator on the omni consensus chain initialized with a self-delegation`,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := cfg.validate(); err != nil {
				return errors.Wrap(err, "verify flags")
			}

			err := CreateValidator(cmd.Context(), cfg)
			if err != nil {
				return errors.Wrap(err, "create-validator")
			}

			return nil
		},
	}

	bindCreateValConfig(cmd, &cfg)

	return cmd
}

// EOAConfig defines the required data to sign and submit evm transactions.
type EOAConfig struct {
	Network        netconf.ID
	ExecutionRPC   string
	ConsensusRPC   string
	PrivateKeyFile string
}

func (v EOAConfig) privateKey() (*ecdsa.PrivateKey, error) {
	if v.PrivateKeyFile == "" {
		return nil, errors.New("required flag --private-key-file not set")
	}

	opPrivKey, err := crypto.LoadECDSA(v.PrivateKeyFile)
	if err != nil {
		return nil, errors.Wrap(err, "load private key")
	}

	return opPrivKey, nil
}

func (v EOAConfig) validate() error {
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

type CreateValConfig struct {
	EOAConfig
	ConsensusPubKeyHex string
	SelfDelegation     uint64
}

func (c CreateValConfig) consensusPublicKey() (*ecdsa.PublicKey, error) {
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

func (c CreateValConfig) validate() error {
	if err := c.EOAConfig.validate(); err != nil {
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

func CreateValidator(ctx context.Context, cfg CreateValConfig) error {
	operatorPriv, err := cfg.privateKey()
	if err != nil {
		return err
	}
	opAddr := crypto.PubkeyToAddress(operatorPriv.PublicKey)

	eth, cprov, backend, err := setupClients(ctx, cfg.EOAConfig, operatorPriv)
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
	txOpts.Value = bi.Ether(cfg.SelfDelegation) // Send self-delegation
	consPubkey, err := cfg.consensusPublicKey()
	if err != nil {
		return err
	}

	digest, err := contract.GetConsPubkeyDigest(bindOpts, opAddr)
	if err != nil {
		return errors.Wrap(err, "get consensus pubkey digest")
	}

	pk := k1.PrivKey(crypto.FromECDSA(operatorPriv))
	sig, err := k1util.Sign(pk, digest)
	if err != nil {
		return err
	}

	tx, err := contract.CreateValidator(txOpts, crypto.CompressPubkey(consPubkey), sig[:])
	if err != nil {
		return errors.Wrap(err, "create validator")
	}

	rec, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	}

	log.Info(ctx, "🎉 Create-validator transaction sent and included on-chain",
		"link", cfg.Network.Static().OmniScanTXURL(tx.Hash()),
		"block", rec.BlockNumber.Uint64(),
	)

	log.Info(ctx, "⏳ Staking events are delayed in the Omni Consensus chain and may take up to 12h to apply.")

	return nil
}

type DelegateConfig struct {
	EOAConfig
	Amount           uint64
	Self             bool
	ValidatorAddress string
}

func (d DelegateConfig) validate() error {
	if !d.Self && d.ValidatorAddress == "" || d.Self && d.ValidatorAddress != "" {
		return errors.New("required either --self or --validator-address flags", "required_value", "true")
	}

	if d.ValidatorAddress != "" && !common.IsHexAddress(d.ValidatorAddress) {
		return errors.New("invalid --validator-address")
	}

	if d.Amount < minDelegation {
		return errors.New("insufficient --amount", "minimum", minDelegation, "amount", d.Amount)
	}

	return d.EOAConfig.validate()
}

func newDelegateCmd() *cobra.Command {
	var cfg DelegateConfig

	cmd := &cobra.Command{
		Use:   "delegate",
		Short: "Delegate Omni tokens to a validator",
		Long:  `Delegate an amount of Omni tokens to a validator from your wallet.`,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := cfg.validate(); err != nil {
				return errors.Wrap(err, "verify flags")
			}

			err := Delegate(cmd.Context(), cfg)
			if err != nil {
				return errors.Wrap(err, "delegate")
			}

			return nil
		},
	}

	bindDelegateConfig(cmd, &cfg)

	return cmd
}

func Delegate(ctx context.Context, cfg DelegateConfig) error {
	delegatorPriv, err := cfg.privateKey()
	if err != nil {
		return err
	}
	delegatorAddr := crypto.PubkeyToAddress(delegatorPriv.PublicKey)

	eth, cprov, backend, err := setupClients(ctx, cfg.EOAConfig, delegatorPriv)
	if err != nil {
		return err
	}

	validatorAddr := delegatorAddr
	if cfg.ValidatorAddress != "" {
		validatorAddr = common.HexToAddress(cfg.ValidatorAddress)
	}

	// check if we already have an existing validator
	if _, ok, err := cprov.SDKValidator(ctx, validatorAddr); err != nil {
		return err
	} else if !ok {
		return &CliError{
			Msg:     "Not a validator address: " + validatorAddr.Hex(),
			Suggest: "Ensure the validator exists (if you're operator, see create-validator command)",
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
			Suggest: "Fund the delegator address with sufficient OMNI for (self-)delegation and gas",
		}
	}

	txOpts, err := backend.BindOpts(ctx, delegatorAddr)
	if err != nil {
		return err
	}
	txOpts.Value = bi.Ether(cfg.Amount) // Send delegation

	callOpts := &bind.CallOpts{Context: ctx}
	ok, err := contract.IsAllowlistEnabled(callOpts)
	if err != nil {
		return errors.Wrap(err, "check allowlist enabled")
	} else if ok {
		ok, err := contract.IsAllowedValidator(&bind.CallOpts{Context: ctx}, validatorAddr)
		if err != nil {
			return err
		} else if !ok {
			return errors.New("active validator not in allowed list [BUG]")
		}
	}

	tx, err := contract.Delegate(txOpts, validatorAddr)
	if err != nil {
		return errors.Wrap(err, "delegate")
	}

	rec, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	}

	log.Info(ctx, "🎉 Delegate transaction sent and included on-chain",
		"link", cfg.Network.Static().OmniScanTXURL(tx.Hash()),
		"block", rec.BlockNumber.Uint64(),
	)

	log.Info(ctx, "⏳ Staking events are delayed in the Omni Consensus chain and may take up to 12h to apply.")

	return nil
}

func newUnjailCmd() *cobra.Command {
	var cfg EOAConfig

	cmd := &cobra.Command{
		Use:   "unjail",
		Short: "Unjail a validator",
		Long: "Sign and broadcast an unjail transaction that unjails a jailed validator. " +
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

func unjailValidator(ctx context.Context, cfg EOAConfig) error {
	opPrivKey, err := crypto.LoadECDSA(cfg.PrivateKeyFile)
	if err != nil {
		return errors.Wrap(err, "load private key")
	}
	opAddr := crypto.PubkeyToAddress(opPrivKey.PublicKey)

	_, cprov, backend, err := setupClients(ctx, cfg, opPrivKey)
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

	log.Info(ctx, "🎉 Unjail transaction sent and included on-chain",
		"link", cfg.Network.Static().OmniScanTXURL(tx.Hash()),
		"block", rec.BlockNumber.Uint64(),
	)

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

// doNotModify is a special CosmosSDK string that indicates the value should not be modified.
const doNotModify = "[do-not-modify]"

type EditValConfig struct {
	EOAConfig
	Moniker                  string
	Identity                 string
	Website                  string
	SecurityContact          string
	Details                  string
	CommissionRatePercentage int32
	MinSelfDelegationEther   int64
}

// NoopEditValConfig returns a default EditValConfig that will not modify existing values.
func NoopEditValConfig() EditValConfig {
	return EditValConfig{
		Moniker:                  doNotModify,
		Identity:                 doNotModify,
		Website:                  doNotModify,
		SecurityContact:          doNotModify,
		Details:                  doNotModify,
		CommissionRatePercentage: -1,
		MinSelfDelegationEther:   -1,
	}
}

func (d EditValConfig) modifiedAttrs() []any {
	var resp []any
	if d.Moniker != doNotModify {
		resp = append(resp, slog.String("moniker", d.Moniker))
	}
	if d.Identity != doNotModify {
		resp = append(resp, slog.String("identity", d.Identity))
	}
	if d.Website != doNotModify {
		resp = append(resp, slog.String("website", d.Website))
	}
	if d.SecurityContact != doNotModify {
		resp = append(resp, slog.String("security-contact", d.SecurityContact))
	}
	if d.Details != doNotModify {
		resp = append(resp, slog.String("details", d.Details))
	}
	if d.CommissionRatePercentage != -1 {
		resp = append(resp, slog.Any("commission-rate", d.CommissionRatePercentage))
	}
	if d.MinSelfDelegationEther != -1 {
		resp = append(resp, slog.Int64("min-self-delegation", d.MinSelfDelegationEther))
	}

	return resp
}

func (d EditValConfig) validate() error {
	if d.CommissionRatePercentage < -1 || d.CommissionRatePercentage > 100 {
		return errors.New("commission-rate not a valid percentage [0,100]", "maximum", 100, "minimum", 0)
	}
	if d.MinSelfDelegationEther == 0 {
		return errors.New("min-self-delegation in OMNI (not wei) must not be zero", "minimum", 1)
	}
	if d.MinSelfDelegationEther < -1 {
		return errors.New("min-self-delegation in OMNI (not wei) must not be negative")
	}
	if len(d.Moniker) > 70 {
		return errors.New("moniker too long", "maximum", 70, "length", len(d.Moniker))
	}
	if len(d.Identity) > 3000 {
		return errors.New("identity too long", "maximum", 3000, "length", len(d.Identity))
	}
	if len(d.Website) > 140 {
		return errors.New("website too long", "maximum", 140, "length", len(d.Website))
	}
	if len(d.SecurityContact) > 140 {
		return errors.New("security-contact too long", "maximum", 140, "length", len(d.SecurityContact))
	}
	if len(d.Details) > 280 {
		return errors.New("details too long", "maximum", 280, "length", len(d.Details))
	}

	return d.EOAConfig.validate()
}

func newEditValCmd() *cobra.Command {
	cfg := NoopEditValConfig()

	cmd := &cobra.Command{
		Use:   "edit-validator",
		Short: "Edit an existing validator metadata",
		Long:  `Sign and broadcast an edit-validator transaction that updates the metadata of an existing validator with 0.1 OMNI fee.`,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := cfg.validate(); err != nil {
				return errors.Wrap(err, "verify flags")
			}

			err := EditVal(cmd.Context(), cfg)
			if err != nil {
				return errors.Wrap(err, "edit validator")
			}

			return nil
		},
	}

	bindEditValConfig(cmd, &cfg)

	return cmd
}

func EditVal(ctx context.Context, cfg EditValConfig) error {
	valPriv, err := cfg.privateKey()
	if err != nil {
		return err
	}
	valAddr := crypto.PubkeyToAddress(valPriv.PublicKey)

	_, cprov, backend, err := setupClients(ctx, cfg.EOAConfig, valPriv)
	if err != nil {
		return err
	}

	// check if we already have an existing validator
	val, ok, err := cprov.SDKValidator(ctx, valAddr)
	if err != nil {
		return err
	} else if !ok {
		return &CliError{
			Msg:     "Operator address is not a validator: " + valAddr.Hex(),
			Suggest: "Ensure operator is already created as validator, see create-validator command",
		}
	}

	modifiedParams := cfg.modifiedAttrs()
	if len(modifiedParams) == 0 {
		return &CliError{
			Msg:     "All flags are default, no update possible",
			Suggest: "Provide at least one flag to update",
		}
	}

	log.Info(ctx, "Modifying the following parameters", modifiedParams...)

	contract, err := bindings.NewStaking(common.HexToAddress(predeploys.Staking), backend)
	if err != nil {
		return err
	}

	fee, err := contract.Fee(&bind.CallOpts{Context: ctx})
	if err != nil {
		return err
	}

	txOpts, err := backend.BindOpts(ctx, valAddr)
	if err != nil {
		return err
	}
	txOpts.Value = fee

	callOpts := &bind.CallOpts{Context: ctx}
	ok, err = contract.IsAllowlistEnabled(callOpts)
	if err != nil {
		return errors.Wrap(err, "check allowlist enabled")
	} else if ok {
		ok, err := contract.IsAllowedValidator(&bind.CallOpts{Context: ctx}, valAddr)
		if err != nil {
			return err
		} else if !ok {
			return errors.New("active validator not in allowed list [BUG]")
		}
	}

	minSelfWei := math.NewInt(-1)
	if cfg.MinSelfDelegationEther != -1 {
		// CLI flag min-self-delegation is in OMNI, convert to wei
		minSelfWei = math.NewInt(cfg.MinSelfDelegationEther).MulRaw(params.Ether)
		if val.MinSelfDelegation.GTE(minSelfWei) {
			return &CliError{
				Msg:     "--min-self-delegation too low",
				Suggest: "Provide a higher value than existing min-self-delegation=" + val.MinSelfDelegation.QuoRaw(params.Ether).String(),
			}
		}
	}

	tx, err := contract.EditValidator(txOpts, bindings.StakingEditValidatorParams{
		Moniker:                  cfg.Moniker,
		Identity:                 cfg.Identity,
		Website:                  cfg.Website,
		SecurityContact:          cfg.SecurityContact,
		Details:                  cfg.Details,
		CommissionRatePercentage: cfg.CommissionRatePercentage,
		MinSelfDelegation:        minSelfWei.BigInt(),
	})
	if err != nil {
		return errors.Wrap(err, "edit validator")
	}

	rec, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	}

	log.Info(ctx, "🎉 Edit validator transaction sent and included on-chain",
		"link", cfg.Network.Static().OmniScanTXURL(tx.Hash()),
		"block", rec.BlockNumber.Uint64(),
	)

	log.Info(ctx, "⏳ Staking events are delayed in the Omni Consensus chain and may take up to 12h to apply.")

	return nil
}

// setupClients is a helper that creates the omni evm client,
// omni consensus client and a backend set with the operator private key.
func setupClients(
	ctx context.Context,
	conf EOAConfig,
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

	eth, err := ethclient.DialContext(ctx, chainMeta.Name, conf.ExecutionRPC)
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
