package cmd

import (
	"encoding/json"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"

	"github.com/cometbft/cometbft/crypto"
	"github.com/cometbft/cometbft/types"
	cmttime "github.com/cometbft/cometbft/types/time"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cosmosstd "github.com/cosmos/cosmos-sdk/std"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	stypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/gogoproto/proto"
)

func MakeGenesis(chainID string, valPubKeys ...crypto.PubKey) (types.GenesisDoc, error) {
	power := cosmostypes.DefaultPowerReduction // Use any non-zero power for this single validator.

	cometVals := make([]types.GenesisValidator, 0, len(valPubKeys))
	cosmosVals := make([]stypes.Validator, 0, len(valPubKeys))
	for _, pubkey := range valPubKeys {
		cometVals = append(cometVals, types.GenesisValidator{
			Address: pubkey.Address(),
			PubKey:  pubkey,
			Power:   power.Int64(),
		})

		addr, err := k1util.PubKeyToAddress(pubkey)
		if err != nil {
			return types.GenesisDoc{}, err
		}
		pb, err := k1util.PubKeyToCosmos(pubkey)
		if err != nil {
			return types.GenesisDoc{}, err
		}
		cosmosVal, err := stypes.NewValidator(addr.Hex(), pb, stypes.Description{})
		if err != nil {
			return types.GenesisDoc{}, errors.Wrap(err, "create cosmos validator")
		}
		cosmosVal.Status = stypes.Bonded
		cosmosVal.Tokens = power

		cosmosVals = append(cosmosVals, cosmosVal)
	}

	stakingGen := stypes.NewGenesisState(stypes.DefaultParams(), cosmosVals, nil)
	_, err := proto.Marshal(stakingGen)
	if err != nil {
		return types.GenesisDoc{}, errors.Wrap(err, "proto marshal staking genesis")
	}

	stakingBZ, err := getCodec().MarshalJSON(stakingGen)
	if err != nil {
		return types.GenesisDoc{}, errors.Wrap(err, "json marshal staking genesis")
	}

	genesisData := map[string]json.RawMessage{"staking": stakingBZ}
	genesisBZ, err := json.Marshal(genesisData)
	if err != nil {
		return types.GenesisDoc{}, errors.Wrap(err, "marshal genesis")
	}

	return types.GenesisDoc{
		ChainID:         chainID,
		GenesisTime:     cmttime.Now(),
		ConsensusParams: DefaultConsensusParams(),
		Validators:      cometVals,
		AppState:        genesisBZ,
	}, nil
}

func getCodec() *codec.ProtoCodec {
	reg := codectypes.NewInterfaceRegistry()
	cosmosstd.RegisterInterfaces(reg)

	return codec.NewProtoCodec(reg)
}
