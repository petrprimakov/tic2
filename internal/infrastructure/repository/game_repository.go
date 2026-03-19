package repository

import (
	"context"
	"fmt"

	"tic2/internal/domain/model"
	apperrors "tic2/internal/errors"

	"github.com/google/uuid"
)

type gameRepository struct {
	storage *gameStorage
}

func NewGameRepository(s *gameStorage) GameRepository {
	return &gameRepository{storage: s}
}

func (r *gameRepository) Save(_ context.Context, game model.Game) error {
	r.storage.store(game)
	return nil
}

func (r *gameRepository) Get(_ context.Context, id uuid.UUID) (model.Game, error) {
	game, ok := r.storage.load(id.String())
	if !ok {
		return model.Game{}, apperrors.NewNotFound(fmt.Sprintf("game %s not found", id))
	}
	return game, nil
}

func (r *gameRepository) Delete(_ context.Context, id uuid.UUID) error {
	r.storage.delete(id.String())
	return nil
}
