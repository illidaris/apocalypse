package xorm

import (
	"fmt"
	"github.com/illidaris/apocalypse/pkg/consts"
	"go.uber.org/zap"
	xLog "xorm.io/xorm/log"
)

type XLogger struct{}

// BeforeSQL implements ContextLogger
func (l *XLogger) BeforeSQL(ctx xLog.LogContext) {}

// AfterSQL implements ContextLogger
func (l *XLogger) AfterSQL(ctx xLog.LogContext) {
	// write context meta data
	fields := FieldsFromCtx(ctx.Ctx)
	// write sql data
	fields = append(fields, SqlFromLogContext(ctx)...)
	// write sql param and result
	msg := MessageFromLogContext(ctx)
	// write exec error
	if ctx.Err != nil {
		fields = append(fields, zap.String(consts.Error.String(), ctx.Err.Error()))
		logger.Error(msg, fields...)
	} else {
		logger.Info(msg, fields...)
	}
}

// Debugf implements ContextLogger
func (l *XLogger) Debugf(format string, v ...interface{}) {
	logger.Debug(fmt.Sprintf(format, v...))
}

// Errorf implements ContextLogger
func (l *XLogger) Errorf(format string, v ...interface{}) {
	logger.Error(fmt.Sprintf(format, v...))
}

// Infof implements ContextLogger
func (l *XLogger) Infof(format string, v ...interface{}) {
	logger.Info(fmt.Sprintf(format, v...))
}

// Warnf implements ContextLogger
func (l *XLogger) Warnf(format string, v ...interface{}) {
	logger.Warn(fmt.Sprintf(format, v...))
}

// Level implements ContextLogger
func (l *XLogger) Level() xLog.LogLevel {
	// TODO: 做一层转化
	return xLog.LOG_INFO
}

// SetLevel implements ContextLogger
func (l *XLogger) SetLevel(lv xLog.LogLevel) {
	// TODO: 做一层转化

}

// ShowSQL implements ContextLogger
func (l *XLogger) ShowSQL(show ...bool) {

}

// IsShowSQL implements ContextLogger
func (l *XLogger) IsShowSQL() bool {
	return true
}
