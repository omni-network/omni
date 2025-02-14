package types

import (
	"context"

	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
)

type TokenDescriptors struct {
	Name   string
	Symbol string
}

type TokenDeployment struct {
	Name    string
	Symbol  string
	ChainID uint64
	Address common.Address
}

type TokenGasLimits struct {
	BridgeNoLockbox   uint64
	BridgeWithLockbox uint64
}

type XToken interface {
	// Name returns the name of the xtoken
	Name() string

	// Symbol returns the symbol of the xtoken
	Symbol() string

	// Wraps returns the token this xtoken wraps
	Wraps() TokenDescriptors

	// GasLimits returns the gas limits for the xtoken
	GasLimits() TokenGasLimits

	// Address returns the xtoken deployment address, consistent across all chains in the network
	Address(ctx context.Context, networkID netconf.ID) (common.Address, error)

	// Canonical returns the canonical deployed of the token this xtoken wraps
	Canonical(ctx context.Context, networkID netconf.ID) (TokenDeployment, error)
}
