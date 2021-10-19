package logger

import (
	"context"
	"fmt"
	"strings"

	"go.uber.org/zap/zapcore"
)

// Level is a logging priority. Higher levels are more important.
type Level int8

// Logging levels.
const (
	DebugLevel Level = iota - 1
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel // Critical exists only for config backward compatibility.
)

var levelStrings = map[Level]string{
	DebugLevel: "debug",
	InfoLevel:  "info",
	WarnLevel:  "warning",
	ErrorLevel: "error",
	FatalLevel: "fatal",
}

var zapLevels = map[Level]zapcore.Level{
	DebugLevel: zapcore.DebugLevel,
	InfoLevel:  zapcore.InfoLevel,
	WarnLevel:  zapcore.WarnLevel,
	ErrorLevel: zapcore.ErrorLevel,
	FatalLevel: zapcore.ErrorLevel,
}

var zapPrints = map[Level]func(ctx context.Context) func(msg string, fields ...zapcore.Field){
	DebugLevel: debugCtxFunc,
	InfoLevel:  infoCtxFunc,
	WarnLevel:  warnCtxFunc,
	ErrorLevel: errorCtxFunc,
	FatalLevel: criticalCtxFunc,
}

// String returns the name of the logging level.
func (l Level) String() string {
	s, found := levelStrings[l]
	if found {
		return s
	}
	return fmt.Sprintf("Level(%d)", l)
}

// Unabled returns true if given level is enabled.
func (l Level) Unabled(level Level) bool {
	return level < l
}

// Unpack unmarshals a level string to a Level. This implements
// ucfg.StringUnpacker.
func (l *Level) Unpack(str string) error {
	str = strings.ToLower(str)
	for level, name := range levelStrings {
		if name == str {
			*l = level
			return nil
		}
	}
	return fmt.Errorf(fmt.Sprintf("invalid level '%v'", str))
}

// zapLevel
/**
 * @Description:
 * @receiver l
 * @return zapcore.Level
 */
func (l Level) zapLevel() zapcore.Level {
	z, found := zapLevels[l]
	if found {
		return z
	}
	return zapcore.InfoLevel
}
