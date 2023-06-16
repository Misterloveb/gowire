package server

import (
	"github.com/Misterloveb/gowire/internal/controller"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"reflect"
)

type RegisterController struct {
	Login  *controller.LoginController
	Loger  *controller.LogerController
	Index  *controller.IndexController
	System *controller.SystemController
}

func (ctl *RegisterController) Register(r *gin.Engine) {
	valuer := reflect.ValueOf(ctl).Elem()
	numField := valuer.NumField()
	for i := 0; i < numField; i++ {
		if obj := valuer.Field(i).MethodByName("RegisterRouter"); obj != reflect.ValueOf(nil) {
			obj.Call([]reflect.Value{reflect.ValueOf(r)})
		}
	}
}

func NewGinEngine(conf *viper.Viper, ctl *RegisterController) *gin.Engine {
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
