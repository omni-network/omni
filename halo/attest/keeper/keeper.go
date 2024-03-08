package keeper

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/halo/comet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/log"

	abci "github.com/cometbft/cometbft/abci/types"
	cmttypes "github.com/cometbft/cometbft/types"

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
//
// TODOs:
//   - Delete attestations after 14 days.
type Keeper struct {
	attTable     AttestationTable
	sigTable     SignatureTable
	cdc          codec.BinaryCodec
	storeService store.KVStoreService
	voter        types.Voter
	skeeper      baseapp.ValidatorStore
	cmtAPI       comet.API
	namer        types.ChainNameFunc
	voteWindow   uint64
	voteExtLimit uint64
}

// New returns a new attestation keeper.
func New(
	cdc codec.BinaryCodec,
	storeSvc store.KVStoreService,
	skeeper baseapp.ValidatorStore,
	namer types.ChainNameFunc,
	voteWindow uint64,
	voteExtLimit uint64,
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
	}

	return k, nil
}

// SetCometAPI sets the comet API client.
func (k *Keeper) SetCometAPI(cmtAPI comet.API) {
	k.cmtAPI = cmtAPI
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

// Add adds the given aggregate votes as pen the store.
// It merges the votes with attestations it already exists.
func (k *Keeper) Add(ctx context.Context, msg *types.MsgAddVotes) error {
	for _, vote := range msg.Votes {
		err := k.addOne(ctx, vote)
		if err != nil {
			return errors.Wrap(err, "add one")
		}
	}

	return nil
}

// addOne adds the given aggregate vote to the store.
// It merges it if the attestation already exists.
func (k *Keeper) addOne(ctx context.Context, agg *types.AggVote) error {
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
			ValidatorsHash:  nil, // Unknown at this point.
		})
		if err != nil {
			return errors.Wrap(err, "insert")
		}
	} else if err != nil {
		return errors.Wrap(err, "by att unique key")
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
func (k *Keeper) Approve(ctx context.Context, valset *cmttypes.ValidatorSet) error {
	if valset == nil {
		return errors.New("validator set cannot be nil")
	}
	pendingIdx := AttestationStatusChainIdHeightIndexKey{}.WithStatus(int32(Status_Pending))
	iter, err := k.attTable.List(ctx, pendingIdx)
	if err != nil {
		return errors.Wrap(err, "list pending")
	}
	defer iter.Close()

	valsByPower := make(map[common.Address]int64)
	for _, val := range valset.Validators {
		addr, err := k1util.PubKeyToAddress(val.PubKey)
		if err != nil {
			return errors.Wrap(err, "pubkey to address")
		}
		valsByPower[addr] = val.VotingPower
	}

	skip := make(map[uint64]bool)           // Skip processing chains as soon as a pending attestation is found.
	edgesByChain := make(map[uint64]uint64) // Track new minimum edges for updated vote windows.
	for iter.Next() {
		att, err := iter.Value()
		if err != nil {
			return errors.Wrap(err, "value")
		}
		if skip[att.GetChainId()] {
			continue
		}

		sigs, err := k.getSigs(ctx, att.GetId())
		if err != nil {
			return errors.Wrap(err, "get att signatures")
		}

		toDelete, ok := isApproved(sigs, valsByPower, valset.TotalVotingPower())
		if !ok {
			skip[att.GetChainId()] = true
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
		att.ValidatorsHash = valset.Hash()
		err = k.attTable.Update(ctx, att)
		if err != nil {
			return errors.Wrap(err, "save")
		}

		edgesByChain[att.GetChainId()] = uintSub(att.GetHeight(), k.voteWindow)
		approvedHeight.WithLabelValues(k.namer(att.GetChainId())).Set(float64(att.GetHeight()))
	}

	count := k.voter.TrimBehind(edgesByChain)
	if count > 0 {
		log.Warn(ctx, "Trimmed votes behind vote-window (expected if node was struggling)", nil, "count", count)
	}

	return nil
}

// attestationFrom returns the subsequent approved attestations from the provided height (inclusive).
func (k *Keeper) attestationFrom(ctx context.Context, chainID uint64, height uint64, max uint64) ([]*types.Attestation, error) {
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
			ValidatorsHash:  att.GetValidatorsHash(),
			AttestationRoot: att.GetAttestationRoot(),
			Signatures:      sigs,
		})
	}

	return resp, nil
}

// latestAttestation returns the latest approved attestation for the given chain or
// false if none is found.
func (k *Keeper) latestAttestation(ctx context.Context, chainID uint64) (*types.Attestation, bool, error) {
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
		ValidatorsHash:  att.GetValidatorsHash(),
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

func (k *Keeper) EndBlock(ctx context.Context) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if sdkCtx.BlockHeight() <= 1 {
		return nil // First block doesn't have any vote extensions to approve.
	}

	// We should technically use the validators from the previous block, but that isn't available immediately
	// after a snapshot restore, so workaround is just to use current set. Only drawback is that the last
	// votes from validators that are no longer in the set will be ignored.
	valset, ok, err := k.cmtAPI.Validators(ctx, sdkCtx.BlockHeight())
	if err != nil {
		return errors.Wrap(err, "fetch validators")
	} else if !ok {
		return errors.New("current validators not available [BUG]")
	}

	return k.Approve(ctx, valset)
}

// ExtendVote extends a vote with application-injected data (vote extensions).
func (k *Keeper) ExtendVote(ctx sdk.Context, _ *abci.RequestExtendVote) (*abci.ResponseExtendVote, error) {
	votes := k.voter.GetAvailable()

	// Filter by vote window and if limited exceeded.
	var filtered []*types.Vote
	for _, vote := range votes {
		if cmp, err := k.windowCompare(ctx, vote.BlockHeader.ChainId, vote.BlockHeader.Height); err != nil {
			return nil, errors.Wrap(err, "windower")
		} else if cmp != 0 {
			continue
		}
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
		log.Info(ctx, "Accepting nil vote extension", err) // This can happen in some edge-cases.
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

// validatorsByAddress returns the validator set by ethereum address for the provided height or false if not available.
func (k *Keeper) validatorsByAddress(ctx context.Context, height int64) (map[common.Address]bool, bool, error) {
	valset, ok, err := k.cmtAPI.Validators(ctx, height)
	if err != nil {
		return nil, false, err
	} else if !ok {
		return nil, false, nil
	}

	resp := make(map[common.Address]bool)
	for _, val := range valset.Validators {
		addr, err := k1util.PubKeyToAddress(val.PubKey)
		if err != nil {
			return nil, false, err
		}
		resp[addr] = true
	}

	return resp, true, nil
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
// - Ensure votes represent at least 2/3 of the total voting power.
func (k *Keeper) verifyAggVotes(ctx context.Context, validators map[common.Address]bool, aggs []*types.AggVote) error {
	for _, agg := range aggs {
		if err := agg.Verify(); err != nil {
			return errors.Wrap(err, "verify aggregate vote")
		}

		// Ensure all votes are from validators in the set
		for _, sig := range agg.Signatures {
			addr := common.BytesToAddress(sig.GetValidatorAddress())
			_, ok := validators[addr]
			if !ok {
				return errors.New("vote from unknown validator")
			}
		}

		// Ensure the block header is in the vote window.
		if resp, err := k.windowCompare(ctx, agg.BlockHeader.ChainId, agg.BlockHeader.Height); err != nil {
			return errors.Wrap(err, "windower")
		} else if resp != 0 {
			return errors.New("vote outside window", "resp", resp)
		}
	}

	return nil
}

// isApproved returns whether the given signatures are approved by the given validators.
// It also returns the signatures to delete (not in the validator set).
func isApproved(sigs []*Signature, valsByPower map[common.Address]int64, total int64) ([]*Signature, bool) {
	var sum int64
	var toDelete []*Signature
	for _, sig := range sigs {
		power, ok := valsByPower[common.BytesToAddress(sig.GetValidatorAddress())]
		if !ok {
			toDelete = append(toDelete, sig)
			continue
		}

		sum += power
	}

	return toDelete, sum > total*2/3
}

func windowCompare(voteWindow uint64, header *types.BlockHeader, latest uint64) int {
	x := header.Height
	mid := latest
	delta := voteWindow

	if x < uintSub(mid, delta) {
		return -1
	} else if x > mid+delta {
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
