package db_test

import (
	"testing"

	"github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/explorer/db/ent/enttest"
	"github.com/omni-network/omni/explorer/db/ent/migrate"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "github.com/mattn/go-sqlite3"
)

func TestEntClient(t *testing.T) {
	t.Parallel()
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
		enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	assert.NotNil(t, client)

	t.Cleanup(func() {
		err := client.Close()
		require.NoError(t, err)
		if err != nil {
			t.Error("failed to close test client", err)
		}
	})
}
