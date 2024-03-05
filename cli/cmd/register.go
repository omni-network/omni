package cmd

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"os"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/avs"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	eigentypes "github.com/Layr-Labs/eigenlayer-cli/pkg/types"
	eigenutils "github.com/Layr-Labs/eigenlayer-cli/pkg/utils"
	eigenecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	eigensdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"gopkg.in/yaml.v3"
)

const l1BlockPeriod = time.Second * 12

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
func Register(ctx context.Context, configFile string, avsAddr string, opts ...regOpt) error {
	// Default dependencies.
	deps := RegDeps{
		Prompter:       eigenutils.NewPrompter(),
		NewBackendFunc: ethbackend.NewBackend,
		VerifyFunc: func(op eigensdktypes.Operator) error {
			return op.Validate() //nolint:wrapcheck // Wrapped below
		},
	}
	for _, opt := range opts {
		opt(&deps)
	}

	eigenCfg, err := readConfig(configFile)
	if err != nil {
		return err
	} else if err := deps.VerifyFunc(eigenCfg.Operator); err != nil {
		return errors.Wrap(err, "config validation failed")
	}

	password, err := deps.Prompter.InputHiddenString("Enter password to decrypt the ecdsa private key:", "",
		func(password string) error {
			return nil
		},
	)
	if err != nil {
		return errors.Wrap(err, "read input")
	}

	privKey, err := eigenecdsa.ReadKey(eigenCfg.PrivateKeyStorePath, password)
	if err != nil {
		return errors.Wrap(err, "read private key", "path", eigenCfg.PrivateKeyStorePath)
	}

	ethCl, err := ethclient.Dial(chainNameFromID(eigenCfg.ChainId), eigenCfg.EthRPCUrl)
	if err != nil {
		return errors.Wrap(err, "dial eth client", "url", eigenCfg.EthRPCUrl)
	}

	avsAddress, err := avsAddressOrDefault(avsAddr, &eigenCfg.ChainId)
	if err != nil {
		return err
	}

	backend, err := deps.NewBackendFunc(chainNameFromID(eigenCfg.ChainId), eigenCfg.ChainId.Uint64(), l1BlockPeriod, ethCl)
	if err != nil {
		return errors.Wrap(err, "create backend")
	}

	contracts, err := makeContracts(ctx, backend, eigenCfg, avsAddress)
	if err != nil {
		return err
	}

	operator, err := backend.AddAccount(privKey)
	if err != nil {
		return errors.Wrap(err, "add account")
	}

	return avs.RegisterOperatorWithAVS(ctx, contracts, backend, operator)
}

// makeContracts returns a avs Contracts struct with the given backend and delegation manager, avs directory, and omni avs contracts.
// Note only those three contracts are populated.
func makeContracts(ctx context.Context, backend *ethbackend.Backend, cfg eigentypes.OperatorConfigNew, avsAddr common.Address) (avs.Contracts, error) {
	if !common.IsHexAddress(cfg.ELDelegationManagerAddress) {
		return avs.Contracts{}, errors.New("invalid delegation manager address")
	}

	delManAddr := common.HexToAddress(cfg.ELDelegationManagerAddress)
	delMan, err := bindings.NewDelegationManager(delManAddr, backend)
	if err != nil {
		return avs.Contracts{}, errors.Wrap(err, "delegation manager")
	}

	omniAVS, err := bindings.NewOmniAVS(avsAddr, backend)
	if err != nil {
		return avs.Contracts{}, errors.Wrap(err, "omni avs")
	}

	avsDirecAddr, err := omniAVS.AvsDirectory(&bind.CallOpts{Context: ctx})
	if err != nil {
		return avs.Contracts{}, errors.Wrap(err, "avs directory")
	}
	avsDirec, err := bindings.NewAVSDirectory(avsDirecAddr, backend)
	if err != nil {
		return avs.Contracts{}, errors.Wrap(err, "avs directory")
	}

	return avs.Contracts{
		DelegationManager: delMan,
		AVSDirectory:      avsDirec,
		OmniAVS:           omniAVS,

		DelegationManagerAddr: delManAddr,
		AVSDirectoryAddr:      avsDirecAddr,
		OmniAVSAddr:           avsAddr,
	}, nil
}

// readConfig returns the eigen-layer operator configuration from the given file.
func readConfig(file string) (eigentypes.OperatorConfigNew, error) {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return eigentypes.OperatorConfigNew{}, errors.Wrap(err, "eigen config file not found", "path", file)
	}

	bz, err := os.ReadFile(file)
	if err != nil {
		return eigentypes.OperatorConfigNew{}, errors.Wrap(err, "read eigen config file", "path", file)
	}

	var config eigentypes.OperatorConfigNew
	if err := yaml.Unmarshal(bz, &config); err != nil {
		return eigentypes.OperatorConfigNew{}, errors.Wrap(err, "unmarshal eigen config file")
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

//nolint:gocritic,unparam,revive // It will be expanded in future.
func avsFromChainID(chainID *big.Int) (common.Address, bool) {
	switch chainID.Int64() {
	// case eigenutils.GoerliChainId:
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
	case eigenutils.GoerliChainId:
		return "goerli"
	case eigenutils.HoleskyChainId:
		return "holesky"
	default:
		return "unknown"
	}
}
