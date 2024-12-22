package app

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"zavad/config"
	"zavad/internal/delivery/http/middleware"
	"zavad/internal/delivery/http/server"
	"zavad/internal/repository"
	"zavad/internal/usecase"
)

func New() *fx.App {
	return fx.New(
		fx.Options(
			repository.New(),
			usecase.New(),
			middleware.New(),
			server.New(),
		),
		fx.Provide(
			context.Background,
			config.NewConfig,
			zap.NewDevelopment,
		),
		fx.WithLogger(
			func(log *zap.Logger) fxevent.Logger {
				return &fxevent.ZapLogger{Logger: log}
			},
		),
	)
}
