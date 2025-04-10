package keeper

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	ptypes "github.com/omni-network/omni/halo/portal/types"
	"github.com/omni-network/omni/halo/registry/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/umath"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"

	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"
	"cosmossdk.io/core/store"
	"cosmossdk.io/orm/model/ormdb"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	emitPortal       ptypes.EmitPortal
	networkTable     NetworkTable
	portalRegAddress common.Address
	portalRegistry   *bindings.PortalRegistryFilterer
	chainNamer       types.ChainNameFunc

	latestCache *cache
}

func NewKeeper(
	emitPortal ptypes.EmitPortal,
	storeService store.KVStoreService,
	namer types.ChainNameFunc,
) (Keeper, error) {
	schema := &ormv1alpha1.ModuleSchemaDescriptor{SchemaFile: []*ormv1alpha1.ModuleSchemaDescriptor_FileEntry{
		{Id: 1, ProtoFileName: File_halo_registry_keeper_registry_proto.Path()},
	}}

	modDB, err := ormdb.NewModuleDB(schema, ormdb.ModuleDBOptions{KVStoreService: storeService})
	if err != nil {
		return Keeper{}, errors.Wrap(err, "create module db")
	}

	registryStore, err := NewRegistryStore(modDB)
	if err != nil {
		return Keeper{}, errors.Wrap(err, "create registry store")
	}

	address := common.HexToAddress(predeploys.PortalRegistry)
	portalReg, err := bindings.NewPortalRegistryFilterer(address, nil) // Passing nil backend if safe since only Parse functions are used.
	if err != nil {
		return Keeper{}, errors.Wrap(err, "new portal registry")
	}

	return Keeper{
		emitPortal:       emitPortal,
		networkTable:     registryStore.NetworkTable(),
		portalRegAddress: address,
		portalRegistry:   portalReg,
		chainNamer:       namer,
		latestCache:      new(cache),
	}, nil
}

// getOrCreateNetwork returns a network created in the current height.
// If one already exists, it will be returned.
// If none already exists, a new one will be created using the previous as base.
// New networks are emitted as cross chain messages to portals.
func (k Keeper) getOrCreateNetwork(ctx context.Context) (*Network, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	createHeight, err := umath.ToUint64(sdkCtx.BlockHeight())
	if err != nil {
		return nil, err
	}

	var lastPortals []*Portal

	latestNetworkID, err := k.networkTable.LastInsertedSequence(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get last network ID")
	} else if latestNetworkID != 0 {
		// Get the latest network
		lastNetwork, err := k.networkTable.Get(ctx, latestNetworkID)
		if err != nil {
			return nil, errors.Wrap(err, "get network")
		} else if lastNetwork.GetCreatedHeight() == createHeight {
			// This network was created in this block, use it as is.
			return lastNetwork, nil
		}

		lastPortals = lastNetwork.GetPortals()
	}

	// Create a new network using the latest network as base
	network := &Network{
		CreatedHeight: createHeight,
		Portals:       lastPortals,
	}
	network.Id, err = k.networkTable.InsertReturningId(ctx, network)
	if err != nil {
		return nil, errors.Wrap(err, "insert next network")
	}

	k.latestCache.Set(network)

	_, err = k.emitPortal.EmitMsg(
		sdkCtx,
		ptypes.MsgTypeNetwork,
		network.GetId(),
		xchain.BroadcastChainID,
		xchain.ShardBroadcast0,
	)
	if err != nil {
		return nil, errors.Wrap(err, "emit portal message")
	}

	return network, nil
}
