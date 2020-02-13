package musicworld

import (
	"go.uber.org/zap"
)

// Spin starts a service
func Spin() {
	// configure logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Info("executing root command")
}
