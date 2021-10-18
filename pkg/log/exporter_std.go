package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type StdExporter struct {
	Core *Std
}

func (e *StdExporter) Encoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
}

func (e *StdExporter) Writer() zapcore.WriteSyncer {
	return zapcore.Lock(os.Stdout)
}

func (e *StdExporter) Level() zapcore.Level {
	return e.Core.GetLevel().zapLevel()
}
