package logger

import "go.uber.org/zap"

var Lg *zap.Logger

func InitLogger() func() error {
	Lg, _ = zap.NewProduction()

	return Lg.Sync
}
