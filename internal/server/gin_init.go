package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
	"github.com/spf13/viper"
	"io"
	"os"
	"strings"
)

func NewGinEngine(conf *viper.Viper, writer io.Writer) *gin.Engine {
	ginMode := conf.GetString("app.gin_mode")
	gin.SetMode(ginMode)
	r := gin.New()
	r.Use(gin.Recovery())
	w := writer
	if ginMode == gin.DebugMode {
		w = io.MultiWriter(os.Stdout, w)
	}
	skippath := strings.Split(conf.GetString("app.nologurl"), ",")
	logger := gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: ginLogFormatter(),
		Output:    w,
		SkipPaths: skippath,
	})
	r.Use(logger)
	r.MaxMultipartMemory = viper.GetInt64("upload.upload_maxsize") << 20 //MB
	return r
}
func ginLogFormatter() gin.LogFormatter {
	return func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s-[%s] \"%s %s %s %d %dms\" \"%s\"\n",
			params.ClientIP,
			params.TimeStamp.Format(carbon.DateMilliLayout),
			params.Method,
			params.Path,
			params.Request.Proto,
			params.StatusCode,
			params.Latency.Milliseconds(),
			params.Request.UserAgent(),
		)
	}
}
