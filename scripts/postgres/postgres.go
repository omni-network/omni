package postgres

import (
	"context"
	"fmt"
	"os/exec"
	"testing"
	"time"

	"github.com/omni-network/omni/lib/tutil"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
)

func Start(ctx context.Context, t *testing.T) string {
	t.Helper()
	port := tutil.RandomAvailablePort(t)
	t.Logf("Starting postgres on port %d", port)

	// docker run -it --rm postgres psql -h some-postgres -U postgres
	cmd := exec.CommandContext(ctx, "docker", "run",
		"--rm",
		"-e", "POSTGRES_PASSWORD=password",
		"-e", "POSTGRES_USER=admin",
		"-e", "POSTGRES_DB=postgres",
		"-p", fmt.Sprintf("%d:5432", port),
		"postgres",
	)
	err := cmd.Start()
	require.NoError(t, err)

	t.Cleanup(func() {
		_ = cmd.Process.Kill()
		_ = cmd.Wait()
	})

	// Wait for the postgres server to start
	dbConn := fmt.Sprintf("postgres://admin:password@localhost:%d/postgres?sslmode=disable", port)

	for i := 0; ; i++ {
		if cmd.ProcessState != nil && cmd.ProcessState.Exited() {
			t.Fatalf("Postgres process exited unexpectedly: %v", cmd.ProcessState)
		}
		cl, err := pgx.Connect(ctx, dbConn)
		if err == nil {
			err = cl.QueryRow(ctx, "SELECT 1").Scan(new(int))
		}
		if err != nil {
			if i > 0 && i%10 == 0 {
				t.Logf("Waiting for postgres to start: %v", err)
			}
			time.Sleep(time.Second)

			continue
		}

		break
	}

	t.Logf("Postgres available at: %s", dbConn)

	return dbConn
}
