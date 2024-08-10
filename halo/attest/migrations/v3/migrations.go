package v3

import (
	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"
	"cosmossdk.io/orm/model/ormdb"
	sdk "github.com/cosmos/cosmos-sdk/types"
	attestkeeper "github.com/omni-network/omni/halo/attest/keeper"
	v2 "github.com/omni-network/omni/halo/attest/migrations/v2"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
)

func MigrateStore(ctx sdk.Context, attKeeper *attestkeeper.Keeper) error {
	log.Info(ctx, "💟💟💟 Migrating attestation store from v2 to v3")
	v2Schema := &ormv1alpha1.ModuleSchemaDescriptor{SchemaFile: []*ormv1alpha1.ModuleSchemaDescriptor_FileEntry{
		{Id: 1, ProtoFileName: v2.File_halo_attest_migrations_v2_attestation_proto.Path()},
	}}

	v2ModDB, err := ormdb.NewModuleDB(v2Schema, ormdb.ModuleDBOptions{KVStoreService: attKeeper.StoreService()})
	if err != nil {
		return errors.Wrap(err, "create module db")
	}

	v2AttStore, err := v2.NewAttestationStore(v2ModDB)
	if err != nil {
		return errors.Wrap(err, "create attestation store")
	}
	v2AttTable := v2AttStore.AttestationTable()
	v3AttTable := attKeeper.AttTable()

	v2Iter, err := v2AttTable.List(ctx, v2.AttestationPrimaryKey{})
	if err != nil {
		return errors.Wrap(err, "list v2 attestation")
	}
	defer v2Iter.Close()

	var count int
	for v2Iter.Next() {
		v2Val, err := v2Iter.Value()
		if err != nil {
			return errors.Wrap(err, "get v2 attestation value")
		}

		if err := v2AttTable.Delete(ctx, v2Val); err != nil {
			return errors.Wrap(err, "delete v2 attestation")
		}

		v3Val := attestkeeper.Attestation{
			ChainId:         v2Val.ChainId,
			ConfLevel:       v2Val.ConfLevel,
			AttestOffset:    v2Val.AttestOffset,
			BlockHeight:     v2Val.BlockHeight,
			BlockHash:       v2Val.BlockHash,
			MsgRoot:         v2Val.MsgRoot,
			AttestationRoot: v2Val.AttestationRoot,
			Status:          v2Val.Status,
			ValidatorSetId:  v2Val.ValidatorSetId,
			CreatedHeight:   v2Val.CreatedHeight,
			FinalizedAttId:  v2Val.FinalizedAttId,
		}
		v3ValID, err := v3AttTable.InsertReturningId(ctx, &v3Val)
		if err != nil {
			return errors.Wrap(err, "save v3 attestation")
		}

		if err := attKeeper.UpgradeSigAttIDs(ctx, v2Val.Id, v3ValID); err != nil {
			return errors.Wrap(err, "upgrade sig att ids")
		}

		count++
	}

	log.Info(ctx, "💟💟💟 Migrated attestations", "count", count)

	return nil
}
