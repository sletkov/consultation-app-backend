package repository

import (
	"context"
	"database/sql"
	"encoding/json"
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

	query := "SELECT * FROM consultations WHERE draft = false"

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		byteStudents := make([]byte, 0)
		students := make([]*models.User, 0)
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
			&consultation.TeacherID,
			&consultation.Draft,
			&byteStudents,
		)

		err = json.Unmarshal(byteStudents, &students)
		if err != nil {
			return nil, err
		}
		fmt.Println(string(byteStudents))
		consultation.Students = students
		fmt.Println(consultation.Students)

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
	byteStudents := make([]byte, 0)
	students := make([]*models.User, 0)
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
		&consultation.TeacherID,
		&consultation.Draft,
		&byteStudents,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(fmt.Sprintf("No consultation with id: %s", consultationID))
		}
	}
	json.Unmarshal(byteStudents, &students)
	consultation.Students = students

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
		byteStudents := make([]byte, 0)
		students := make([]*models.User, 0)
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
			&consultation.TeacherID,
			&consultation.Draft,
			&byteStudents,
		)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(byteStudents, &students)
		if err != nil {
			return nil, err
		}
		consultation.Students = students
		fmt.Println(consultation.Students)
		consultations = append(consultations, consultation)
	}

	return consultations, nil
}

func (r *ConsultationRepository) CreateConsultation(ctx context.Context, cons *models.Consultation) error {
	query := "INSERT INTO consultations (id, title, description, consultation_format, consultation_type, teacher_name, teacher_id, consultation_date, consultation_time, campus, classroom, link, students_limit, students_count, draft, students) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14, $15, $16)"
	fmt.Println(cons)

	byteStudents, err := json.Marshal(cons.Students)
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(
		ctx,
		query,
		cons.ID,
		cons.Title,
		cons.Description,
		cons.Format,
		cons.Type,
		cons.TeacherName,
		cons.TeacherID,
		cons.Date,
		cons.Time,
		cons.Campus,
		cons.Classroom,
		cons.Link,
		cons.Limit,
		cons.StudentsCount,
		cons.Draft,
		byteStudents,
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

func (r *ConsultationRepository) SignupConsultation(ctx context.Context, student *models.User, consultation *models.Consultation) error {
	query1 := "INSERT INTO students_consultations (student_id, consultation_id) VALUES($1,$2)"
	query2 := "UPDATE consultations SET students_count = consultations.students_count + 1, students = $1 WHERE id = $2"

	students := consultation.Students
	students = append(students, student)

	jsonStudents, err := json.Marshal(students)
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, query1, student.ID, consultation.ID)
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, query2, jsonStudents, consultation.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *ConsultationRepository) DeleteConsultation(ctx context.Context, consultationID string) error {
	query1 := "DELETE FROM consultations WHERE id = $1"
	query2 := "DELETE FROM teachers_consultations WHERE consultation_id = $1"
	query3 := "DELETE FROM students_consultations WHERE consultation_id = $1"

	_, err := r.db.ExecContext(ctx, query1, consultationID)
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, query2, consultationID)
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, query3, consultationID)
	if err != nil {
		return err
	}

	return nil
}

func (r *ConsultationRepository) UpdateConsultation(ctx context.Context, consultation *models.Consultation, consultationID string) error {
	query := "UPDATE consultations SET title = $1, description = $2, consultation_format = $3, consultation_type = $4, consultation_date = $5, campus = $6, classroom = $7, students_limit = $8, link = $9, consultation_time = $10 WHERE id = $11"

	_, err := r.db.ExecContext(
		ctx,
		query,
		consultation.Title,
		consultation.Description,
		consultation.Format,
		consultation.Type,
		consultation.Date,
		consultation.Campus,
		consultation.Classroom,
		consultation.Limit,
		consultation.Link,
		consultation.Time,
		consultationID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *ConsultationRepository) GetStudentsByConsultationID(ctx context.Context, consultationID string) ([]*models.User, error) {
	fmt.Println("[GET_STUDENTS]")
	var students []*models.User
	studentsIDs := make([]string, 0)

	query1 := "SELECT * FROM students_consultations WHERE consultation_id = $1"

	rows, err := r.db.QueryContext(ctx, query1, consultationID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var studentID string
		var consultationID string
		err = rows.Scan(
			&studentID,
			&consultationID,
		)
		studentsIDs = append(studentsIDs, studentID)
	}
	fmt.Println("studentsIDs:", studentsIDs)

	query2 := "SELECT * FROM users WHERE id = ANY($1) AND role = 'student'"
	rows, err = r.db.QueryContext(ctx, query2, pq.Array(studentsIDs))
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		student := &models.User{}
		err = rows.Scan(
			&student.ID,
			&student.FullName,
			&student.Role,
			&student.Email,
			&student.Password,
		)
		student.RemoveSensitiveFields()
		students = append(students, student)
	}
	fmt.Println("students", students)

	return students, nil
}
