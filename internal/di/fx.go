package di

import (
	"go.uber.org/fx"

	"tic2/internal/application/service"
	"tic2/internal/config"
	"tic2/internal/datasource/repository"
	"tic2/internal/server"
	"tic2/internal/web/handler"
)

var Module = fx.Options(
	fx.Provide(config.Load),
	fx.Provide(repository.NewStorage),
	fx.Provide(repository.NewGameRepository),
	fx.Provide(service.NewGameService),
	fx.Provide(handler.NewGameHandler),
	fx.Provide(server.New),
)
