package log

import "go.uber.org/zap"

var _logger *zap.Logger

func InitLogger() error {
	zap.New()
}
