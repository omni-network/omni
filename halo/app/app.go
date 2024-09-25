package app

import (
	"context"
	"fmt"
	"strings"

	attestkeeper "github.com/omni-network/omni/halo/attest/keeper"
	atypes "github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/halo/comet"
	"github.com/omni-network/omni/halo/evmslashing"
	"github.com/omni-network/omni/halo/evmstaking"
	"github.com/omni-network/omni/halo/evmupgrade"
	registrykeeper "github.com/omni-network/omni/halo/registry/keeper"
	rtypes "github.com/omni-network/omni/halo/registry/types"
	valsynckeeper "github.com/omni-network/omni/halo/valsync/keeper"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	evmengkeeper "github.com/omni-network/omni/octane/evmengine/keeper"
	etypes "github.com/omni-network/omni/octane/evmengine/types"

	abci "github.com/cometbft/cometbft/abci/types"

	"cosmossdk.io/depinject"
	sdklog "cosmossdk.io/log"
	evidencekeeper "cosmossdk.io/x/evidence/keeper"
	upgradekeeper "cosmossdk.io/x/upgrade/keeper"
	utypes "cosmossdk.io/x/upgrade/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	consensuskeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"

	_ "cosmossdk.io/api/cosmos/tx/config/v1"          // import for side-effects
	_ "cosmossdk.io/x/evidence"                       // import for side-effects
	_ "cosmossdk.io/x/upgrade"                        // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/auth"           // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/auth/tx/config" // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/bank"           // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/consensus"      // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/distribution"   // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/genutil"        // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/mint"           // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/slashing"       // import for side-effects
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
	SlashingKeeper        slashingkeeper.Keeper
	DistrKeeper           distrkeeper.Keeper
	ConsensusParamsKeeper consensuskeeper.Keeper
	EVMEngKeeper          *evmengkeeper.Keeper
	AttestKeeper          *attestkeeper.Keeper
	ValSyncKeeper         *valsynckeeper.Keeper
	RegistryKeeper        registrykeeper.Keeper
	EvidenceKeeper        evidencekeeper.Keeper
	UpgradeKeeper         *upgradekeeper.Keeper

	SlashingEventProc evmslashing.EventProcessor
	StakingEventProc  evmstaking.EventProcessor
	UpgradeEventProc  evmupgrade.EventProcessor
}

// newApp returns a reference to an initialized App.
func newApp(
	logger sdklog.Logger,
	db dbm.DB,
	engineCl ethclient.EngineClient,
	voter atypes.Voter,
	chainVerNamer atypes.ChainVerNameFunc,
	chainNamer rtypes.ChainNameFunc,
	feeRecProvider etypes.FeeRecipientProvider,
	appOpts servertypes.AppOptions,
	baseAppOpts ...func(*baseapp.BaseApp),
) (*App, error) {
	depCfg := depinject.Configs(
		appConfig,
		depinject.Provide(diProviders...),
		depinject.Supply(
			logger,
			engineCl,
			chainVerNamer,
			chainNamer,
			voter,
			feeRecProvider,
			appOpts,
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
		&app.SlashingKeeper,
		&app.DistrKeeper,
		&app.ConsensusParamsKeeper,
		&app.EVMEngKeeper,
		&app.AttestKeeper,
		&app.ValSyncKeeper,
		&app.RegistryKeeper,
		&app.EvidenceKeeper,
		&app.UpgradeKeeper,
		&app.SlashingEventProc,
		&app.StakingEventProc,
		&app.UpgradeEventProc,
	); err != nil {
		return nil, errors.Wrap(err, "dep inject")
	}

	// Wire provider.
	app.EVMEngKeeper.SetVoteProvider(app.AttestKeeper)
	app.AttestKeeper.SetValidatorProvider(app.ValSyncKeeper)
	app.AttestKeeper.SetPortalRegistry(app.RegistryKeeper)

	baseAppOpts = append(baseAppOpts, func(bapp *baseapp.BaseApp) {
		// Use evm engine to create block proposals.
		// Note that we do not check MaxTxBytes since all EngineEVM transaction MUST be included since we cannot
		// postpone them to the next block. Nit: we could drop some vote extensions though...?
		bapp.SetPrepareProposal(app.EVMEngKeeper.PrepareProposal)

		// Route proposed messages to keepers for verification and external state updates.
		bapp.SetProcessProposal(makeProcessProposalHandler(makeProcessProposalRouter(app), app.txConfig))

		// Use attest keeper to extend votes.
		bapp.SetExtendVoteHandler(app.AttestKeeper.ExtendVote)
		bapp.SetVerifyVoteExtensionHandler(app.AttestKeeper.VerifyVoteExtension)
	})

	app.App = appBuilder.Build(db, nil, baseAppOpts...)

	// Blocker overrides
	{
		// Workaround for official endblockers since valsync replaces staking endblocker, but cosmos panics if it's not there.
		app.ModuleManager.OrderEndBlockers = endBlockers
		app.SetEndBlocker(app.EndBlocker)

		// Also dump upgrade info to disk if binary is too old.
		// Cosmos upgrade only dumps when doing the upgrade.
		app.SetPreBlocker(func(ctx sdk.Context, req *abci.RequestFinalizeBlock) (*sdk.ResponsePreBlock, error) {
			resp, err := app.PreBlocker(ctx, req)
			if lastAppliedUpgrade, ok := isErrWrongVersion(err); ok {
				height := sdk.UnwrapSDKContext(ctx).BlockHeight()
				err := app.UpgradeKeeper.DumpUpgradeInfoToDisk(height, utypes.Plan{Name: lastAppliedUpgrade})
				if err != nil {
					// Best effort just log
					log.Warn(ctx, "Failed writing upgrade info", err)
				}
			}

			return resp, err //nolint:wrapcheck // Don't wrap this cosmos error.
		})
	}

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
	a.EVMEngKeeper.SetCometAPI(api)
}

// ClientContext returns a new client context with the app's codec and tx config.
func (a App) ClientContext(ctx context.Context) client.Context {
	return client.Context{}.
		WithInterfaceRegistry(a.interfaceRegistry).
		WithTxConfig(a.txConfig).
		WithChainID(a.ChainID()).
		WithCmdContext(ctx).
		WithCodec(a.appCodec)
}

// isErrWrongVersion returns the last applied upgrade if the error is due to a wrong app version, else it returns false.
func isErrWrongVersion(err error) (string, bool) {
	if err == nil {
		return "", false
	}

	// Get the root cause of the error.
	cause := err
	for {
		next := errors.Unwrap(cause)
		if next == nil {
			break
		}
		cause = next
	}

	var ignore int
	var lastAppliedUpgrade string
	i, err := fmt.Fscanf(
		strings.NewReader(cause.Error()),
		"wrong app version %d, upgrade handler is missing for %s upgrade plan",
		&ignore, &lastAppliedUpgrade)
	if err != nil || i != 2 {
		return "", false
	}

	return lastAppliedUpgrade, true
}

var (
	_ runtime.AppI            = (*App)(nil)
	_ servertypes.Application = (*App)(nil)
)
