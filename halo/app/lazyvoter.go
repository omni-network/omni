package app

import (
	"context"
	"sync"
	"time"

	atypes "github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/halo/attest/voter"
	"github.com/omni-network/omni/halo/comet"
	vtypes "github.com/omni-network/omni/halo/valsync/types"
	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
	xprovider "github.com/omni-network/omni/lib/xchain/provider"

	"github.com/cometbft/cometbft/crypto"

	"github.com/ethereum/go-ethereum/common"
)

var _ atypes.Voter = (*voterLoader)(nil)

var _ atypes.VoterDeps = voteDeps{}

type voteDeps struct {
	comet.API
	cchain.Provider
}

// voterLoader wraps a voter instances that is lazy loaded from the on-chain registry.
// It is basically a noop while not loaded.
type voterLoader struct {
	mu         sync.Mutex
	voter      *voter.Voter
	proposed   map[xchain.AttestHeader]*atypes.AttestHeader // Dedup cached headers, since cache contains multiple rounds of data.
	committed  map[xchain.AttestHeader]*atypes.AttestHeader
	lastValSet *vtypes.ValidatorSetResponse
	isVal      bool
	localAddr  common.Address
}

func newVoterLoader(privKey crypto.PrivKey) (*voterLoader, error) {
	localAddr, err := k1util.PubKeyToAddress(privKey.PubKey())
	if err != nil {
		return nil, err
	}

	return &voterLoader{
		localAddr: localAddr,
		proposed:  make(map[xchain.AttestHeader]*atypes.AttestHeader),
		committed: make(map[xchain.AttestHeader]*atypes.AttestHeader),
	}, nil
}

// LazyLoad blocks until the network config can be loaded from the on-chain registry, then it initializes and starts
// the voter instance and binds it to the lazy wrapper.
//
//nolint:nestif // 2 levels is not that bad
func (l *voterLoader) LazyLoad(
	ctx context.Context,
	netID netconf.ID,
	omniEVMCl ethclient.Client,
	endpoints xchain.RPCEndpoints,
	cprov cchain.Provider,
	privKey crypto.PrivKey,
	voterStateFile string,
	cmtAPI comet.API,
	asyncAbort chan<- error,
) error {
	if len(endpoints) == 0 {
		log.Warn(ctx, "Flag --xchain-evm-rpc-endpoints empty. The app will crash if it becomes a validator since it cannot perform xchain voting duties", nil)
	}

	if !l.isValidator() {
		log.Info(ctx, "Local halo node is not a validator")
	}

	// Wait until this node becomes a validator before initializing voter.
	// This mitigates crashes due to invalid rpc endpoint config in non-validator nodes.
	backoff := expbackoff.New(ctx, expbackoff.WithPeriodicConfig(time.Second))
	for !l.isValidator() {
		backoff()
		if ctx.Err() != nil {
			return errors.Wrap(ctx.Err(), "lazy loading canceled")
		}
	}

	log.Info(ctx, "🫡 Local halo node is a validator, starting voter")

	if len(endpoints) == 0 {
		// Note that this negatively affects chain liveness, but xchain liveness already negatively affected so rather
		// highlight the issue to the operator by crashing. #allornothing
		return errors.New("flag --xchain-evm-rpc-endpoints empty so cannot perform xchain voting duties")
	}

	// Use the RPCEndpoints config as the list of expected chains to load from the registry.
	// This is required for fresh genesis chains since portals are registered one at a time.
	// So netconf.AwaitOnConsensusChain can wait for all to be registered before returning.
	//
	// For existing chains however, clear expected, since we take what we get on-chain
	// and avoid a dependency on possibly mismatching/incorrect RPCEndpoints config.
	//
	// TODO(corver): Dynamic reloading of voter when on-chain registry is updated.
	expected := solvernet.OnlyCoreEndpoints(endpoints).Keys()
	const day = 100_000 // At least a day old
	if height, err := omniEVMCl.BlockNumber(ctx); err == nil && height > day {
		expected = nil
	}

	network, err := netconf.AwaitOnConsensusChain(ctx, netID, cprov, expected)
	if err != nil {
		return err
	}

	var xprov xchain.Provider
	if netID == netconf.Simnet {
		omni, ok := network.OmniConsensusChain()
		if !ok {
			return errors.New("omni chain not found in network")
		}

		xprov, err = xprovider.NewMock(omni.BlockPeriod*8/10, omni.ID, cprov)
		if err != nil {
			return err
		}
	} else {
		ethClients := make(map[uint64]ethclient.Client)
		for _, chain := range network.EVMChains() {
			// Use EngineAPI as omni_evm RPC client.
			if netconf.IsOmniExecution(netID, chain.ID) {
				ethClients[chain.ID] = omniEVMCl
				continue
			}

			rpc, err := endpoints.ByNameOrID(chain.Name, chain.ID)
			if err != nil {
				return err
			}

			ethCl, err := ethclient.DialContext(ctx, chain.Name, rpc)
			if err != nil {
				return err
			}
			go ethCl.CloseIdleConnectionsForever(ctx)

			ethClients[chain.ID] = ethCl
		}

		xprov = xprovider.New(network, ethClients, cprov)
	}

	deps := voteDeps{
		API:      cmtAPI,
		Provider: cprov,
	}

	v, err := voter.LoadVoter(privKey, voterStateFile, xprov, deps, network, asyncAbort)
	if err != nil {
		return errors.Wrap(err, "create voter")
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	// Process all cached values
	if err := v.SetProposed(ctx, mapValues(l.proposed)); err != nil {
		return errors.Wrap(err, "set cached proposed")
	}
	if err := v.SetCommitted(ctx, mapValues(l.committed)); err != nil {
		return errors.Wrap(err, "set cached committed")
	}
	if l.lastValSet != nil {
		if err := v.UpdateValidatorSet(l.lastValSet); err != nil {
			return errors.Wrap(err, "update validator set")
		}
	}

	// Clear all cached values
	l.proposed = nil
	l.committed = nil
	l.lastValSet = nil

	// Set voter and start it
	l.voter = v
	v.Start(ctx)

	return nil
}

func (l *voterLoader) getVoter() (*voter.Voter, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.voter, l.voter != nil
}

func (l *voterLoader) isValidator() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.isVal
}

func (l *voterLoader) GetAvailable() []*atypes.Vote {
	if v, ok := l.getVoter(); ok {
		return v.GetAvailable()
	}

	return nil // Return empty list if voter not available yet.
}

func (l *voterLoader) SetProposed(ctx context.Context, headers []*atypes.AttestHeader) error {
	if v, ok := l.getVoter(); ok {
		return v.SetProposed(ctx, headers)
	}

	// Cache (and dedup) these headers to provider to voter once available.
	// This could be votes we sent right before a restart.
	l.mu.Lock()
	defer l.mu.Unlock()
	mergeHeaders(l.proposed, headers)

	return nil
}

func (l *voterLoader) SetCommitted(ctx context.Context, headers []*atypes.AttestHeader) error {
	if v, ok := l.getVoter(); ok {
		return v.SetCommitted(ctx, headers)
	}

	// Cache (and dedup) these headers to provider to voter once available.
	// This could be votes we sent right before a restart.
	l.mu.Lock()
	defer l.mu.Unlock()
	mergeHeaders(l.committed, headers)

	return nil
}

func (l *voterLoader) LocalAddress() common.Address {
	if v, ok := l.getVoter(); ok {
		return v.LocalAddress()
	}

	return l.localAddr
}

func (l *voterLoader) TrimBehind(minsByChain map[xchain.ChainVersion]uint64) int {
	if v, ok := l.getVoter(); ok {
		return v.TrimBehind(minsByChain)
	}

	return 0
}

func (l *voterLoader) UpdateValidatorSet(valset *vtypes.ValidatorSetResponse) error {
	isVal, err := valset.IsValidator(l.localAddr)
	if err != nil {
		return err
	}
	setConstantGauge(cometValidator, isVal)

	if v, ok := l.getVoter(); ok {
		return v.UpdateValidatorSet(valset)
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	l.isVal = isVal
	l.lastValSet = valset

	return nil
}

func (l *voterLoader) WaitDone() {
	if v, ok := l.getVoter(); ok {
		v.WaitDone()
		return
	}
}

// mergeHeaders merges the provided headers into the map.
func mergeHeaders(m map[xchain.AttestHeader]*atypes.AttestHeader, headers []*atypes.AttestHeader) {
	for _, header := range headers {
		m[header.ToXChain()] = header
	}
}

// mapValues returns the map values (in random order).
func mapValues[K comparable, V any](m map[K]V) []V {
	resp := make([]V, 0, len(m))
	for _, v := range m {
		resp = append(resp, v)
	}

	return resp
}
