package app_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/omni-network/omni/lib/tutil"
	"github.com/omni-network/omni/lib/uni"
	"github.com/omni-network/omni/scripts/postgres"
	"github.com/omni-network/omni/scripts/trade/app"
	"github.com/omni-network/omni/scripts/trade/config"
	"github.com/omni-network/omni/scripts/trade/users"
	usersclient "github.com/omni-network/omni/scripts/trade/users/client"
	usersdb "github.com/omni-network/omni/scripts/trade/users/db"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
)

func TestSmoke(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	dbConn := postgres.Start(ctx, t)

	conn, err := pgx.Connect(ctx, dbConn)
	require.NoError(t, err)

	// Create schema
	require.NoError(t, usersdb.CreateSchema(ctx, conn))

	port := tutil.RandomAvailablePort(t)

	cfg := config.DefaultConfig()
	cfg.RPCListen = fmt.Sprintf(":%d", port)
	cfg.DBConn = dbConn

	errStart, stopFunc, err := app.Start(ctx, cfg)
	require.NoError(t, err)
	defer func() {
		select {
		case err := <-errStart:
			require.NoError(t, err, "app should not return an error")
		default:
		}
		require.NoError(t, stopFunc(ctx), "stop app")
	}()

	addr := fmt.Sprintf("http://localhost:%d", port)
	testUsers(ctx, t, addr)
}

func testUsers(ctx context.Context, t *testing.T, addr string) {
	t.Helper()
	cl := usersclient.New(addr)

	var all []users.User
	for i := range 3 {
		u0, err := cl.Create(ctx, users.RequestCreate{
			ID:      uuid.New(),
			PrivyID: fmt.Sprintf("privy-%d", i),
			Address: uni.EVMAddress(tutil.RandomAddress()),
		})
		tutil.RequireNoError(t, err)
		all = append(all, u0)

		u1, err := cl.GetByID(ctx, u0.ID)
		require.NoError(t, err)
		require.Equal(t, u0, u1)

		u2, err := cl.GetByPrivyID(ctx, u0.PrivyID)
		require.NoError(t, err)
		require.Equal(t, u0, u2)

		u3, err := cl.GetByAddress(ctx, u0.Address)
		require.NoError(t, err)
		require.Equal(t, u0, u3)
	}

	allUsers, err := cl.ListAll(ctx)
	require.NoError(t, err)
	require.Len(t, allUsers, len(all))
	for _, u := range all {
		require.Contains(t, allUsers, u)
	}
}
