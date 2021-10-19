package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type StdExporter struct {
	Core *Std
}

// Encoder
/**
 * @Description:
 * @receiver e
 * @return zapcore.Encoder
 */
func (e *StdExporter) Encoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
}

// Writer
/**
 * @Description:
 * @receiver e
 * @return zapcore.WriteSyncer
 */
func (e *StdExporter) Writer() zapcore.WriteSyncer {
	return zapcore.Lock(os.Stdout)
}

// Level
/**
 * @Description:
 * @receiver e
 * @return zapcore.Level
 */
func (e *StdExporter) Level() zapcore.Level {
	return e.Core.GetLevel().zapLevel()
}
