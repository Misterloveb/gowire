package server

import (
	"fmt"
	"github.com/Misterloveb/gowire/internal/middleware"
	"github.com/Misterloveb/gowire/pkg/writer"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
	"github.com/spf13/viper"
	"io"
	"os"
	"strings"
)

func registerMiddleware(r *gin.Engine, conf *viper.Viper) {
	r.Use(gin.Recovery())
	r.Use(ginLogger(conf))
	//注册session
	r.Use(middleware.Session(conf, middleware.NewCookieStore(conf)))
	//自定义404页面
	r.Use(middleware.NotFound404())
}

func ginLogger(conf *viper.Viper) gin.HandlerFunc {
	var writeor io.Writer

	if conf.GetString("app.gin_mode") == gin.DebugMode {
		writeor = os.Stdout
	} else {
		writeor = writer.NewWriter(conf, "httplog")
	}
	skippath := strings.Split(conf.GetString("app.nologurl"), ",")
	return gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: ginLogFormatter(),
		Output:    writeor,
		SkipPaths: skippath,
	})
}

func ginLogFormatter() gin.LogFormatter {
	return func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s-[%s] \"%s %s %d %dms\"\n",
			params.ClientIP,
			params.TimeStamp.Format(carbon.DateTimeMilliLayout),
			params.Method,
			params.Path,
			params.StatusCode,
			params.Latency.Milliseconds(),
		)
	}
}
