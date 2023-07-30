package logger

import (
	"context"
	"errors"
	"log"
	"syscall"
)

type (
	SyncWaiter struct {
		logger *Logger
	}
)

func newSyncWaiter(logger *Logger) *SyncWaiter {
	return &SyncWaiter{logger: logger}
}

func (sw *SyncWaiter) Runner(ctx context.Context) func() error {
	return sw.logger.syncWait(ctx)
}

func (l *Logger) syncWait(ctx context.Context) func() error {
	return func() error {
		<-ctx.Done()

		if err := l.Sync(); err != nil && !errors.Is(err, syscall.ENOTTY) {
			log.Printf("sync logger error: %v", err)

			return err
		}

		return nil
	}
}

func (sw *SyncWaiter) Name() string {
	return "LoggerSyncWaiter"
}
