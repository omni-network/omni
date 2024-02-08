package cosmos

import (
	"context"

	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"github.com/ethereum/go-ethereum/common"

	"cosmossdk.io/core/address"
	"cosmossdk.io/core/store"
	cosmoslog "cosmossdk.io/log"
	"cosmossdk.io/store/rootmulti"
	storetypes "cosmossdk.io/store/types"
	cosmosdb "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	akeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	atypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	skeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
)

// SDKContext returns a new sdk.Context to use when interacting with cosmos modules.
// The header is the current cometBFT header.
func SDKContext(ctx context.Context, db cosmosdb.DB, header cmtproto.Header) sdk.Context {
	sdkctx := sdk.NewContext(
		rootmulti.NewStore(db, cosmosLog, nil),
		header,
		false,
		cosmosLog,
	)
	sdkctx.WithContext(ctx)

	return sdkctx
}

type Modules struct {
	Account akeeper.AccountKeeper
	Bank    bkeeper.BaseKeeper
	Staking *skeeper.Keeper
}

func MakeModules() Modules {
	// This will use standard protobuf registry
	cdc := codec.NewProtoCodec(types.NewInterfaceRegistry())
	var addrCodec ethAddreCodec // Our eth address codec

	// Simple prefix key
	newService := func(key string) store.KVStoreService {
		return runtime.NewKVStoreService(storetypes.NewKVStoreKey(key))
	}

	ak := akeeper.NewAccountKeeper(
		cdc,
		newService("account"),
		func() sdk.AccountI { return new(atypes.BaseAccount) },
		nil,
		addrCodec,
		"",
		"",
	)

	bk := bkeeper.NewBaseKeeper(
		cdc,
		newService("bank"),
		ak,
		nil,
		"",
		cosmosLog,
	)

	sk := skeeper.NewKeeper(
		cdc,
		newService("staking"),
		ak,
		bk,
		"",
		addrCodec,
		addrCodec,
	)

	return Modules{
		Account: ak,
		Bank:    bk,
		Staking: sk,
	}
}

var _ address.Codec = ethAddreCodec{}

type ethAddreCodec struct{}

func (ethAddreCodec) StringToBytes(text string) ([]byte, error) {
	return common.HexToAddress(text).Bytes(), nil
}

func (ethAddreCodec) BytesToString(bz []byte) (string, error) {
	return common.BytesToAddress(bz).String(), nil
}

// See cometLog how we can adapt this.
//
//nolint:gochecknoglobals // wip
var cosmosLog cosmoslog.Logger
