package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sletkov/consultation-app-backend/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (r *UserRepository) SaveUser(ctx context.Context, u *models.User) error {
	query := "INSERT INTO users (id, fullname, role, email, password) VALUES($1,$2,$3,$4,$5)"

	//Если во время выполения запроса к базе данных возникла ошибка, возвращаем ошибку
	_, err := r.db.ExecContext(ctx, query, u.ID, u.FullName, u.Role, u.Email, u.EncryptedPassword)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user := &models.User{}

	query := "SELECT * FROM users WHERE email = $1"
	err := r.db.
		QueryRowContext(
			ctx,
			query,
			email).
		Scan(
			&user.ID,
			&user.FullName,
			&user.Role,
			&user.Email,
			&user.EncryptedPassword,
		)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	user := &models.User{}
	fmt.Println("id is", id)
	query := "SELECT * FROM users WHERE id = $1"
	err := r.db.QueryRowContext(
		ctx,
		query,
		id,
	).Scan(
		&user.ID,
		&user.FullName,
		&user.Role,
		&user.Email,
		&user.EncryptedPassword,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
