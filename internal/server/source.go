package server

import (
	"html/template"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func registerSource(conf *viper.Viper, r *gin.Engine) {
	//注册自定义模板函数
	setTplFun(r)
	//模板文件
	r.LoadHTMLGlob(conf.GetString("app.template_path"))
	//静态资源
	r.Static("/static", conf.GetString("app.static_path"))
	r.Static("/upload", conf.GetString("upload.upload_path"))
	r.StaticFile("/favicon.ico", conf.GetString("app.static_path")+"/favicon.ico")
}

func setTplFun(r *gin.Engine) {
	r.SetFuncMap(template.FuncMap{
		"splitTpl": splitTpl,
	})
}
func splitTpl(str string) []string {
	return strings.Split(str, "|")
}
