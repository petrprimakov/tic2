package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"tic2/internal/domain/model"
	apperrors "tic2/internal/errors"
)

type UserRepository interface {
	Create(ctx context.Context, user model.User) error
	GetByLogin(ctx context.Context, login string) (model.User, error)
}

type userRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) UserRepository {
	return &userRepository{pool: pool}
}

func (r *userRepository) Create(ctx context.Context, u model.User) error {
	const q = `INSERT INTO users (id, login, password_hash) VALUES ($1, $2, $3)`
	_, err := r.pool.Exec(ctx, q, u.ID, u.Login, u.PasswordHash)
	if err != nil {
		// нарушение UNIQUE по login
		return fmt.Errorf("create user: %w", err)
	}
	return nil
}

func (r *userRepository) GetByLogin(ctx context.Context, login string) (model.User, error) {
	const q = `SELECT id, login, password_hash FROM users WHERE login = $1`

	var u model.User
	err := r.pool.QueryRow(ctx, q, login).Scan(&u.ID, &u.Login, &u.PasswordHash)
	if errors.Is(err, pgx.ErrNoRows) {
		return model.User{}, apperrors.ErrNotFound
	}
	if err != nil {
		return model.User{}, fmt.Errorf("get user by login: %w", err)
	}
	return u, nil
}
