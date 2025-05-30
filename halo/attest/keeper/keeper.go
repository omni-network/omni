package keeper

import (
	"bytes"
	"context"

	"github.com/omni-network/omni/halo/attest/types"
	rtypes "github.com/omni-network/omni/halo/registry/types"
	vtypes "github.com/omni-network/omni/halo/valsync/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/umath"
	"github.com/omni-network/omni/lib/xchain"

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

// initialAttestOffset is the first attest offset to attest to for all chains.
const initialAttestOffset uint64 = 1

var _ sdk.ExtendVoteHandler = (*Keeper)(nil).ExtendVote
var _ sdk.VerifyVoteExtensionHandler = (*Keeper)(nil).VerifyVoteExtension

// Keeper is the attestation keeper.
// It keeps tracks of all attestations included on-chain and detects when they are approved.
type Keeper struct {
	attTable       AttestationTable
	sigTable       SignatureTable
	cdc            codec.BinaryCodec
	storeService   store.KVStoreService
	skeeper        baseapp.ValidatorStore
	valProvider    vtypes.ValidatorProvider
	portalRegistry rtypes.PortalRegistry
	namer          types.ChainVerNameFunc
	voter          types.Voter

	voteWindowUp   uint64 // Vote window upper bound delta
	voteWindowDown uint64 // Vote window lower bound delta
	voteExtLimit   uint64
	trimLag        uint64 // Non-consensus chain trim lag
	cTrimLag       uint64 // Consensus chain trim lag

	valAddrCache *valAddrCache
}

// New returns a new attestation keeper.
func New(
	cdc codec.BinaryCodec,
	storeSvc store.KVStoreService,
	skeeper baseapp.ValidatorStore,
	namer types.ChainVerNameFunc,
	voter types.Voter,
	voteWindowUp uint64,
	voteWindowDown uint64,
	voteExtLimit uint64,
	trimLag uint64,
	cTrimLag uint64,
) (*Keeper, error) {
	schema := &ormv1alpha1.ModuleSchemaDescriptor{SchemaFile: []*ormv1alpha1.ModuleSchemaDescriptor_FileEntry{
		{Id: 1, ProtoFileName: File_halo_attest_keeper_attestation_proto.Path()},
	}}

	modDB, err := ormdb.NewModuleDB(schema, ormdb.ModuleDBOptions{KVStoreService: storeSvc})
	if err != nil {
		return nil, errors.Wrap(err, "create module db")
	}

	attStore, err := NewAttestationStore(modDB)
	if err != nil {
		return nil, errors.Wrap(err, "create attestation store")
	}

	if cTrimLag < trimLag {
		return nil, errors.New("consensus trim lag must be greater than or equal to trim lag")
	}

	k := &Keeper{
		attTable:       attStore.AttestationTable(),
		sigTable:       attStore.SignatureTable(),
		cdc:            cdc,
		storeService:   storeSvc,
		skeeper:        skeeper,
		namer:          namer,
		voter:          voter,
		voteWindowUp:   voteWindowUp,
		voteWindowDown: voteWindowDown,
		voteExtLimit:   voteExtLimit,
		trimLag:        trimLag,
		cTrimLag:       cTrimLag,
		portalRegistry: stubPortalRegistry{},
		valAddrCache:   new(valAddrCache),
	}

	return k, nil
}

// SetValidatorProvider sets the validator provider.
func (k *Keeper) SetValidatorProvider(valProvider vtypes.ValidatorProvider) {
	k.valProvider = valProvider
}

// SetPortalRegistry sets the portal registry.
func (k *Keeper) SetPortalRegistry(portalRegistry rtypes.PortalRegistry) {
	k.portalRegistry = portalRegistry
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

	countsByChainVer := make(map[xchain.ChainVersion]int)
	for _, aggVote := range msg.Votes {
		countsByChainVer[aggVote.AttestHeader.XChainVersion()]++

		// Sanity check that all votes are from prev block validators.
		for _, sig := range aggVote.Signatures {
			sigTup, err := sig.ToXChain()
			if err != nil {
				return err
			}
			if !valset.Contains(sigTup.ValidatorAddress) {
				return errors.New("vote from unknown validator [BUG]")
			}
		}

		err := k.addOne(ctx, aggVote, valset.ID)
		if err != nil {
			return errors.Wrap(err, "add one")
		}
	}

	for chainVer, count := range countsByChainVer {
		votesProposed.WithLabelValues(k.namer(chainVer)).Observe(float64(count))
	}

	return nil
}

// addOne adds the given aggregate vote to the store.
// It merges it if the attestation already exists.
//
//nolint:nestif // Ignore for now.
func (k *Keeper) addOne(ctx context.Context, agg *types.AggVote, valSetID uint64) error {
	defer latency("add_one")()

	header := agg.AttestHeader
	attRoot, err := agg.AttestationRoot()
	if err != nil {
		return errors.Wrap(err, "attestation root")
	}

	height, err := umath.ToUint64(sdk.UnwrapSDKContext(ctx).BlockHeight())
	if err != nil {
		return err
	}

	// Get existing attestation (by unique key) or insert new one.
	var attID uint64
	existing, err := k.attTable.GetByAttestationRoot(ctx, attRoot[:])
	insertAtt := ormerrors.IsNotFound(err)
	if insertAtt {
		// Insert new attestation
		attID, err = k.attTable.InsertReturningId(ctx, &Attestation{
			ChainId:         agg.AttestHeader.SourceChainId,
			ConfLevel:       agg.AttestHeader.ConfLevel,
			AttestOffset:    agg.AttestHeader.AttestOffset,
			BlockHeight:     agg.BlockHeader.BlockHeight,
			BlockHash:       agg.BlockHeader.BlockHash,
			MsgRoot:         agg.MsgRoot,
			AttestationRoot: attRoot[:],
			Status:          uint32(Status_Pending),
			ValidatorSetId:  0, // Unknown at this point.
			CreatedHeight:   height,
			FinalizedAttId:  0, // No finalized override yet.
		})
		if err != nil {
			return errors.Wrap(err, "insert")
		}
	} else if err != nil {
		return errors.Wrap(err, "by att unique key")
	} else if existing.GetFinalizedAttId() != 0 {
		log.Debug(ctx, "Ignoring vote for attestation with finalized override",
			"agg_id", attID,
			"chain", k.namer(header.XChainVersion()),
			"attest_offset", header.AttestOffset,
		)

		return nil
	} else if isApprovedByDifferentSet(existing, valSetID) {
		log.Debug(ctx, "Ignoring vote for attestation approved by different validator set",
			"att_id", existing.GetId(),
			"existing_valset_id", existing.GetValidatorSetId(),
			"vote_valset_id", valSetID,
			"chain", k.namer(header.XChainVersion()),
			"attest_offset", header.AttestOffset,
			"sigs", len(agg.Signatures),
		)
		// Technically these new votes could be from validators also in that previous set, but
		// we don't have consistent access to historical validator sets.

		return nil
	} else {
		attID = existing.GetId()
	}

	// Insert signatures
	var insertSig bool
	for _, sig := range agg.Signatures {
		sigTup, err := sig.ToXChain()
		if err != nil {
			return err
		}

		err = k.sigTable.Insert(ctx, &Signature{
			Signature:        sig.GetSignature(),
			ValidatorAddress: sig.GetValidatorAddress(),
			AttId:            attID,
			ChainId:          agg.AttestHeader.GetSourceChainId(),
			ConfLevel:        agg.AttestHeader.GetConfLevel(),
			AttestOffset:     agg.AttestHeader.GetAttestOffset(),
		})

		if errors.Is(err, ormerrors.UniqueKeyViolation) {
			attrs := []any{"agg_id", attID, "chain", k.namer(header.XChainVersion()), "attest_offset", header.AttestOffset, log.Hex7("validator", sig.ValidatorAddress)}
			if ok, err := k.isDoubleSign(ctx, attID, agg, sig); err != nil {
				return err
			} else if ok {
				doubleSignCounter.WithLabelValues(sigTup.ValidatorAddress.Hex()).Inc()
				log.Warn(ctx, "🚨 Ignoring duplicate slashable vote", nil, attrs...)
			} else {
				// Ignore identical duplicate. See https://github.com/omni-network/omni/issues/2286.
				log.Debug(ctx, "Ignoring duplicate vote", attrs...)
			}
		} else if err != nil {
			return errors.Wrap(err, "insert signature")
		} else {
			insertSig = true
		}
	}

	if insertAtt && !insertSig {
		// Revert insertion of new attestations without valid signatures.
		if err = k.attTable.DeleteBy(ctx, AttestationPrimaryKey{}.WithId(attID)); err != nil {
			return errors.Wrap(err, "delete attestation")
		}
	}

	return nil
}

// Approve approves any pending attestations that have quorum signatures from the provided set.
func (k *Keeper) Approve(ctx context.Context, valset ValSet) error {
	defer latency("approve")()

	pendingIdx := AttestationStatusChainIdConfLevelAttestOffsetIndexKey{}.WithStatus(uint32(Status_Pending))
	iter, err := k.attTable.List(ctx, pendingIdx)
	if err != nil {
		return errors.Wrap(err, "list pending")
	}
	defer iter.Close()

	approvedByChain := make(map[xchain.ChainVersion]uint64) // Cache the latest approved attestation offset by chain version.
	for iter.Next() {
		att, err := iter.Value()
		if err != nil {
			return errors.Wrap(err, "value")
		}
		chainVer := att.XChainVersion()
		chainVerName := k.namer(chainVer)

		// Ensure we approve sequentially, not skipping any heights.
		{
			// Populate the cache if not already.
			if _, ok := approvedByChain[chainVer]; !ok {
				latest, found, err := k.latestAttestation(ctx, att.XChainVersion())
				if err != nil {
					return errors.Wrap(err, "latest approved")
				} else if found {
					approvedByChain[chainVer] = latest.GetAttestOffset()
				}
			}
			head, ok := approvedByChain[chainVer]
			if !ok && att.GetAttestOffset() != initialAttestOffset {
				// Only start attesting from offset==1
				continue
			} else if ok && head+1 != att.GetAttestOffset() {
				// This isn't the next attestation to approve, so we can't approve it yet.
				continue
			}
		}

		sigs, err := k.getSigs(ctx, att.GetId())
		if err != nil {
			return errors.Wrap(err, "get att signatures")
		}

		{
			// Calculate pending blocks; safe to ignore errors since metrics is non-critical.
			current, _ := umath.ToUint64(sdk.UnwrapSDKContext(ctx).BlockHeight())
			delta := umath.SubtractOrZero(current, att.GetCreatedHeight())
			pendingBlocks.WithLabelValues(chainVerName).Set(float64(delta))
		}

		setMetrics := func(att *Attestation) {
			approvedHeight.WithLabelValues(chainVerName).Set(float64(att.GetBlockHeight()))
			approvedOffset.WithLabelValues(chainVerName).Set(float64(att.GetAttestOffset()))
		}

		toDelete, ok, err := isApproved(sigs, valset)
		if err != nil {
			return err
		} else if !ok {
			// Check if there is a finalized attestation that overrides this one.
			if ok, err := k.maybeOverrideFinalized(ctx, att); err != nil {
				return err
			} else if ok {
				setMetrics(att)
				approvedByChain[chainVer] = att.GetAttestOffset()
			}

			continue
		}

		for _, sig := range toDelete {
			addr, err := sig.ValidatorEthAddress()
			if err != nil {
				return err
			}

			discardedVotesCounter.WithLabelValues(addr.Hex(), chainVerName).Inc()

			if err = k.sigTable.Delete(ctx, sig); err != nil {
				return errors.Wrap(err, "delete sig")
			}
		}

		// Update status
		att.Status = uint32(Status_Approved)
		att.ValidatorSetId = valset.ID
		err = k.attTable.Update(ctx, att)
		if err != nil {
			return errors.Wrap(err, "save")
		}

		setMetrics(att)
		approvedByChain[chainVer] = att.GetAttestOffset()

		log.Debug(ctx, "📬 Approved attestation",
			"chain", chainVerName,
			"attest_offset", att.GetAttestOffset(),
			"height", att.GetBlockHeight(),
			log.Hex7("hash", att.GetBlockHash()),
		)
	}

	// Trim votes behind minimum vote-window
	minVoteWindows := make(map[xchain.ChainVersion]uint64)
	for chainVer, head := range approvedByChain {
		minVoteWindows[chainVer] = umath.SubtractOrZero(head, k.voteWindowDown)
	}

	count := k.voter.TrimBehind(minVoteWindows)
	if count > 0 {
		log.Warn(ctx, "Trimmed votes behind vote-window (expected if node was struggling)", nil, "count", count)
	}

	return nil
}

// ListAttestationsFrom returns the subsequent approved attestations from the provided offset (inclusive).
func (k *Keeper) ListAttestationsFrom(ctx context.Context, chainID uint64, confLevel uint32, offset uint64, max uint64) ([]*types.Attestation, error) {
	defer latency("attestations_from")()
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	consensusID, err := netconf.ConsensusChainIDStr2Uint64(sdkCtx.ChainID())
	if err != nil {
		return nil, errors.Wrap(err, "parse chain id")
	}

	from := AttestationStatusChainIdConfLevelAttestOffsetIndexKey{}.WithStatusChainIdConfLevelAttestOffset(uint32(Status_Approved), chainID, confLevel, offset)
	to := AttestationStatusChainIdConfLevelAttestOffsetIndexKey{}.WithStatusChainIdConfLevelAttestOffset(uint32(Status_Approved), chainID, confLevel, offset+max)

	iter, err := k.attTable.ListRange(ctx, from, to)
	if err != nil {
		return nil, errors.Wrap(err, "list range")
	}
	defer iter.Close()

	var resp []*types.Attestation
	next := offset
	for iter.Next() {
		att, err := iter.Value()
		if err != nil {
			return nil, errors.Wrap(err, "value")
		}

		if att.GetAttestOffset() != next {
			break
		}
		next++

		// If this attestation is overridden by a finalized attestation, use that instead.
		if att.GetFinalizedAttId() != 0 {
			att, err = k.attTable.Get(ctx, att.GetFinalizedAttId())
			if err != nil {
				return nil, errors.Wrap(err, "get finalized attestation")
			}
		}

		sigs, err := k.getSigTuples(ctx, att.GetId())
		if err != nil {
			return nil, errors.Wrap(err, "get att sigs")
		}

		resp = append(resp, toProto(att, sigs, consensusID))
	}

	return resp, nil
}

// maybeOverrideFinalized returns the approved finalized attestation and true for the provided fuzzy attestation if it exists.
func (k *Keeper) maybeOverrideFinalized(ctx context.Context, att *Attestation) (bool, error) {
	if att.GetStatus() != uint32(Status_Pending) {
		return false, errors.New("attestation not pending [BUG]")
	}

	if att.GetConfLevel() == uint32(xchain.ConfFinalized) {
		return false, nil // Only fuzzy attestations are overwritten with finalized attestations.
	}

	finalizedIdx := AttestationStatusChainIdConfLevelAttestOffsetIndexKey{}.WithStatusChainIdConfLevelAttestOffset(uint32(Status_Approved), att.GetChainId(), uint32(xchain.ConfFinalized), att.GetAttestOffset())
	iter, err := k.attTable.List(ctx, finalizedIdx)
	if err != nil {
		return false, errors.Wrap(err, "list finalized")
	}
	defer iter.Close()

	if !iter.Next() {
		// No finalized attestation found.
		return false, nil
	}

	finalized, err := iter.Value()
	if err != nil {
		return false, errors.Wrap(err, "value finalized")
	}

	if iter.Next() {
		return false, errors.New("multiple finalized attestation found [BUG]")
	} else if finalized.GetFinalizedAttId() != 0 {
		return false, errors.New("finalized attestation has finalized attestation [BUG]")
	}

	att.FinalizedAttId = finalized.GetId()
	att.Status = uint32(Status_Approved)
	if err = k.attTable.Update(ctx, att); err != nil {
		return false, errors.Wrap(err, "update attestation")
	}

	log.Debug(ctx, "📬 Fuzzy attestation overridden by finalized",
		"chain", k.namer(att.XChainVersion()),
		"attest_offset", att.GetAttestOffset(),
		"height", att.GetBlockHeight(),
		log.Hex7("hash", att.GetBlockHash()),
	)

	return true, nil
}

// latestAttestation returns the latest approved attestation for the given chain or
// false if none is found.
func (k *Keeper) latestAttestation(ctx context.Context, version xchain.ChainVersion) (*Attestation, bool, error) {
	defer latency("latest_attestation")()

	idx := AttestationStatusChainIdConfLevelAttestOffsetIndexKey{}.WithStatusChainIdConfLevel(uint32(Status_Approved), version.ID, uint32(version.ConfLevel))
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
		return nil, false, errors.New("multiple attestations found")
	}

	// If this attestation is overridden by a finalized attestation, return that instead.
	if att.GetFinalizedAttId() != 0 {
		att, err := k.attTable.Get(ctx, att.GetFinalizedAttId())
		if err != nil {
			return nil, false, errors.Wrap(err, "get finalized attestation")
		}

		return att, true, nil
	}

	return att, true, nil
}

// earliestAttestation returns the earliest approved attestation for the given chain currently found in state,
// or false if none is found.
func (k *Keeper) earliestAttestation(ctx context.Context, version xchain.ChainVersion) (*Attestation, bool, error) {
	defer latency("earliest_attestation")()

	idx := AttestationStatusChainIdConfLevelAttestOffsetIndexKey{}.WithStatusChainIdConfLevel(uint32(Status_Approved), version.ID, uint32(version.ConfLevel))
	iter, err := k.attTable.List(ctx, idx, ormlist.DefaultLimit(1))
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
		return nil, false, errors.New("multiple attestations found")
	}

	// If this attestation is overridden by a finalized attestation, return that instead.
	if att.GetFinalizedAttId() != 0 {
		att, err := k.attTable.Get(ctx, att.GetFinalizedAttId())
		if err != nil {
			return nil, false, errors.Wrap(err, "get finalized attestation")
		}

		return att, true, nil
	}

	return att, true, nil
}

// listAllAttestations returns all attestations for the given chain and status and attestOffset up to a maximum of 100.
func (k *Keeper) listAllAttestations(ctx context.Context, version xchain.ChainVersion, status Status, attestOffset uint64) ([]*types.Attestation, error) {
	defer latency("list_all_attestations")()
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	consensusID, err := netconf.ConsensusChainIDStr2Uint64(sdkCtx.ChainID())
	if err != nil {
		return nil, errors.Wrap(err, "get consensus chain id")
	}

	const limit = 100

	start := AttestationStatusChainIdConfLevelAttestOffsetIndexKey{}.WithStatusChainIdConfLevelAttestOffset(status.Uint32(), version.ID, uint32(version.ConfLevel), attestOffset)
	end := AttestationStatusChainIdConfLevelAttestOffsetIndexKey{}.WithStatusChainIdConfLevel(status.Uint32(), version.ID, uint32(version.ConfLevel))
	iter, err := k.attTable.ListRange(ctx, start, end)
	if err != nil {
		return nil, errors.Wrap(err, "list")
	}
	defer iter.Close()

	var resp []*types.Attestation
	for iter.Next() {
		att, err := iter.Value()
		if err != nil {
			return nil, errors.Wrap(err, "value")
		}

		sigs, err := k.getSigTuples(ctx, att.GetId())
		if err != nil {
			return nil, errors.Wrap(err, "get att sigs")
		}

		resp = append(resp, toProto(att, sigs, consensusID))

		if len(resp) >= limit {
			break
		}
	}

	return resp, nil
}

// getSigs returns the signatures for the given attestation ID.
func (k *Keeper) getSigs(ctx context.Context, attID uint64) ([]*Signature, error) {
	attIDIdx := SignatureAttIdValidatorAddressIndexKey{}.WithAttId(attID)
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

// getSigTuples returns the signature tuples for the given attestation ID.
func (k *Keeper) getSigTuples(ctx context.Context, attID uint64) ([]*types.SigTuple, error) {
	pbsigs, err := k.getSigs(ctx, attID)
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

	return sigs, nil
}

func (k *Keeper) BeginBlock(ctx context.Context) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	consensusID, err := netconf.ConsensusChainIDStr2Uint64(sdkCtx.ChainID())
	if err != nil {
		return errors.Wrap(err, "parse chain id")
	}

	head, err := umath.ToUint64(sdkCtx.BlockHeight())
	if err != nil {
		return err
	}

	before := umath.SubtractOrZero(head, k.trimLag)
	cBefore := umath.SubtractOrZero(head, k.cTrimLag)

	return k.deleteBefore(ctx, before, consensusID, cBefore)
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
	ctx = ctx.WithContext(log.WithCtx(ctx.Context(), "height", ctx.BlockHeight()))

	cChainID, err := netconf.ConsensusChainIDStr2Uint64(ctx.ChainID())
	if err != nil {
		return nil, errors.Wrap(err, "parse chain id")
	}

	votes := k.voter.GetAvailable()

	// Filter by vote window and if limit exceeded.
	countsByChainVer := make(map[xchain.ChainVersion]int)
	duplicate := make(map[xchain.AttestHeader]bool)
	var filtered []*types.Vote
	for _, vote := range votes {
		if err := vote.Verify(); err != nil {
			return nil, errors.Wrap(err, "verify vote")
		}
		if err := verifyHeaderChains(ctx, cChainID, k.portalRegistry, vote.AttestHeader, vote.BlockHeader); err != nil {
			return nil, errors.Wrap(err, "check supported chain")
		}

		if duplicate[vote.AttestHeader.ToXChain()] {
			doubleSignCounter.WithLabelValues(k.voter.LocalAddress().Hex()).Inc()
			log.Warn(ctx, "🚨 Skipping own duplicate slashable vote [BUG]", nil, "chain", k.namer(vote.AttestHeader.XChainVersion()), "attest_offset", vote.AttestHeader.AttestOffset)

			continue
		}
		duplicate[vote.AttestHeader.ToXChain()] = true

		if cmp, err := k.windowCompare(ctx, vote.AttestHeader.XChainVersion(), vote.AttestHeader.AttestOffset); err != nil {
			return nil, errors.Wrap(err, "windower")
		} else if cmp != 0 {
			// Skip votes not in the window
			continue
		}
		countsByChainVer[vote.AttestHeader.XChainVersion()]++
		filtered = append(filtered, vote)

		if umath.Len(filtered) >= k.voteExtLimit {
			break
		}
	}

	bz, err := proto.Marshal(&types.Votes{
		Votes: filtered,
	})
	if err != nil {
		return nil, errors.Wrap(err, "marshal atts")
	}

	for chainVer, count := range countsByChainVer {
		votesExtended.WithLabelValues(k.namer(chainVer)).Observe(float64(count))
	}

	log.Info(ctx, "Voter extending vote", types.VoteLogs(filtered)...)

	return &abci.ResponseExtendVote{
		VoteExtension: bz,
	}, nil
}

// VerifyVoteExtension verifies a vote extension.
//
// Note this code assumes that cometBFT will only call this function for an active validator in the current set.
func (k *Keeper) VerifyVoteExtension(ctx sdk.Context, req *abci.RequestVerifyVoteExtension) (
	*abci.ResponseVerifyVoteExtension, error,
) {
	respAccept := &abci.ResponseVerifyVoteExtension{
		Status: abci.ResponseVerifyVoteExtension_ACCEPT,
	}
	respReject := &abci.ResponseVerifyVoteExtension{
		Status: abci.ResponseVerifyVoteExtension_REJECT,
	}

	height, err := umath.ToUint64(ctx.BlockHeight())
	if err != nil {
		return nil, err
	}

	_, ok, err := k.parseAndVerifyVoteExtension(ctx, req.ValidatorAddress, req.VoteExtension, height)
	if err != nil {
		log.Warn(ctx, "Rejecting vote extension", err, log.Hex7("validator", req.ValidatorAddress))
		return respReject, nil
	} else if !ok {
		log.Warn(ctx, "Rejecting vote extension containing vote behind window", nil, log.Hex7("validator", req.ValidatorAddress))
		return respReject, nil
	}

	return respAccept, nil
}

// parseAndVerifyVoteExtension returns a list of valid vote extensions and true if all votes are valid,
// or an error if any validation failed (except "vote-behind-window").
//
// vote-behind-window:
// - is not valid,
// - but is not always considered an error (as it is expected in PrepareProposal).
// - false is returned in this case,
// - to indicate that not all votes are valid (nor returned).
func (k *Keeper) parseAndVerifyVoteExtension(
	ctx sdk.Context,
	valAddr []byte,
	voteExt []byte,
	voteHeight uint64,
) ([]*types.Vote, bool, error) {
	cChainID, err := netconf.ConsensusChainIDStr2Uint64(ctx.ChainID())
	if err != nil {
		return nil, false, errors.Wrap(err, "parse chain id")
	}

	// Get the ethereum address of the validator
	ethAddr, err := k.getValEthAddr(ctx, valAddr, voteHeight)
	if err != nil {
		return nil, false, err // This error should never occur
	}

	votes, ok, err := votesFromExtension(voteExt)
	if err != nil {
		return nil, false, errors.Wrap(err, "votes from extension")
	} else if !ok {
		return nil, true, nil // Empty vote extension is fine
	} else if umath.Len(votes.Votes) > k.voteExtLimit {
		return nil, false, errors.New("vote extension limit exceeded", "count", len(votes.Votes), "limit", k.voteExtLimit)
	}

	duplicate := make(map[common.Hash]bool)          // Detect identical duplicate votes (same AttestationRoot)
	doubleSign := make(map[xchain.AttestHeader]bool) // Detect double sign votes (same AttestHeader)
	var resp []*types.Vote
	for _, vote := range votes.Votes {
		if err := vote.Verify(); err != nil {
			return nil, false, errors.Wrap(err, "verify vote")
		}

		attRoot, err := vote.AttestationRoot()
		if err != nil {
			return nil, false, errors.Wrap(err, "att root")
		}
		if duplicate[attRoot] {
			return nil, false, errors.New("duplicate identical vote")
		}
		duplicate[attRoot] = true

		if doubleSign[vote.AttestHeader.ToXChain()] {
			doubleSignCounter.WithLabelValues(ethAddr.Hex()).Inc()

			return nil, false, errors.New("duplicate slashable vote")
		}
		doubleSign[vote.AttestHeader.ToXChain()] = true

		// Ensure the votes are from the requesting validator itself.
		if !bytes.Equal(vote.Signature.ValidatorAddress, ethAddr[:]) {
			return nil, false, errors.New("mismatching vote and req validator address", "vote", ethAddr, "req", vote.Signature.ValidatorAddress)
		}

		if err := verifyHeaderChains(ctx, cChainID, k.portalRegistry, vote.AttestHeader, vote.BlockHeader); err != nil {
			return nil, false, errors.Wrap(err, "verify chain headers", "chain", k.namer(vote.AttestHeader.XChainVersion()))
		}

		if cmp, err := k.windowCompare(ctx, vote.AttestHeader.XChainVersion(), vote.AttestHeader.AttestOffset); err != nil {
			return nil, false, errors.Wrap(err, "window compare")
		} else if cmp > 0 {
			return nil, false, errors.New("vote ahead of window")
		} else if cmp < 0 {
			// Vote-behind-window is expected in PrepareProposal, just don't add to response.
			continue
		}

		resp = append(resp, vote)
	}

	return resp, len(votes.Votes) == len(resp), nil
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
	height, err := umath.ToUint64(sdk.UnwrapSDKContext(ctx).BlockHeight())
	if err != nil {
		return ValSet{}, err
	}

	prevBlock := umath.SubtractOrZero(height, 1) // This is only called AFTER genesis, so safe to subtract 1.

	resp, err := k.valProvider.ActiveSetByHeight(ctx, prevBlock)
	if err != nil {
		return ValSet{}, err
	}

	valsByPower := make(map[common.Address]int64)
	for _, val := range resp.Validators {
		ethAddr, err := val.EthereumAddress()
		if err != nil {
			return ValSet{}, err
		}
		valsByPower[ethAddr] = val.Power
	}

	return ValSet{
		ID:   resp.Id,
		Vals: valsByPower,
	}, nil
}

func (k *Keeper) windowCompare(ctx context.Context, chainVer xchain.ChainVersion, offset uint64) (int, error) {
	latest, exists, err := k.latestAttestation(ctx, chainVer)
	if err != nil {
		return 0, err
	}

	latestOffset := initialAttestOffset // Use initial offset if attestation doesn't exist.
	if exists {
		latestOffset = latest.GetAttestOffset()
	}

	return windowCompare(k.voteWindowDown, k.voteWindowUp, latestOffset, offset), nil
}

// verifyAggVotes verifies the given aggregates votes:
// - Ensure all aggregate votes are valid.
// - Ensure all votes are for supported chains.
// - Ensure all aggregation is valid; no duplicate aggregate votes.
// - Ensure the vote extension limit is not exceeded per validator.
// - Ensure all votes are from validators in the provided set.
// - Ensure all votes are unique per validator.
// - Ensure the vote block header is in the vote window.
func (k *Keeper) verifyAggVotes(
	ctx context.Context,
	cChainID uint64,
	valset ValSet,
	aggs []*types.AggVote,
	windowCompareFunc func(context.Context, xchain.ChainVersion, uint64) (int, error), // Aliased for testing
) error {
	duplicate := make(map[common.Hash]bool)         // Detects duplicate aggregate votes.
	countsPerVal := make(map[common.Address]uint64) // Enforce vote extension limit.
	for _, agg := range aggs {
		if err := agg.Verify(); err != nil {
			return errors.Wrap(err, "verify aggregate vote")
		}
		errAttrs := []any{"chain", k.namer(agg.AttestHeader.XChainVersion()), "attest_offset", agg.AttestHeader.AttestOffset}

		if err := verifyHeaderChains(ctx, cChainID, k.portalRegistry, agg.AttestHeader, agg.BlockHeader); err != nil {
			return errors.Wrap(err, "check supported chain")
		}

		attRoot, err := agg.AttestationRoot()
		if err != nil {
			return errors.Wrap(err, "attestation root")
		}

		if duplicate[attRoot] {
			return errors.New("invalid duplicate aggregate votes", errAttrs...) // Note this is duplicate aggregates, which may contain non-overlapping votes so not technically slashable.
		}
		duplicate[attRoot] = true

		// Ensure all votes are from unique validators in the set
		duplicateSig := make(map[common.Address]bool) // Enforce vote extension limit.
		for _, sig := range agg.Signatures {
			addr, err := sig.ValidatorEthAddress()
			if err != nil {
				return err
			}

			if duplicateSig[addr] {
				return errors.New("duplicate validator vote", append(errAttrs, "validator", addr)...)
			}
			duplicateSig[addr] = true

			if !valset.Contains(addr) {
				return errors.New("vote from unknown validator", append(errAttrs, "validator", addr)...)
			}

			countsPerVal[addr]++
			if countsPerVal[addr] > k.voteExtLimit {
				return errors.New("vote extension limit exceeded", append(errAttrs, "validator", addr)...)
			}
		}

		// Ensure the block header is in the vote window.
		if resp, err := windowCompareFunc(ctx, agg.AttestHeader.XChainVersion(), agg.AttestHeader.AttestOffset); err != nil {
			return errors.Wrap(err, "windower")
		} else if resp != 0 {
			errAttrs = append(errAttrs, "resp", resp)

			return errors.New("vote outside window", errAttrs...)
		}
	}

	return nil
}

// deleteBefore deletes all attestations and signatures before the given height (inclusive).
// Consensus chain attestations are compared against cHeight (inclusive).
// Note this always deletes block 0, but genesis block doesn't contain any attestations.
func (k *Keeper) deleteBefore(ctx context.Context, height uint64, consensusID uint64, cHeight uint64) error {
	defer latency("delete_before")()

	// Create latest- and earliest- read-through caches to mitigate DB reads.
	latestOffset := newLatestLookupCache(k)
	earliestOffset := newEarliestLookupCache(k)

	// Get all supported confirmation levels.
	confLevels, err := k.portalRegistry.ConfLevels(ctx)
	if err != nil {
		return errors.Wrap(err, "conf levels")
	}

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
		} else if att.GetChainId() == consensusID && att.GetCreatedHeight() > cHeight {
			// Consensus chain attestations are deleted much later, since they have possible valset update dependencies.
			continue
		}

		// Wait for the attestation to exit the vote window before deleting it.
		// This includes pending attestations.
		if latest, ok, err := latestOffset(ctx, att.XChainVersion()); err != nil {
			return err
		} else if !ok || windowCompare(k.voteWindowDown, k.voteWindowUp, latest, att.GetAttestOffset()) >= 0 {
			continue
		}

		var fuzzyDepsFound bool
		for _, fuzzyVer := range fuzzyDependents(att.XChainVersion(), confLevels) {
			// For finalized attestations (with fuzzy dependents), only delete them AFTER all fuzzy attestations have been deleted.
			// This avoids deleting finalized overrides for current or future fuzzy attestations.
			earliestFuzzy, ok, err := earliestOffset(ctx, fuzzyVer)
			if err != nil {
				return err
			} else if !ok || att.GetAttestOffset() >= earliestFuzzy {
				// If !ok, then no fuzzy attestations have been created, so we can't delete finalized.
				// The earliest fuzzy must be AFTER the finalized att.
				fuzzyDepsFound = true
				break
			}
		}
		if fuzzyDepsFound {
			continue // Skip deleting finalized attestation.
		}

		if err := k.instrumentVotes(ctx, att); err != nil {
			return errors.Wrap(err, "instrument votes")
		}

		// Delete signatures
		if err := k.sigTable.DeleteBy(ctx, SignatureAttIdValidatorAddressIndexKey{}.WithAttId(att.GetId())); err != nil {
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

// instrumentVotes tracks basic voter performance by instrumenting votes.
// It tracks whether validators are voting vs voting late vs not voting.
func (k *Keeper) instrumentVotes(ctx context.Context, att *Attestation) error {
	// Discard votes if attestation never approved or if overridden by finalized.
	discardVotes := att.GetStatus() != uint32(Status_Approved) || att.GetFinalizedAttId() != 0

	chainVerName := k.namer(att.XChainVersion())
	sigs, err := k.getSigs(ctx, att.GetId())
	if err != nil {
		return errors.Wrap(err, "get att sigs")
	}

	included := make(map[common.Address]bool)
	for _, sig := range sigs {
		addr, err := sig.ValidatorEthAddress()
		if err != nil {
			return err
		}

		included[addr] = true
		if discardVotes {
			discardedVotesCounter.WithLabelValues(addr.Hex(), chainVerName).Inc()
		}
	}

	// Do not set other metrics if we are discarding votes.
	if discardVotes {
		return nil
	}

	if att.GetValidatorSetId() == 0 { // Sanity check
		return errors.New("missing validator set id for approved non-overridden attestation [BUG]", "att_id", att.GetId())
	}

	valset, err := k.valProvider.ValidatorSet(ctx, &vtypes.ValidatorSetRequest{Id: att.GetValidatorSetId()})
	if err != nil {
		return errors.Wrap(err, "validator set", "id", att.GetValidatorSetId())
	}

	for _, val := range valset.Validators {
		addr, err := val.EthereumAddress()
		if err != nil {
			return errors.Wrap(err, "validator address")
		}

		expectedVotesCounter.WithLabelValues(addr.Hex(), chainVerName).Inc()
		approvedVotesCounter.WithLabelValues(addr.Hex(), chainVerName).Add(boolToFloat(included[addr]))
		missingVotesCounter.WithLabelValues(addr.Hex(), chainVerName).Add(boolToFloat(!included[addr]))
		discardedVotesCounter.WithLabelValues(addr.Hex(), chainVerName).Add(0)
	}

	return nil
}

// isDoubleSign returns true if the vote qualifies as a slashable double sign.
func (k *Keeper) isDoubleSign(ctx context.Context, attID uint64, agg *types.AggVote, sig *types.SigTuple) (bool, error) {
	// Check if this is a duplicate of an existing vote
	if identicalVote, err := k.sigTable.GetByAttIdValidatorAddress(ctx, attID, sig.ValidatorAddress); err == nil {
		// Sanity check that this is indeed an identical vote
		if !bytes.Equal(identicalVote.GetSignature(), sig.GetSignature()) {
			return false, errors.New("different signature for identical vote [BUG]")
		}

		return false, nil
	} else if !errors.Is(err, ormerrors.NotFound) {
		return false, errors.Wrap(err, "get identical vote")
	} // else identical vote doesn't exist

	doubleSign, err := k.sigTable.HasByChainIdConfLevelAttestOffsetValidatorAddress(ctx, agg.BlockHeader.ChainId, agg.AttestHeader.ConfLevel, agg.AttestHeader.AttestOffset, sig.ValidatorAddress)
	if err != nil {
		return false, errors.Wrap(err, "check double sign")
	} else if !doubleSign {
		return false, errors.New("duplicate vote neither identical nor double sign [BUG]")
	} // else double sign

	return true, nil
}

// isApproved returns whether the given signatures are approved by the given validators.
// It also returns the signatures to delete (not in the validator set).
func isApproved(sigs []*Signature, valset ValSet) ([]*Signature, bool, error) {
	var sum int64
	var toDelete []*Signature
	for _, sig := range sigs {
		addr, err := sig.ValidatorEthAddress()
		if err != nil {
			return nil, false, err
		}

		power, ok := valset.Vals[addr]
		if !ok {
			toDelete = append(toDelete, sig)
			continue
		}

		sum += power
	}

	isApproved := sum > valset.TotalPower()*2/3

	return toDelete, isApproved, nil
}

func verifyHeaderChains(ctx context.Context, cChainID uint64, registry rtypes.PortalRegistry, attHeader *types.AttestHeader, blockHeader *types.BlockHeader) error {
	if attHeader.SourceChainId != blockHeader.ChainId {
		return errors.New("mismatching chain id", "block", blockHeader.ChainId, "att", attHeader.SourceChainId)
	}

	if attHeader.ConsensusChainId != cChainID {
		return errors.New("invalid consensus chain id", "expected", cChainID, "got", attHeader.ConsensusChainId)
	}

	chainConfLevels, err := registry.ConfLevels(ctx)
	if err != nil {
		return errors.Wrap(err, "supported conf levels")
	}

	confLevels, ok := chainConfLevels[blockHeader.ChainId]
	if !ok {
		return errors.New("missing conf levels", "chain", blockHeader.ChainId)
	}

	var found bool
	for _, confLevel := range confLevels {
		if uint32(confLevel) == attHeader.ConfLevel {
			found = true
			break
		}
	}
	if !found {
		return errors.New("unsupported conf level", "chain", blockHeader.ChainId, "conf", xchain.ConfLevel(attHeader.ConfLevel).String())
	}

	return nil
}

// windowCompare returns -1 if x < mid-voteWindowDown, 1 if x > mid+voteWindowUp, else 0.
func windowCompare(voteWindowDown uint64, voteWindowUp, mid uint64, x uint64) int {
	if x < umath.SubtractOrZero(mid, voteWindowDown) {
		return -1
	} else if x > mid+voteWindowUp {
		return 1
	}

	return 0
}

// isApprovedByDifferentSet returns true if the attestation is approved by a different validator set.
func isApprovedByDifferentSet(att *Attestation, valSetID uint64) bool {
	if att.GetStatus() != uint32(Status_Approved) {
		return false
	}

	return att.GetValidatorSetId() != valSetID
}

// toProto converts from the keeper.Attestation type to the types.Attestation type.
func toProto(att *Attestation, sigs []*types.SigTuple, cChainID uint64) *types.Attestation {
	return &types.Attestation{
		AttestHeader: &types.AttestHeader{
			ConsensusChainId: cChainID,
			SourceChainId:    att.GetChainId(),
			ConfLevel:        att.GetConfLevel(),
			AttestOffset:     att.GetAttestOffset(),
		},
		BlockHeader: &types.BlockHeader{
			ChainId:     att.GetChainId(),
			BlockHeight: att.GetBlockHeight(),
			BlockHash:   att.GetBlockHash(),
		},
		ValidatorSetId: att.GetValidatorSetId(),
		MsgRoot:        att.GetMsgRoot(),
		Signatures:     sigs,
	}
}

// stubPortalRegistry is a stub implementation of the portal registry.
// It doesn't support any chains.
type stubPortalRegistry struct{}

func (stubPortalRegistry) ConfLevels(context.Context) (map[uint64][]xchain.ConfLevel, error) {
	return map[uint64][]xchain.ConfLevel{}, nil
}
