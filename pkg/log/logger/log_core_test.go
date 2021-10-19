package logger

import (
	"context"
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestInfo(t *testing.T) {
	New(nil)
	Debug("debug msg")
	Info("info msg")
	Warn("warn msg")
	Error("error msg")
	Fatal("fatal msg")
}

func TestInfoCtx(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()
	New(nil)
	ctx := context.Background()
	ctx = NewContext(ctx, []zapcore.Field{
		{Key: "test1", Type: zapcore.StringType, String: "test123456"},
	}...)
	funcDebug(ctx)
	funcInfo(ctx)
	funcWarn(ctx)
	funcError(ctx)
	funcFatal(ctx)
	t.Log("finish")
}
