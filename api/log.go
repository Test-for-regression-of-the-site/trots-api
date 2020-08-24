package api

import (
	"log"

	"go.uber.org/zap"
)

// Logger is the generic logger interface.
type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

func stdErrorLog(lg Logger) *log.Logger {
	if sugar, isZap := lg.(*zap.SugaredLogger); isZap {
		var z, _ = zap.NewStdLogAt(sugar.Desugar(), zap.ErrorLevel)
		return z
	}
	var ch = &chopper{printf: lg.Errorf}
	return log.New(ch, "", 0)
}

type chopper struct {
	printf func(format string, args ...interface{})
}

func (ch *chopper) Write(data []byte) (int, error) {
	ch.printf("%s", data)
	return len(data), nil
}
