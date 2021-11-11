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
