package app

import (
	"github.com/omni-network/omni/lib/errors"

	"cosmossdk.io/depinject"
	"cosmossdk.io/log"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	consensuskeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"

	_ "cosmossdk.io/api/cosmos/tx/config/v1"          // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/auth"           // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/auth/tx/config" // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/bank"           // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/consensus"      // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/distribution"   // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/genutil"        // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/mint"           // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/staking"        // import for side-effects
)

const Name = "halo"

// App extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type App struct {
	*runtime.App

	appCodec          codec.Codec
	txConfig          client.TxConfig
	interfaceRegistry codectypes.InterfaceRegistry

	// keepers
	AccountKeeper         authkeeper.AccountKeeper
	BankKeeper            bankkeeper.Keeper
	StakingKeeper         *stakingkeeper.Keeper
	DistrKeeper           distrkeeper.Keeper
	ConsensusParamsKeeper consensuskeeper.Keeper
}

// newApp returns a reference to an initialized App.
func newApp(
	logger log.Logger,
	db dbm.DB,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) (*App, error) {
	depCfg := depinject.Configs(
		DepConfig(),
		depinject.Supply(
			appOpts, logger,
		),
	)

	var (
		app        = new(App)
		appBuilder = new(runtime.AppBuilder)
	)
	if err := depinject.Inject(depCfg,
		&appBuilder,
		&app.appCodec,
		&app.txConfig,
		&app.interfaceRegistry,
		&app.AccountKeeper,
		&app.BankKeeper,
		&app.StakingKeeper,
		&app.DistrKeeper,
		&app.ConsensusParamsKeeper,
	); err != nil {
		return nil, errors.Wrap(err, "dep inject")
	}

	app.App = appBuilder.Build(db, nil, baseAppOptions...)

	app.SetInitChainer(nil)
	if err := app.Load(true); err != nil {
		return nil, errors.Wrap(err, "load app")
	}

	return app, nil
}

func (App) LegacyAmino() *codec.LegacyAmino {
	return nil
}

func (App) ExportAppStateAndValidators(_ bool, _, _ []string) (servertypes.ExportedApp, error) {
	return servertypes.ExportedApp{}, errors.New("not implemented")
}

// SimulationManager implements the SimulationApp interface.
func (App) SimulationManager() *module.SimulationManager {
	return nil
}

var (
	_ runtime.AppI            = (*App)(nil)
	_ servertypes.Application = (*App)(nil)
)
