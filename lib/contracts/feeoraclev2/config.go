package feeoraclev2

import (
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/tokens"
)

const defaultBaseGasLimitGwei = 100_000
const defaultBaseBytes = 100
const defaultGasPerByteGwei = 16

type feeConfig struct {
	ChainID      uint64
	DataCostID   uint64
	BaseGasLimit uint32
}

type dataCostConfig struct {
	ID         uint64
	BaseBytes  uint32
	GasPerByte uint64
	GasToken   uint16
}

var (
	feeConfigs = map[uint64]feeConfig{
		// Mainnets.
		evmchain.IDEthereum: {
			ChainID:      evmchain.IDEthereum,
			DataCostID:   evmchain.IDEthereum,
			BaseGasLimit: defaultBaseGasLimitGwei,
		},
		evmchain.IDOmniMainnet: {
			ChainID:      evmchain.IDOmniMainnet,
			DataCostID:   evmchain.IDOmniMainnet,
			BaseGasLimit: defaultBaseGasLimitGwei,
		},
		evmchain.IDArbitrumOne: {
			ChainID:      evmchain.IDArbitrumOne,
			DataCostID:   evmchain.IDEthereum,
			BaseGasLimit: defaultBaseGasLimitGwei,
		},
		evmchain.IDOptimism: {
			ChainID:      evmchain.IDOptimism,
			DataCostID:   evmchain.IDEthereum,
			BaseGasLimit: defaultBaseGasLimitGwei,
		},
		evmchain.IDBase: {
			ChainID:      evmchain.IDBase,
			DataCostID:   evmchain.IDEthereum,
			BaseGasLimit: defaultBaseGasLimitGwei,
		},

		// Testnets.
		evmchain.IDOmniStaging: {
			ChainID:      evmchain.IDOmniStaging,
			DataCostID:   evmchain.IDOmniStaging,
			BaseGasLimit: defaultBaseGasLimitGwei,
		},
		evmchain.IDOmniOmega: {
			ChainID:      evmchain.IDOmniOmega,
			DataCostID:   evmchain.IDOmniOmega,
			BaseGasLimit: defaultBaseGasLimitGwei,
		},
		evmchain.IDHolesky: {
			ChainID:      evmchain.IDHolesky,
			DataCostID:   evmchain.IDHolesky,
			BaseGasLimit: defaultBaseGasLimitGwei,
		},
		evmchain.IDSepolia: {
			ChainID:      evmchain.IDSepolia,
			DataCostID:   evmchain.IDSepolia,
			BaseGasLimit: defaultBaseGasLimitGwei,
		},
		evmchain.IDArbSepolia: {
			ChainID:      evmchain.IDArbSepolia,
			DataCostID:   evmchain.IDSepolia,
			BaseGasLimit: defaultBaseGasLimitGwei,
		},
		evmchain.IDOpSepolia: {
			ChainID:      evmchain.IDOpSepolia,
			DataCostID:   evmchain.IDSepolia,
			BaseGasLimit: defaultBaseGasLimitGwei,
		},
		evmchain.IDBaseSepolia: {
			ChainID:      evmchain.IDBaseSepolia,
			DataCostID:   evmchain.IDSepolia,
			BaseGasLimit: defaultBaseGasLimitGwei,
		},

		// Ephemeral chains.
		evmchain.IDOmniDevnet: {
			ChainID:      evmchain.IDOmniDevnet,
			DataCostID:   evmchain.IDOmniDevnet,
			BaseGasLimit: defaultBaseGasLimitGwei,
		},
		evmchain.IDMockL1: {
			ChainID:      evmchain.IDMockL1,
			DataCostID:   evmchain.IDMockL1,
			BaseGasLimit: defaultBaseGasLimitGwei,
		},
		evmchain.IDMockL2: {
			ChainID:      evmchain.IDMockL2,
			DataCostID:   evmchain.IDMockL2,
			BaseGasLimit: defaultBaseGasLimitGwei,
		},
		evmchain.IDMockOp: {
			ChainID:      evmchain.IDMockOp,
			DataCostID:   evmchain.IDMockOp,
			BaseGasLimit: defaultBaseGasLimitGwei,
		},
		evmchain.IDMockArb: {
			ChainID:      evmchain.IDMockArb,
			DataCostID:   evmchain.IDMockArb,
			BaseGasLimit: defaultBaseGasLimitGwei,
		},
	}

	dataCostConfigs = map[uint64]dataCostConfig{
		// Mainnets.
		evmchain.IDEthereum: {
			ID:         evmchain.IDEthereum,
			BaseBytes:  defaultBaseBytes,
			GasPerByte: defaultGasPerByteGwei,
			GasToken:   gasTokenIDs[tokens.ETH],
		},
		evmchain.IDOmniMainnet: {
			ID:         evmchain.IDOmniMainnet,
			BaseBytes:  defaultBaseBytes,
			GasPerByte: defaultGasPerByteGwei,
			GasToken:   gasTokenIDs[tokens.NOM],
		},

		// Testnets.
		evmchain.IDOmniStaging: {
			ID:         evmchain.IDOmniStaging,
			BaseBytes:  defaultBaseBytes,
			GasPerByte: defaultGasPerByteGwei,
			GasToken:   gasTokenIDs[tokens.NOM],
		},
		evmchain.IDOmniOmega: {
			ID:         evmchain.IDOmniOmega,
			BaseBytes:  defaultBaseBytes,
			GasPerByte: defaultGasPerByteGwei,
			GasToken:   gasTokenIDs[tokens.NOM],
		},
		evmchain.IDHolesky: {
			ID:         evmchain.IDHolesky,
			BaseBytes:  defaultBaseBytes,
			GasPerByte: defaultGasPerByteGwei,
			GasToken:   gasTokenIDs[tokens.ETH],
		},
		evmchain.IDSepolia: {
			ID:         evmchain.IDSepolia,
			BaseBytes:  defaultBaseBytes,
			GasPerByte: defaultGasPerByteGwei,
			GasToken:   gasTokenIDs[tokens.ETH],
		},

		// Ephemeral chains.
		evmchain.IDOmniDevnet: {
			ID:         evmchain.IDOmniDevnet,
			BaseBytes:  defaultBaseBytes,
			GasPerByte: defaultGasPerByteGwei,
			GasToken:   gasTokenIDs[tokens.NOM],
		},
		evmchain.IDMockL1: {
			ID:         evmchain.IDMockL1,
			BaseBytes:  defaultBaseBytes,
			GasPerByte: defaultGasPerByteGwei,
			GasToken:   gasTokenIDs[tokens.ETH],
		},
		evmchain.IDMockL2: {
			ID:         evmchain.IDMockL2,
			BaseBytes:  defaultBaseBytes,
			GasPerByte: defaultGasPerByteGwei,
			GasToken:   gasTokenIDs[tokens.ETH],
		},
		evmchain.IDMockOp: {
			ID:         evmchain.IDMockOp,
			BaseBytes:  defaultBaseBytes,
			GasPerByte: defaultGasPerByteGwei,
			GasToken:   gasTokenIDs[tokens.ETH],
		},
		evmchain.IDMockArb: {
			ID:         evmchain.IDMockArb,
			BaseBytes:  defaultBaseBytes,
			GasPerByte: defaultGasPerByteGwei,
			GasToken:   gasTokenIDs[tokens.ETH],
		},
	}
)

func getFeeConfig(chainID uint64) (feeConfig, bool) {
	cfg, ok := feeConfigs[chainID]
	return cfg, ok
}

func getDataCostConfig(chainID uint64) (dataCostConfig, bool) {
	cfg, ok := dataCostConfigs[chainID]
	return cfg, ok
}
