package db

import (
	"context"
	"database/sql"

	"github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"

	_ "github.com/jackc/pgx/v5/stdlib" // importing postgres sql driver
	_ "github.com/mattn/go-sqlite3"    // importing sqlite3 sql driver
)

type EntServerClient interface {
	CreateNewEntClient(databaseURL string) (*ent.Client, error)
	CreateSQLLiteEntClient() (*ent.Client, error)
	CreateNewEntClientWithSchema(ctx context.Context, databaseURL string) (*ent.Client, error)
}

type Client struct {
	EntServerClient
}

func NewClient() *Client {
	client := Client{}

	return &client
}

// Creates an ORM client without creating the tables.
func (Client) CreateNewEntClient(databaseURL string) (*ent.Client, error) {
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open postgres connection")
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)

	client := ent.NewClient(ent.Driver(drv))

	return client, nil
}

// Creates an ORM client and creates the tables on initialization.
func (Client) CreateNewEntClientWithSchema(ctx context.Context, databaseURL string) (*ent.Client, error) {
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open postgres connection")
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)

	client := ent.NewClient(ent.Driver(drv))

	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		return nil, errors.Wrap(err, "failed to do schema creation for postgresql db")
	}

	return client, nil
}

// Creates an in-memory ORM client with tables.
func (Client) CreateSQLiteEntClient(ctx context.Context) (*ent.Client, error) {
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		return nil, errors.Wrap(err, "failed to open sqlite connection")
	}
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {
			log.Error(ctx, "Closing ent client", err)
		}
	}(client)

	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		return nil, errors.Wrap(err, "failed to do schema creation for sqlite db")
	}

	return client, nil
}
