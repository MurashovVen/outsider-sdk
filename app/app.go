package app

import (
	"context"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/MurashovVen/outsider-sdk/app/logger"
	"github.com/MurashovVen/outsider-sdk/app/termination"
)

type (
	App struct {
		works []Work

		logger *logger.Logger
	}

	Work interface {
		Runner(context.Context) func() error
		Name() string
	}
)

func New(logger *logger.Logger, options ...Option) *App {
	app := &App{
		works: []Work{
			termination.NewWaiter(
				termination.WaiterWithLogger(logger),
			),
			logger.SyncWaiter(),
		},
		logger: logger,
	}

	for _, opt := range options {
		opt(app)
	}

	return app
}

func (a *App) Run(ctx context.Context) error {
	eg, egCtx := errgroup.WithContext(ctx)

	for _, work := range a.works {
		a.logger.Info("start work", zap.String("name", work.Name()))

		eg.Go(work.Runner(egCtx))
	}

	return eg.Wait()
}
