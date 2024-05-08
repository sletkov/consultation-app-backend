package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"github.com/sletkov/consultation-app-backend/internal/models"
)

type ConsultationRepository struct {
	db *sql.DB
}

func NewConsultationRepository(db *sql.DB) *ConsultationRepository {
	return &ConsultationRepository{
		db,
	}
}

func (r *ConsultationRepository) GetAll(ctx context.Context) ([]*models.Consultation, error) {
	var consultations []*models.Consultation

	query := "SELECT * FROM consultations"

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		consultation := &models.Consultation{}
		err := rows.Scan(
			&consultation.ID,
			&consultation.Title,
			&consultation.Description,
			&consultation.Format,
			&consultation.Type,
			&consultation.TeacherName,
			&consultation.Date,
			&consultation.Campus,
			&consultation.Classroom,
			&consultation.Limit,
			&consultation.StudentsCount,
			&consultation.Link,
			&consultation.Time,
		)

		if err != nil {
			return nil, err
		}
		consultations = append(consultations, consultation)
	}

	return consultations, nil
}

func (r *ConsultationRepository) GetConsultationByID(ctx context.Context, consultationID string) (*models.Consultation, error) {
	consultation := &models.Consultation{}

	query := "SELECT * FROM consultations WHERE id = $1"

	if err := r.db.QueryRowContext(ctx,
		query,
		consultationID,
	).Scan(
		&consultation.ID,
		&consultation.Title,
		&consultation.Description,
		&consultation.Format,
		&consultation.Type,
		&consultation.TeacherName,
		&consultation.Date,
		&consultation.Campus,
		&consultation.Classroom,
		&consultation.Limit,
		&consultation.StudentsCount,
		&consultation.Link,
		&consultation.Time,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(fmt.Sprintf("No consultation with id: %s", consultationID))
		}
	}
	return consultation, nil
}

func (r *ConsultationRepository) GetStudentConsultationsIDs(ctx context.Context, studentID string) ([]string, error) {
	consultationsIDs := make([]string, 0)

	query := "SELECT * FROM students_consultations WHERE student_id = $1"

	rows, err := r.db.QueryContext(ctx, query, studentID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var studentID string
		var consultationID string
		err := rows.Scan(
			&studentID,
			&consultationID,
		)
		if err != nil {
			return nil, err
		}

		consultationsIDs = append(consultationsIDs, consultationID)
	}
	return consultationsIDs, nil
}

func (r *ConsultationRepository) GetTeacherConsultationsIDs(ctx context.Context, teacherID string) ([]string, error) {
	consultationsIDs := make([]string, 0)

	query := "SELECT * FROM teachers_consultations WHERE teacher_id = $1"

	rows, err := r.db.QueryContext(ctx, query, teacherID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var teacherID string
		var consultationID string
		err := rows.Scan(
			&teacherID,
			&consultationID,
		)
		if err != nil {
			return nil, err
		}

		consultationsIDs = append(consultationsIDs, consultationID)
	}
	return consultationsIDs, nil
}

func (r *ConsultationRepository) GetUserConsultations(ctx context.Context, consultationsIDs []string) ([]*models.Consultation, error) {
	consultations := make([]*models.Consultation, 0)

	query := "SELECT * FROM consultations WHERE id = ANY($1)"

	rows, err := r.db.QueryContext(ctx, query, pq.Array(consultationsIDs))
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		consultation := &models.Consultation{}
		fmt.Printf("cons before scan: %v\n", consultation)
		err := rows.Scan(
			&consultation.ID,
			&consultation.Title,
			&consultation.Description,
			&consultation.Format,
			&consultation.Type,
			&consultation.TeacherName,
			&consultation.Date,
			&consultation.Campus,
			&consultation.Classroom,
			&consultation.Limit,
			&consultation.StudentsCount,
			&consultation.Link,
			&consultation.Time,
		)

		if err != nil {
			return nil, err
		}
		consultations = append(consultations, consultation)
	}

	return consultations, nil
}

func (r *ConsultationRepository) CreateConsultation(ctx context.Context, cons *models.Consultation) error {
	query := "INSERT INTO consultations (id, title, description, consultation_format, consultation_type, teacher_name, consultation_date, consultation_time, campus, classroom, link, students_limit, students_count) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)"
	fmt.Println(cons)
	_, err := r.db.ExecContext(
		ctx,
		query,
		cons.ID,
		cons.Title,
		cons.Description,
		cons.Format,
		cons.Type,
		cons.TeacherName,
		cons.Date,
		cons.Time,
		cons.Campus,
		cons.Classroom,
		cons.Link,
		cons.Limit,
		cons.StudentsCount,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *ConsultationRepository) AddTeacherConsultation(ctx context.Context, teacherID, consultationID string) error {
	query := "INSERT INTO teachers_consultations (teacher_id, consultation_id) VALUES($1,$2)"

	_, err := r.db.ExecContext(
		ctx,
		query,
		teacherID,
		consultationID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *ConsultationRepository) SignupConsultation(ctx context.Context, studentID, ConsultationID string) error {
	query1 := "INSERT INTO students_consultations (student_id, consultation_id) VALUES($1,$2)"

	_, err := r.db.ExecContext(ctx, query1, studentID, ConsultationID)
	if err != nil {
		return err
	}

	return nil
}
