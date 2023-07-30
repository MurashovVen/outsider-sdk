package logger

import (
	"errors"

	"go.uber.org/zap"

	"github.com/MurashovVen/outsider-sdk/app/configuration"
)

type Logger struct {
	*zap.Logger
}

func MustCreateLogger(environment configuration.Environment) *Logger {
	l, err := CreateLogger(environment)
	if err != nil {
		panic("creating logger: " + err.Error())
	}

	return l
}

func CreateLogger(environment configuration.Environment) (*Logger, error) {
	switch environment {
	case configuration.DevEnv:
		l, err := zap.NewDevelopment()
		return &Logger{Logger: l}, err

	default:
		//nolint:goerr113
		return nil, errors.New("unknown environment")
	}
}

func NewNop() *Logger {
	return &Logger{
		Logger: zap.NewNop(),
	}
}

func (l *Logger) SyncWaiter() *SyncWaiter {
	return newSyncWaiter(l)
}

func (l *Logger) Named(name string) *Logger {
	return &Logger{
		Logger: l.Logger.Named(name),
	}
}
