// Package logz provides utils for the zap logger.
package logz

import (
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewDev creates a new development logger.
func NewDev(wr io.Writer) *zap.SugaredLogger {
	var encoderCfg = zap.NewDevelopmentEncoderConfig()
	var consoleLg = zapcore.NewConsoleEncoder(encoderCfg)
	var core = zapcore.NewCore(consoleLg, zapcore.AddSync(wr), zap.DebugLevel)
	return zap.New(core).Sugar()
}

// NewProd creates a new production logger.
func NewProd(wr io.Writer) *zap.SugaredLogger {
	var encoderCfg = zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	var console = zapcore.NewConsoleEncoder(encoderCfg)
	var core = zapcore.NewCore(console, zapcore.AddSync(wr), zap.InfoLevel)

	return zap.New(core).Sugar()
}
