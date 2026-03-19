package repository

import (
	"context"
	"tictactoe/internal/domain/model"

	"github.com/google/uuid"
)

type GameRepository interface {
	Save(ctx context.Context, game model.Game) error
	Get(ctx context.Context, id uuid.UUID) (model.Game, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
