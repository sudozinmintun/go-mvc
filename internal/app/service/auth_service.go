package service

import (
	"errors"
	"fmt"
	"pmsys/internal/app/models"
	"pmsys/internal/app/repository"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(email, password string) (*models.User, error)
	Login(email, password string) (*models.User, error)
}

type authService struct {
	users repository.UserRepository
}

func NewAuthService(users repository.UserRepository) AuthService {
	return &authService{users: users}
}

func (s *authService) Register(email, password string) (*models.User, error) {
	_, err := s.users.FindByEmail(email)
	if err == nil {
		return nil, errors.New("email already registered")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u := &models.User{
		Email:    email,
		Password: string(hashed),
	}

	if err := s.users.Create(u); err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	return u, nil
}

func (s *authService) Login(email, password string) (*models.User, error) {
	u, err := s.users.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) != nil {
		return nil, errors.New("invalid credentials")
	}

	return u, nil
}
