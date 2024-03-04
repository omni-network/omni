package cmd

import (
	"context"
	"math/big"
	"os"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/avs"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/test/e2e/backend"

	"github.com/ethereum/go-ethereum/common"

	eigentypes "github.com/Layr-Labs/eigenlayer-cli/pkg/types"
	eigenutils "github.com/Layr-Labs/eigenlayer-cli/pkg/utils"
	eigenecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"gopkg.in/yaml.v3"
)

const l1BlockPeriod = time.Second * 12

// register registers the operator with the omni AVS contract.
//
// It assumes that the operator is already registered with the Eigen-Layer
// and that the eigen-layer configuration file (and ecdsa keystore) is present on disk.
func register(ctx context.Context, configFile string, prompter eigenutils.Prompter, avsAddr string) error {
	cfg, err := readConfig(configFile)
	if err != nil {
		return err
	}

	password, err := prompter.InputHiddenString("Enter password to decrypt the ecdsa private key:", "",
		func(password string) error {
			return nil
		},
	)
	if err != nil {
		return errors.Wrap(err, "read input")
	}

	privKey, err := eigenecdsa.ReadKey(cfg.PrivateKeyStorePath, password)
	if err != nil {
		return errors.Wrap(err, "read private key", "path", cfg.PrivateKeyStorePath)
	}

	ethCl, err := ethclient.Dial(chainNameFromID(cfg.ChainId), cfg.EthRPCUrl)
	if err != nil {
		return errors.Wrap(err, "dial eth client", "url", cfg.EthRPCUrl)
	}

	var avsAddress common.Address
	if avsAddr != "" {
		if common.IsHexAddress(avsAddr) {
			return errors.New("invalid avs address")
		}
		avsAddress = common.HexToAddress(avsAddr)
	} else if addr, ok := avsFromChainID(&cfg.ChainId); ok {
		avsAddress = addr
	} else {
		return errors.New("avs address not provided and no default for chain found", "chain_id", cfg.ChainId.String())
	}

	backend, err := backend.NewBackend(chainNameFromID(cfg.ChainId), cfg.ChainId.Uint64(), l1BlockPeriod, ethCl)
	if err != nil {
		return errors.Wrap(err, "create backend")
	}

	contracts, err := makeContracts(backend, cfg, avsAddress)
	if err != nil {
		return err
	}

	operator, err := backend.AddAccount(privKey)
	if err != nil {
		return errors.Wrap(err, "add account")
	}

	return avs.RegisterOperatorWithAVS(ctx, contracts, backend, operator)
}

// makeContracts returns a avs Contracts struct with the given backend and delegation manager and omni avs.
// Note only those two contracts are populated.
func makeContracts(backend backend.Backend, cfg eigentypes.OperatorConfigNew, avsAddr common.Address) (avs.Contracts, error) {
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

	return avs.Contracts{
		DelegationManager: delMan,
		OmniAVS:           omniAVS,

		DelegationManagerAddr: delManAddr,
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

	if err := config.Operator.Validate(); err != nil {
		return eigentypes.OperatorConfigNew{}, errors.Wrap(err, "validate operator config")
	}

	return config, nil
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
