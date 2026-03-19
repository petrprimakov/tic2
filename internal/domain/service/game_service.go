package service

import (
	"context"
	"errors"

	"tictactoe/internal/domain/model"
	"tictactoe/internal/repository"

	"github.com/google/uuid"
)

type gameService struct {
	repo repository.GameRepository
}

func NewGameService(repo repository.GameRepository) GameService {
	return &gameService{repo: repo}
}

func (s *gameService) CreateGame(ctx context.Context) (model.Game, error) {
	game := model.NewGame()
	if err := s.repo.Save(ctx, game); err != nil {
		return model.Game{}, err
	}
	return game, nil
}

func (s *gameService) DeleteGame(ctx context.Context, id uuid.UUID) error {
	if _, err := s.repo.Get(ctx, id); err != nil {
		return err
	}
	return s.repo.Delete(ctx, id)
}

func (s *gameService) ComputeNextMove(ctx context.Context, game model.Game) (model.Game, error) {
	persisted, err := s.repo.Get(ctx, game.ID)
	if err != nil {
		return model.Game{}, err
	}

	if err := s.ValidateBoard(ctx, persisted.Board, game.Board); err != nil {
		return model.Game{}, err
	}

	if finished, _ := s.CheckGameOver(ctx, game.Board); finished {
		return game, nil
	}

	col, row := s.bestMove(game.Board)
	if col == -1 {
		return game, errors.New("no available moves")
	}
	game.Board[row][col] = 2

	if finished, _ := s.CheckGameOver(ctx, game.Board); finished {
		_ = s.repo.Delete(ctx, game.ID)
		return game, nil
	}

	if err := s.repo.Save(ctx, game); err != nil {
		return model.Game{}, err
	}
	return game, nil
}

func (s *gameService) ValidateBoard(_ context.Context, persisted, incoming model.Board) error {
	newMoves := 0
	for r := range persisted {
		for c := range persisted[r] {
			if persisted[r][c] != 0 && persisted[r][c] != incoming[r][c] {
				return errors.New("invalid board: previously made move was altered")
			}
			if persisted[r][c] == 0 && incoming[r][c] == 1 {
				newMoves++
			}
			if persisted[r][c] == 0 && incoming[r][c] == 2 {
				return errors.New("invalid board: player cannot place computer's symbol")
			}
		}
	}
	if newMoves != 1 {
		return errors.New("invalid board: player must make exactly one move")
	}
	return nil
}

func (s *gameService) CheckGameOver(_ context.Context, board model.Board) (bool, int) {
	if w := s.winner(board); w != 0 {
		return true, w
	}
	for r := range board {
		for c := range board[r] {
			if board[r][c] == 0 {
				return false, 0
			}
		}
	}
	return true, 0
}

func (s *gameService) bestMove(board model.Board) (bestCol, bestRow int) {
	bestScore := -1 << 30
	bestCol, bestRow = -1, -1
	for r := range board {
		for c := range board[r] {
			if board[r][c] != 0 {
				continue
			}
			board[r][c] = 2
			score := s.minimax(board, false, -1<<30, 1<<30, 0)
			board[r][c] = 0
			if score > bestScore {
				bestScore = score
				bestCol, bestRow = c, r
			}
		}
	}
	return
}

func (s *gameService) minimax(board model.Board, isComputerTurn bool, alpha, beta, depth int) int {
	if w := s.winner(board); w != 0 {
		if w == 2 {
			return 10 - depth
		}
		return -10 + depth
	}
	empty := 0
	for r := range board {
		for c := range board[r] {
			if board[r][c] == 0 {
				empty++
			}
		}
	}
	if empty == 0 {
		return 0
	}

	if isComputerTurn {
		best := -1 << 30
		for r := range board {
			for c := range board[r] {
				if board[r][c] != 0 {
					continue
				}
				board[r][c] = 2
				best = max(best, s.minimax(board, false, alpha, beta, depth+1))
				board[r][c] = 0
				alpha = max(alpha, best)
				if beta <= alpha {
					return best
				}
			}
		}
		return best
	}

	best := 1 << 30
	for r := range board {
		for c := range board[r] {
			if board[r][c] != 0 {
				continue
			}
			board[r][c] = 1
			best = min(best, s.minimax(board, true, alpha, beta, depth+1))
			board[r][c] = 0
			beta = min(beta, best)
			if beta <= alpha {
				return best
			}
		}
	}
	return best
}

func (s *gameService) winner(b model.Board) int {
	win_lines := [8][3][2]int{
		{{0, 0}, {0, 1}, {0, 2}}, {{1, 0}, {1, 1}, {1, 2}}, {{2, 0}, {2, 1}, {2, 2}},
		{{0, 0}, {1, 0}, {2, 0}}, {{0, 1}, {1, 1}, {2, 1}}, {{0, 2}, {1, 2}, {2, 2}},
		{{0, 0}, {1, 1}, {2, 2}}, {{0, 2}, {1, 1}, {2, 0}},
	}
	for _, l := range win_lines {
		a1, a2, a3 := b[l[0][0]][l[0][1]], b[l[1][0]][l[1][1]], b[l[2][0]][l[2][1]]
		if a1 != 0 && a1 == a2 && a2 == a3 {
			return a1
		}
	}
	return 0
}
