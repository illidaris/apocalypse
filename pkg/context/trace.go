package context

import (
	"context"
	"github.com/illidaris/apocalypse/pkg/consts"
)

func SetTraceID(ctx context.Context, v string) context.Context {
	return context.WithValue(ctx, consts.TraceID, v)
}

func GetTraceID(ctx context.Context) string {
	return GetString(ctx, consts.TraceID)
}
