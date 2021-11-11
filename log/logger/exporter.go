package logger

import (
	"github.com/illidaris/apocalypse/pkg/consts"
	"go.uber.org/zap/zapcore"
)

type IExporter interface {
	Encoder() zapcore.Encoder
	Writer() zapcore.WriteSyncer
	Level() zapcore.Level
}

func NewExporters() []IExporter {
	return []IExporter{
		&FileExporter{},
		&StdExporter{},
	}
}

// fmtEncoder choose format eg: json/console
func fmtEncoder(format string, cfg zapcore.EncoderConfig) zapcore.Encoder {
	switch format {
	case "console":
		return zapcore.NewConsoleEncoder(cfg)
	case "json":
		return zapcore.NewJSONEncoder(cfg)
	default:
		return zapcore.NewConsoleEncoder(cfg)
	}
}

// configEncoder config Encoder
func configEncoder() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        consts.Datetime.String(),
		LevelKey:       consts.LevelKey.String(),
		NameKey:        consts.NameKey.String(),
		CallerKey:      consts.Caller.String(),
		FunctionKey:    consts.FunctionKey.String(),
		MessageKey:     consts.Message.String(),
		StacktraceKey:  consts.StacktraceKey.String(),
		LineEnding:     consts.LineEnding.String(),
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}
