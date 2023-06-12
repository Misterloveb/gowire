package router

import (
	"gindemo/middleware"
	"html/template"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func splitTpl(str string) []string {
	return strings.Split(str, "|")
}
func setTplFun(r *gin.Engine) {
	r.SetFuncMap(template.FuncMap{
		"splitTpl": splitTpl,
	})
}
func setMiddleware(r *gin.Engine) {
	//自定义404页面
	//注册session
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("PHPSESSIONID", store))
	r.Use(middleware.NotFound404)
}

func InitRoutes() *gin.Engine {
	gin.SetMode(viper.GetString("webapp.gin_mode"))
	r := gin.Default()
	r.MaxMultipartMemory = viper.GetInt64("upload.upload_maxsize") << 20 //MB
	//注册中间件
	setMiddleware(r)
	//注册自定义模板函数
	setTplFun(r)
	//模板文件
	r.LoadHTMLGlob(viper.GetString("webapp.template_path"))
	//静态资源
	r.Static("/static", viper.GetString("webapp.static_path"))
	r.Static("/upload", viper.GetString("upload.upload_path"))
	r.StaticFile("/favicon.ico", viper.GetString("webapp.static_path")+"/favicon.ico")
	//注册业务路由
	registerAppRoute(r)
	return r
}
func registerAppRoute(r *gin.Engine) {
	loginRoute(r)
	indexRoute(r)
	logerRoute(r)
	systemRoute(r)
}
