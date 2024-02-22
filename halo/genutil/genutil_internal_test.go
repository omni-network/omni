package genutil

import (
	"testing"

	atypes "github.com/omni-network/omni/halo/attest/types"
	etypes "github.com/omni-network/omni/halo/evmengine/types"

	"github.com/cometbft/cometbft/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/require"
)

func TestDefaultConsensusParams(t *testing.T) {
	t.Parallel()
	cons := defaultConsensusGenesis()
	require.EqualValues(t, cons.Params.ABCI.VoteExtensionsEnableHeight, 1)
	require.EqualValues(t, cons.Params.Validator.PubKeyTypes[0], types.ABCIPubKeyTypeSecp256k1)
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

	cdc := getCodec()
	txConfig := authtx.NewTxConfig(cdc, nil)

	b := txConfig.NewTxBuilder()
	err := b.SetMsgs(msgs...)
	require.NoError(t, err)

	tx := b.GetTx()

	require.Len(t, tx.GetMsgs(), 2)
	msgsV2, err := tx.GetMsgsV2()
	require.NoError(t, err)
	require.Len(t, msgsV2, 2)
}
