package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"tic2/internal/domain/model"
	apperrors "tic2/internal/errors"
)

type AuthService struct {
	users *UserService
}

func NewAuthService(users *UserService) *AuthService {
	return &AuthService{users: users}
}

// SignUp регистрирует нового пользователя.
func (s *AuthService) SignUp(ctx context.Context, req model.SignUpRequest) error {
	if err := validateCredentials(req.Login, req.Password); err != nil {
		return err
	}
	_, err := s.users.CreateUser(ctx, req)
	return err
}

// SignIn проверяет логин/пароль и возвращает UUID пользователя.
func (s *AuthService) SignIn(ctx context.Context, login, password string) (uuid.UUID, error) {
	u, err := s.users.FindByLogin(ctx, login)
	if err != nil {
		return uuid.Nil, apperrors.ErrUnauthorized
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)); err != nil {
		return uuid.Nil, apperrors.ErrUnauthorized
	}
	return u.ID, nil
}

// validateCredentials — простая проверка длины/пустоты.
func validateCredentials(login, password string) error {
	if len(login) < 2 {
		return fmt.Errorf("login too short")
	}
	if len(password) < 2 {
		return fmt.Errorf("password too short")
	}
	return nil
}
