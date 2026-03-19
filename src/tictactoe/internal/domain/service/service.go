package service

import (
	"context"

	"tictactoe/internal/domain/model"

	"github.com/google/uuid"
)

type GameService interface {
	CreateGame(ctx context.Context) (model.Game, error)
	ComputeNextMove(ctx context.Context, game model.Game) (model.Game, error)
	ValidateBoard(ctx context.Context, persisted, incoming model.Board) error
	CheckGameOver(ctx context.Context, board model.Board) (finished bool, winner int)
	DeleteGame(ctx context.Context, id uuid.UUID) error
}
