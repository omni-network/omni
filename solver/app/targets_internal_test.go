package app

import (
	"testing"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/solver/targets"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/stretchr/testify/require"
)

func TestCallAllower(t *testing.T) {
	t.Parallel()

	targets.InitStatic()

	// actual address does not matter
	middlemanAddr := common.HexToAddress("0x1234567890123456789012345678901234567890")

	tests := []struct {
		name     string
		network  netconf.ID
		chainID  uint64
		target   common.Address
		calldata []byte
		allowed  bool
	}{
		{
			name:     "staging flowgen",
			network:  netconf.Staging,
			chainID:  evmchain.IDSepolia,
			target:   eoa.MustAddress(netconf.Staging, eoa.RoleFlowgen),
			calldata: nil,
			allowed:  true,
		},
		{
			name:     "omega flowgen",
			network:  netconf.Omega,
			chainID:  evmchain.IDSepolia,
			target:   eoa.MustAddress(netconf.Omega, eoa.RoleFlowgen),
			calldata: nil,
			allowed:  true,
		},
		{
			name:     "mainnet flowgen",
			network:  netconf.Mainnet,
			chainID:  evmchain.IDSepolia,
			target:   eoa.MustAddress(netconf.Mainnet, eoa.RoleFlowgen),
			calldata: nil,
			allowed:  true,
		},
		{
			name:     "holesky eigen",
			network:  netconf.Omega,
			chainID:  evmchain.IDHolesky,
			target:   targets.EigenHoleskyStrategyManager,
			calldata: []byte{0x01, 0x02, 0x03}, // doesn't matter,
			allowed:  true,
		},
		{
			name:     "holesky mainnet",
			network:  netconf.Mainnet,
			chainID:  evmchain.IDEthereum,
			target:   targets.EigenMainnetStrategyManager,
			calldata: []byte{0x01, 0x02, 0x03}, // doesn't matter,
			allowed:  true,
		},
		{
			name:     "staging staking",
			network:  netconf.Staging,
			chainID:  evmchain.IDOmniStaging,
			target:   common.HexToAddress(predeploys.Staking),
			calldata: []byte{0x01, 0x02, 0x03}, // doesn't matter,
			allowed:  true,
		},
		{
			name:     "omega staking",
			network:  netconf.Omega,
			chainID:  evmchain.IDOmniOmega,
			target:   common.HexToAddress(predeploys.Staking),
			calldata: []byte{0x01, 0x02, 0x03}, // doesn't matter,
			allowed:  true,
		},
		{
			name:     "mainnet staking",
			network:  netconf.Mainnet,
			chainID:  evmchain.IDOmniMainnet,
			target:   common.HexToAddress(predeploys.Staking),
			calldata: []byte{0x01, 0x02, 0x03}, // doesn't matter,
			allowed:  true,
		},
		{
			name:    "middleman allowed call",
			network: netconf.Mainnet,
			chainID: evmchain.IDEthereum,
			target:  middlemanAddr,
			// executeAndTransfer call to allowed target (eigen strategy manager)
			calldata: hexutil.MustDecode("0xfebe2c2c000000000000000000000000b82381a3fbd3fafa77b3a7be693342618240067b000000000000000000000000e3481474b23f88a8917dbcb4cbc55efcf0f68cc7000000000000000000000000858646372cc42e1a627fce94aa7a7033e7cf075a0000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000004447e7ef24000000000000000000000000e3481474b23f88a8917dbcb4cbc55efcf0f68cc70000000000000000000000000000000000000000000000008ac7230489e8000000000000000000000000000000000000000000000000000000000000"),
			allowed:  true,
		},
		{
			name:    "middleman disallowed call",
			network: netconf.Omega,
			chainID: evmchain.IDSepolia,
			target:  middlemanAddr,
			// executeAndTransfer call to disallowed target
			calldata: hexutil.MustDecode("0xfebe2c2c000000000000000000000000b82381a3fbd3fafa77b3a7be693342618240067b0000000000000000000000006415d3b5fc615d4a00c71f4044dec24c141ebff8000000000000000000000000e3481474b23f88a8917dbcb4cbc55efcf0f68cc70000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000004447e7ef240000000000000000000000006415d3b5fc615d4a00c71f4044dec24c141ebff80000000000000000000000000000000000000000000000008ac7230489e8000000000000000000000000000000000000000000000000000000000000"),
			allowed:  false,
		},
		{
			name:    "middleman invalid params",
			network: netconf.Omega,
			chainID: evmchain.IDSepolia,
			target:  middlemanAddr,
			// executeAndTransfer, invalid params (only selector)
			calldata: hexutil.MustDecode("0xfebe2c2c"),
			allowed:  false,
		},
		{
			name:    "middleman invalid calldata",
			network: netconf.Omega,
			chainID: evmchain.IDSepolia,
			target:  middlemanAddr,
			// invalid calldata (not executeAndTransfer)
			calldata: hexutil.MustDecode("0x12345678"),
			allowed:  false,
		},
		{
			name:     "devnet, calls not restricted",
			network:  netconf.Devnet,
			chainID:  evmchain.IDSepolia,
			target:   common.HexToAddress("0xe3481474b23f88a8917DbcB4cBC55Efcf0f68CC7"), // doesn't matter
			calldata: []byte{0x01, 0x02, 0x03},                                          // doesn't matter,
			allowed:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			allow := newCallAllower(tt.network, middlemanAddr)
			allowed := allow(tt.chainID, tt.target, tt.calldata)

			require.Equal(t, tt.allowed, allowed)
		})
	}
}

func TestParseMiddlemanCall(t *testing.T) {
	t.Parallel()

	// executeAndTransfer call to allowed target (0x6415D3B5fc615D4a00C71f4044dEc24C141EBFf8 - sepolia symbiotic wstETH vault)
	calldata := hexutil.MustDecode("0xfebe2c2c000000000000000000000000b82381a3fbd3fafa77b3a7be693342618240067b000000000000000000000000e3481474b23f88a8917dbcb4cbc55efcf0f68cc70000000000000000000000006415d3b5fc615d4a00c71f4044dec24c141ebff80000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000004447e7ef24000000000000000000000000e3481474b23f88a8917dbcb4cbc55efcf0f68cc70000000000000000000000000000000000000000000000008ac7230489e8000000000000000000000000000000000000000000000000000000000000")
	target, proxiedData, err := parseMiddlemanCall(calldata)
	require.NoError(t, err)
	require.Equal(t, common.HexToAddress("0x6415D3B5fc615D4a00C71f4044dEc24C141EBFf8"), target)
	require.Equal(t, "0x47e7ef24000000000000000000000000e3481474b23f88a8917dbcb4cbc55efcf0f68cc70000000000000000000000000000000000000000000000008ac7230489e80000", hexutil.Encode(proxiedData))

	// executeAndTransfer call to disallowed target (0xe3481474b23f88a8917DbcB4cBC55Efcf0f68CC7)
	calldata = hexutil.MustDecode("0xfebe2c2c000000000000000000000000b82381a3fbd3fafa77b3a7be693342618240067b0000000000000000000000006415d3b5fc615d4a00c71f4044dec24c141ebff8000000000000000000000000e3481474b23f88a8917dbcb4cbc55efcf0f68cc70000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000004447e7ef240000000000000000000000006415d3b5fc615d4a00c71f4044dec24c141ebff80000000000000000000000000000000000000000000000008ac7230489e8000000000000000000000000000000000000000000000000000000000000")
	target, proxiedData, err = parseMiddlemanCall(calldata)
	require.NoError(t, err)
	require.Equal(t, common.HexToAddress("0xe3481474b23f88a8917DbcB4cBC55Efcf0f68CC7"), target)
	require.Equal(t, "0x47e7ef240000000000000000000000006415d3b5fc615d4a00c71f4044dec24c141ebff80000000000000000000000000000000000000000000000008ac7230489e80000", hexutil.Encode(proxiedData))
}
