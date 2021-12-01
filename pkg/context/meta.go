package context

import (
	"context"
)

func GetString(ctx context.Context, key interface{}) string {
	v := ctx.Value(key)
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}

func GetInt64(ctx context.Context, key interface{}) int64 {
	v := ctx.Value(key)
	if s, ok := v.(int64); ok {
		return s
	}
	return -1
}
