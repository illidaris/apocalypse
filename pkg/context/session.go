package context

import (
	"context"
	"github.com/illidaris/apocalypse/pkg/consts"
)

func SetSessionID(ctx context.Context, v string) context.Context {
	return context.WithValue(ctx, consts.SessionID, v)
}

func GetSessionID(ctx context.Context) string {
	return GetString(ctx, consts.SessionID)
}

func SetSessionBirth(ctx context.Context, v int64) context.Context {
	return context.WithValue(ctx, consts.SessionBirth, v)
}

func GetSessionBirth(ctx context.Context) int64 {
	return GetInt64(ctx, consts.SessionBirth)
}
