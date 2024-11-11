package cursors_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/xchain"
	"github.com/omni-network/omni/relayer/app/cursors"

	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"
	"cosmossdk.io/core/store"
	"cosmossdk.io/orm/model/ormdb"
	"cosmossdk.io/orm/types/ormerrors"
	db "github.com/cosmos/cosmos-db"
	"github.com/stretchr/testify/require"
)

// Test_StreamConfirmation tests how cursors are being confirmed.
func Test_StreamConfirmation(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	stream := xchain.StreamID{
		SourceChainID: 1,
		DestChainID:   2,
		ShardID:       0,
	}

	c := func(offset uint64, firstMsg uint64, lastMsg uint64, confirmed bool) *cursors.Cursor {
		return buildCursor(offset, firstMsg, lastMsg, confirmed, stream)
	}

	tests := []struct {
		desc       string
		cursors    []*cursors.Cursor    // db cursors
		final      *xchain.SubmitCursor // on-chain submitted cursor result
		confOffset uint64               // offset up to which cursors should be confirmed
	}{{
		desc:       "single non-empty attestation",
		cursors:    []*cursors.Cursor{c(1, 1, 2, false)},
		final:      &xchain.SubmitCursor{StreamID: stream, MsgOffset: 2, AttestOffset: 1},
		confOffset: 1,
	}, {
		desc: "non-empty attestation followed by empty",
		cursors: []*cursors.Cursor{
			c(1, 1, 2, false),
			c(2, 0, 0, false),
			c(3, 0, 0, false),
		},
		final:      &xchain.SubmitCursor{StreamID: stream, MsgOffset: 2, AttestOffset: 1},
		confOffset: 3,
	}, {
		desc: "non-empty attestation followed by empty, followed by non-empty but non-confirmed, followed by empty",
		cursors: []*cursors.Cursor{
			c(1, 1, 2, false),
			c(2, 0, 0, false),
			c(3, 0, 0, false),
			c(4, 3, 4, false),
			c(5, 0, 0, false),
			c(6, 0, 0, false),
		},
		final:      &xchain.SubmitCursor{StreamID: stream, MsgOffset: 2, AttestOffset: 1}, // only first non-empty confirmed
		confOffset: 3,
	}, {
		desc: "empty confirmed followed by empty non-confirmed",
		cursors: []*cursors.Cursor{
			c(1, 0, 0, true),
			c(2, 0, 0, false),
			c(3, 0, 0, false),
		},
		confOffset: 3,
	}, {
		desc: "empty confirmed followed by empty non-confirmed, followed by non-empty",
		cursors: []*cursors.Cursor{
			c(1, 0, 0, true),
			c(2, 0, 0, false),
			c(3, 0, 0, false),
			c(4, 3, 5, false),
		},
		final:      &xchain.SubmitCursor{StreamID: stream, MsgOffset: 4, AttestOffset: 4}, // not all messages are confirmed
		confOffset: 3,
	}, {
		desc: "all empty and non-confirmed",
		cursors: []*cursors.Cursor{
			c(1, 0, 0, false),
			c(2, 0, 0, false),
			c(3, 0, 0, false),
		},
		confOffset: 0,
	}}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			cursorTable, err := newMemCursorTable()
			require.NoError(t, err)
			provider := &mockProvider{}
			watcher, err := cursors.NewWatcher(cursorTable, provider)
			require.NoError(t, err)

			for _, c := range test.cursors {
				require.NoError(t, cursorTable.Insert(ctx, c))
			}

			var getSubmittedCalled int
			shouldCallGetSubmitted := test.final != nil

			if shouldCallGetSubmitted {
				provider.getSubmittedCursorFunc = func(
					ctx context.Context,
					ref xchain.Ref,
					s xchain.StreamID,
				) (xchain.SubmitCursor, bool, error) {
					require.Equal(t, xchain.FinalizedRef, ref)
					require.Equal(t, stream, s)
					getSubmittedCalled++

					return *test.final, true, nil
				}
			}

			err = watcher.ConfirmStream(ctx, stream)
			require.NoError(t, err)
			// make sure we only call it once to optimize resource use,
			// and only if we do have non-empty attestation
			require.Equal(t, shouldCallGetSubmitted, getSubmittedCalled == 1)

			for _, c := range test.cursors {
				confirmed, err := cursorTable.Get(ctx, c.GetSrcChainId(), c.GetConfLevel(), c.GetDstChainId(), c.GetAttestOffset())
				require.NoError(t, err)
				// make sure all cursors before confirmed index are indeed confirmed
				shouldBeConfirmed := c.GetAttestOffset() <= test.confOffset
				require.Equal(t, shouldBeConfirmed, confirmed.GetConfirmed(), fmt.Sprintf("offset %d", c.GetAttestOffset()))
			}
		})
	}
}

// Test_StreamTrimming tests how confirmed cursors are being removed from storage.
func Test_StreamTrimming(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	stream := xchain.StreamID{
		SourceChainID: 1,
		DestChainID:   2,
		ShardID:       0,
	}

	c := func(offset uint64, firstMsg uint64, lastMsg uint64, confirmed bool) *cursors.Cursor {
		return buildCursor(offset, firstMsg, lastMsg, confirmed, stream)
	}

	tests := []struct {
		desc         string
		cursors      []*cursors.Cursor // db cursors
		deleteOffset uint64            // offset up to which the cursors were deleted
	}{{
		desc:         "single non-empty attestation",
		cursors:      []*cursors.Cursor{c(1, 1, 2, true)},
		deleteOffset: 0,
	}, {
		desc: "multiple confirmed",
		cursors: []*cursors.Cursor{
			c(1, 1, 2, true),
			c(2, 0, 0, true),
			c(3, 0, 0, true),
		},
		deleteOffset: 2,
	}, {
		desc: "multiple confirmed, followed by non-confirmed",
		cursors: []*cursors.Cursor{
			c(1, 1, 2, true),
			c(2, 0, 0, true),
			c(3, 0, 0, true),
			c(4, 0, 0, false),
			c(5, 0, 0, false),
		},
		deleteOffset: 2,
	}, {
		desc: "multiple empty non-confirmed, followed by non-empty confirmed",
		cursors: []*cursors.Cursor{
			c(1, 0, 0, false),
			c(2, 0, 0, false),
			c(3, 1, 2, true),
			c(5, 0, 0, false),
		},
		deleteOffset: 2,
	}, {
		desc: "multiple empty non-confirmed",
		cursors: []*cursors.Cursor{
			c(1, 0, 0, false),
			c(2, 0, 0, false),
			c(3, 0, 0, false),
		},
		deleteOffset: 0,
	}}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			cursorTable, err := newMemCursorTable()
			require.NoError(t, err)
			watcher, err := cursors.NewWatcher(cursorTable, &mockProvider{})
			require.NoError(t, err)

			for _, c := range test.cursors {
				require.NoError(t, cursorTable.Insert(ctx, c))
			}

			err = watcher.TrimStream(ctx, stream)
			require.NoError(t, err)

			for _, c := range test.cursors {
				_, err := cursorTable.Get(ctx, c.GetSrcChainId(), c.GetConfLevel(), c.GetDstChainId(), c.GetAttestOffset())
				deleted := ormerrors.IsNotFound(err)
				shouldDelete := c.GetAttestOffset() <= test.deleteOffset
				require.Equal(t, shouldDelete, deleted, fmt.Sprintf("offset %d", c.GetAttestOffset()))
			}
		})
	}
}

type mockProvider struct {
	xchain.Provider
	getSubmittedCursorFunc func(ctx context.Context, ref xchain.Ref, stream xchain.StreamID) (xchain.SubmitCursor, bool, error)
}

func (m *mockProvider) GetSubmittedCursor(
	ctx context.Context,
	ref xchain.Ref,
	stream xchain.StreamID,
) (xchain.SubmitCursor, bool, error) {
	if m.getSubmittedCursorFunc == nil {
		panic("not defined getSubmittedCursorFunc")
	}

	return m.getSubmittedCursorFunc(ctx, ref, stream)
}

func buildCursor(
	offset uint64,
	firstMsg uint64,
	lastMsg uint64,
	confirmed bool,
	stream xchain.StreamID,
) *cursors.Cursor {
	return &cursors.Cursor{
		SrcChainId:      stream.SourceChainID,
		DstChainId:      stream.DestChainID,
		ConfLevel:       uint32(stream.ShardID),
		AttestOffset:    offset,
		Confirmed:       confirmed,
		FirstXmsgOffset: firstMsg,
		LastXmsgOffset:  lastMsg,
	}
}

func newMemCursorTable() (cursors.CursorTable, error) {
	schema := &ormv1alpha1.ModuleSchemaDescriptor{SchemaFile: []*ormv1alpha1.ModuleSchemaDescriptor_FileEntry{
		{Id: 1, ProtoFileName: cursors.File_relayer_app_cursors_cursors_proto.Path()},
	}}

	modDB, err := ormdb.NewModuleDB(schema, ormdb.ModuleDBOptions{KVStoreService: dbStoreService{DB: db.NewMemDB()}})
	if err != nil {
		return nil, errors.Wrap(err, "create ormdb module db")
	}

	dbStore, err := cursors.NewCursorsStore(modDB)
	if err != nil {
		return nil, errors.Wrap(err, "create store")
	}

	return dbStore.CursorTable(), nil
}

type dbStoreService struct {
	db.DB
}

func (db dbStoreService) OpenKVStore(context.Context) store.KVStore {
	return db.DB
}
