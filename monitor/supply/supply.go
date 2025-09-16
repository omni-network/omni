package supply

import (
	"context"
	"math/big"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
)

func MonitorForever(ctx context.Context, cprov cchain.Provider, network netconf.Network, ethCls map[uint64]ethclient.Client) {
	if _, ok := netconf.EthereumChainID(network.ID); !ok {
		return // No L1 chain, nothing to monitor
	}

	timer := time.NewTimer(0)
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			timer.Reset(time.Hour)

			if err := instrSupplies(ctx, cprov, network, ethCls); err != nil {
				log.Warn(ctx, "Token supply instrumentation failed", err)
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

	cosmosSupplyWei, err := stakeAmount(response.Supply)
	if err != nil {
		return errors.Wrap(err, "stake amount")
	}

	cChainSupply.Set(bi.ToEtherF64(cosmosSupplyWei))

	addrs, err := contracts.GetAddresses(ctx, network.ID)
	if err != nil {
		return errors.Wrap(err, "get addresses")
	}

	l1, ok := netconf.EthereumChainID(network.ID)
	if !ok {
		return errors.New("no L1 chain for network")
	}
	l1Client, ok := ethCls[l1]
	if !ok {
		return errors.New("ethereum client")
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
	l1Erc20Supply.Set(bi.ToEtherF64(l1TokenSupplyWei))

	l1BridgeBalanceWei, err := l1Token.BalanceOf(callOpts, addrs.L1Bridge)
	if err != nil {
		return errors.Wrap(err, "l1 bridge balance")
	}
	bridgeBalance.Set(bi.ToEtherF64(l1BridgeBalanceWei))

	return nil
}

func stakeAmount(coins sdk.Coins) (*big.Int, error) {
	if len(coins) != 1 {
		return nil, errors.New("unexpected number of coins")
	}
	ok, coin := coins.Find(sdk.DefaultBondDenom)
	if !ok {
		return nil, errors.New("missing default bond denom")
	}

	return coin.Amount.BigInt(), nil
}
