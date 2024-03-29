package db

import (
	"context"
	"database/sql"

	"github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/lib/errors"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"

	_ "github.com/jackc/pgx/v5/stdlib"                       // importing postgres sql driver
	_ "github.com/mattn/go-sqlite3"                          // importing sqlite driver
	_ "github.com/omni-network/omni/explorer/db/ent/runtime" // import the runtime package to register hooks
)

func NewPostgressClient(dbURL string) (*ent.Client, error) {
	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		return nil, errors.Wrap(err, "opening postgres connection")
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)

	return ent.NewClient(ent.Driver(drv)), nil
}

func NewSQLiteClient() (*ent.Client, error) {
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		return nil, errors.Wrap(err, "failed to open sqlite connection")
	}

	return client, nil
}

// CreateSchema creates all schema resources and runs all required migrations.
func CreateSchema(ctx context.Context, client *ent.Client) error {
	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		return errors.Wrap(err, "schema creation for postgresql db")
	}

	return nil
}
