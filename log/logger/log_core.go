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
	printCtx(context.TODO(), DebugLevel, msg)
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
	printCtx(context.TODO(), InfoLevel, msg)
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
	printCtx(context.TODO(), WarnLevel, msg)
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
	printCtx(context.TODO(), ErrorLevel, msg)
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
	printCtx(context.TODO(), FatalLevel, msg)
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
