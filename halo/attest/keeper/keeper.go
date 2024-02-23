package keeper

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/halo/comet"
	"github.com/omni-network/omni/lib/engine"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/log"

	abci "github.com/cometbft/cometbft/abci/types"
	cmttypes "github.com/cometbft/cometbft/types"

	"github.com/ethereum/go-ethereum/common"

	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"
	"cosmossdk.io/core/store"
	"cosmossdk.io/orm/model/ormdb"
	"cosmossdk.io/orm/types/ormerrors"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	skeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
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
	ethCl        engine.API
	voter        types.Voter
	skeeper      *skeeper.Keeper // TODO(corver): Define a interface for the methods we use.
	cmtAPI       comet.API
	namer        types.ChainNameFunc
}

// NewKeeper returns a new attestation keeper.
func NewKeeper(
	cdc codec.BinaryCodec,
	storeSvc store.KVStoreService,
	ethCl engine.API,
	skeeper *skeeper.Keeper,
	voter types.Voter,
	namer types.ChainNameFunc,
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

	return &Keeper{
		attTable:     attstore.AttestationTable(),
		sigTable:     attstore.SignatureTable(),
		cdc:          cdc,
		storeService: storeSvc,
		ethCl:        ethCl,
		skeeper:      skeeper,
		voter:        voter,
		namer:        namer,
	}, nil
}

// SetCometAPI sets the comet API client.
func (k *Keeper) SetCometAPI(cmtAPI comet.API) {
	k.cmtAPI = cmtAPI
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

	var attID uint64
	exiting, err := k.attTable.GetByChainIdHeightHash(ctx, header.ChainId, header.Height, header.Hash)
	if ormerrors.IsNotFound(err) {
		// Insert new attestation
		attID, err = k.attTable.InsertReturningId(ctx, &Attestation{
			ChainId:        agg.BlockHeader.ChainId,
			Height:         agg.BlockHeader.Height,
			Hash:           agg.BlockHeader.Hash,
			BlockRoot:      agg.BlockRoot,
			Status:         int32(Status_Pending),
			ValidatorsHash: nil, // Unknown at this point.
		})
		if err != nil {
			return errors.Wrap(err, "insert")
		}
	} else if err != nil {
		return errors.Wrap(err, "by block header")
	} else if !bytes.Equal(exiting.GetBlockRoot(), agg.BlockRoot) {
		return errors.New("mismatching block root")
	} else {
		attID = exiting.GetId()
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
				"chain_id", agg.BlockHeader.ChainId,
				"height", agg.BlockHeader.Height,
				log.Hex7("validator", sig.ValidatorAddress),
			)
		} else if err != nil {
			return errors.Wrap(err, "insert signature")
		}
	}

	return nil
}

// Approve approves any pending attestations that have quorum signatures form the provided set.
func (k *Keeper) Approve(ctx context.Context, valset *cmttypes.ValidatorSet) error {
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

	skip := make(map[uint64]bool) // Skip processing chains as soon as a pending attestation is found.
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
		err = k.attTable.Save(ctx, att)
		if err != nil {
			return errors.Wrap(err, "save")
		}

		approvedHeight.WithLabelValues(k.namer(att.GetChainId())).Set(float64(att.GetHeight()))
	}

	return nil
}

// attestationFrom returns the subsequent approved attestations from the provided height (inclusive).
func (k *Keeper) attestationFrom(ctx context.Context, chainID uint64, height uint64, max uint64,
) ([]*types.Attestation, error) {
	from := AttestationStatusChainIdHeightIndexKey{}.WithStatusChainIdHeight(
		int32(Status_Approved), chainID, height)
	to := AttestationStatusChainIdHeightIndexKey{}.WithStatusChainIdHeight(
		int32(Status_Approved), chainID, height+max)

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
			ValidatorsHash: att.GetValidatorsHash(),
			BlockRoot:      att.GetBlockRoot(),
			Signatures:     sigs,
		})
	}

	return resp, nil
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

	valset, err := k.cmtAPI.Validators(ctx, sdkCtx.BlockHeight()-1) // Get the validators for the previous block.
	if err != nil {
		return errors.Wrap(err, "fetch validators")
	}

	return k.Approve(ctx, valset)
}

// ExtendVote extends a vote with application-injected data (vote extensions).
func (k *Keeper) ExtendVote(ctx sdk.Context, _ *abci.RequestExtendVote) (*abci.ResponseExtendVote, error) {
	votes := k.voter.GetAvailable()

	// TODO(corver): Only include attestations in window, also ensure max size.
	bz, err := proto.Marshal(&types.Votes{
		Votes: votes,
	})
	if err != nil {
		return nil, errors.Wrap(err, "marshal atts")
	}

	// Make nice logs
	heights := make(map[uint64][]uint64)
	for _, vote := range votes {
		heights[vote.BlockHeader.ChainId] = append(heights[vote.BlockHeader.ChainId], vote.BlockHeader.Height)
	}
	attrs := []any{
		slog.Int("votes", len(votes)),
		log.Hex7("validator", k.voter.LocalAddress().Bytes()),
	}
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
func (k *Keeper) VerifyVoteExtension(_ sdk.Context, _ *abci.RequestVerifyVoteExtension) (
	*abci.ResponseVerifyVoteExtension, error,
) {
	// TODO(corver): Figure out what to verify. E.g. outside window or too big.
	return &abci.ResponseVerifyVoteExtension{
		Status: abci.ResponseVerifyVoteExtension_ACCEPT,
	}, nil
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
