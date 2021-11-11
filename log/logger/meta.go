package logger

import "go.uber.org/zap"

var (
	config *Config // config

	ctxLogger  *zap.Logger // log core key-value from context
	funcLogger *zap.Logger // log core without context
)
