package logger

import (
	"context"
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestWithContext(t *testing.T) {
	New(nil)
	ctx := context.Background()
	ctx = NewContext(ctx, []zapcore.Field{
		{Key: "test1", Type: zapcore.StringType, String: "test123456"},
	}...)
	ctx = NewContext(ctx, []zapcore.Field{
		{Key: "test2", Type: zapcore.StringType, String: "test1234567"},
	}...)
	funcDebug(ctx)
	funcInfo(ctx)
	Info("xx")
}

func funcDebug(ctx context.Context) {
	DebugCtx(ctx, "debug")
}
func funcInfo(ctx context.Context) {
	InfoCtx(ctx, "info")
}
func funcWarn(ctx context.Context) {
	WarnCtx(ctx, "warn")
}
func funcError(ctx context.Context) {
	ErrorCtx(ctx, "error")
}
func funcFatal(ctx context.Context) {
	FatalCtx(ctx, "fatal")
}
