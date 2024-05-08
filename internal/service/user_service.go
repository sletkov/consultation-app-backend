package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/sletkov/consultation-app-backend/internal/models"
)

type UserRepository interface {
	SaveUser(ctx context.Context, u *models.User) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByID(ctx context.Context, id string) (*models.User, error)
}

type UserService struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) *UserService {
	return &UserService{
		userRepository,
	}
}

func (service *UserService) SaveUser(ctx context.Context, user *models.User) (string, error) {
	//check if user already exists
	u, err := service.userRepository.GetUserByEmail(ctx, user.Email)
	if u != nil {
		return "", errors.New(fmt.Sprintf("user with email %s already exists", u.Email))
	}

	id := uuid.New().String()
	user.ID = id

	user.EncryptPassword()
	user.Sanitize()

	err = service.userRepository.SaveUser(ctx, user)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (service *UserService) Login(ctx context.Context, email string, password string) (*models.User, error) {
	user, err := service.userRepository.GetUserByEmail(ctx, email)
	fmt.Println("user by email:", user)
	if err != nil || !user.ComparePassword(password) {
		return nil, err
	}
	return user, nil
}

func (service *UserService) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	user, err := service.userRepository.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
