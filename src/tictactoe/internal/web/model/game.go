package model

import (
	domain "tictactoe/internal/domain/model"

	"github.com/google/uuid"
)

type GameRequest struct {
	Cells [3][3]int `json:"cells"`
}

type GameResponse struct {
	ID       string    `json:"id"`
	Cells    [3][3]int `json:"cells"`
	Finished bool      `json:"finished"`
	Winner   int       `json:"winner,omitempty"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func (r GameRequest) ToDomain(id uuid.UUID) domain.Game {
	return domain.Game{
		ID:    id,
		Board: domain.Board(r.Cells),
	}
}

func GameResponseFromDomain(g domain.Game, finished bool, winner int) GameResponse {
	return GameResponse{
		ID:       g.ID.String(),
		Cells:    [3][3]int(g.Board),
		Finished: finished,
		Winner:   winner,
	}
}
