package db

import (
	"context"

	_ "embed"
)

//go:embed schema.sql
var schemaSQL []byte

func CreateSchema(ctx context.Context, db DBTX) error {
	// Create the schema if it does not exist
	_, err := db.Exec(ctx, string(schemaSQL))
	if err != nil {
		return err
	}

	return nil
}
