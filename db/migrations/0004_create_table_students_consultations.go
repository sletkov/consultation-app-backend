package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateTableStudentsConsultations, downCreateTableStudentsConsultations)
}

func upCreateTableStudentsConsultations(ctx context.Context, tx *sql.Tx) error {
	query := `
		CREATE TABLE IF NOT EXISTS students_consultations(
			student_id varchar not null,
			consultation_id varchar not null,
			PRIMARY KEY(student_id, consultation_id)
		);	
	`
	_, err := tx.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

func downCreateTableStudentsConsultations(ctx context.Context, tx *sql.Tx) error {
	query := "DROP TABLE IF EXISTS students_consultations"

	_, err := tx.Exec(query)

	if err != nil {
		return err
	}

	return nil
}
