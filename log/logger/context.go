package logger

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewContext
/**
 * @Description:
 * @param ctx
 * @param fields
 * @return context.Context
 */
func NewContext(ctx context.Context, fields ...zapcore.Field) context.Context {
	return context.WithValue(ctx, "_CtxLogger", WithContext(ctx).With(fields...))
}

// WithContext
/**
 * @Description:
 * @param ctx
 * @return *zap.Logger
 */
func WithContext(ctx context.Context) *zap.Logger {
	if ctx == nil {
		return ctxLogger
	}
	if l, ok := ctx.Value("_CtxLogger").(*zap.Logger); ok {
		return l
	}
	return ctxLogger
}

// debugCtxFunc
/**
 * @Description:
 * @param ctx
 * @return func(msg string, fields ...zapcore.Field)
 */
func debugCtxFunc(ctx context.Context) func(msg string, fields ...zapcore.Field) {
	return WithContext(ctx).Debug
}

// infoCtxFunc
/**
 * @Description:
 * @param ctx
 * @return func(msg string, fields ...zapcore.Field)
 */
func infoCtxFunc(ctx context.Context) func(msg string, fields ...zapcore.Field) {
	return WithContext(ctx).Info
}

// warnCtxFunc
/**
 * @Description:
 * @param ctx
 * @return func(msg string, fields ...zapcore.Field)
 */
func warnCtxFunc(ctx context.Context) func(msg string, fields ...zapcore.Field) {
	return WithContext(ctx).Warn
}

// errorCtxFunc
/**
 * @Description:
 * @param ctx
 * @return func(msg string, fields ...zapcore.Field)
 */
func errorCtxFunc(ctx context.Context) func(msg string, fields ...zapcore.Field) {
	return WithContext(ctx).Error
}

// criticalCtxFunc
/**
 * @Description:
 * @param ctx
 * @return func(msg string, fields ...zapcore.Field)
 */
func criticalCtxFunc(ctx context.Context) func(msg string, fields ...zapcore.Field) {
	return WithContext(ctx).Fatal
}

// printCtx
/**
 * @Description:
 * @param ctx
 * @param level
 * @param msg
 */
func printCtx(ctx context.Context, level Level, msg string) {
	zapPrints[level](ctx)(msg)
}
