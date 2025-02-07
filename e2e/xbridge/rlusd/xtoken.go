package rlusd

import (
	"context"

	"github.com/omni-network/omni/e2e/xbridge/types"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
)

type rlusd struct{}

var _ types.XToken = (*rlusd)(nil)

func XToken() types.XToken          { return rlusd{} }
func Name() string                  { return xtoken.Name }
func Symbol() string                { return xtoken.Symbol }
func Wraps() types.TokenDescriptors { return wraps }

func (rlusd) Name() string                  { return xtoken.Name }
func (rlusd) Symbol() string                { return xtoken.Symbol }
func (rlusd) Wraps() types.TokenDescriptors { return wraps }

// Address returns the address of the RLUSD.e token on the given network.
func (rlusd) Address(ctx context.Context, networkID netconf.ID) (common.Address, error) {
	addr, err := contracts.Create3Address(ctx, networkID, saltID(xtoken))
	if err != nil {
		return common.Address{}, errors.Wrap(err, "salt")
	}

	return addr, nil
}

// Canonical returns the canonical RLUSD deployment by network.
func (rlusd) Canonical(ctx context.Context, networkID netconf.ID) (types.TokenDeployment, error) {
	canonical, ok := canonicals[networkID]
	if ok {
		return canonical, nil
	}

	// if no canonical deployments in static, we deploy a stand in mock

	addr, err := contracts.Create3Address(ctx, networkID, saltID(wraps))
	if err != nil {
		return types.TokenDeployment{}, errors.Wrap(err, "salt")
	}

	return types.TokenDeployment{
		Name:    wraps.Name,
		Symbol:  wraps.Symbol,
		ChainID: netconf.EthereumChainID(networkID),
		Address: addr,
	}, nil
}
