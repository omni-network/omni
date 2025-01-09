package keeper_test

import (
	"context"
	"testing"

	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/feature"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"
	"github.com/omni-network/omni/octane/evmengine/keeper"
	"github.com/omni-network/omni/octane/evmengine/types"

	k1 "github.com/cometbft/cometbft/crypto/secp256k1"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

func TestFetchProcEvents(t *testing.T) {
	t.Parallel()

	// Abridged addresses
	s := common.HexToAddress(predeploys.Staking)
	p := common.HexToAddress(predeploys.PortalRegistry)

	fetch := func(ctx context.Context) []types.EVMEvent {
		hash := tutil.RandomHash()
		ethStake := int64(7)
		privKey := k1.GenPrivKey()
		ethCl, err := ethclient.NewEngineMock(
			ethclient.WithPortalRegister(netconf.SimnetNetwork()),        // 3 events
			ethclient.WithMockValidatorCreation(privKey.PubKey()),        // 1 event
			ethclient.WithMockSelfDelegation(privKey.PubKey(), ethStake), // 1 event
		)
		require.NoError(t, err)

		procs := []types.EvmEventProcessor{
			testProc{Address: s},
			testProc{Address: p},
			testProc{Address: common.HexToAddress(predeploys.Slashing)},
		}

		events, err := keeper.FetchProcEvents(ctx, ethCl, hash, procs...)
		require.NoError(t, err)

		return events
	}

	ctx := context.Background()

	// Legacy ordering by address > topics > data
	events := fetch(ctx)
	assertOrder(t, events, p, p, p, s, s)

	// Simple ordering by index (see MockEngineClient for indexes)
	ctx = feature.WithFlag(ctx, feature.FlagSimpleEVMEvents)
	events = fetch(ctx)
	assertOrder(t, events, s, p, p, p, s)
}

func assertOrder(t *testing.T, events []types.EVMEvent, addresses ...common.Address) {
	t.Helper()
	require.Len(t, events, len(addresses))
	for i, ev := range events {
		require.Equal(t, addresses[i], common.BytesToAddress(ev.Address))
	}
}

type testProc struct {
	types.EvmEventProcessor
	Address common.Address
}

func (p testProc) FilterParams() ([]common.Address, [][]common.Hash) {
	return []common.Address{p.Address}, nil
}
