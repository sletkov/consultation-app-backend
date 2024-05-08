package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateTableTeachersConsultations, downCreateTableTeachersConsultations)
}

func upCreateTableTeachersConsultations(ctx context.Context, tx *sql.Tx) error {
	query := `
		CREATE TABLE IF NOT EXISTS teachers_consultations(
			teacher_id varchar not null,
			consultation_id varchar not null,
			PRIMARY KEY(teacher_id, consultation_id)
		);	
	`
	_, err := tx.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

func downCreateTableTeachersConsultations(ctx context.Context, tx *sql.Tx) error {
	query := "DROP TABLE IF EXISTS teachers_consultations"

	_, err := tx.Exec(query)

	if err != nil {
		return err
	}

	return nil
}
