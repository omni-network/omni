package app

import (
	"context"
	"regexp"

	"github.com/omni-network/omni/halo/app/upgrades"
	attestkeeper "github.com/omni-network/omni/halo/attest/keeper"
	atypes "github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/halo/comet"
	"github.com/omni-network/omni/halo/evmdistribution"
	evmredenomkeeper "github.com/omni-network/omni/halo/evmredenom/keeper"
	evmredenomsubmit "github.com/omni-network/omni/halo/evmredenom/submit"
	"github.com/omni-network/omni/halo/evmslashing"
	evmstakinkeeper "github.com/omni-network/omni/halo/evmstaking/keeper"
	"github.com/omni-network/omni/halo/evmupgrade"
	registrykeeper "github.com/omni-network/omni/halo/registry/keeper"
	rtypes "github.com/omni-network/omni/halo/registry/types"
	valsynckeeper "github.com/omni-network/omni/halo/valsync/keeper"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	evmengkeeper "github.com/omni-network/omni/octane/evmengine/keeper"
	etypes "github.com/omni-network/omni/octane/evmengine/types"

	abci "github.com/cometbft/cometbft/abci/types"

	"cosmossdk.io/depinject"
	sdklog "cosmossdk.io/log"
	evidencekeeper "cosmossdk.io/x/evidence/keeper"
	upgradekeeper "cosmossdk.io/x/upgrade/keeper"
	upgradetypes "cosmossdk.io/x/upgrade/types"
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
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
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
	AccountKeeper          authkeeper.AccountKeeper
	BankKeeper             bankkeeper.Keeper
	StakingKeeper          *stakingkeeper.Keeper
	SlashingKeeper         slashingkeeper.Keeper
	DistrKeeper            distrkeeper.Keeper
	ConsensusParamsKeeper  consensuskeeper.Keeper
	EVMEngKeeper           *evmengkeeper.Keeper
	AttestKeeper           *attestkeeper.Keeper
	ValSyncKeeper          *valsynckeeper.Keeper
	EVMStakingKeeper       *evmstakinkeeper.Keeper
	RegistryKeeper         registrykeeper.Keeper
	EvidenceKeeper         evidencekeeper.Keeper
	UpgradeKeeper          *upgradekeeper.Keeper
	MintKeeper             mintkeeper.Keeper
	EVMRedenomKeeper       *evmredenomkeeper.Keeper
	EVMRedenomSubmitConfig evmredenomsubmit.Config

	SlashingEventProc     evmslashing.EventProcessor
	UpgradeEventProc      evmupgrade.EventProcessor
	DistributionEventProc evmdistribution.EventProcessor
}

// newApp returns a reference to an initialized App.
func newApp(
	ctx context.Context,
	network netconf.ID,
	logger sdklog.Logger,
	db dbm.DB,
	engineCl ethclient.EngineClient,
	voter atypes.Voter,
	chainVerNamer atypes.ChainVerNameFunc,
	chainNamer rtypes.ChainNameFunc,
	feeRecProvider etypes.FeeRecipientProvider,
	evmRedenomSubmitConfig evmredenomsubmit.Config,
	appOpts servertypes.AppOptions,
	asyncAbort chan<- error,
	baseAppOpts ...func(*baseapp.BaseApp),
) (*App, error) {
	depCfg := depinject.Configs(
		depinject.Provide(diProviders...),
		depinject.Invoke(diInvokers...),
		depinject.Configs(bankWrapperBindings...),
		appConfig(ctx, network),
		depinject.Supply(
			logger,
			engineCl,
			chainVerNamer,
			chainNamer,
			voter,
			feeRecProvider,
			appOpts,
			evmRedenomSubmitConfig,
		),
	)

	var (
		app        = new(App)
		appBuilder = new(runtime.AppBuilder)
	)
	dependencies := []any{
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
		&app.EVMStakingKeeper,
		&app.MintKeeper,
		&app.SlashingEventProc,
		&app.UpgradeEventProc,
		&app.DistributionEventProc,
		&app.EVMRedenomKeeper,
		&app.EVMRedenomSubmitConfig,
	}

	if err := depinject.Inject(depCfg, dependencies...); err != nil {
		return nil, errors.Wrap(err, "dep inject")
	}

	// Wire provider.
	app.EVMEngKeeper.SetVoteProvider(app.AttestKeeper)
	app.AttestKeeper.SetValidatorProvider(app.ValSyncKeeper)
	app.AttestKeeper.SetPortalRegistry(app.RegistryKeeper)

	baseAppOpts = append(baseAppOpts, func(bapp *baseapp.BaseApp) { //nolint:contextcheck // False positive wrt ctx
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
		app.ModuleManager.OrderEndBlockers = endBlockers(ctx)
		app.SetEndBlocker(app.EndBlocker)

		// Wrap upgrade module preblocker and do immediate shutdown if upgrade is needed.
		app.SetPreBlocker(func(ctx sdk.Context, req *abci.RequestFinalizeBlock) (*sdk.ResponsePreBlock, error) { //nolint:contextcheck // False positive wrt ctx
			resp, err := app.PreBlocker(ctx, req)
			if upgrade, ok := isErrOldBinary(err); ok {
				// Dump last applied upgrade info to disk so cosmovisor can auto upgrade.
				if err := dumpLastAppliedUpgradeInfo(ctx, app.UpgradeKeeper); err != nil {
					log.Error(ctx, "Failed writing last applied upgrade info", err)
				}
				asyncAbort <- errors.Wrap(err, "⛔️ network already upgraded, switch halo binary", "upgrade", upgrade)
				<-ctx.Done() // Wait for shutdown.
			} else if upgrade, ok := isErrUpgradeNeeded(err); ok {
				asyncAbort <- errors.Wrap(err, "⛔️ network upgrade needed now, switch halo binary", "upgrade", upgrade)
				<-ctx.Done() // Wait for shutdown.
			}

			return resp, err //nolint:wrapcheck // Don't wrap this cosmos error.
		})
	}

	// setUpgradeHandlers should be called before `Load()`
	// because StoreLoad is sealed after that
	if err := app.setUpgradeHandlers(ctx); err != nil {
		return nil, errors.Wrap(err, "set upgrade handlers")
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

func (a App) setUpgradeHandlers(ctx context.Context) error {
	for _, u := range upgrades.Upgrades {
		a.UpgradeKeeper.SetUpgradeHandler(u.Name, u.HandlerFunc(a))
	}

	upgradeInfo, err := a.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		return errors.Wrap(err, "read upgrade info from disk")
	} else if upgradeInfo.Name == "" {
		return nil // No upgrade info found
	}

	for _, u := range upgrades.Upgrades {
		if u.Name != upgradeInfo.Name {
			continue
		}

		a.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, u.Store(ctx)))

		return nil
	}

	return errors.New("unknown upgrade info [BUG]", "name", upgradeInfo.Name)
}

func (a App) GetModuleManager() *module.Manager {
	return a.ModuleManager
}

func (a App) GetModuleConfigurator() module.Configurator {
	return a.Configurator()
}

func (a App) GetSlashingKeeper() slashingkeeper.Keeper {
	return a.SlashingKeeper
}

func (a App) GetMintKeeper() mintkeeper.Keeper {
	return a.MintKeeper
}

func (a App) GetAccountKeeper() authkeeper.AccountKeeper {
	return a.AccountKeeper
}

func (a App) GetEVMEngineKeeper() *evmengkeeper.Keeper {
	return a.EVMEngKeeper
}

func (a App) GetEVMRedenomKeeper() *evmredenomkeeper.Keeper {
	return a.EVMRedenomKeeper
}

func (a App) GetEVMRedenomSubmitConfig() evmredenomsubmit.Config {
	return a.EVMRedenomSubmitConfig
}

// dumpLastAppliedUpgradeInfo dumps the last applied upgrade info to disk.
// This is a workaround for halovisor to auto upgrade binaries
// after snapsyncing to a post-upgrade state using a pre-upgrade (old) binary.
func dumpLastAppliedUpgradeInfo(ctx sdk.Context, keeper *upgradekeeper.Keeper) error {
	name, height, err := keeper.GetLastCompletedUpgrade(ctx)
	if err != nil {
		return errors.Wrap(err, "get last completed upgrade")
	}

	// Note that we need to ensure that the next binary doesn't actually run any
	// store loader upgrades on startup, it was already done during the upgrade.
	// We therefore ensure height isn't current.

	current := sdk.UnwrapSDKContext(ctx).BlockHeight()
	if height >= current { // Sanity check that the upgrade was in the past.
		return errors.New("unexpected last upgrade height [BUG]")
	}

	err = keeper.DumpUpgradeInfoToDisk(height, upgradetypes.Plan{
		Name:   name,
		Height: height,
	})
	if err != nil {
		return errors.Wrap(err, "dump upgrade info")
	}

	return nil
}

// isErrOldBinary returns the last applied upgrade and true if the error is due to the
// upgrade module detecting the binary is too old.
func isErrOldBinary(err error) (string, bool) {
	if err == nil {
		return "", false
	}

	cause := errors.Cause(err)

	reg := regexp.MustCompile(`wrong app version \d+, upgrade handler is missing for (.+) upgrade plan`)
	if !reg.MatchString(cause.Error()) {
		return "", false
	}

	matches := reg.FindStringSubmatch(cause.Error())
	if len(matches) != 2 {
		return "", false
	}

	return matches[1], true
}

func isErrUpgradeNeeded(err error) (string, bool) {
	if err == nil {
		return "", false
	}

	cause := errors.Cause(err)

	reg := regexp.MustCompile(`UPGRADE "(.+)" NEEDED .+`)
	if !reg.MatchString(cause.Error()) {
		return "", false
	}

	matches := reg.FindStringSubmatch(cause.Error())
	if len(matches) != 2 {
		return "", false
	}

	return matches[1], true
}

var (
	_ runtime.AppI            = (*App)(nil)
	_ servertypes.Application = (*App)(nil)
)
