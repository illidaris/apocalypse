package logger

import (
	"context"
)

// Debug
/**
 * @Description:
 * @param msg
 */
func Debug(msg string) {
	funcLogger.Debug(msg)
}

// DebugCtx
/**
 * @Description:
 * @param ctx
 * @param msg
 */
func DebugCtx(ctx context.Context, msg string) {
	printCtx(ctx, DebugLevel, msg)
}

// Info
/**
 * @Description:
 * @param msg
 */
func Info(msg string) {
	funcLogger.Info(msg)
}

// InfoCtx
/**
 * @Description:
 * @param ctx
 * @param msg
 */
func InfoCtx(ctx context.Context, msg string) {
	printCtx(ctx, InfoLevel, msg)
}

// Warn
/**
 * @Description:
 * @param msg
 */
func Warn(msg string) {
	funcLogger.Warn(msg)
}

// WarnCtx
/**
 * @Description:
 * @param ctx
 * @param msg
 */
func WarnCtx(ctx context.Context, msg string) {
	printCtx(ctx, WarnLevel, msg)
}

// Error
/**
 * @Description:
 * @param msg
 */
func Error(msg string) {
	funcLogger.Error(msg)
}

// ErrorCtx
/**
 * @Description:
 * @param ctx
 * @param msg
 */
func ErrorCtx(ctx context.Context, msg string) {
	printCtx(ctx, ErrorLevel, msg)
}

// Fatal
/**
 * @Description:
 * @param msg
 */
func Fatal(msg string) {
	funcLogger.Fatal(msg)
}

// FatalCtx
/**
 * @Description:
 * @param ctx
 * @param msg
 */
func FatalCtx(ctx context.Context, msg string) {
	printCtx(ctx, FatalLevel, msg)
}
