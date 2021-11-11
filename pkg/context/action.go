package context

import (
	"context"
	"github.com/illidaris/apocalypse/pkg/consts"
)

func SetAction(ctx context.Context, v string) context.Context {
	return context.WithValue(ctx, consts.Action, v)
}

func GetAction(ctx context.Context) string {
	return GetString(ctx, consts.Action)
}

func SetStep(ctx context.Context, v string) context.Context {
	return context.WithValue(ctx, consts.Step, v)
}

func GetStep(ctx context.Context) string {
	return GetString(ctx, consts.Step)
}
