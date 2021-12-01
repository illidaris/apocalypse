package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/illidaris/apocalypse/log/logger"
	"github.com/illidaris/apocalypse/pkg/consts"
	myContext "github.com/illidaris/apocalypse/pkg/context"
	"go.uber.org/zap"
	"time"
)

type HttpMetaData consts.MetaData

const (
	HTTPStatusCode  HttpMetaData = "statusCode"
	HTTPContentType HttpMetaData = "contentType"
	HTTPMethod      HttpMetaData = "httpMethod"
	HTTPPath        HttpMetaData = "httpPath"
	HTTPQuery       HttpMetaData = "httpQuery"
	HTTPClientIp    HttpMetaData = "httpClientIp"
	HTTPUserAgent   HttpMetaData = "httpUserAgent"
)

// WithTrace add trace log context
func WithTrace(c *gin.Context, birth time.Time) *gin.Context {
	// trace
	traceID := c.GetHeader("X-Request-ID")
	// session
	uuid, _ := uuid.NewUUID()
	sID := uuid.String()
	if traceID == "" {
		traceID = sID
	}
	sessionBirth := birth.UTC().UnixNano()
	// assembly trace & session
	ctx := c.Request.Context()
	ctx = myContext.SetTraceID(ctx, traceID)
	ctx = myContext.SetSessionID(ctx, sID)
	ctx = myContext.SetSessionBirth(ctx, sessionBirth)
	ctx = logger.NewContext(ctx,
		zap.String(consts.TraceID.String(), traceID),
		zap.String(consts.SessionID.String(), sID),
		zap.Int64(consts.SessionBirth.String(), sessionBirth))
	// instead of request
	c.Request = c.Request.WithContext(ctx)
	return c
}
