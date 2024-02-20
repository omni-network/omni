package keeper

import (
	"fmt"

	"github.com/omni-network/omni/halo/evmengine/types"
	"github.com/omni-network/omni/lib/engine"

	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	grpc1 "github.com/cosmos/gogoproto/grpc"
)

type Keeper struct {
	cdc          codec.BinaryCodec
	storeService store.KVStoreService
	logger       log.Logger
	ethCl        engine.API
	txConfig     client.TxConfig
	providers    []types.CPayloadProvider
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	ethCl engine.API,
	txConfig client.TxConfig,
	providers []types.CPayloadProvider,
) Keeper {
	return Keeper{
		cdc:          cdc,
		storeService: storeService,
		logger:       logger,
		ethCl:        ethCl,
		txConfig:     txConfig,
		providers:    providers,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// RegisterProposalService registers the proposal service on the provided router.
// This implements abci.ProcessProposal verification of new proposals.
func (k Keeper) RegisterProposalService(server grpc1.Server) {
	types.RegisterMsgServiceServer(server, NewProposalServer(k))
}
