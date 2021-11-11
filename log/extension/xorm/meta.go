package xorm

import (
	"context"
	"fmt"
	"github.com/illidaris/apocalypse/pkg/consts"
	context2 "github.com/illidaris/apocalypse/pkg/context"
	"go.uber.org/zap"
	xLog "xorm.io/xorm/log"
)

var (
	logger *zap.Logger // log core without context
)

// FieldsFromCtx write context meta data
func FieldsFromCtx(ctx context.Context) []zap.Field {
	traceID := context2.GetTraceID(ctx)
	sessionID := context2.GetSessionID(ctx)
	actionID := context2.GetAction(ctx)
	stepID := context2.GetStep(ctx)
	return []zap.Field{
		zap.String(consts.TraceID.String(), traceID),
		zap.String(consts.SessionID.String(), sessionID),
		zap.String(consts.Action.String(), actionID),
		zap.String(consts.Step.String(), stepID),
	}
}

// SqlFromLogContext write sql data
func SqlFromLogContext(ctx xLog.LogContext) []zap.Field {
	// write sql cost
	return []zap.Field{
		zap.String(consts.Category.String(), "SQL"),
		zap.Duration(consts.Duration.String(), ctx.ExecuteTime),
	}
}

// MessageFromLogContext build message
func MessageFromLogContext(ctx xLog.LogContext) string {
	var queryString string
	if len(ctx.SQL) > 2048 {
		queryString = ctx.SQL[0:2048]
	} else {
		queryString = ctx.SQL
	}
	args := make([]interface{}, 0)
	if len(ctx.Args) > 100 {
		args = ctx.Args[0:100]
	}
	var rows int64
	var rowErr error
	if ctx.Result != nil {
		rows, rowErr = ctx.Result.RowsAffected()
	}
	return fmt.Sprintf("[sql]%s,[args]%v,[rows]%d,[err]%s", queryString, args, rows, rowErr)
}
