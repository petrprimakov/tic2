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
		INSERT INTO games (id, board, player_x, player_o, state, turn_player, winner, vs_computer)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (id) DO UPDATE SET
			board       = EXCLUDED.board,
			player_x    = EXCLUDED.player_x,
			player_o    = EXCLUDED.player_o,
			state       = EXCLUDED.state,
			turn_player = EXCLUDED.turn_player,
			winner      = EXCLUDED.winner,
			vs_computer = EXCLUDED.vs_computer
	`

	_, err = r.pool.Exec(ctx, q,
		game.ID,
		boardJSON,
		game.PlayerX,
		game.PlayerO,
		game.State,
		game.TurnPlayer,
		game.Winner,
		game.VsComputer,
	)
	if err != nil {
		return fmt.Errorf("exec save game: %w", err)
	}
	return nil
}

func (r *gameRepository) Get(ctx context.Context, id uuid.UUID) (model.Game, error) {
	const q = `
		SELECT id, board, player_x, player_o, state, turn_player, winner, vs_computer
		FROM games
		WHERE id = $1
	`

	g, err := scanGame(r.pool.QueryRow(ctx, q, id))
	if errors.Is(err, pgx.ErrNoRows) {
		return model.Game{}, apperrors.ErrNotFound
	}
	if err != nil {
		return model.Game{}, fmt.Errorf("get game: %w", err)
	}
	return g, nil
}

func (r *gameRepository) Delete(ctx context.Context, id uuid.UUID) error {
	const q = `DELETE FROM games WHERE id = $1`
	if _, err := r.pool.Exec(ctx, q, id); err != nil {
		return fmt.Errorf("delete game: %w", err)
	}
	return nil
}

func (r *gameRepository) ListWaiting(ctx context.Context) ([]model.Game, error) {
	const q = `
		SELECT id, board, player_x, player_o, state, turn_player, winner, vs_computer
		FROM games
		WHERE state = $1
		ORDER BY id
	`

	rows, err := r.pool.Query(ctx, q, model.StateWaitingForPlayers)
	if err != nil {
		return nil, fmt.Errorf("query waiting games: %w", err)
	}
	defer rows.Close()

	var games []model.Game
	for rows.Next() {
		g, err := scanGame(rows)
		if err != nil {
			return nil, fmt.Errorf("scan waiting game: %w", err)
		}
		games = append(games, g)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate waiting games: %w", err)
	}
	return games, nil
}

// scanGame читает одну строку games во всех методах одинаково.
// Принимает pgx.Row — подходит и для QueryRow, и для rows внутри итерации.
func scanGame(row pgx.Row) (model.Game, error) {
	var (
		g         model.Game
		boardJSON []byte
	)
	err := row.Scan(
		&g.ID,
		&boardJSON,
		&g.PlayerX,
		&g.PlayerO,
		&g.State,
		&g.TurnPlayer,
		&g.Winner,
		&g.VsComputer,
	)
	if err != nil {
		return model.Game{}, err
	}
	if err := json.Unmarshal(boardJSON, &g.Board); err != nil {
		return model.Game{}, fmt.Errorf("unmarshal board: %w", err)
	}
	return g, nil
}
