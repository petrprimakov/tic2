// internal/domain/model/game.go
package model

import "github.com/google/uuid"

type Game struct {
	ID            uuid.UUID `json:"id"`
	Board         Board     `json:"board"`
	CurrentPlayer int       `json:"currentPlayer"` // 1 = человек, 2 = компьютер
	Status        string    `json:"status"`        // "active", "won", "draw"
	Winner        int       `json:"winner,omitempty"`
}

func (b Board) IsEmpty(r, c int) bool { return b[r][c] == 0 }

func NewGame() Game {
	return Game{ID: uuid.New()}
}
