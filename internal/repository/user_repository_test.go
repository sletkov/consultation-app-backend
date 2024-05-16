package repository

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sletkov/consultation-app-backend/internal/models"
	"log"
	"testing"
)

func TestUserRepository_SaveUser(t *testing.T) {
	db, m, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	type args struct {
		ctx  context.Context
		user *models.User
	}

	testCases := []struct {
		name         string
		args         args
		mockBehavior func()
		wantErr      bool
	}{
		{
			name: "OK",
			args: args{},
			mockBehavior: func() {
				query := "INSERT INTO users (id, fullname, role, email, password) VALUES($1,$2,$3,$4,$5)"
				m.ExpectExec(query).WithArgs().WillReturnResult(nil)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockBehavior()
			userRepo := NewUserRepository(db)
			if err := userRepo.SaveUser(tc.args.ctx, tc.args.user); err != nil {
				t.Error(err)
			}
		})

	}
}

func TestUserRepository_GetUserByEmail(t *testing.T) {

}

func TestUserRepository_GetUserByID(t *testing.T) {

}
