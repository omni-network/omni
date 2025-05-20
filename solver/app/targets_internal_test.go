package app

import (
	"testing"

	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/stretchr/testify/require"
)

func TestCallAllower(t *testing.T) {
	t.Parallel()

	executorAddr := tutil.RandomAddress()

	tests := []struct {
		name     string
		network  netconf.ID
		chainID  uint64
		target   common.Address
		calldata []byte
		allowed  bool
	}{
		{
			name:     "devnet calls not restricted",
			network:  netconf.Devnet,
			chainID:  evmchain.IDSepolia,
			target:   common.HexToAddress("0xe3481474b23f88a8917DbcB4cBC55Efcf0f68CC7"), // doesn't matter
			calldata: []byte{0x01, 0x02, 0x03},                                          // doesn't matter
			allowed:  true,
		},
		{
			name:     "staging calls not restricted",
			network:  netconf.Staging,
			chainID:  evmchain.IDSepolia,
			target:   common.HexToAddress("0xe3481474b23f88a8917DbcB4cBC55Efcf0f68CC7"), // doesn't matter
			calldata: []byte{0x01, 0x02, 0x03},                                          // doesn't matter
			allowed:  true,
		},
		{
			name:     "omega calls not restricted",
			network:  netconf.Omega,
			chainID:  evmchain.IDSepolia,
			target:   common.HexToAddress("0xe3481474b23f88a8917DbcB4cBC55Efcf0f68CC7"), // doesn't matter
			calldata: []byte{0x01, 0x02, 0x03},                                          // doesn't matter
			allowed:  true,
		},
		{
			name:     "mainnet calls not restricted",
			network:  netconf.Mainnet,
			chainID:  evmchain.IDEthereum,
			target:   common.HexToAddress("0xe3481474b23f88a8917DbcB4cBC55Efcf0f68CC7"), // doesn't matter
			calldata: []byte{0x01, 0x02, 0x03},                                          // doesn't matter
			allowed:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			allow := newCallAllower(tt.network, executorAddr)
			allowed := allow(tt.chainID, tt.target, tt.calldata)

			require.Equal(t, tt.allowed, allowed)
		})
	}
}

func TestParseExecutorCall(t *testing.T) {
	t.Parallel()

	// executeAndTransfer call to allowed target (0x6415D3B5fc615D4a00C71f4044dEc24C141EBFf8 - sepolia symbiotic wstETH vault)
	calldata := hexutil.MustDecode("0xfebe2c2c000000000000000000000000b82381a3fbd3fafa77b3a7be693342618240067b000000000000000000000000e3481474b23f88a8917dbcb4cbc55efcf0f68cc70000000000000000000000006415d3b5fc615d4a00c71f4044dec24c141ebff80000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000004447e7ef24000000000000000000000000e3481474b23f88a8917dbcb4cbc55efcf0f68cc70000000000000000000000000000000000000000000000008ac7230489e8000000000000000000000000000000000000000000000000000000000000")
	target, proxiedData, err := parseExecutorCall(calldata)
	require.NoError(t, err)
	require.Equal(t, common.HexToAddress("0x6415D3B5fc615D4a00C71f4044dEc24C141EBFf8"), target)
	require.Equal(t, "0x47e7ef24000000000000000000000000e3481474b23f88a8917dbcb4cbc55efcf0f68cc70000000000000000000000000000000000000000000000008ac7230489e80000", hexutil.Encode(proxiedData))

	// executeAndTransfer call to disallowed target (0xe3481474b23f88a8917DbcB4cBC55Efcf0f68CC7)
	calldata = hexutil.MustDecode("0xfebe2c2c000000000000000000000000b82381a3fbd3fafa77b3a7be693342618240067b0000000000000000000000006415d3b5fc615d4a00c71f4044dec24c141ebff8000000000000000000000000e3481474b23f88a8917dbcb4cbc55efcf0f68cc70000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000004447e7ef240000000000000000000000006415d3b5fc615d4a00c71f4044dec24c141ebff80000000000000000000000000000000000000000000000008ac7230489e8000000000000000000000000000000000000000000000000000000000000")
	target, proxiedData, err = parseExecutorCall(calldata)
	require.NoError(t, err)
	require.Equal(t, common.HexToAddress("0xe3481474b23f88a8917DbcB4cBC55Efcf0f68CC7"), target)
	require.Equal(t, "0x47e7ef240000000000000000000000006415d3b5fc615d4a00c71f4044dec24c141ebff80000000000000000000000000000000000000000000000008ac7230489e80000", hexutil.Encode(proxiedData))
}
