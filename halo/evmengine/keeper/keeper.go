package keeper

import (
	"context"
	"sync"
	"time"

	"github.com/omni-network/omni/halo/comet"
	"github.com/omni-network/omni/halo/evmengine/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/k1util"

	eengine "github.com/ethereum/go-ethereum/beacon/engine"

	"cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	grpc1 "github.com/cosmos/gogoproto/grpc"
)

type Keeper struct {
	cdc             codec.BinaryCodec
	storeService    store.KVStoreService
	engineCl        ethclient.EngineClient
	txConfig        client.TxConfig
	voteProvider    types.VoteExtensionProvider
	eventProcs      []types.EvmEventProcessor
	cmtAPI          comet.API
	addrProvider    types.AddressProvider
	buildDelay      time.Duration
	buildOptimistic bool

	// mutablePayload contains the previous optimistically triggered payload.
	// It is optimistic because the validator set can change,
	// so we might not actually be the next proposer.
	mutablePayload struct {
		sync.Mutex
		ID        *eengine.PayloadID
		Height    uint64
		UpdatedAt time.Time
	}
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	engineCl ethclient.EngineClient,
	txConfig client.TxConfig,
) *Keeper {
	return &Keeper{
		cdc:          cdc,
		storeService: storeService,
		engineCl:     engineCl,
		txConfig:     txConfig,
	}
}

// TODO(corver): Figure out how to use depinject for this.
func (k *Keeper) AddEventProcessor(p types.EvmEventProcessor) {
	k.eventProcs = append(k.eventProcs, p)
}

func (k *Keeper) SetVoteProvider(p types.VoteExtensionProvider) {
	k.voteProvider = p
}

// SetCometAPI sets the comet API client.
func (k *Keeper) SetCometAPI(c comet.API) {
	k.cmtAPI = c
}

func (k *Keeper) SetAddressProvider(p types.AddressProvider) {
	k.addrProvider = p
}

// SetBuildDelay sets the build delay parameter.
func (k *Keeper) SetBuildDelay(d time.Duration) {
	k.buildDelay = d
}

// SetBuildOptimistic sets the optimistic build parameter.
func (k *Keeper) SetBuildOptimistic(b bool) {
	k.buildOptimistic = b
}

// RegisterProposalService registers the proposal service on the provided router.
// This implements abci.ProcessProposal verification of new proposals.
func (k *Keeper) RegisterProposalService(server grpc1.Server) {
	types.RegisterMsgServiceServer(server, NewProposalServer(k))
}

// isNextProposer returns true if the local node is the proposer
// for the next block. It also returns the next block height.
//
// Note that the validator set can change, so this is an optimistic check.
func (k *Keeper) isNextProposer(ctx context.Context) (bool, uint64, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	header := sdkCtx.BlockHeader()
	nextHeight := header.Height + 1

	valset, ok, err := k.cmtAPI.Validators(ctx, header.Height)
	if err != nil {
		return false, 0, err
	} else if !ok {
		return false, 0, errors.New("validators not available")
	}

	idx, _ := valset.GetByAddress(header.ProposerAddress)
	if idx < 0 {
		return false, 0, errors.New("proposer not in validator set")
	}

	nextIdx := int(idx+1) % len(valset.Validators)
	nextProposer := valset.Validators[nextIdx]
	nextAddr, err := k1util.PubKeyToAddress(nextProposer.PubKey)
	if err != nil {
		return false, 0, err
	}

	isNextProposer := nextAddr == k.addrProvider.LocalAddress()

	return isNextProposer, uint64(nextHeight), nil
}

func (k *Keeper) setOptimisticPayload(id *eengine.PayloadID, height uint64) {
	k.mutablePayload.Lock()
	defer k.mutablePayload.Unlock()

	k.mutablePayload.ID = id
	k.mutablePayload.Height = height
	k.mutablePayload.UpdatedAt = time.Now()
}

func (k *Keeper) getOptimisticPayload() (*eengine.PayloadID, uint64, time.Time) {
	k.mutablePayload.Lock()
	defer k.mutablePayload.Unlock()

	return k.mutablePayload.ID, k.mutablePayload.Height, k.mutablePayload.UpdatedAt
}
