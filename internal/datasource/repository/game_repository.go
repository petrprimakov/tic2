package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"tic2/internal/domain/model"
	apperrors "tic2/internal/errors"
)

type gameRepository struct {
	pool *pgxpool.Pool
}

func NewGameRepository(pool *pgxpool.Pool) GameRepository {
	return &gameRepository{pool: pool}
}

func (r *gameRepository) Save(ctx context.Context, game model.Game) error {
	boardJSON, err := json.Marshal(game.Board)
	if err != nil {
		return fmt.Errorf("marshal board: %w", err)
	}

	const q = `
        INSERT INTO games (id, board, current_player, status, winner)
        VALUES ($1, $2, $3, $4, $5)
        ON CONFLICT (id) DO UPDATE SET
            board          = EXCLUDED.board,
            current_player = EXCLUDED.current_player,
            status         = EXCLUDED.status,
            winner         = EXCLUDED.winner
    `
	_, err = r.pool.Exec(ctx, q,
		game.ID, boardJSON, game.CurrentPlayer, game.Status, game.Winner)
	if err != nil {
		return fmt.Errorf("exec save: %w", err)
	}
	return nil
}

func (r *gameRepository) Get(ctx context.Context, id uuid.UUID) (model.Game, error) {
	const q = `SELECT board, current_player, status, winner FROM games WHERE id = $1`

	var (
		boardJSON []byte
		g         model.Game
	)
	g.ID = id

	err := r.pool.QueryRow(ctx, q, id).Scan(
		&boardJSON, &g.CurrentPlayer, &g.Status, &g.Winner,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return model.Game{}, apperrors.ErrNotFound
	}
	if err != nil {
		return model.Game{}, fmt.Errorf("scan game: %w", err)
	}
	if err := json.Unmarshal(boardJSON, &g.Board); err != nil {
		return model.Game{}, fmt.Errorf("unmarshal board: %w", err)
	}
	return g, nil
}

func (r *gameRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.pool.Exec(ctx, `DELETE FROM games WHERE id = $1`, id)
	return err
}
