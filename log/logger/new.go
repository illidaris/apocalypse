package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(cfg *Config) {
	if cfg == nil {
		cfg = defaultConfig()
	}
	config = cfg
	exps := NewExporters()
	NewLogger(exps...)
}

func NewLogger(exporters ...IExporter) {
	coreTree := make([]zapcore.Core, 0)
	for _, exp := range exporters {
		coreTree = append(coreTree, zapcore.NewCore(exp.Encoder(), exp.Writer(), exp.Level()))
	}
	cores := zapcore.NewTee(coreTree...)
	// build log
	lg := zap.New(cores, zap.AddCaller(), zap.AddCallerSkip(0))
	zap.ReplaceGlobals(lg)

	ctxLogger = zap.L().WithOptions(zap.AddCallerSkip(2))
	funcLogger = zap.L().WithOptions(zap.AddCallerSkip(0))
}
