package di

import (
	"go.uber.org/fx"

	"tictactoe/internal/config"
	"tictactoe/internal/domain/service"
	"tictactoe/internal/repository"
	"tictactoe/internal/server"
	"tictactoe/internal/web/handler"
)

var Module = fx.Options(
	fx.Provide(config.Load),
	fx.Provide(repository.NewStorage),
	fx.Provide(repository.NewGameRepository),
	fx.Provide(service.NewGameService),
	fx.Provide(handler.NewGameHandler),
	fx.Provide(server.New),
)
