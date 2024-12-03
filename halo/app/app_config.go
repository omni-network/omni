package app

import (
	"context"

	attestmodule "github.com/omni-network/omni/halo/attest/module"
	attesttypes "github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/halo/evmslashing"
	"github.com/omni-network/omni/halo/evmstaking"
	evmstaking2module "github.com/omni-network/omni/halo/evmstaking2/module"
	evmstaking2types "github.com/omni-network/omni/halo/evmstaking2/types"
	"github.com/omni-network/omni/halo/evmupgrade"
	portalmodule "github.com/omni-network/omni/halo/portal/module"
	portaltypes "github.com/omni-network/omni/halo/portal/types"
	registrymodule "github.com/omni-network/omni/halo/registry/module"
	registrytypes "github.com/omni-network/omni/halo/registry/types"
	"github.com/omni-network/omni/halo/sdk"
	valsyncmodule "github.com/omni-network/omni/halo/valsync/module"
	valsynctypes "github.com/omni-network/omni/halo/valsync/types"
	"github.com/omni-network/omni/lib/feature"
	engevmmodule "github.com/omni-network/omni/octane/evmengine/module"
	engevmtypes "github.com/omni-network/omni/octane/evmengine/types"

	runtimev1alpha1 "cosmossdk.io/api/cosmos/app/runtime/v1alpha1"
	appv1alpha1 "cosmossdk.io/api/cosmos/app/v1alpha1"
	authmodulev1 "cosmossdk.io/api/cosmos/auth/module/v1"
	bankmodulev1 "cosmossdk.io/api/cosmos/bank/module/v1"
	consensusmodulev1 "cosmossdk.io/api/cosmos/consensus/module/v1"
	distrmodulev1 "cosmossdk.io/api/cosmos/distribution/module/v1"
	evidencemodulev1 "cosmossdk.io/api/cosmos/evidence/module/v1"
	genutilmodulev1 "cosmossdk.io/api/cosmos/genutil/module/v1"
	slashingmodulev1 "cosmossdk.io/api/cosmos/slashing/module/v1"
	stakingmodulev1 "cosmossdk.io/api/cosmos/staking/module/v1"
	txconfigv1 "cosmossdk.io/api/cosmos/tx/config/v1"
	upgrademodulev1 "cosmossdk.io/api/cosmos/upgrade/module/v1"
	"cosmossdk.io/core/appconfig"
	"cosmossdk.io/depinject"
	evidencetypes "cosmossdk.io/x/evidence/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

const (
	// TODO(corver): Maybe move these to genesis itself.
	genesisVoteWindowUp   uint64 = 64 // Allow early votes for <latest attestation - 64>
	genesisVoteWindowDown uint64 = 2  // Only allow late votes for <latest attestation - 2>
	genesisVoteExtLimit   uint64 = 256
	genesisTrimLag        uint64 = 1      // Allow deleting attestations in block after approval.
	genesisCTrimLag       uint64 = 72_000 // Delete consensus attestations state after +-1 day (given a period of 1.2s).
)

//nolint:gochecknoglobals // Cosmos-style
var (
	genesisModuleOrder = []string{
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		stakingtypes.ModuleName,
		slashingtypes.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		upgradetypes.ModuleName,
		valsynctypes.ModuleName,
		engevmtypes.ModuleName,
	}

	beginBlockers = []string{
		distrtypes.ModuleName, // Note: slashing happens after distr.BeginBlocker
		slashingtypes.ModuleName,
		stakingtypes.ModuleName, // Note: staking module is required if HistoricalEntries param > 0
		evidencetypes.ModuleName,
		attesttypes.ModuleName,
	}

	endBlockers = []string{
		attesttypes.ModuleName,
		valsynctypes.ModuleName, // Wraps staking module end blocker (must come after attest module)
		upgradetypes.ModuleName,
	}

	// blocked account addresses.
	blockAccAddrs = []string{
		authtypes.FeeCollectorName,
		distrtypes.ModuleName,
		stakingtypes.BondedPoolName,
		stakingtypes.NotBondedPoolName,
		// TODO(christian): rename package, the rest can stay because names are the same
		evmstaking.ModuleName,
	}

	moduleAccPerms = []*authmodulev1.ModuleAccountPermission{
		{Account: authtypes.FeeCollectorName},
		{Account: distrtypes.ModuleName},
		{Account: stakingtypes.BondedPoolName, Permissions: []string{authtypes.Burner, stakingtypes.ModuleName}},
		{Account: stakingtypes.NotBondedPoolName, Permissions: []string{authtypes.Burner, stakingtypes.ModuleName}},
		// TODO(christian): rename package, the rest can stay because names are the same
		{Account: evmstaking.ModuleName, Permissions: []string{authtypes.Burner, authtypes.Minter}},
	}

	// appConfig application configuration (used by depinject).
	appConfig = func(ctx context.Context) depinject.Config {
		return appconfig.Compose(&appv1alpha1.Config{
			Modules: func() []*appv1alpha1.ModuleConfig {
				configs := []*appv1alpha1.ModuleConfig{
					{
						Name: runtime.ModuleName,
						Config: appconfig.WrapAny(&runtimev1alpha1.Module{
							AppName:       Name,
							BeginBlockers: beginBlockers,
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
						Name:   engevmtypes.ModuleName,
						Config: appconfig.WrapAny(&engevmmodule.Module{}),
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
				}

				// TODO(christian): integrate into the list above
				if feature.FlagEVMStakingModule.Enabled(ctx) {
					configs = append(configs, &appv1alpha1.ModuleConfig{
						Name:   evmstaking2types.ModuleName,
						Config: appconfig.WrapAny(&evmstaking2module.Module{}),
					})
				}

				return configs
			}(),
		})
	}

	// diProviders defines a list of depinject provider functions.
	// These are non-cosmos module constructors used in halo's app wiring.
	diProviders = []any{
		evmslashing.DIProvide,
		// TODO(christian): remove later, but seems like it can stay here even if feature is enabled
		evmstaking.DIProvide,
		evmupgrade.DIProvide,
	}
)
