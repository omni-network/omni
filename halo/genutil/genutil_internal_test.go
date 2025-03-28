package genutil

import (
	"testing"

	haloapp "github.com/omni-network/omni/halo/app"
	atypes "github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/lib/netconf"
	etypes "github.com/omni-network/omni/octane/evmengine/types"

	"github.com/cometbft/cometbft/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/require"
)

func TestDefaultConsensusParams(t *testing.T) {
	t.Parallel()
	cons := defaultConsensusGenesis()
	require.EqualValues(t, 1, cons.Params.ABCI.VoteExtensionsEnableHeight)
	require.Equal(t, types.ABCIPubKeyTypeSecp256k1, cons.Params.Validator.PubKeyTypes[0])
	require.EqualValues(t, -1, cons.Params.Block.MaxBytes)
	require.EqualValues(t, -1, cons.Params.Block.MaxGas)
}

func TestEncodeTXs(t *testing.T) {
	t.Parallel()
	msgs := []sdk.Msg{
		&etypes.MsgExecutionPayload{
			Authority: authtypes.NewModuleAddress("evm").String(),
		},
		&atypes.MsgAddVotes{
			Authority: authtypes.NewModuleAddress("evm").String(),
		},
	}

	encConf, err := haloapp.ClientEncodingConfig(t.Context(), netconf.Simnet)
	require.NoError(t, err)

	b := encConf.TxConfig.NewTxBuilder()
	err = b.SetMsgs(msgs...)
	require.NoError(t, err)

	tx := b.GetTx()

	require.Len(t, tx.GetMsgs(), 2)
	msgsV2, err := tx.GetMsgsV2()
	require.NoError(t, err)
	require.Len(t, msgsV2, 2)
}
