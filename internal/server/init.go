package server

import (
	"github.com/Misterloveb/gowire/internal/controller"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func NewGinEngine(conf *viper.Viper, ctl *controller.RegisterController) *gin.Engine {
	gin.SetMode(conf.GetString("app.gin_mode"))
	r := gin.New()
	//注册全局中间件
	registerMiddleware(r, conf)
	r.MaxMultipartMemory = viper.GetInt64("upload.upload_maxsize") << 20 //MB
	//初始化系统资源
	registerSource(conf, r)
	//注册业务路由
	ctl.Register(r)
	return r
}
