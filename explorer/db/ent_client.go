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
)

type EntServerClient interface {
	CreateNewEntClient(databaseURL string) (*ent.Client, error)
	CreateSQLLiteEntClient() (*ent.Client, error)
}

type ClientImpl struct {
	EntServerClient
}

func NewClient() *ClientImpl {
	client := ClientImpl{}

	return &client
}

func (ClientImpl) CreateNewEntClient(ctx context.Context, databaseURL string) (*ent.Client, error) {
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		log.Error(ctx, "Failed to connect to postgres database: %v", err)
		return nil, errors.Wrap(err, "failed to open postgres connection")
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)

	client := ent.NewClient(ent.Driver(drv))

	return client, nil
}

func (ClientImpl) CreateNewEntClientWithSchema(ctx context.Context, databaseURL string) (*ent.Client, error) {
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		log.Error(ctx, "Failed to connect to postgres database: %v", err)
		return nil, errors.Wrap(err, "failed to open postgres connection")
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)

	client := ent.NewClient(ent.Driver(drv))

	if err := client.Schema.Create(ctx); err != nil {
		log.Error(ctx, "Failed creating schema resources: %v", err)
		return nil, errors.Wrap(err, "failed to do schema creation for postgresql db")
	}

	return client, nil
}

func (ClientImpl) CreateSQLiteEntClient(ctx context.Context) (*ent.Client, error) {
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Error(ctx, "Failed opening connection to sqlite: %v", err)
		return nil, errors.Wrap(err, "failed to open sqlite connection")
	}
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {
			log.Error(ctx, "Failed closing ent client: %v", err)
		}
	}(client)

	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		log.Error(ctx, "Failed creating schema resources: %v", err)
		return nil, errors.Wrap(err, "failed to do schema creation for sqlite db")
	}

	return client, nil
}
