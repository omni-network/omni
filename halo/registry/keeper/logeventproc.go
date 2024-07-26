package keeper

import (
	"bytes"
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/halo/registry/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	evmenginetypes "github.com/omni-network/omni/octane/evmengine/types"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ evmenginetypes.EvmEventProcessor = (*Keeper)(nil)

var (
	portalRegABI   = mustGetABI(bindings.PortalRegistryMetaData)
	portalRegEvent = mustGetEvent(portalRegABI, "PortalRegistered")
)

func (Keeper) Name() string {
	return types.ModuleName
}

func (k Keeper) Addresses() []common.Address {
	return []common.Address{k.portalRegAdress}
}

// Prepare returns all omni portal registry contract EVM event logs from the provided block hash.
func (k Keeper) Prepare(ctx context.Context, blockHash common.Hash) ([]*evmenginetypes.EVMEvent, error) {
	logs, err := k.ethCl.FilterLogs(ctx, ethereum.FilterQuery{
		BlockHash: &blockHash,
		Addresses: k.Addresses(),
		Topics:    [][]common.Hash{{portalRegEvent.ID}},
	})
	if err != nil {
		return nil, errors.Wrap(err, "filter logs")
	}

	resp := make([]*evmenginetypes.EVMEvent, 0, len(logs))
	for _, l := range logs {
		topics := make([][]byte, 0, len(l.Topics))
		for _, t := range l.Topics {
			topics = append(topics, t.Bytes())
		}
		resp = append(resp, &evmenginetypes.EVMEvent{
			Address: l.Address.Bytes(),
			Topics:  topics,
			Data:    l.Data,
		})
	}

	return resp, nil
}

// Deliver processes a omni portal registry events.
func (k Keeper) Deliver(ctx context.Context, _ common.Hash, elog *evmenginetypes.EVMEvent) error {
	ethlog := elog.ToEthLog()

	switch ethlog.Topics[0] {
	case portalRegEvent.ID:
		reg, err := k.portalRegistry.ParsePortalRegistered(ethlog)
		if err != nil {
			return errors.Wrap(err, "parse create validator")
		}

		return k.addPortal(ctx, &Portal{
			ChainId:      reg.ChainId,
			Address:      reg.Addr.Bytes(),
			DeployHeight: reg.DeployHeight,
			ShardIds:     reg.Shards,
		})
	default:
		return errors.New("unknown event")
	}
}

// addPortal adds the portal to the network config, creating a new epoch and network if necessary.
func (k Keeper) addPortal(ctx context.Context, portal *Portal) error {
	if err := portal.Verify(); err != nil {
		return errors.Wrap(err, "verify portal")
	}

	network, err := k.getOrCreateNetwork(ctx)
	if err != nil {
		return errors.Wrap(err, "get or create network")
	} else if network.GetId() == 0 {
		return errors.New("invalid existing network")
	}

	// Add new portal to the network
	network.Portals, err = mergePortal(network.GetPortals(), portal)
	if err != nil {
		return errors.Wrap(err, "merge portal", "chain_id", portal.GetChainId())
	}

	if err := k.updateNetwork(ctx, network); err != nil {
		return errors.Wrap(err, "insert network")
	}

	log.Info(ctx, "🔭 Added network portal",
		"network_id", network.GetId(),
		"chain", k.chainNamer(portal.GetChainId()),
		"shards", shardLabels(portal.GetShardIds()),
		"height", sdk.UnwrapSDKContext(ctx).BlockHeight(),
	)

	return nil
}

// mergePortal merges the new portal with the existing list.
func mergePortal(existing []*Portal, portal *Portal) ([]*Portal, error) {
	for i, e := range existing {
		if e.GetChainId() != portal.GetChainId() {
			continue
		}

		// Merge new shads with an existing portal
		if !bytes.Equal(e.GetAddress(), portal.GetAddress()) {
			return nil, errors.New("cannot merge existing portal with mismatching address",
				"existing", e.GetAddress(), "new", portal.GetAddress())
		} else if e.GetDeployHeight() != portal.GetDeployHeight() {
			return nil, errors.New("cannot merge existing portal with mismatching deploy height",
				"existing", e.GetDeployHeight(), "new", portal.GetDeployHeight())
		}

		toMerge := newShards(e.GetShardIds(), portal.GetShardIds())
		if len(toMerge) == 0 {
			return nil, errors.New("cannot merge existing portal with no new shards",
				"existing", e.GetShardIds(), "new", portal.GetShardIds())
		}

		existing[i].ShardIds = append(existing[i].ShardIds, toMerge...)

		return existing, nil
	}

	return append(existing, portal), nil // New chain, just append
}

// newShards returns the new shards that are not in the existing list.
func newShards(existing []uint64, shards []uint64) []uint64 {
	exists := func(s uint64) bool {
		for _, e := range existing {
			if e == s {
				return true
			}
		}

		return false
	}

	var resp []uint64
	for _, s := range shards {
		if !exists(s) {
			resp = append(resp, s)
		}
	}

	return resp
}

// mustGetABI returns the metadata's ABI as an abi.ABI type.
// It panics on error.
func mustGetABI(metadata *bind.MetaData) *abi.ABI {
	abi, err := metadata.GetAbi()
	if err != nil {
		panic(err)
	}

	return abi
}

// mustGetEvent returns the event with the given name from the ABI.
// It panics if the event is not found.
func mustGetEvent(abi *abi.ABI, name string) abi.Event {
	event, ok := abi.Events[name]
	if !ok {
		panic("event not found")
	}

	return event
}
