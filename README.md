# apocalypse
> this is a log frame.
#### project struct
```
apocalypse
├─log
│  ├─extension
│  │  ├─gin
│  │  └─xorm
│  └─logger
└─pkg
    ├─consts
    └─context
```
#### log

```go
package main

import (
	"context"
	"go.uber.org/zap/zapcore"
	"github.com/illidaris/apocalypse/log/logger"
)

func main() {
	logger.New(nil)
	
	ctx:=context.TODO()
	logger.InfoCtx(ctx,"info log")
	logger.WithContext(ctx).Info("info log")
	logger.Info("info log")
}

```

#### log with gin

gin logger will record gin log

```go
package main

import (
	"net/http"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
	ginExt "github.com/illidaris/apocalypse/log/gin"
	"github.com/illidaris/apocalypse/log/logger"
)

func main() {
	// init log core
	logger.New(nil)
	// init gin
	router := gin.Default()
	router.Use(ginExt.LoggerHandler())
	router.Use(ginExt.RecoverHandler())
	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "success")
	})
	router.GET("/error", func(c *gin.Context) {
		panic(errors.New("this is an error"))
	})
}

```

#### log with XOrm

xorm logger will record xorm log

```go
package main

import (
	"xorm.io/xorm"
	_ "github.com/go-sql-driver/mysql"
	xormExt "github.com/illidaris/apocalypse/log/gin"
	"github.com/illidaris/apocalypse/log/logger"
)

func main() {
	// init log core
	logger.New(nil)
	eng, err := xorm.NewEngine("mysql", "root:123@/test?charset=utf8")
	if err != nil {
		panic(err)
	}
	eng.ShowSQL(true)
	// assembly xorm log
	eng.SetLogger(xormExt.NewXLogger())
}

```