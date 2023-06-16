package server

import (
	"fmt"
	"gindemo/internal/middleware"
	"gindemo/pkg/writer"
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
	writeor = writer.NewWriter(conf, "httplog")
	if conf.GetString("app.gin_mode") == gin.DebugMode {
		writeor = io.MultiWriter(os.Stdout, writeor)
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
