package logger

import "go.uber.org/zap/zapcore"

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
