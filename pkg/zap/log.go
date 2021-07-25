package zap

import (
	"go.uber.org/zap"
)

func NewLogger() (*zap.SugaredLogger, error)  {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	return sugar, nil
}