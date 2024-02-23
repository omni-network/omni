package app

import (
	attestkeeper "github.com/omni-network/omni/halo/attest/keeper"
	atypes "github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/halo/comet"
	evmengkeeper "github.com/omni-network/omni/halo/evmengine/keeper"
	"github.com/omni-network/omni/lib/engine"
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
	EVMEngKeeper          *evmengkeeper.Keeper
	AttestKeeper          *attestkeeper.Keeper
}

// newApp returns a reference to an initialized App.
func newApp(
	logger log.Logger,
	db dbm.DB,
	ethCl engine.API,
	voter atypes.Voter,
	baseAppOpts ...func(*baseapp.BaseApp),
) (*App, error) {
	depCfg := depinject.Configs(
		DepConfig(),
		depinject.Supply(
			logger, ethCl, voter,
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
		&app.EVMEngKeeper,
		&app.AttestKeeper,
	); err != nil {
		return nil, errors.Wrap(err, "dep inject")
	}

	// Add CProviders to evmengine keeper.
	app.EVMEngKeeper.AddProvider(app.AttestKeeper)

	proposalHandler := makeProcessProposalHandler(app)

	baseAppOpts = append(baseAppOpts, func(bapp *baseapp.BaseApp) {
		// Use evm engine to create block proposals.
		// Note that we do not check MaxTxBytes since all EngineEVM transaction MUST be included since we cannot
		// postpone them to the next block. Nit: we could drop some vote extensions though...?
		bapp.SetPrepareProposal(app.EVMEngKeeper.PrepareProposal)

		// Route proposed messaged to keepers for verification and external state updates.
		bapp.SetProcessProposal(proposalHandler)

		// Use attest keeper to extend votes.
		bapp.SetExtendVoteHandler(app.AttestKeeper.ExtendVote)
		bapp.SetVerifyVoteExtensionHandler(app.AttestKeeper.VerifyVoteExtension)
	})

	app.App = appBuilder.Build(db, nil, baseAppOpts...)

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

// SetCometAPI sets the comet API client.
// TODO(corver): Figure out how to use depinject to set this.
func (a App) SetCometAPI(api comet.API) {
	a.AttestKeeper.SetCometAPI(api)
	a.EVMEngKeeper.SetCometAPI(api)
}

var (
	_ runtime.AppI            = (*App)(nil)
	_ servertypes.Application = (*App)(nil)
)
