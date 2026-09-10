// internal/domain/model/game.go
package model

import "github.com/google/uuid"

type Game struct {
	ID            uuid.UUID `json:"id" db:"id"`
	Board         Board     `json:"board" db:"board"`
	CurrentPlayer int       `json:"currentPlayer"  db:"current_player"` // 1 = человек, 2 = компьютер
	Status        string    `json:"status"  db:"status"`                // "active", "won", "draw"
	Winner        int       `json:"winner,omitempty"  db:"winner"`
}

func (b Board) IsEmpty(r, c int) bool { return b[r][c] == 0 }

func NewGame() Game {
	return Game{ID: uuid.New()}
}
