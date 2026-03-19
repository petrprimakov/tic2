// internal/domain/model/board.go
package model

import "errors"

// Board представляет игровое поле 3x3
// 0 = пусто, 1 = игрок (X), 2 = компьютер (O)
type Board [3][3]int

// NewBoard создаёт пустое поле
func NewBoard() Board {
	return Board{}
}

// MakeMove делает ход (row, col, player)
// Возвращает ошибку, если ход некорректен
func (b *Board) MakeMove(row, col, player int) error {
	if row < 0 || row > 2 || col < 0 || col > 2 {
		return errors.New("некорректные координаты")
	}
	if b[row][col] != 0 {
		return errors.New("клетка занята")
	}
	b[row][col] = player
	return nil
}

// IsFull проверяет, заполнено ли поле
func (b Board) IsFull() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

// GetWinner возвращает победителя (1 или 2) или 0, если победителя нет
func (b Board) GetWinner() int {
	// Проверка строк, столбцов, диагоналей
	// (реализуй сам или я помогу в следующем сообщении)
	return 0 // пока заглушка
}
