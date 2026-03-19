// internal/domain/service/game_service.go
package service

import "tic2/internal/domain/model"

type GameService interface {
	// MakeComputerMove делает ход компьютера с помощью Минимакса
	MakeComputerMove(game *model.Game) error

	// ValidateBoard проверяет, что поле корректно (не изменены предыдущие ходы)
	ValidateBoard(game *model.Game, incomingBoard model.Board) error

	// CheckGameOver проверяет, закончилась ли игра
	// Возвращает (finished, winner) — winner = 0 если ничья, 1 или 2 если победа
	CheckGameOver(game *model.Game) (bool, int)
}
