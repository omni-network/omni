package admin

import (
	"context"
	"reflect"
	"slices"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// portalSpec defines portal specs per network, with chain specific overrides.
// To update portal spec, update this map.
// Then run `ensure-portal-spec` to apply the changes.
var portalSpec = map[netconf.ID]NetworkPortalSpec{
	netconf.Devnet:  {Global: DefaultPortalSpec()},
	netconf.Staging: {Global: DefaultPortalSpec()},
	netconf.Omega:   {Global: DefaultPortalSpec()},
	netconf.Mainnet: {Global: DefaultPortalSpec()},
}

// PortalSpec is the specification for the OmniPortal contract.
type PortalSpec struct {
	// PauseAll indicates that all actions on the portal should be paused.
	PauseAll bool

	// PauseXCall indicates that all xcalls should be paused.
	PauseXCall bool

	// PauseXCallTo indicates that xcalls to specific chains should be paused.
	PauseXCallTo []uint64

	// PauseXSubmit indicates that all xsubmits should be paused.
	PauseXSubmit bool

	// PauseXSubmitFrom indicates that xsubmits from specific chains should be paused.
	PauseXSubmitFrom []uint64
}

// PortalDirectives defines updates required on a portal to match a spec.
type PortalDirectives struct {
	// PauseAll indicates that all actions on the portal should be paused.
	PauseAll bool

	// UnpauseAll indicates that all actions on the portal should be unpaused.
	UnpauseAll bool

	// PauseXCall indicates that all xcalls should be paused.
	PauseXCall bool

	// UnpauseXCall indicates that all xcalls should be unpaused.
	UnpauseXCall bool

	// PauseXCallTo indicates that xcalls to specific chains should be paused.
	PauseXCallTo []uint64

	// UnpauseXCallTo indicates that xcalls to specific chains should be unpaused.
	UnpauseXCallTo []uint64

	// PauseXSubmit indicates that all xsubmits should be paused.
	PauseXSubmit bool

	// UnpauseXSubmit indicates that all xsubmits should be unpaused.
	UnpauseXSubmit bool

	// PauseXSubmitFrom indicates that xsubmits from specific chains should be paused.
	PauseXSubmitFrom []uint64

	// UnpauseXSubmitFrom indicates that xsubmits from specific chains should be unpaused.
	UnpauseXSubmitFrom []uint64
}

type NetworkPortalSpec struct {
	// Global is the portal spec, maintained on all chains.
	Global PortalSpec

	// ChainOverrides overrides the global spec for specific chains. Must specify full spec, not just diff.
	ChainOverrides map[uint64]*PortalSpec
}

func DefaultPortalSpec() PortalSpec {
	return PortalSpec{
		PauseAll:         false,
		PauseXCall:       false,
		PauseXCallTo:     nil,
		PauseXSubmit:     false,
		PauseXSubmitFrom: nil,
	}
}

// EnsurePortalSpec ensures that live portal contracts are configured as per the local spec.
func EnsurePortalSpec(ctx context.Context, def app.Definition, cfg Config, localSpecOverride *PortalSpec) error {
	return setup(def, cfg).run(ctx, func(ctx context.Context, s shared, c chain) error {
		local, err := localPortalSpec(s.testnet.Network, c.ChainID)
		if err != nil {
			return errors.Wrap(err, "get local portal spec", "chain", c.Name)
		}

		if localSpecOverride != nil {
			local = *localSpecOverride
		}

		ethCl, err := ethclient.DialContext(ctx, c.Name, c.RPCEndpoint)
		if err != nil {
			return errors.Wrap(err, "dial eth client", "chain", c.Name)
		}

		live, err := livePortalSpec(ctx, s.testnet.EVMChains(), c.EVMChain, c.PortalAddress, ethCl)
		if err != nil {
			return errors.Wrap(err, "get live portal spec", "chain", c.Name)
		}

		directives, err := makePortalDirectives(local, live)
		if err != nil {
			return errors.Wrap(err, "make directives", "chain", c.Name)
		}

		return runPortalDirectives(ctx, s, c, directives)
	})
}

// makePortalDirectives returns the directives required to bring a portal in line with the local spec.
func makePortalDirectives(local, live PortalSpec) (PortalDirectives, error) {
	if err := local.Verify(); err != nil {
		return PortalDirectives{}, errors.Wrap(err, "verify local spec")
	}

	if err := live.Verify(); err != nil {
		return PortalDirectives{}, errors.Wrap(err, "verify live spec")
	}

	pauseAll := local.PauseAll && !live.PauseAll
	unpauseAll := !local.PauseAll && live.PauseAll
	pauseXCall := local.PauseXCall && !live.PauseXCall
	unpauseXCall := !local.PauseXCall && live.PauseXCall
	pauseXSubmit := local.PauseXSubmit && !live.PauseXSubmit
	unpauseXSubmit := !local.PauseXSubmit && live.PauseXSubmit

	var pauseXCallTo []uint64
	var unpauseXCallTo []uint64
	var pauseXSumbitFrom []uint64
	var unpauseXSubmitFrom []uint64

	for _, chain := range local.PauseXCallTo {
		if !contains(live.PauseXCallTo, chain) {
			pauseXCallTo = append(pauseXCallTo, chain)
		}
	}

	for _, chain := range live.PauseXCallTo {
		if !contains(local.PauseXCallTo, chain) {
			unpauseXCallTo = append(unpauseXCallTo, chain)
		}
	}

	for _, chain := range local.PauseXSubmitFrom {
		if !contains(live.PauseXSubmitFrom, chain) {
			pauseXSumbitFrom = append(pauseXSumbitFrom, chain)
		}
	}

	for _, chain := range live.PauseXSubmitFrom {
		if !contains(local.PauseXSubmitFrom, chain) {
			unpauseXSubmitFrom = append(unpauseXSubmitFrom, chain)
		}
	}

	return PortalDirectives{
		PauseAll:           pauseAll,
		UnpauseAll:         unpauseAll,
		PauseXCall:         pauseXCall,
		UnpauseXCall:       unpauseXCall,
		PauseXSubmit:       pauseXSubmit,
		UnpauseXSubmit:     unpauseXSubmit,
		PauseXCallTo:       pauseXCallTo,
		UnpauseXCallTo:     unpauseXCallTo,
		PauseXSubmitFrom:   pauseXSumbitFrom,
		UnpauseXSubmitFrom: unpauseXSubmitFrom,
	}, nil
}

// Verify checks that there are no conflicting specifications.
func (s PortalSpec) Verify() error {
	if s.PauseAll && (s.PauseXCall || s.PauseXSubmit || len(s.PauseXCallTo) > 0 || len(s.PauseXSubmitFrom) > 0) {
		return errors.New("if pause-all, no other actions should be specified")
	}

	if s.PauseXCall && len(s.PauseXCallTo) > 0 {
		return errors.New("if pause-xcall, pause-xcall-to should be empty")
	}

	if s.PauseXSubmit && len(s.PauseXSubmitFrom) > 0 {
		return errors.New("if pause-xsubmit, pause-xsubmit-from should be empty")
	}

	return nil
}

// Verify checks that there are no conflicting directives.
func (d PortalDirectives) Verify() error {
	if d.PauseAll && d.UnpauseAll {
		return errors.New("cannot pause and unpause all")
	}

	if d.PauseAll && (d.PauseXCall || d.PauseXSubmit || len(d.PauseXCallTo) > 0 || len(d.PauseXSubmitFrom) > 0) {
		return errors.New("if pause-all, no other actions should be specified")
	}

	if d.PauseXCall && d.UnpauseXCall {
		return errors.New("cannot pause and unpause xcall")
	}

	if d.PauseXCall && len(d.PauseXCallTo) > 0 {
		return errors.New("if pause-xcall, pause-xcall-to should be empty")
	}

	if d.PauseXSubmit && d.UnpauseXSubmit {
		return errors.New("cannot pause and unpause xsubmit")
	}

	if d.PauseXSubmit && len(d.PauseXSubmitFrom) > 0 {
		return errors.New("if pause-xsubmit, pause-xsubmit-from should be empty")
	}

	for _, chain := range d.PauseXCallTo {
		if contains(d.UnpauseXCallTo, chain) {
			return errors.New("cannot pause and unpause xcall to same chain", "chain", chain)
		}
	}

	for _, chain := range d.PauseXSubmitFrom {
		if contains(d.UnpauseXSubmitFrom, chain) {
			return errors.New("cannot pause and unpause xsubmit from same chain", "chain", chain)
		}
	}

	return nil
}

// localPortalSpec returns the configured, local portal spec for a chain.
func localPortalSpec(network netconf.ID, chainID uint64) (PortalSpec, error) {
	net, ok := portalSpec[network]
	if !ok {
		return PortalSpec{}, errors.New("no spec for network", "network", network)
	}

	spec := net.Global
	if net.ChainOverrides != nil && net.ChainOverrides[chainID] != nil {
		spec = *net.ChainOverrides[chainID]
	}

	if err := spec.Verify(); err != nil {
		return PortalSpec{}, errors.Wrap(err, "verify spec", "network", network, "chain", chainID)
	}

	return spec, nil
}

// livePortalSpec returns the live, on-chain portal spec for a chain.
func livePortalSpec(ctx context.Context, chains []types.EVMChain, c types.EVMChain, portalAddr common.Address, ethCl ethclient.Client) (PortalSpec, error) {
	portal, err := bindings.NewOmniPortal(portalAddr, ethCl)
	if err != nil {
		return PortalSpec{}, errors.Wrap(err, "new portal contract", "chain", c.Name)
	}

	log.Info(ctx, "Fetching portal spec", "chain", c.Name, "address", portalAddr)

	paused, err := portal.IsPaused1(&bind.CallOpts{Context: ctx})
	if err != nil {
		return PortalSpec{}, errors.Wrap(err, "is paused", "chain", c.Name)
	}

	// If the portal is paused, we don't need to check further.
	if paused {
		return PortalSpec{PauseAll: true}, nil
	}

	actionXSubmit, err := portal.ActionXSubmit(&bind.CallOpts{Context: ctx})
	if err != nil {
		return PortalSpec{}, errors.Wrap(err, "action xsubmit", "chain", c.Name)
	}

	actionXCall, err := portal.ActionXCall(&bind.CallOpts{Context: ctx})
	if err != nil {
		return PortalSpec{}, errors.Wrap(err, "action xcall", "chain", c.Name)
	}

	xcallPaused, err := portal.IsPaused(&bind.CallOpts{Context: ctx}, actionXCall)
	if err != nil {
		return PortalSpec{}, errors.Wrap(err, "is xcall paused", "chain", c.Name)
	}

	xsubmitPaused, err := portal.IsPaused(&bind.CallOpts{Context: ctx}, actionXSubmit)
	if err != nil {
		return PortalSpec{}, errors.Wrap(err, "is xsubmit paused", "chain", c.Name)
	}

	var xCallPausedTo []uint64
	var xSubmitPausedFrom []uint64

	for _, chain := range chains {
		isXSubmitPausedFrom, err := portal.IsPaused0(&bind.CallOpts{Context: ctx}, actionXSubmit, chain.ChainID)
		if err != nil {
			return PortalSpec{}, errors.Wrap(err, "is xsubmit paused from", "chain", c.Name, "from", chain.Name)
		}

		isXCallPausedTo, err := portal.IsPaused0(&bind.CallOpts{Context: ctx}, actionXCall, chain.ChainID)
		if err != nil {
			return PortalSpec{}, errors.Wrap(err, "is xcall paused to", "chain", c.Name, "to", chain.Name)
		}

		if isXCallPausedTo {
			xCallPausedTo = append(xCallPausedTo, chain.ChainID)
		}

		if isXSubmitPausedFrom {
			xSubmitPausedFrom = append(xSubmitPausedFrom, chain.ChainID)
		}
	}

	spec := PortalSpec{
		PauseXCall:   xcallPaused,
		PauseXSubmit: xsubmitPaused,
	}

	// only specify PauseXSubmitFrom if PauseXSubmit==false
	if !xsubmitPaused && len(xSubmitPausedFrom) > 0 {
		spec.PauseXSubmitFrom = xSubmitPausedFrom
	}

	// only specify PauseXCallTo if PauseXCall==false
	if !xcallPaused && len(xCallPausedTo) > 0 {
		spec.PauseXCallTo = xCallPausedTo
	}

	return spec, nil
}

// runPortalDirectives applics PortalDirectives on chain.
func runPortalDirectives(ctx context.Context, s shared, c chain, directives PortalDirectives) error {
	if err := directives.Verify(); err != nil {
		return errors.Wrap(err, "verify directives", "chain", c.Name)
	}

	if isEmpty(directives) {
		log.Info(ctx, "No directives to apply", "chain", c.Name)
		return nil
	}

	if directives.PauseAll {
		err := pausePortal(ctx, s, c)
		if err != nil {
			return errors.Wrap(err, "pause portal", "chain", c.Name)
		}
	}

	if directives.UnpauseAll {
		err := unpausePortal(ctx, s, c)
		if err != nil {
			return errors.Wrap(err, "unpause portal", "chain", c.Name)
		}
	}

	if directives.PauseXCall {
		err := pauseXCall(ctx, s, c)
		if err != nil {
			return errors.Wrap(err, "pause xcall", "chain", c.Name)
		}
	}

	if directives.UnpauseXCall {
		err := unpauseXCall(ctx, s, c)
		if err != nil {
			return errors.Wrap(err, "unpause xcall", "chain", c.Name)
		}
	}

	for _, to := range directives.PauseXCallTo {
		err := pauseXCallTo(ctx, s, c, to)
		if err != nil {
			return errors.Wrap(err, "pause xcall to", "chain", c.Name, "to", to)
		}
	}

	for _, to := range directives.UnpauseXCallTo {
		err := unpauseXCallTo(ctx, s, c, to)
		if err != nil {
			return errors.Wrap(err, "unpause xcall to", "chain", c.Name, "to", to)
		}
	}

	if directives.PauseXSubmit {
		err := pauseXSubmit(ctx, s, c)
		if err != nil {
			return errors.Wrap(err, "pause xsubmit", "chain", c.Name)
		}
	}

	if directives.UnpauseXSubmit {
		err := unpauseXSubmit(ctx, s, c)
		if err != nil {
			return errors.Wrap(err, "unpause xsubmit", "chain", c.Name)
		}
	}

	for _, from := range directives.PauseXSubmitFrom {
		err := pauseXSubmitFrom(ctx, s, c, from)
		if err != nil {
			return errors.Wrap(err, "pause xsubmit from", "chain", c.Name, "from", from)
		}
	}

	for _, from := range directives.UnpauseXSubmitFrom {
		err := unpauseXSubmitFrom(ctx, s, c, from)
		if err != nil {
			return errors.Wrap(err, "unpause xsubmit from", "chain", c.Name, "from", from)
		}
	}

	return nil
}

func isEmpty(v any) bool {
	rv := reflect.ValueOf(v)

	// If the value is invalid (e.g., nil interface), return true
	if !rv.IsValid() {
		return true
	}

	// Report len 0 slices, maps, and channels as empty
	switch rv.Kind() {
	case reflect.Slice, reflect.Map, reflect.Chan:
		return rv.Len() == 0
	default:
		return rv.IsZero()
	}
}

func contains[T comparable](ts []T, t T) bool {
	return slices.Contains(ts, t)
}
