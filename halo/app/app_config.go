package app

import (
	"context"

	attestmodule "github.com/omni-network/omni/halo/attest/module"
	attesttypes "github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/halo/evmdistribution"
	evmredenommodule "github.com/omni-network/omni/halo/evmredenom/module"
	evmredenomtypes "github.com/omni-network/omni/halo/evmredenom/types"
	"github.com/omni-network/omni/halo/evmslashing"
	evmstakingmodule "github.com/omni-network/omni/halo/evmstaking/module"
	evmstakingtypes "github.com/omni-network/omni/halo/evmstaking/types"
	"github.com/omni-network/omni/halo/evmupgrade"
	portalmodule "github.com/omni-network/omni/halo/portal/module"
	portaltypes "github.com/omni-network/omni/halo/portal/types"
	registrymodule "github.com/omni-network/omni/halo/registry/module"
	registrytypes "github.com/omni-network/omni/halo/registry/types"
	"github.com/omni-network/omni/halo/sdk"
	valsyncmodule "github.com/omni-network/omni/halo/valsync/module"
	valsynctypes "github.com/omni-network/omni/halo/valsync/types"
	"github.com/omni-network/omni/halo/withdraw"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/netconf"
	engevmmodule "github.com/omni-network/omni/octane/evmengine/module"
	engevmtypes "github.com/omni-network/omni/octane/evmengine/types"

	k1 "github.com/cometbft/cometbft/crypto/secp256k1"

	runtimev1alpha1 "cosmossdk.io/api/cosmos/app/runtime/v1alpha1"
	appv1alpha1 "cosmossdk.io/api/cosmos/app/v1alpha1"
	authmodulev1 "cosmossdk.io/api/cosmos/auth/module/v1"
	bankmodulev1 "cosmossdk.io/api/cosmos/bank/module/v1"
	consensusmodulev1 "cosmossdk.io/api/cosmos/consensus/module/v1"
	distrmodulev1 "cosmossdk.io/api/cosmos/distribution/module/v1"
	evidencemodulev1 "cosmossdk.io/api/cosmos/evidence/module/v1"
	genutilmodulev1 "cosmossdk.io/api/cosmos/genutil/module/v1"
	mintmodulev1 "cosmossdk.io/api/cosmos/mint/module/v1"
	slashingmodulev1 "cosmossdk.io/api/cosmos/slashing/module/v1"
	stakingmodulev1 "cosmossdk.io/api/cosmos/staking/module/v1"
	txconfigv1 "cosmossdk.io/api/cosmos/tx/config/v1"
	upgrademodulev1 "cosmossdk.io/api/cosmos/upgrade/module/v1"
	"cosmossdk.io/core/appconfig"
	"cosmossdk.io/depinject"
	evidencetypes "cosmossdk.io/x/evidence/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// Note that these constants form part of consensus logic and can only be changed during network upgrades.
const (
	// TODO(corver): Maybe move these to genesis itself.
	genesisVoteWindowUp   uint64 = 64 // Allow early votes for <latest attestation - 64>
	genesisVoteWindowDown uint64 = 2  // Only allow late votes for <latest attestation - 2>
	genesisVoteExtLimit   uint64 = 256
	genesisTrimLag        uint64 = 1      // Allow deleting attestations in block after approval.
	genesisCTrimLag       uint64 = 72_000 // Delete consensus attestations state after +-1 day (given a period of 1.2s).

	deliverIntervalProtected = 1 // Disable batching in protected networks.
	deliverIntervalEphemeral = 2 // Fast updates while testing

	maxWithdrawalsPerBlock uint64 = 32 // The maximum number of withdrawals included in one block.
)

//nolint:gochecknoglobals // Cosmos-style
var (
	genesisModuleOrder = []string{
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		stakingtypes.ModuleName,
		slashingtypes.ModuleName,
		minttypes.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		upgradetypes.ModuleName,
		valsynctypes.ModuleName,
		engevmtypes.ModuleName,
	}

	beginBlockers = func(context.Context) []string {
		modules := []string{
			minttypes.ModuleName,  // Mint module is generally first
			distrtypes.ModuleName, // Note: slashing happens after distr.BeginBlocker
			slashingtypes.ModuleName,
			stakingtypes.ModuleName, // Note: staking module is required if HistoricalEntries param > 0
			evidencetypes.ModuleName,
			attesttypes.ModuleName,
		}

		return modules
	}

	endBlockers = func(context.Context) []string {
		return []string{
			attesttypes.ModuleName,
			evmstakingtypes.ModuleName,
			valsynctypes.ModuleName, // Wraps staking module end blocker (must come after attest module)
			upgradetypes.ModuleName,
		}
	}

	// blocked account addresses.
	blockAccAddrs = []string{
		authtypes.FeeCollectorName,
		distrtypes.ModuleName,
		stakingtypes.BondedPoolName,
		stakingtypes.NotBondedPoolName,
		evmstakingtypes.ModuleName,
		minttypes.ModuleName,
	}

	moduleAccPerms = []*authmodulev1.ModuleAccountPermission{
		{Account: authtypes.FeeCollectorName},
		{Account: distrtypes.ModuleName, Permissions: []string{authtypes.Burner}},
		{Account: stakingtypes.BondedPoolName, Permissions: []string{authtypes.Burner, authtypes.Staking}},
		{Account: stakingtypes.NotBondedPoolName, Permissions: []string{authtypes.Burner, authtypes.Staking}},
		{Account: evmstakingtypes.ModuleName, Permissions: []string{authtypes.Burner, authtypes.Minter}},
		{Account: minttypes.ModuleName, Permissions: []string{authtypes.Minter}},
	}

	// bankWrapperBindings returns a list of depinject.Configs that binds the withdraw.BankWrapper
	// to all x/bank Keeper interfaces.
	bankWrapperBindings = withdraw.BindInterfaces()

	// appConfig application configuration (used by depinject).
	appConfig = func(ctx context.Context, network netconf.ID) depinject.Config {
		return appconfig.Compose(&appv1alpha1.Config{
			Modules: func() []*appv1alpha1.ModuleConfig {
				return []*appv1alpha1.ModuleConfig{
					{
						Name: runtime.ModuleName,
						Config: appconfig.WrapAny(&runtimev1alpha1.Module{
							AppName:       Name,
							BeginBlockers: beginBlockers(ctx),
							PreBlockers:   []string{upgradetypes.ModuleName},
							// Setting endblockers in newApp since valsync replaces staking endblocker.
							InitGenesis: genesisModuleOrder,
							OverrideStoreKeys: []*runtimev1alpha1.StoreKeyConfig{
								{
									ModuleName: authtypes.ModuleName,
									KvStoreKey: "acc",
								},
							},
						}),
					},
					{
						Name: authtypes.ModuleName,
						Config: appconfig.WrapAny(&authmodulev1.Module{
							ModuleAccountPermissions: moduleAccPerms,
							Bech32Prefix:             sdk.Bech32HRP,
						}),
					},
					{
						Name: "tx",
						Config: appconfig.WrapAny(&txconfigv1.Config{
							SkipAnteHandler: true, // Disable ante handler (since we don't have proper txs).
							SkipPostHandler: true,
						}),
					},
					{
						Name: banktypes.ModuleName,
						Config: appconfig.WrapAny(&bankmodulev1.Module{
							BlockedModuleAccountsOverride: blockAccAddrs,
						}),
					},
					{
						Name:   consensustypes.ModuleName,
						Config: appconfig.WrapAny(&consensusmodulev1.Module{}),
					},
					{
						Name:   distrtypes.ModuleName,
						Config: appconfig.WrapAny(&distrmodulev1.Module{}),
					},
					{
						Name:   evmredenomtypes.ModuleName,
						Config: appconfig.WrapAny(&evmredenommodule.Module{}),
					},
					{
						Name:   genutiltypes.ModuleName,
						Config: appconfig.WrapAny(&genutilmodulev1.Module{}),
					},
					{
						Name:   stakingtypes.ModuleName,
						Config: appconfig.WrapAny(&stakingmodulev1.Module{}),
					},
					{
						Name:   slashingtypes.ModuleName,
						Config: appconfig.WrapAny(&slashingmodulev1.Module{}),
					},
					{
						Name:   evidencetypes.ModuleName,
						Config: appconfig.WrapAny(&evidencemodulev1.Module{}),
					},
					{
						Name: upgradetypes.ModuleName,
						Config: appconfig.WrapAny(&upgrademodulev1.Module{
							Authority: evmupgrade.ModuleName,
						}),
					},
					{
						Name: engevmtypes.ModuleName,
						Config: appconfig.WrapAny(&engevmmodule.Module{
							MaxWithdrawalsPerBlock: maxWithdrawalsPerBlock,
						}),
					},
					{
						Name: attesttypes.ModuleName,
						Config: appconfig.WrapAny(&attestmodule.Module{
							VoteWindowUp:       genesisVoteWindowUp,
							VoteWindowDown:     genesisVoteWindowDown,
							VoteExtensionLimit: genesisVoteExtLimit,
							TrimLag:            genesisTrimLag,
							ConsensusTrimLag:   genesisCTrimLag,
						}),
					},
					{
						Name:   valsynctypes.ModuleName,
						Config: appconfig.WrapAny(&valsyncmodule.Module{}),
					},
					{
						Name:   portaltypes.ModuleName,
						Config: appconfig.WrapAny(&portalmodule.Module{}),
					},
					{
						Name:   registrytypes.ModuleName,
						Config: appconfig.WrapAny(&registrymodule.Module{}),
					},
					{
						Name: evmstakingtypes.ModuleName,
						Config: appconfig.WrapAny(&evmstakingmodule.Module{
							DeliverInterval: deliverInterval(network),
						}),
					},
					{
						Name:   minttypes.ModuleName,
						Config: appconfig.WrapAny(&mintmodulev1.Module{}),
					},
				}
			}(),
		})
	}

	// diInvokers defines a list of depinject invoke functions.
	// These are non-cosmos-module invokers used in halo's app wiring.
	diInvokers = []any{
		withdraw.DIInvoke,
		evmredenommodule.DIInvoke,
	}

	// diProviders defines a list of depinject provider functions.
	// These are non-cosmos-module providers used in halo's app wiring.
	diProviders = []any{
		evmslashing.DIProvide,
		evmupgrade.DIProvide,
		evmdistribution.DIProvide,
		withdraw.DIProvide,
	}
)

func deliverInterval(network netconf.ID) int64 {
	if network.IsProtected() {
		return deliverIntervalProtected
	}

	return deliverIntervalEphemeral
}

type EncodingConfig struct {
	InterfaceRegistry types.InterfaceRegistry
	Codec             codec.Codec
	TxConfig          client.TxConfig
}

// ClientEncodingConfig returns the client encoding configuration for the given network.
// Note this should not be used by halo itself. It mocks dependencies to load modules
// only in order to register interfaces and protos for encoding/decoding, not for actual logic/state.
func ClientEncodingConfig(ctx context.Context, network netconf.ID) (EncodingConfig, error) {
	noopVoter, err := newVoterLoader(k1.GenPrivKey())
	if err != nil {
		return EncodingConfig{}, err
	}

	engineCl := struct {
		ethclient.EngineClient
	}{}

	depCfg := depinject.Configs(
		appConfig(ctx, network),
		depinject.Provide(diProviders...),
		depinject.Configs(bankWrapperBindings...),
		depinject.Invoke(diInvokers...),
		depinject.Supply(
			newSDKLogger(ctx),
			attesttypes.ChainVerNameFunc(netconf.ChainVersionNamer(network)),
			registrytypes.ChainNameFunc(netconf.ChainNamer(network)),
			noopVoter,
			engineCl,
			engevmtypes.FeeRecipientProvider(burnEVMFees{}),
		),
	)

	var resp EncodingConfig

	if err := depinject.Inject(depCfg, []any{
		&resp.InterfaceRegistry,
		&resp.Codec,
		&resp.TxConfig,
	}...); err != nil {
		return EncodingConfig{}, errors.Wrap(err, "dep inject")
	}

	return resp, nil
}
