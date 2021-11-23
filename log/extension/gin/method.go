package gin

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/illidaris/apocalypse/log/logger"
	"github.com/illidaris/apocalypse/pkg/consts"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

// LoggerHandler record log
func LoggerHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// trace
		traceID := c.GetHeader("X-Request-ID")
		// content type
		contentType := c.GetHeader("Content-Type")
		// session
		sessionID, _ := uuid.NewUUID()
		sessionBirth := time.Now()
		// assembly trace & session
		rawCtx := c.Request.Context()
		ctx := logger.NewContext(rawCtx,
			zap.String(consts.TraceID.String(), traceID),
			zap.String(consts.SessionID.String(), sessionID.String()),
			zap.Int64(consts.SessionBirth.String(), sessionBirth.UTC().UnixNano()))
		// instead of request
		c.Request = c.Request.WithContext(ctx)
		// before
		c.Next()
		// after
		cost := time.Since(sessionBirth)
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		curCtx := c.Request.Context()
		curCtx = logger.NewContext(curCtx,
			zap.String(string(HTTPMethod), c.Request.Method),
			zap.String(string(HTTPContentType), contentType),
			zap.String(string(HTTPPath), path),
			zap.String(string(HTTPQuery), query),
			zap.String(string(HTTPClientIp), c.ClientIP()),
			zap.String(string(HTTPUserAgent), c.Request.UserAgent()),
			zap.Int(string(HTTPStatusCode), c.Writer.Status()),
			zap.Duration(consts.Duration.String(), cost),
		)
		logger.InfoCtx(curCtx, c.Errors.ByType(gin.ErrorTypePrivate).String())
	}
}

// RecoverHandler recover from panic
func RecoverHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				fmt.Println(err)
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					var se *os.SyscallError
					if ok := errors.Is(ne.Err, se); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.WithContext(c.Request.Context()).Error(c.Request.URL.Path, zap.Any(consts.Error.String(), err), zap.String("HttpRequest", string(httpRequest)))
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: error check
					c.Abort()
					return
				}

				logger.WithContext(c.Request.Context()).Error("recover from panic",
					zap.Any(consts.Error.String(), err),
					zap.String("HttpRequest", string(httpRequest)),
					zap.String("stack", string(debug.Stack())))

				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
