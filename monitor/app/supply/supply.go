package supply

import (
	"context"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/params"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
)

func MonitorForever(ctx context.Context, cprov cchain.Provider, network netconf.Network, ethCls map[uint64]ethclient.Client) {
	ticker := time.NewTicker(time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := instrSupplies(ctx, cprov, network, ethCls); err != nil {
				log.Warn(ctx, "Token supply intrumentation failed", err)
			}
		}
	}
}

// instrSupplies instruments the supplies of OMNI token on the L1 (without the bridge balances)
// and on the consensus chain.
func instrSupplies(ctx context.Context, cprov cchain.Provider, network netconf.Network, ethCls map[uint64]ethclient.Client) error {
	response, err := cprov.QueryClients().Bank.TotalSupply(ctx, &types.QueryTotalSupplyRequest{})
	if err != nil {
		return errors.Wrap(err, "total supply query")
	}
	var cosmosSupplyWei uint64
	for _, coin := range response.Supply {
		if coin.Denom == sdk.DefaultBondDenom {
			cosmosSupplyWei += coin.Amount.Uint64()
		}
	}

	cChainSupply.Set(toGwei(float64(cosmosSupplyWei)))

	addrs, err := contracts.GetAddresses(ctx, network.ID)
	if err != nil {
		panic(err)
	}

	ethChainID := netconf.EthereumChainID(network.ID)
	l1Client, ok := ethCls[ethChainID]
	if !ok {
		return errors.Wrap(err, "ethereum client")
	}
	l1Token, err := bindings.NewOmni(addrs.Token, l1Client)
	if err != nil {
		return errors.Wrap(err, "contract bindings")
	}

	callOpts := &bind.CallOpts{Context: ctx}

	l1TokenSupplyWei, err := l1Token.TotalSupply(callOpts)
	if err != nil {
		return errors.Wrap(err, "l1 token supply")
	}
	eChainSupply.Set(toGwei(float64(l1TokenSupplyWei.Uint64())))

	l1BridgeBalanceWei, err := l1Token.BalanceOf(callOpts, addrs.L1Bridge)
	if err != nil {
		return errors.Wrap(err, "l1 bridge balance")
	}
	bridgeBalance.Set(toGwei(float64(l1BridgeBalanceWei.Uint64())))

	return nil
}

func toGwei(wei float64) float64 {
	return wei / params.GWei
}
