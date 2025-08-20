package admin

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// bridgeSpec defines the bridge spec for each network.
// To update bridge spec, update this map.
// Then run `ensure-bridge-spec` to apply the changes.
var bridgeSpec = map[netconf.ID]NetworkBridgeSpec{
	netconf.Devnet:  DefaultBridgeSpec(),
	netconf.Staging: DefaultBridgeSpec(),
	netconf.Omega: {
		Native: BridgeSpec{
			PauseAll:      false,
			PauseWithdraw: false,
			PauseBridge:   true,
		},
		L1: BridgeSpec{
			PauseAll:      false,
			PauseWithdraw: false,
			PauseBridge:   true,
		},
	},
	netconf.Mainnet: {
		Native: BridgeSpec{
			PauseAll:      false,
			PauseWithdraw: false,
			PauseBridge:   true,
		},
		L1: BridgeSpec{
			PauseAll:      false,
			PauseWithdraw: false,
			PauseBridge:   true,
		},
	},
}

// BridgeSpec is the specification for a bridge contract (native or L1).
type BridgeSpec struct {
	PauseAll      bool
	PauseWithdraw bool
	PauseBridge   bool
}

// BridgeDirectives define updates required for a bridge contract to match the spec.
type BridgeDirectives struct {
	PauseAll        bool
	UnpauseAll      bool
	PauseWithdraw   bool
	UnpauseWithdraw bool
	PauseBridge     bool
	UnpauseBridge   bool
}

// NetworkBridgeSpec defines the bridge spec for a network, both native and L1.
type NetworkBridgeSpec struct {
	Native BridgeSpec
	L1     BridgeSpec
}

// NetworkBridgeDirectives defines the bridge directives for a network, both native and L1.
type NetworkBridgeDirectives struct {
	Native BridgeDirectives
	L1     BridgeDirectives
}

func DefaultBridgeSpec() NetworkBridgeSpec {
	return NetworkBridgeSpec{
		Native: BridgeSpec{
			PauseAll:      false,
			PauseWithdraw: false,
			PauseBridge:   false,
		},
		L1: BridgeSpec{
			PauseAll:      false,
			PauseWithdraw: false,
			PauseBridge:   false,
		},
	}
}

// Verify checks that there are no conflicting specifications.
func (s BridgeSpec) Verify() error {
	if s.PauseAll && (s.PauseWithdraw && s.PauseBridge) {
		return errors.New("if pause all, no other pause flags should be set")
	}

	return nil
}

// EnsureBridgeSpec ensures that live bridge contracts are configured as per the local spec.
func EnsureBridgeSpec(ctx context.Context, def app.Definition, cfg Config, specOverride *NetworkBridgeSpec) error {
	s := setup(def, cfg)

	l1Chain, ok := s.testnet.EthereumChain()
	if !ok {
		return errors.New("no ethereum chain")
	}

	var nativeSpecOverride *BridgeSpec
	if specOverride != nil {
		nativeSpecOverride = &specOverride.Native
	}

	var l1SpecOverride *BridgeSpec
	if specOverride != nil {
		l1SpecOverride = &specOverride.L1
	}

	omniEVM, err := setupChain(ctx, s, omniEVMName)
	if err != nil {
		return errors.Wrap(err, "setup omni evm")
	}

	if err := ensureNativeBridgeSpec(ctx, s, omniEVM, nativeSpecOverride); err != nil {
		return errors.Wrap(err, "ensure native bridge spec")
	}

	l1, err := setupChain(ctx, s, l1Chain.Name)
	if err != nil {
		return errors.Wrap(err, "setup l1")
	}

	if err := ensureL1BridgeSpec(ctx, s, l1, l1SpecOverride); err != nil {
		return errors.Wrap(err, "ensure l1 bridge spec")
	}

	return nil
}

// ensureNativeBridgeSpec ensures that the live native bridge contract is configured as per the local spec.
func ensureNativeBridgeSpec(ctx context.Context, s shared, c chain, specOverride *BridgeSpec) error {
	local := bridgeSpec[s.testnet.Network].Native
	if specOverride != nil {
		local = *specOverride
	}

	ethCl, err := ethclient.DialContext(ctx, c.Name, c.RPCEndpoint)
	if err != nil {
		return errors.Wrap(err, "dial eth client")
	}

	addr := common.HexToAddress(predeploys.OmniBridgeNative)

	contract, err := bindings.NewOmniBridgeNative(addr, ethCl)
	if err != nil {
		return errors.Wrap(err, "new omni bridge native")
	}

	live, err := liveBridgeSpec(ctx, contract)
	if err != nil {
		return errors.Wrap(err, "live native spec")
	}

	directives, err := makeBridgeDirectives(local, live)
	if err != nil {
		return errors.Wrap(err, "make bridge directives")
	}

	return runBridgeDirectives(ctx, s, c, addr, contract, directives)
}

// ensureL1BridgeSpec ensures that the live L1 bridge contract is configured as per the local spec.
func ensureL1BridgeSpec(ctx context.Context, s shared, c chain, specOverride *BridgeSpec) error {
	local := bridgeSpec[s.testnet.Network].L1
	if specOverride != nil {
		local = *specOverride
	}

	ethCl, err := ethclient.DialContext(ctx, c.Name, c.RPCEndpoint)
	if err != nil {
		return errors.Wrap(err, "dial eth client")
	}

	addrs, err := contracts.GetAddresses(ctx, s.testnet.Network)
	if err != nil {
		return errors.Wrap(err, "get addrs")
	}

	contract, err := bindings.NewOmniBridgeL1(addrs.L1Bridge, ethCl)
	if err != nil {
		return errors.Wrap(err, "new omni bridge l1")
	}

	live, err := liveBridgeSpec(ctx, contract)
	if err != nil {
		return errors.Wrap(err, "live l1 spec")
	}

	directives, err := makeBridgeDirectives(local, live)
	if err != nil {
		return errors.Wrap(err, "make bridge directives")
	}

	return runBridgeDirectives(ctx, s, c, addrs.L1Bridge, contract, directives)
}

// BridgeCommon is a common interface between native and L1 bridge contracts. It lists a subset of OmniBridgeCommon views.
type BridgeCommon interface {
	// KeyPauseAll returns the key for pausing all actions (PauseableUpgradeable.KeyPauseAll)
	KeyPauseAll(opts *bind.CallOpts) ([32]byte, error)

	// ACTIONWITHDRAW returns the key for pausing withdraws (OmniBridgeCommon.ACTION_WITHDRAW)
	ACTIONWITHDRAW(opts *bind.CallOpts) ([32]byte, error)

	// ACTIONBRIDGE returns the key for pausing bridge actions (OmniBridgeCommon.ACTION_BRIDGE)
	ACTIONBRIDGE(opts *bind.CallOpts) ([32]byte, error)

	// IsPaused returns true if the action is paused
	IsPaused(opts *bind.CallOpts, action [32]byte) (bool, error)
}

// liveBridgeSpec returns the live spec for a bridge contract.
func liveBridgeSpec(ctx context.Context, contract BridgeCommon) (BridgeSpec, error) {
	callopts := &bind.CallOpts{Context: ctx}

	keys, err := getBridgePauseKeys(ctx, contract)
	if err != nil {
		return BridgeSpec{}, errors.Wrap(err, "get pause keys")
	}

	paused, err := contract.IsPaused(callopts, keys.all)
	if err != nil {
		return BridgeSpec{}, errors.Wrap(err, "is paused all")
	}

	if paused {
		return BridgeSpec{PauseAll: true}, nil
	}

	withdrawPaused, err := contract.IsPaused(callopts, keys.withdraw)
	if err != nil {
		return BridgeSpec{}, errors.Wrap(err, "is paused withdraw")
	}

	bridgePaused, err := contract.IsPaused(callopts, keys.bridge)
	if err != nil {
		return BridgeSpec{}, errors.Wrap(err, "is paused bridge")
	}

	return BridgeSpec{
		PauseAll:      false,
		PauseWithdraw: withdrawPaused,
		PauseBridge:   bridgePaused,
	}, nil
}

// makeBridgeDirectives returns directives necessary to match live spec with local.
func makeBridgeDirectives(local, live BridgeSpec) (BridgeDirectives, error) {
	if err := local.Verify(); err != nil {
		return BridgeDirectives{}, errors.Wrap(err, "verify local spec")
	}

	if err := live.Verify(); err != nil {
		return BridgeDirectives{}, errors.Wrap(err, "verify live spec")
	}

	return BridgeDirectives{
		PauseAll:        local.PauseAll && !live.PauseAll,
		UnpauseAll:      !local.PauseAll && live.PauseAll,
		PauseWithdraw:   local.PauseWithdraw && !live.PauseWithdraw,
		UnpauseWithdraw: !local.PauseWithdraw && live.PauseWithdraw,
		PauseBridge:     local.PauseBridge && !live.PauseBridge,
		UnpauseBridge:   !local.PauseBridge && live.PauseBridge,
	}, nil
}

// define labels for pause actions, used for logging.
const (
	labelAll      = "all"
	labelWithdraw = "withdraw"
	labelBridge   = "bridge"
)

// runBridgeDirectives applies directives to a bridge contract.
func runBridgeDirectives(
	ctx context.Context,
	s shared,
	c chain,
	addr common.Address,
	contract BridgeCommon,
	directives BridgeDirectives,
) error {
	keys, err := getBridgePauseKeys(ctx, contract)
	if err != nil {
		return errors.Wrap(err, "get pause keys")
	}

	if directives.PauseAll {
		if err := pauseBridge(ctx, s, c, addr, keys.all, labelAll); err != nil {
			return errors.Wrap(err, "pause all")
		}
	}

	if directives.UnpauseAll {
		if err := unpauseBridge(ctx, s, c, addr, keys.all, labelAll); err != nil {
			return errors.Wrap(err, "unpause all")
		}
	}

	if directives.PauseWithdraw {
		if err := pauseBridge(ctx, s, c, addr, keys.withdraw, labelWithdraw); err != nil {
			return errors.Wrap(err, "pause withdraw")
		}
	}

	if directives.UnpauseWithdraw {
		if err := unpauseBridge(ctx, s, c, addr, keys.withdraw, labelWithdraw); err != nil {
			return errors.Wrap(err, "unpause withdraw")
		}
	}

	if directives.PauseBridge {
		if err := pauseBridge(ctx, s, c, addr, keys.bridge, labelBridge); err != nil {
			return errors.Wrap(err, "pause bridge")
		}
	}

	if directives.UnpauseBridge {
		if err := unpauseBridge(ctx, s, c, addr, keys.bridge, labelBridge); err != nil {
			return errors.Wrap(err, "unpause bridge")
		}
	}

	return nil
}

type bridgePauseKeys struct {
	all      [32]byte
	withdraw [32]byte
	bridge   [32]byte
}

// getBridgePauseKeys returns the pause keys for a bridge contract.
func getBridgePauseKeys(ctx context.Context, contract BridgeCommon) (bridgePauseKeys, error) {
	callopts := &bind.CallOpts{Context: ctx}

	all, err := contract.KeyPauseAll(callopts)
	if err != nil {
		return bridgePauseKeys{}, errors.Wrap(err, "key pause all")
	}

	withdraw, err := contract.ACTIONWITHDRAW(callopts)
	if err != nil {
		return bridgePauseKeys{}, errors.Wrap(err, "key pause withdraw")
	}

	bridge, err := contract.ACTIONBRIDGE(callopts)
	if err != nil {
		return bridgePauseKeys{}, errors.Wrap(err, "key pause bridge")
	}

	// sanity check
	if all == withdraw || all == bridge || withdraw == bridge {
		return bridgePauseKeys{}, errors.New("pause keys are not unique")
	}

	return bridgePauseKeys{
		all:      all,
		withdraw: withdraw,
		bridge:   bridge,
	}, nil
}
