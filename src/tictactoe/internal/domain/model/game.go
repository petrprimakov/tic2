package model

import "github.com/google/uuid"

type Board [3][3]int

func (b Board) IsEmpty(r, c int) bool { return b[r][c] == 0 }

type Game struct {
	ID    uuid.UUID
	Board Board
}

func NewGame() Game {
	return Game{ID: uuid.New()}
}
