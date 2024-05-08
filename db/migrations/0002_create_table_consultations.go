package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateTableConsultations, downCreateTableConsultations)
}

func upCreateTableConsultations(ctx context.Context, tx *sql.Tx) error {
	query := `
		CREATE TABLE IF NOT EXISTS consultations(
			id varchar primary key not null,
			title varchar not null,
			description varchar,
			consultation_format varchar not null,
			consultation_type varchar not null,
			teacher_name varchar not null,
			consultation_date varchar not null,
			campus varchar not null,
			classroom varchar not null,
			students_limit integer,
			students_count integer,
			link varchar,
			consultation_time varchar
		);	
	`
	_, err := tx.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

func downCreateTableConsultations(ctx context.Context, tx *sql.Tx) error {
	query := "DROP TABLE IF EXISTS consultations"

	_, err := tx.Exec(query)

	if err != nil {
		return err
	}

	return nil
}
