package aggregate

import (
	"bytes"
	"context"

	pb "github.com/omni-network/omni/halo/aggregate/v1"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/xchain"

	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/ethereum/go-ethereum/common"

	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"
	"cosmossdk.io/core/store"
	"cosmossdk.io/orm/model/ormdb"
	"cosmossdk.io/orm/types/ormerrors"
)

// Keeper is the aggregate attestation keeper.
// It keeps tracks of all attestations included on-chain and detects when they are approved.
type Keeper struct {
	aggTable pb.AggAttestationTable
	sigTable pb.AttSignatureTable
}

// NewKeeper returns a new aggregate attestation keeper.
func NewKeeper(storeSvc store.KVStoreService) (Keeper, error) {
	schema := &ormv1alpha1.ModuleSchemaDescriptor{SchemaFile: []*ormv1alpha1.ModuleSchemaDescriptor_FileEntry{
		{Id: 1, ProtoFileName: pb.File_halo_aggregate_v1_aggregate_proto.Path()},
	}}

	modDB, err := ormdb.NewModuleDB(schema, ormdb.ModuleDBOptions{KVStoreService: storeSvc})
	if err != nil {
		return Keeper{}, errors.Wrap(err, "create module db")
	}

	aggstore, err := pb.NewAggregateStore(modDB)
	if err != nil {
		return Keeper{}, errors.Wrap(err, "create aggregate store")
	}

	return Keeper{
		aggTable: aggstore.AggAttestationTable(),
		sigTable: aggstore.AttSignatureTable(),
	}, nil
}

// Add adds the given aggregate attestations to the store.
// It merges the aggregate if it already exists.
func (a Keeper) Add(ctx context.Context, aggs []xchain.AggAttestation) error {
	for _, agg := range aggs {
		var aggID uint64
		exiting, err := a.aggTable.GetByChainIdHeightHash(ctx, agg.SourceChainID, agg.BlockHeight, agg.BlockHash[:])
		if ormerrors.IsNotFound(err) {
			// Insert new aggregate
			aggID, err = a.aggTable.InsertReturningId(ctx, &pb.AggAttestation{
				ChainId:   agg.SourceChainID,
				Height:    agg.BlockHeight,
				Hash:      agg.BlockHash[:],
				BlockRoot: agg.BlockRoot[:],
				Status:    pb.AggStatus_Pending,
			})
			if err != nil {
				return errors.Wrap(err, "insert")
			}
		} else if err != nil {
			return errors.Wrap(err, "by block header")
		} else if !bytes.Equal(exiting.GetBlockRoot(), agg.BlockRoot[:]) {
			return errors.New("mismatching block root")
		} else {
			aggID = exiting.GetId()
		}

		// Insert signatures
		for _, sig := range agg.Signatures {
			err := a.sigTable.Insert(ctx, &pb.AttSignature{
				Signature:        sig.Signature[:],
				ValidatorAddress: sig.ValidatorAddress.Bytes(),
				AggId:            aggID,
			})
			if err != nil {
				return errors.Wrap(err, "insert signature")
			}
		}
	}

	return nil
}

// Approve approves any pending aggregate attestations that have quorum signatures form the provided set.
func (a Keeper) Approve(ctx context.Context, valSetID uint64, validators abci.ValidatorUpdates) error {
	approvedIdx := pb.AggAttestationStatusChainIdHeightIndexKey{}.WithStatus(pb.AggStatus_Approved)
	iter, err := a.aggTable.List(ctx, approvedIdx)
	if err != nil {
		return errors.Wrap(err, "list pending")
	}
	defer iter.Close()

	for iter.Next() {
		agg, err := iter.Value()
		if err != nil {
			return errors.Wrap(err, "value")
		}

		sigs, err := a.getAggSigs(ctx, agg.GetId())
		if err != nil {
			return errors.Wrap(err, "get agg validators")
		}

		toDelete, ok := isApproved(validators, sigs)
		if !ok {
			continue
		}

		for _, sig := range toDelete {
			err := a.sigTable.Delete(ctx, sig)
			if err != nil {
				return errors.Wrap(err, "delete sig")
			}
		}

		// Update status
		agg.Status = pb.AggStatus_Approved
		agg.ValidatorSetId = valSetID
		err = a.aggTable.Save(ctx, agg)
		if err != nil {
			return errors.Wrap(err, "save")
		}
	}

	return nil
}

// ApprovedFrom returns the subsequent approved attestations from the provided height (inclusive).
func (a Keeper) ApprovedFrom(ctx context.Context, chainID uint64, height uint64, max uint64,
) ([]xchain.AggAttestation, error) {
	from := pb.AggAttestationStatusChainIdHeightIndexKey{}.WithStatusChainIdHeight(
		pb.AggStatus_Approved, chainID, height)
	to := pb.AggAttestationStatusChainIdHeightIndexKey{}.WithStatusChainIdHeight(
		pb.AggStatus_Approved, chainID, height+max)

	iter, err := a.aggTable.ListRange(ctx, from, to)
	if err != nil {
		return nil, errors.Wrap(err, "list range")
	}
	defer iter.Close()

	var resp []xchain.AggAttestation
	next := height
	for iter.Next() {
		agg, err := iter.Value()
		if err != nil {
			return nil, errors.Wrap(err, "value")
		}

		if agg.GetHeight() != next {
			break
		}
		next++

		pbsigs, err := a.getAggSigs(ctx, agg.GetId())
		if err != nil {
			return nil, errors.Wrap(err, "get agg sigs")
		}

		var sigs []xchain.SigTuple
		for _, pbsig := range pbsigs {
			sigs = append(sigs, xchain.SigTuple{
				ValidatorAddress: common.Address(pbsig.GetValidatorAddress()),
				Signature:        [65]byte(pbsig.GetSignature()),
			})
		}

		resp = append(resp, xchain.AggAttestation{
			BlockHeader: xchain.BlockHeader{
				SourceChainID: agg.GetChainId(),
				BlockHeight:   agg.GetHeight(),
				BlockHash:     [32]byte(agg.GetHash()),
			},
			ValidatorSetID: agg.GetValidatorSetId(),
			BlockRoot:      [32]byte(agg.GetBlockRoot()),
			Signatures:     sigs,
		})
	}

	return resp, nil
}

// getAggSigs returns the signatures for the given aggregate ID.
func (a Keeper) getAggSigs(ctx context.Context, aggID uint64) ([]*pb.AttSignature, error) {
	aggIDIdx := pb.AttSignatureAggIdIndexKey{}.WithAggId(aggID)
	sigIter, err := a.sigTable.List(ctx, aggIDIdx)
	if err != nil {
		return nil, errors.Wrap(err, "list sig")
	}
	defer sigIter.Close()

	var sigs []*pb.AttSignature
	for sigIter.Next() {
		sig, err := sigIter.Value()
		if err != nil {
			return nil, errors.Wrap(err, "value sig")
		}

		sigs = append(sigs, sig)
	}

	return sigs, nil
}

// isApproved returns whether the given signatures are approved by the given validators.
// It also returns the signatures to delete (not in the validator set).
func isApproved(validators abci.ValidatorUpdates, sigs []*pb.AttSignature) ([]*pb.AttSignature, bool) {
	valSet := make(map[common.Address]int64)
	var total int64
	for _, val := range validators {
		addr, err := k1util.PubKeyPBToAddress(val.PubKey)
		if err != nil {
			return nil, false
		}

		total += val.Power
		valSet[addr] = val.Power
	}

	var sum int64
	var toDelete []*pb.AttSignature
	for _, sig := range sigs {
		power, ok := valSet[common.Address(sig.GetValidatorAddress())]
		if !ok {
			toDelete = append(toDelete, sig)
			continue
		}

		sum += power
	}

	return toDelete, sum > total*2/3
}
