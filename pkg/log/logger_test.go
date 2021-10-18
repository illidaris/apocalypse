package log

import (
	"go.uber.org/zap"
	"testing"
)

func TestNew(t *testing.T) {
	New(nil)
	zap.L().Info("123")
}
