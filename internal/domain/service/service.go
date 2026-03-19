// internal/domain/service/game_service.go

package service

import (
	"context"

	"tic2/internal/domain/model"

	"github.com/google/uuid"
)

type GameService interface {
	CreateGame(ctx context.Context) (model.Game, error)
	DeleteGame(ctx context.Context, id uuid.UUID) error
	ComputeNextMove(ctx context.Context, game model.Game) (model.Game, error)
	ValidateBoard(ctx context.Context, persisted, incoming model.Board) error
	CheckGameOver(ctx context.Context, board model.Board) (bool, int)
}
