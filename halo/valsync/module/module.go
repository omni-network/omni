package module

import (
	"context"
	"encoding/json"

	atypes "github.com/omni-network/omni/halo/attest/types"
	ptypes "github.com/omni-network/omni/halo/portal/types"
	"github.com/omni-network/omni/halo/valsync/keeper"
	"github.com/omni-network/omni/halo/valsync/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"

	abci "github.com/cometbft/cometbft/abci/types"

	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/core/store"
	"cosmossdk.io/depinject"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

var (
	_ module.AppModuleBasic  = (*AppModule)(nil)
	_ module.HasGenesis      = (*AppModule)(nil)
	_ module.HasABCIEndBlock = (*AppModule)(nil)
	_ appmodule.AppModule    = (*AppModule)(nil)
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
func (AppModuleBasic) RegisterInterfaces(cdctypes.InterfaceRegistry) {}

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

func (m AppModule) EndBlock(ctx context.Context) ([]abci.ValidatorUpdate, error) {
	return m.keeper.EndBlock(ctx)
}

func (m AppModule) InitGenesis(ctx sdk.Context, _ codec.JSONCodec, _ json.RawMessage) {
	if err := m.keeper.InsertGenesisSet(ctx); err != nil {
		panic(errors.Wrap(err, "insert genesis valset"))
	}
}

func (AppModule) ExportGenesis(sdk.Context, codec.JSONCodec) json.RawMessage {
	return nil
}

// DefaultGenesis returns default genesis state as raw bytes for the bank
// module.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesisState())
}

// ValidateGenesis performs genesis state validation for the bank module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, _ client.TxEncodingConfig, bz json.RawMessage) error {
	var data types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &data); err != nil {
		return errors.Wrap(err, "unmarshal genesis state")
	}

	return nil
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

// RegisterServices registers a gRPC query service to respond to the module-specific gRPC queries.
func (m AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterQueryServer(cfg.QueryServer(), m.keeper)
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

	StoreService store.KVStoreService
	Cdc          codec.Codec
	Config       *Module
	SKeeper      types.StakingKeeper
	AKeeper      atypes.AttestKeeper
	Subscriber   types.ValSetSubscriber
	Portal       ptypes.EmitPortal
	EthClient    ethclient.Client
}

type ModuleOutputs struct {
	depinject.Out

	Keeper *keeper.Keeper
	Module appmodule.AppModule
}

func ProvideModule(in ModuleInputs) (ModuleOutputs, error) {
	k, err := keeper.NewKeeper(
		in.StoreService,
		in.SKeeper,
		in.AKeeper,
		in.Subscriber,
		in.Portal,
		in.EthClient,
	)
	if err != nil {
		return ModuleOutputs{}, err
	}

	m := NewAppModule(
		in.Cdc,
		k,
	)

	return ModuleOutputs{
		Keeper: k,
		Module: m,
	}, nil
}
