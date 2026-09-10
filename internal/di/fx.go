package di

import (
	"tic2/internal/application/service"
	"tic2/internal/config"
	"tic2/internal/datasource/postgres"
	"tic2/internal/datasource/repository"
	"tic2/internal/middleware"
	"tic2/internal/server"
	"tic2/internal/web/handler"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		config.Load,
		postgres.NewPool,

		repository.NewGameRepository,
		repository.NewUserRepository,

		service.NewGameService,
		service.NewUserService,
		service.NewAuthService,

		handler.NewGameHandler,
		handler.NewAuthHandler,

		middleware.NewUserAuthenticator,

		server.New,
	),
)
