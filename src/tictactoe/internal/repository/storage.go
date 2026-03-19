package repository

import (
	"sync"
	"tictactoe/internal/domain/model"
)

type gameStorage struct {
	m sync.Map
}

func NewStorage() *gameStorage {
	return newGameStorage()
}

func newGameStorage() *gameStorage {
	return &gameStorage{}
}

func (s *gameStorage) store(g model.Game) {
	s.m.Store(g.ID.String(), g)
}

func (s *gameStorage) load(id string) (model.Game, bool) {
	v, ok := s.m.Load(id)
	if !ok {
		return model.Game{}, false
	}
	return v.(model.Game), true
}

func (s *gameStorage) delete(id string) {
	s.m.Delete(id)
}
