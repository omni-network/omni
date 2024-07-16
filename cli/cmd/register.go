package cmd

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"os"
	"time"

	"github.com/omni-network/omni/lib/contracts/avs"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"

	eigentypes "github.com/Layr-Labs/eigenlayer-cli/pkg/types"
	eigenutils "github.com/Layr-Labs/eigenlayer-cli/pkg/utils"
	eigenecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	eigensdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"gopkg.in/yaml.v3"
)

const l1BlockPeriod = time.Second * 12

type RegConfig struct {
	ConfigFile string
	AVSAddr    string
}

// RegDeps contains the Register dependencies that are abstracted for testing.
type RegDeps struct {
	Prompter       eigenutils.Prompter
	NewBackendFunc func(chainName string, chainID uint64, blockPeriod time.Duration, ethCl ethclient.Client, privateKeys ...*ecdsa.PrivateKey) (*ethbackend.Backend, error)
	VerifyFunc     func(eigensdktypes.Operator) error
}

type regOpt func(*RegDeps)

// Register registers the operator with the omni AVS contract.
//
// It assumes that the operator is already registered with the Eigen-Layer
// and that the eigen-layer configuration file (and ecdsa keystore) is present on disk.
func Register(ctx context.Context, cfg RegConfig, opts ...regOpt) error {
	// Default dependencies.
	deps := RegDeps{
		Prompter:       eigenutils.NewPrompter(),
		NewBackendFunc: ethbackend.NewBackend,
		VerifyFunc: func(op eigensdktypes.Operator) error {
			return op.Validate()
		},
	}
	for _, opt := range opts {
		opt(&deps)
	}

	eigenCfg, err := readConfig(cfg.ConfigFile)
	if err != nil {
		return err
	} else if err := deps.VerifyFunc(eigenCfg.Operator); err != nil {
		return errors.Wrap(err, "config validation failed")
	}

	password, err := deps.Prompter.InputHiddenString("Enter password to decrypt the ecdsa private key:", "",
		func(string) error {
			return nil
		},
	)
	if err != nil {
		return errors.Wrap(err, "read input")
	}

	privKey, err := eigenecdsa.ReadKey(eigenCfg.SignerConfig.PrivateKeyStorePath, password)
	if err != nil {
		return errors.Wrap(err, "read private key", "path", eigenCfg.SignerConfig.PrivateKeyStorePath)
	}

	ethCl, err := ethclient.Dial(chainNameFromID(eigenCfg.ChainId), eigenCfg.EthRPCUrl)
	if err != nil {
		return errors.Wrap(err, "dial eth client", "url", eigenCfg.EthRPCUrl)
	}

	avsAddress, err := avsAddressOrDefault(cfg.AVSAddr, &eigenCfg.ChainId)
	if err != nil {
		return err
	}

	backend, err := deps.NewBackendFunc(chainNameFromID(eigenCfg.ChainId), eigenCfg.ChainId.Uint64(), l1BlockPeriod, ethCl)
	if err != nil {
		return errors.Wrap(err, "create backend")
	}

	operator, err := backend.AddAccount(privKey)
	if err != nil {
		return errors.Wrap(err, "add account")
	}

	err = avs.RegisterOperatorWithAVS(ctx, avsAddress, backend, operator)
	if err != nil {
		// Parse solidity returned reason from CanRegister.
		switch err.Error() {
		case "already registered":
			return &CliError{Msg: "operator address already registered"}
		case "not an operator":
			return &CliError{
				Msg:     "not an eigen layer operator",
				Suggest: "Have you registered as an operator with Eigen-Layer?",
			}
		case "not in allowlist":
			return &CliError{Msg: "operator address not in Omni AVS allow-list"}
		case "max operators reached":
			return &CliError{Msg: "maximum number of operators in Omni AVS reached"}
		case "min stake not met":
			return &CliError{
				Msg:     "minimum stake requirement not met",
				Suggest: "Delegate more stake with Eigen-Layer.",
			}
		case "invalid delegation manager address":
			return &CliError{
				Msg:     "invalid Eigen-Layer delegation manager address",
				Suggest: "Is el_delegation_manager set correctly in your operator.yaml?",
			}
		case "no contract code at given address":
			return &CliError{
				Msg:     "no contract code at given address",
				Suggest: "Is eth_rpc_url set correctly in your operator.yaml?",
			}
		default:
			return err
		}
	}

	log.Info(ctx, "✅ Registration successful", "operator", operator.Hex())

	return nil
}

// readConfig returns the eigen-layer operator configuration from the given file.
func readConfig(file string) (eigentypes.OperatorConfig, error) {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return eigentypes.OperatorConfig{}, errors.Wrap(err, "eigen config file not found", "path", file)
	}

	bz, err := os.ReadFile(file)
	if err != nil {
		return eigentypes.OperatorConfig{}, errors.Wrap(err, "read eigen config file", "path", file)
	}

	var config eigentypes.OperatorConfig
	if err := yaml.Unmarshal(bz, &config); err != nil {
		return eigentypes.OperatorConfig{}, errors.Wrap(err, "unmarshal eigen config file")
	}

	return config, nil
}

func avsAddressOrDefault(avsAddr string, chainID *big.Int) (common.Address, error) {
	var resp common.Address
	if avsAddr != "" {
		if !common.IsHexAddress(avsAddr) {
			return common.Address{}, errors.New("invalid avs address", "address", avsAddr)
		}
		resp = common.HexToAddress(avsAddr)
	} else if addr, ok := avsFromChainID(chainID); ok {
		resp = addr
	} else {
		return common.Address{}, errors.New("avs address not provided and no default for chain found", "chain_id", chainID.Uint64())
	}

	return resp, nil
}

func avsFromChainID(chainID *big.Int) (common.Address, bool) {
	switch chainID.Int64() {
	case eigenutils.HoleskyChainId:
		return netconf.Omega.Static().AVSContractAddress, true
	case eigenutils.MainnetChainId:
		return netconf.Mainnet.Static().AVSContractAddress, true
	// TODO(corver): We need to publish our AVS addresses somewhere
	default:
		return common.Address{}, false
	}
}

// chainNameFromID is a best-effort attempt to map a chain ID to a human-readable name.
func chainNameFromID(id big.Int) string {
	switch id.Int64() {
	case eigenutils.MainnetChainId:
		return "mainnet"
	case eigenutils.HoleskyChainId:
		return "holesky"
	default:
		return "unknown"
	}
}
