package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateTableUsers, downCreateTableUsers)
}

func upCreateTableUsers(ctx context.Context, tx *sql.Tx) error {
	query := `
		CREATE TABLE IF NOT EXISTS users(
			id varchar primary key not null,
			fullname varchar not null,
			role varchar not null,
			email varchar not null,
			password varchar not null
		);	
	`
	_, err := tx.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

func downCreateTableUsers(ctx context.Context, tx *sql.Tx) error {
	query := "DROP TABLE IF EXISTS users"

	_, err := tx.Exec(query)

	if err != nil {
		return err
	}

	return nil
}
