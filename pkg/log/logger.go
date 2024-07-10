package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/xerrors"
)

var (
	Logger *zap.SugaredLogger
)

func init() {
	Logger, _ = NewLogger(false)
}

func InitLogger(debug bool) error {
	var err error
	Logger, err = NewLogger(debug)
	if err != nil {
		return xerrors.Errorf("failed to initialize logger: %w", err)
	}

	return nil
}

func NewLogger(debug bool) (*zap.SugaredLogger, error) {
	config := zap.NewProductionConfig()
	if debug {
		config = zap.NewDevelopmentConfig()
	}

	config.Encoding = "console"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}
	return logger.Sugar(), nil
}
