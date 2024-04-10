package keeper

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/omni-network/omni/halo/attest/types"
	vtypes "github.com/omni-network/omni/halo/valsync/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/ethereum/go-ethereum/common"

	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"
	"cosmossdk.io/core/store"
	"cosmossdk.io/orm/model/ormdb"
	"cosmossdk.io/orm/model/ormlist"
	"cosmossdk.io/orm/types/ormerrors"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	"github.com/cosmos/gogoproto/proto"
)

var _ sdk.ExtendVoteHandler = (*Keeper)(nil).ExtendVote
var _ sdk.VerifyVoteExtensionHandler = (*Keeper)(nil).VerifyVoteExtension

// Keeper is the attestation keeper.
// It keeps tracks of all attestations included on-chain and detects when they are approved.
type Keeper struct {
	attTable     AttestationTable
	sigTable     SignatureTable
	cdc          codec.BinaryCodec
	storeService store.KVStoreService
	voter        types.Voter
	skeeper      baseapp.ValidatorStore
	valProvider  vtypes.ValidatorProvider
	namer        types.ChainNameFunc
	voteWindow   uint64
	voteExtLimit uint64
	trimLag      uint64
}

// New returns a new attestation keeper.
func New(
	cdc codec.BinaryCodec,
	storeSvc store.KVStoreService,
	skeeper baseapp.ValidatorStore,
	namer types.ChainNameFunc,
	voteWindow uint64,
	voteExtLimit uint64,
	trimLag uint64,
) (*Keeper, error) {
	schema := &ormv1alpha1.ModuleSchemaDescriptor{SchemaFile: []*ormv1alpha1.ModuleSchemaDescriptor_FileEntry{
		{Id: 1, ProtoFileName: File_halo_attest_keeper_attestation_proto.Path()},
	}}

	modDB, err := ormdb.NewModuleDB(schema, ormdb.ModuleDBOptions{KVStoreService: storeSvc})
	if err != nil {
		return nil, errors.Wrap(err, "create module db")
	}

	attstore, err := NewAttestationStore(modDB)
	if err != nil {
		return nil, errors.Wrap(err, "create attestation store")
	}

	k := &Keeper{
		attTable:     attstore.AttestationTable(),
		sigTable:     attstore.SignatureTable(),
		cdc:          cdc,
		storeService: storeSvc,
		skeeper:      skeeper,
		namer:        namer,
		voteWindow:   voteWindow,
		voteExtLimit: voteExtLimit,
		trimLag:      trimLag,
	}

	return k, nil
}

// SetValidatorProvider sets the validator provider.
func (k *Keeper) SetValidatorProvider(valProvider vtypes.ValidatorProvider) {
	k.valProvider = valProvider
}

// SetVoter sets the voter.
func (k *Keeper) SetVoter(voter types.Voter) {
	k.voter = voter
}

// RegisterProposalService registers the proposal service on the provided router.
// This implements abci.ProcessProposal verification of new proposals.
func (k *Keeper) RegisterProposalService(server grpc1.Server) {
	types.RegisterMsgServiceServer(server, NewProposalServer(k))
}

// Add adds the given aggregate votes as pending attestations to the store.
// It merges the votes with attestations it already exists.
func (k *Keeper) Add(ctx context.Context, msg *types.MsgAddVotes) error {
	valset, err := k.prevBlockValSet(ctx)
	if err != nil {
		return errors.Wrap(err, "fetch validators")
	}

	countsByChain := make(map[uint64]int)
	for _, aggVote := range msg.Votes {
		countsByChain[aggVote.BlockHeader.ChainId]++

		// Sanity check that all votes are from prev block validators.
		for _, sig := range aggVote.Signatures {
			if !valset.Contains(common.BytesToAddress(sig.ValidatorAddress)) {
				return errors.New("vote from unknown validator [BUG]")
			}
		}

		err := k.addOne(ctx, aggVote, valset.ID)
		if err != nil {
			return errors.Wrap(err, "add one")
		}
	}

	for chainID, count := range countsByChain {
		votesProposed.WithLabelValues(k.namer(chainID)).Observe(float64(count))
	}

	return nil
}

// addOne adds the given aggregate vote to the store.
// It merges it if the attestation already exists.
func (k *Keeper) addOne(ctx context.Context, agg *types.AggVote, valSetID uint64) error {
	defer latency("add_one")()

	header := agg.BlockHeader

	// Get existing attestation (by unique key) or insert new one.
	var attID uint64
	existing, err := k.attTable.GetByChainIdHeightHashAttestationRoot(ctx,
		header.ChainId, header.Height, header.Hash, agg.AttestationRoot)
	if ormerrors.IsNotFound(err) {
		// Insert new attestation
		attID, err = k.attTable.InsertReturningId(ctx, &Attestation{
			ChainId:         agg.BlockHeader.ChainId,
			Height:          agg.BlockHeader.Height,
			Hash:            agg.BlockHeader.Hash,
			AttestationRoot: agg.AttestationRoot,
			Status:          int32(Status_Pending),
			ValidatorSetId:  0, // Unknown at this point.
			CreatedHeight:   uint64(sdk.UnwrapSDKContext(ctx).BlockHeight()),
		})
		if err != nil {
			return errors.Wrap(err, "insert")
		}
	} else if err != nil {
		return errors.Wrap(err, "by att unique key")
	} else if isApprovedByDifferentSet(existing, valSetID) {
		log.Debug(ctx, "Ignoring vote for attestation approved by different validator set",
			"agg_id", attID,
			"chain", k.namer(header.ChainId),
			"height", header.Height,
		)
		// Technically these new votes could be from validators also in that previous set, but
		// we don't have consistent access to historical validator sets.

		return nil
	} else {
		attID = existing.GetId()
	}

	// Insert signatures
	for _, sig := range agg.Signatures {
		err := k.sigTable.Insert(ctx, &Signature{
			Signature:        sig.Signature,
			ValidatorAddress: sig.ValidatorAddress,
			AttId:            attID,
		})

		if errors.Is(err, ormerrors.UniqueKeyViolation) {
			// TODO(corver): We should prevent this from happening earlier.
			log.Warn(ctx, "Ignoring duplicate vote", nil,
				"agg_id", attID,
				"chain", k.namer(header.ChainId),
				"height", header.Height,
				log.Hex7("validator", sig.ValidatorAddress),
			)
		} else if err != nil {
			return errors.Wrap(err, "insert signature")
		}
	}

	return nil
}

// Approve approves any pending attestations that have quorum signatures from the provided set.
func (k *Keeper) Approve(ctx context.Context, valset ValSet) error {
	defer latency("approve")()

	pendingIdx := AttestationStatusChainIdHeightIndexKey{}.WithStatus(int32(Status_Pending))
	iter, err := k.attTable.List(ctx, pendingIdx)
	if err != nil {
		return errors.Wrap(err, "list pending")
	}
	defer iter.Close()

	approvedByChain := make(map[uint64]uint64) // Cache the latest approved attestation by chain.
	for iter.Next() {
		att, err := iter.Value()
		if err != nil {
			return errors.Wrap(err, "value")
		}

		// Ensure we approve sequentially, not skipping any heights.
		{
			// Populate the cache if not already.
			if _, ok := approvedByChain[att.GetChainId()]; ok {
				latest, found, err := k.latestAttestation(ctx, att.GetChainId())
				if err != nil {
					return errors.Wrap(err, "latest approved")
				} else if found {
					approvedByChain[att.GetChainId()] = latest.BlockHeader.Height
				}
			}
			head, ok := approvedByChain[att.GetChainId()]
			if ok && head+1 != att.GetHeight() {
				// This isn't the next attestation to approve, so we can't approve it yet.
				continue
			}
			// TODO(corver): We should ensure we start approving from portal deploy height.
		}

		sigs, err := k.getSigs(ctx, att.GetId())
		if err != nil {
			return errors.Wrap(err, "get att signatures")
		}

		toDelete, ok := isApproved(sigs, valset)
		if !ok {
			continue
		}

		for _, sig := range toDelete {
			err := k.sigTable.Delete(ctx, sig)
			if err != nil {
				return errors.Wrap(err, "delete sig")
			}
		}

		// Update status
		att.Status = int32(Status_Approved)
		att.ValidatorSetId = valset.ID
		err = k.attTable.Update(ctx, att)
		if err != nil {
			return errors.Wrap(err, "save")
		}

		approvedHeight.WithLabelValues(k.namer(att.GetChainId())).Set(float64(att.GetHeight()))
		approvedByChain[att.GetChainId()] = att.GetHeight()
	}

	// Trim votes behind minimum vote-window
	minVoteWindows := make(map[uint64]uint64)
	for chainID, head := range approvedByChain {
		minVoteWindows[chainID] = uintSub(head, k.voteWindow)
	}

	count := k.voter.TrimBehind(minVoteWindows)
	if count > 0 {
		log.Warn(ctx, "Trimmed votes behind vote-window (expected if node was struggling)", nil, "count", count)
	}

	return nil
}

// ListAttestationsFrom returns the subsequent approved attestations from the provided height (inclusive).
func (k *Keeper) ListAttestationsFrom(ctx context.Context, chainID uint64, height uint64, max uint64) ([]*types.Attestation, error) {
	defer latency("attestations_from")()

	from := AttestationStatusChainIdHeightIndexKey{}.WithStatusChainIdHeight(int32(Status_Approved), chainID, height)
	to := AttestationStatusChainIdHeightIndexKey{}.WithStatusChainIdHeight(int32(Status_Approved), chainID, height+max)

	iter, err := k.attTable.ListRange(ctx, from, to)
	if err != nil {
		return nil, errors.Wrap(err, "list range")
	}
	defer iter.Close()

	var resp []*types.Attestation
	next := height
	for iter.Next() {
		att, err := iter.Value()
		if err != nil {
			return nil, errors.Wrap(err, "value")
		}

		if att.GetHeight() != next {
			break
		}
		next++

		pbsigs, err := k.getSigs(ctx, att.GetId())
		if err != nil {
			return nil, errors.Wrap(err, "get att sigs")
		}

		var sigs []*types.SigTuple
		for _, pbsig := range pbsigs {
			sigs = append(sigs, &types.SigTuple{
				ValidatorAddress: pbsig.GetValidatorAddress(),
				Signature:        pbsig.GetSignature(),
			})
		}

		resp = append(resp, &types.Attestation{
			BlockHeader: &types.BlockHeader{
				ChainId: att.GetChainId(),
				Height:  att.GetHeight(),
				Hash:    att.GetHash(),
			},
			ValidatorSetId:  att.GetValidatorSetId(),
			AttestationRoot: att.GetAttestationRoot(),
			Signatures:      sigs,
		})
	}

	return resp, nil
}

// latestAttestation returns the latest approved attestation for the given chain or
// false if none is found.
func (k *Keeper) latestAttestation(ctx context.Context, chainID uint64) (*types.Attestation, bool, error) {
	defer latency("latest_attestation")()

	idx := AttestationStatusChainIdHeightIndexKey{}.WithStatusChainId(int32(Status_Approved), chainID)
	iter, err := k.attTable.List(ctx, idx, ormlist.Reverse(), ormlist.DefaultLimit(1))
	if err != nil {
		return nil, false, errors.Wrap(err, "list")
	}
	defer iter.Close()

	if !iter.Next() {
		return nil, false, nil
	}

	att, err := iter.Value()
	if err != nil {
		return nil, false, errors.Wrap(err, "value")
	}

	if iter.Next() {
		return nil, false, errors.New("multiple attestation found")
	}

	pbsigs, err := k.getSigs(ctx, att.GetId())
	if err != nil {
		return nil, false, errors.Wrap(err, "get att sigs")
	}

	var sigs []*types.SigTuple
	for _, pbsig := range pbsigs {
		sigs = append(sigs, &types.SigTuple{
			ValidatorAddress: pbsig.GetValidatorAddress(),
			Signature:        pbsig.GetSignature(),
		})
	}

	return &types.Attestation{
		BlockHeader: &types.BlockHeader{
			ChainId: att.GetChainId(),
			Height:  att.GetHeight(),
			Hash:    att.GetHash(),
		},
		ValidatorSetId:  att.GetValidatorSetId(),
		AttestationRoot: att.GetAttestationRoot(),
		Signatures:      sigs,
	}, true, nil
}

// getSigs returns the signatures for the given attestation ID.
func (k *Keeper) getSigs(ctx context.Context, attID uint64) ([]*Signature, error) {
	attIDIdx := SignatureAttIdIndexKey{}.WithAttId(attID)
	sigIter, err := k.sigTable.List(ctx, attIDIdx)
	if err != nil {
		return nil, errors.Wrap(err, "list sig")
	}
	defer sigIter.Close()

	var sigs []*Signature
	for sigIter.Next() {
		sig, err := sigIter.Value()
		if err != nil {
			return nil, errors.Wrap(err, "value sig")
		}

		sigs = append(sigs, sig)
	}

	return sigs, nil
}

func (k *Keeper) BeginBlock(ctx context.Context) error {
	head := uint64(sdk.UnwrapSDKContext(ctx).BlockHeight())

	return k.deleteBefore(ctx, uintSub(head, k.trimLag))
}

func (k *Keeper) EndBlock(ctx context.Context) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if sdkCtx.BlockHeight() <= 1 {
		return nil // First block doesn't have any vote extensions to approve.
	}

	valset, err := k.prevBlockValSet(ctx)
	if err != nil {
		return errors.Wrap(err, "fetch validators")
	}

	return k.Approve(ctx, valset)
}

// ExtendVote extends a vote with application-injected data (vote extensions).
func (k *Keeper) ExtendVote(ctx sdk.Context, _ *abci.RequestExtendVote) (*abci.ResponseExtendVote, error) {
	votes := k.voter.GetAvailable()

	// Filter by vote window and if limited exceeded.
	countsByChain := make(map[uint64]int)
	var filtered []*types.Vote
	for _, vote := range votes {
		if cmp, err := k.windowCompare(ctx, vote.BlockHeader.ChainId, vote.BlockHeader.Height); err != nil {
			return nil, errors.Wrap(err, "windower")
		} else if cmp != 0 {
			// Skip votes no in the window
			continue
		}
		countsByChain[vote.BlockHeader.ChainId]++
		filtered = append(filtered, vote)

		if len(filtered) >= int(k.voteExtLimit) {
			break
		}
	}

	bz, err := proto.Marshal(&types.Votes{
		Votes: filtered,
	})
	if err != nil {
		return nil, errors.Wrap(err, "marshal atts")
	}

	for chainID, count := range countsByChain {
		votesExtended.WithLabelValues(k.namer(chainID)).Observe(float64(count))
	}

	// Make nice logs
	const limit = 5
	heights := make(map[uint64][]string)
	for _, vote := range votes {
		hs := heights[vote.BlockHeader.ChainId]
		if len(hs) < limit {
			hs = append(hs, strconv.FormatUint(vote.BlockHeader.Height, 10))
		} else if len(hs) == limit {
			hs = append(hs, "...")
		} else {
			continue
		}
		heights[vote.BlockHeader.ChainId] = hs
	}
	attrs := []any{slog.Int("votes", len(votes))}
	for cid, hs := range heights {
		attrs = append(attrs, slog.String(
			strconv.FormatUint(cid, 10),
			fmt.Sprint(hs),
		))
	}

	log.Info(ctx, "Voted for rollup blocks", attrs...)

	return &abci.ResponseExtendVote{
		VoteExtension: bz,
	}, nil
}

// VerifyVoteExtension verifies a vote extension.
func (k *Keeper) VerifyVoteExtension(ctx sdk.Context, req *abci.RequestVerifyVoteExtension) (
	*abci.ResponseVerifyVoteExtension, error,
) {
	respAccept := &abci.ResponseVerifyVoteExtension{
		Status: abci.ResponseVerifyVoteExtension_ACCEPT,
	}
	respReject := &abci.ResponseVerifyVoteExtension{
		Status: abci.ResponseVerifyVoteExtension_REJECT,
	}

	// Adding logging attributes to sdk context is a bit tricky
	ctx = ctx.WithContext(log.WithCtx(ctx, log.Hex7("validator", req.ValidatorAddress)))

	votes, ok, err := votesFromExtension(req.VoteExtension)
	if err != nil {
		log.Warn(ctx, "Rejecting invalid vote extension", err)
		return respReject, nil
	} else if !ok {
		log.Info(ctx, "Accepting nil vote extension") // This can happen in some edge-cases.
		return respAccept, nil
	} else if len(votes.Votes) > int(k.voteExtLimit) {
		log.Warn(ctx, "Rejecting vote extension exceeding limit", nil, "count", len(votes.Votes), "limit", k.voteExtLimit)
		return respReject, nil
	}

	for _, vote := range votes.Votes {
		if err := vote.Verify(); err != nil {
			log.Warn(ctx, "Rejecting invalid vote", err)
			return respReject, nil
		}
		if cmp, err := k.windowCompare(ctx, vote.BlockHeader.ChainId, vote.BlockHeader.Height); err != nil {
			return nil, errors.Wrap(err, "windower")
		} else if cmp != 0 {
			log.Warn(ctx, "Rejecting out-of-window vote", nil, "cmp", cmp)
			return respReject, nil
		}
	}

	return respAccept, nil
}

type ValSet struct {
	ID   uint64
	Vals map[common.Address]int64
}

func (s ValSet) Contains(addr common.Address) bool {
	_, ok := s.Vals[addr]
	return ok
}

func (s ValSet) TotalPower() int64 {
	var resp int64
	for _, power := range s.Vals {
		resp += power
	}

	return resp
}

// prevBlockValSet returns the previous blocks active validator set.
// Previous block is used since vote extensions are delayed by one block.
func (k *Keeper) prevBlockValSet(ctx context.Context) (ValSet, error) {
	prevBlock := sdk.UnwrapSDKContext(ctx).BlockHeight() - 1
	resp, err := k.valProvider.ActiveSetByHeight(ctx, uint64(prevBlock))
	if err != nil {
		return ValSet{}, err
	}

	valsByPower := make(map[common.Address]int64)
	for _, val := range resp.Validators {
		valsByPower[common.BytesToAddress(val.Address)] = val.Power
	}

	return ValSet{
		ID:   resp.Id,
		Vals: valsByPower,
	}, nil
}

func (k *Keeper) windowCompare(ctx context.Context, chainID uint64, height uint64) (int, error) {
	latest, exists, err := k.latestAttestation(ctx, chainID)
	if err != nil {
		return 0, err
	} else if !exists {
		// TODO(corver): Use netconf deploy height to use as initial window.
		return 0, nil // Allow any height while no approved attestation exists.
	}

	return windowCompare(k.voteWindow, latest.BlockHeader, height), nil
}

// verifyAggVotes verifies the given aggregates votes:
// - Ensure all votes are from validators in the provided set.
// - Ensure the vote block header is in the vote window.
// - Ensure votes represent at least 2/3 of the total voting power. <- This isn't done?
func (k *Keeper) verifyAggVotes(ctx context.Context, valset ValSet, aggs []*types.AggVote) error {
	for _, agg := range aggs {
		if err := agg.Verify(); err != nil {
			return errors.Wrap(err, "verify aggregate vote")
		}
		errAttrs := []any{"chain_id", agg.BlockHeader.ChainId, "height", agg.BlockHeader.Height, log.Hex7("val0", agg.Signatures[0].ValidatorAddress)}

		// Ensure all votes are from validators in the set
		for _, sig := range agg.Signatures {
			addr := common.BytesToAddress(sig.GetValidatorAddress())
			if !valset.Contains(addr) {
				errAttrs = append(errAttrs, log.Hex7("validator", sig.GetValidatorAddress()))

				return errors.New("vote from unknown validator", errAttrs...)
			}
		}

		// Ensure the block header is in the vote window.
		if resp, err := k.windowCompare(ctx, agg.BlockHeader.ChainId, agg.BlockHeader.Height); err != nil {
			return errors.Wrap(err, "windower")
		} else if resp != 0 {
			errAttrs = append(errAttrs, "resp", resp)

			return errors.New("vote outside window", errAttrs...)
		}
	}

	return nil
}

// deleteBefore deletes all attestations and signatures before the given height (inclusive).
// Note this always deletes block 0, but genesis block doesn't contain any attestations.
func (k *Keeper) deleteBefore(ctx context.Context, height uint64) error {
	defer latency("delete_before")()

	start := AttestationCreatedHeightIndexKey{}
	end := AttestationCreatedHeightIndexKey{}.WithCreatedHeight(height)
	iter, err := k.attTable.ListRange(ctx, start, end)
	if err != nil {
		return errors.Wrap(err, "list atts")
	}
	defer iter.Close()

	for iter.Next() {
		att, err := iter.Value()
		if err != nil {
			return errors.Wrap(err, "value att")
		} else if att.GetCreatedHeight() > height {
			return errors.New("query sanity check [BUG]")
		}

		// Delete signatures
		if err := k.sigTable.DeleteBy(ctx, SignatureAttIdIndexKey{}.WithAttId(att.GetId())); err != nil {
			return errors.Wrap(err, "delete sigs")
		}

		// Delete attestation
		err = k.attTable.Delete(ctx, att)
		if err != nil {
			return errors.Wrap(err, "delete att")
		}
	}

	return nil
}

// isApproved returns whether the given signatures are approved by the given validators.
// It also returns the signatures to delete (not in the validator set).
func isApproved(sigs []*Signature, valset ValSet) ([]*Signature, bool) {
	var sum int64
	var toDelete []*Signature
	for _, sig := range sigs {
		power, ok := valset.Vals[common.BytesToAddress(sig.GetValidatorAddress())]
		if !ok {
			toDelete = append(toDelete, sig)
			continue
		}

		sum += power
	}

	return toDelete, sum > valset.TotalPower()*2/3
}

// windowCompare returns -1 if x < mid-voteWindow, 1 if x > mid+voteWindow, else 0.
func windowCompare(voteWindow uint64, midHeader *types.BlockHeader, x uint64) int {
	mid := midHeader.Height

	if x < uintSub(mid, voteWindow) {
		return -1
	} else if x > mid+voteWindow {
		return 1
	}

	return 0
}

// uintSub returns a - b if a > b, else 0.
// Subtracting uints can result in underflow, so we need to check for that.
func uintSub(a, b uint64) uint64 {
	if a <= b {
		return 0
	}

	return a - b
}

// isApprovedByDifferentSet returns true if the attestation is approved by a different validator set.
func isApprovedByDifferentSet(att *Attestation, valSetID uint64) bool {
	if att.GetStatus() != int32(Status_Approved) {
		return false
	}

	return att.GetValidatorSetId() != valSetID
}
