package app

import (
	"testing"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokenmeta"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

func TestClaimants(t *testing.T) {
	t.Parallel()

	golden := map[string]map[netconf.ID]common.Address{}

	for _, tkn := range tokenmeta.All() {
		for _, netID := range netconf.All() {
			claimant, ok := claimants[tkn][netID]
			if !ok {
				continue
			}

			if _, ok := golden[tkn.Symbol]; !ok {
				golden[tkn.Symbol] = map[netconf.ID]common.Address{}
			}

			golden[tkn.Symbol][netID] = claimant
		}
	}

	tutil.RequireGoldenJSON(t, golden)
}

func TestGetClaimants(t *testing.T) {
	t.Parallel()

	mustGet := func(network netconf.ID, order Order) common.Address {
		claimant, ok, err := getClaimant(network, order)
		require.NoError(t, err)
		require.True(t, ok)

		return claimant
	}

	// omega wsteth claimant
	require.Equal(t,
		common.HexToAddress("0x521786BE8A0f455700c25FB25F94A4B626E460Ec"),
		mustGet(netconf.Omega, Order{
			Status:        solvernet.StatusFilled,
			SourceChainID: evmchain.IDHolesky,
			filledData: FilledData{
				MinReceived: []bindings.IERC7683Output{{
					ChainId: bi.N(evmchain.IDHolesky),
					Token:   toBz32(common.HexToAddress("0x8d09a4502cc8cf1547ad300e066060d043f6982d")),
					Amount:  bi.Ether(1),
				}},
			},
		}),
		"omega wsteth claimment",
	)

	// mainnet wsteth claimant
	require.Equal(t,
		common.HexToAddress("0x79Ef4d1224a055Ad4Ee5e2226d0cb3720d929AE7"),
		mustGet(netconf.Mainnet, Order{
			Status:        solvernet.StatusFilled,
			SourceChainID: evmchain.IDBase,
			filledData: FilledData{
				MinReceived: []bindings.IERC7683Output{{
					ChainId: bi.N(evmchain.IDBase),
					Token:   toBz32(common.HexToAddress("0xc1cba3fcea344f92d9239c08c0568f6f2f0ee452")),
					Amount:  bi.Ether(1),
				}},
			},
		}),
		"mainnet wsteth claimment",
	)

	// eth claimed by solver
	order := Order{
		Status:        solvernet.StatusFilled,
		SourceChainID: evmchain.IDBase,
		filledData: FilledData{
			MinReceived: []bindings.IERC7683Output{{
				ChainId: bi.N(evmchain.IDBase),
				Amount:  bi.Ether(1),
			}},
		},
	}

	claimant, ok, err := getClaimant(netconf.Mainnet, order)
	require.NoError(t, err)
	require.False(t, ok)
	require.Equal(t, common.Address{}, claimant)
}
