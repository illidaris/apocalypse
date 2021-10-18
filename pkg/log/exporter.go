package log

import "go.uber.org/zap/zapcore"

type IExporter interface {
	Encoder() zapcore.Encoder
	Writer() zapcore.WriteSyncer
	Level() zapcore.Level
}
