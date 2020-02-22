package musicworld

import (
	"go.uber.org/zap"
)

// InitializeLogger creates a logger block
func InitializeLogger(verbose bool) (*zap.SugaredLogger, error) {
	var logger *zap.SugaredLogger
	switch verbose {
	case true:
		tempLogger, err := zap.NewDevelopment()
		if err != nil {
			return nil, err
		}
		logger = tempLogger.Sugar()
		defer tempLogger.Sync()
	case false:
		tempLogger := zap.NewNop()
		logger = tempLogger.Sugar()
		defer logger.Sync()
	}
	return logger, nil
}
