package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type StdExporter struct{}

// Encoder
/**
 * @Description:
 * @receiver e
 * @return zapcore.Encoder
 */
func (e *StdExporter) Encoder() zapcore.Encoder {
	return fmtEncoder(config.StdFormat, zap.NewDevelopmentEncoderConfig())
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
	return config.GetStdLevel().zapLevel()
}
