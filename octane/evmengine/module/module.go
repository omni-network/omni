package module

import (
	"encoding/json"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/octane/evmengine/keeper"
	"github.com/omni-network/omni/octane/evmengine/types"

	"github.com/ethereum/go-ethereum/common"

	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/core/store"
	"cosmossdk.io/depinject"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

var (
	_ module.AppModuleBasic = (*AppModule)(nil)
	_ appmodule.AppModule   = (*AppModule)(nil)
	_ module.HasGenesis     = (*AppModule)(nil)
)

// ----------------------------------------------------------------------------
// AppModuleBasic
// ----------------------------------------------------------------------------

// AppModuleBasic implements the AppModuleBasic interface that defines the
// independent methods a Cosmos SDK module needs to implement.
type AppModuleBasic struct {
	cdc codec.BinaryCodec
}

func NewAppModuleBasic(cdc codec.BinaryCodec) AppModuleBasic {
	return AppModuleBasic{cdc: cdc}
}

// Name returns the name of the module as a string.
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// RegisterLegacyAminoCodec registers the amino codec for the module, which is used
// to marshal and unmarshal structs to/from []byte in order to persist them in the module's KVStore.
func (AppModuleBasic) RegisterLegacyAminoCodec(*codec.LegacyAmino) {}

// RegisterInterfaces registers a module's interface types and their concrete implementations as proto.Message.
func (AppModuleBasic) RegisterInterfaces(reg cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(reg)
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(client.Context, *runtime.ServeMux) {}

// ----------------------------------------------------------------------------
// AppModule
// ----------------------------------------------------------------------------

// AppModule implements the AppModule interface that defines the inter-dependent methods that modules need to implement.
type AppModule struct {
	AppModuleBasic

	keeper *keeper.Keeper
}

func NewAppModule(
	cdc codec.Codec,
	keeper *keeper.Keeper,
) AppModule {
	return AppModule{
		AppModuleBasic: NewAppModuleBasic(cdc),
		keeper:         keeper,
	}
}

func (m AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, raw json.RawMessage) {
	var data types.GenesisState
	cdc.MustUnmarshalJSON(raw, &data)

	if err := m.keeper.InsertGenesisHead(ctx, data.ExecutionBlockHash); err != nil {
		panic(errors.Wrap(err, "insert genesis head"))
	}
}

func (AppModule) ExportGenesis(sdk.Context, codec.JSONCodec) json.RawMessage {
	return nil
}

func (AppModuleBasic) DefaultGenesis(codec.JSONCodec) json.RawMessage {
	panic("not supported")
}

// ValidateGenesis performs genesis state validation for the bank module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, _ client.TxEncodingConfig, bz json.RawMessage) error {
	var data types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &data); err != nil {
		return errors.Wrap(err, "unmarshal genesis state")
	} else if len(data.ExecutionBlockHash) != common.HashLength {
		return errors.New("invalid execution block hash length")
	}

	return nil
}

// RegisterServices registers a gRPC query service to respond to the module-specific gRPC queries.
func (m AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServiceServer(cfg.MsgServer(), keeper.NewMsgServerImpl(m.keeper))
}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (AppModule) IsOnePerModuleType() {}

// IsAppModule implements the appmodule.AppModule interface.
func (AppModule) IsAppModule() {}

// ----------------------------------------------------------------------------
// App Wiring Setup
// ----------------------------------------------------------------------------

//nolint:gochecknoinits // Cosmos-style
func init() {
	appmodule.Register(
		&Module{},
		appmodule.Provide(
			ProvideModule,
		),
	)
}

type ModuleInputs struct {
	depinject.In

	StoreService   store.KVStoreService
	Cdc            codec.Codec
	Config         *Module
	TXConfig       client.TxConfig
	EngineCl       ethclient.EngineClient
	AddrProvider   types.AddressProvider
	FeeRecProvider types.FeeRecipientProvider
}

type ModuleOutputs struct {
	depinject.Out

	EngEVMKeeper *keeper.Keeper
	Module       appmodule.AppModule
	Hooks        staking.StakingHooksWrapper
}

func ProvideModule(in ModuleInputs) (ModuleOutputs, error) {
	k, err := keeper.NewKeeper(
		in.Cdc,
		in.StoreService,
		in.EngineCl,
		in.TXConfig,
		in.AddrProvider,
		in.FeeRecProvider,
	)
	if err != nil {
		return ModuleOutputs{}, err
	}

	m := NewAppModule(
		in.Cdc,
		k,
	)

	return ModuleOutputs{
		EngEVMKeeper: k,
		Module:       m,
		Hooks:        staking.StakingHooksWrapper{StakingHooks: keeper.Hooks{}},
	}, nil
}
