package app

import (
	"context"
	"docs_flow/config"
	"docs_flow/internal/delivery/http/middleware"
	"docs_flow/internal/delivery/http/server"
	"docs_flow/internal/repository"
	"docs_flow/internal/usecase"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
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
