package keeper

import (
	"context"
	"math/big"
	"testing"

	"cosmossdk.io/x/tx/signing"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cosmosstd "github.com/cosmos/cosmos-sdk/std"
	sdk "github.com/cosmos/cosmos-sdk/types"
	signing2 "github.com/cosmos/cosmos-sdk/types/tx/signing"
	authcodec "github.com/cosmos/cosmos-sdk/x/auth/codec"
	atypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	btypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	dtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/gogoproto/proto"
	eengine "github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	attesttypes "github.com/omni-network/omni/halo/attest/types"
	etypes "github.com/omni-network/omni/halo/evmengine/types"
	"github.com/omni-network/omni/lib/engine"
)

func TestKeeper_PrepareProposal(t *testing.T) {

	keeper := NewKeeper(getCodec(), nil, MockEngineAPI{}, MockTxConfig{}, MockAddressProvider{})
	_ = keeper
}

func getCodec() *codec.ProtoCodec {
	// TODO(corver): Use depinject to get all of this.
	sdkConfig := sdk.GetConfig()
	reg, err := codectypes.NewInterfaceRegistryWithOptions(codectypes.InterfaceRegistryOptions{
		ProtoFiles: proto.HybridResolver,
		SigningOptions: signing.Options{
			AddressCodec:          authcodec.NewBech32Codec(sdkConfig.GetBech32AccountAddrPrefix()),
			ValidatorAddressCodec: authcodec.NewBech32Codec(sdkConfig.GetBech32ValidatorAddrPrefix()),
		},
	})
	if err != nil {
		panic(err)
	}

	cosmosstd.RegisterInterfaces(reg)
	atypes.RegisterInterfaces(reg)
	stypes.RegisterInterfaces(reg)
	btypes.RegisterInterfaces(reg)
	dtypes.RegisterInterfaces(reg)
	etypes.RegisterInterfaces(reg)
	attesttypes.RegisterInterfaces(reg)

	return codec.NewProtoCodec(reg)
}

var _ engine.API = (*MockEngineAPI)(nil)
var _ client.TxConfig = (*MockTxConfig)(nil)
var _ etypes.AddressProvider = (*MockAddressProvider)(nil)

type MockEngineAPI struct{}
type MockTxConfig struct{}
type MockAddressProvider struct{}

func (m MockAddressProvider) LocalAddress() common.Address {
	//TODO implement me
	panic("implement me")
}

func (m MockTxConfig) TxEncoder() sdk.TxEncoder {
	//TODO implement me
	panic("implement me")
}

func (m MockTxConfig) TxDecoder() sdk.TxDecoder {
	//TODO implement me
	panic("implement me")
}

func (m MockTxConfig) TxJSONEncoder() sdk.TxEncoder {
	//TODO implement me
	panic("implement me")
}

func (m MockTxConfig) TxJSONDecoder() sdk.TxDecoder {
	//TODO implement me
	panic("implement me")
}

func (m MockTxConfig) MarshalSignatureJSON(v2s []signing2.SignatureV2) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockTxConfig) UnmarshalSignatureJSON(bytes []byte) ([]signing2.SignatureV2, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockTxConfig) NewTxBuilder() client.TxBuilder {
	//TODO implement me
	panic("implement me")
}

func (m MockTxConfig) WrapTxBuilder(tx sdk.Tx) (client.TxBuilder, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockTxConfig) SignModeHandler() *signing.HandlerMap {
	//TODO implement me
	panic("implement me")
}

func (m MockTxConfig) SigningContext() *signing.Context {
	//TODO implement me
	panic("implement me")
}

func (m MockEngineAPI) BlockNumber(ctx context.Context) (uint64, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockEngineAPI) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockEngineAPI) NewPayloadV2(ctx context.Context, params eengine.ExecutableData) (eengine.PayloadStatusV1, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockEngineAPI) NewPayloadV3(ctx context.Context, params eengine.ExecutableData, versionedHashes []common.Hash, beaconRoot *common.Hash) (eengine.PayloadStatusV1, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockEngineAPI) ForkchoiceUpdatedV2(ctx context.Context, update eengine.ForkchoiceStateV1, payloadAttributes *eengine.PayloadAttributes) (eengine.ForkChoiceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockEngineAPI) ForkchoiceUpdatedV3(ctx context.Context, update eengine.ForkchoiceStateV1, payloadAttributes *eengine.PayloadAttributes) (eengine.ForkChoiceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockEngineAPI) GetPayloadV2(ctx context.Context, payloadID eengine.PayloadID) (*eengine.ExecutionPayloadEnvelope, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockEngineAPI) GetPayloadV3(ctx context.Context, payloadID eengine.PayloadID) (*eengine.ExecutionPayloadEnvelope, error) {
	//TODO implement me
	panic("implement me")
}
