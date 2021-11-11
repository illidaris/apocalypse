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
