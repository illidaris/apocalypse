package logger

import (
	"context"
	"fmt"
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestDebug(t *testing.T) {
	methodProxy(func() {
		Debug("debug")
	})
}

func BenchmarkDebug(b *testing.B) {
	num := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out := fmt.Sprintf("%d", num)
		methodProxy(func() {
			Debug(out)
		})
	}
}

func ExampleDebug() {
	Debug("example")
}

func TestDebugCtx(t *testing.T) {
	methodProxyWithCtx(func(ctx context.Context) {
		DebugCtx(ctx, "debug with ctx")
		WithContext(ctx).Info("info log")
	})
}

func BenchmarkDebugCtx(b *testing.B) {
	num := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out := fmt.Sprintf("%d", num)
		methodProxyWithCtx(func(ctx context.Context) {
			DebugCtx(ctx, out)
		})
	}
}

func ExampleDebugCtx() {
	DebugCtx(context.TODO(), "example")
}

func methodProxyWithCtx(f func(ctx context.Context)) {
	New(nil)
	ctx := context.Background()
	ctx = NewContext(ctx, []zapcore.Field{
		{Key: "test1", Type: zapcore.StringType, String: "test123456"},
	}...)
	ctx = NewContext(ctx, []zapcore.Field{
		{Key: "test2", Type: zapcore.StringType, String: "test1234567"},
	}...)
	f(ctx)
}

func methodProxy(f func()) {
	New(nil)
	f()
}
