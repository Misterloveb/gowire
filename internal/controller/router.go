package controller

import (
	"github.com/Misterloveb/gowire/internal/middleware"
	"github.com/gin-gonic/gin"
	"reflect"
)

type RegisterController struct {
	Login  *LoginController
	Loger  *LogerController
	Index  *IndexController
	System *SystemController
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
func (ctl *LoginController) RegisterRouter(r *gin.Engine) {
	r.StaticFile("/login", "./internal/view/login/login.html")
	r.StaticFile("/register", "./internal/view/login/reg.html")
	r.POST("/register", ctl.Register)
	r.POST("/captcha", ctl.GetCaptcha)
	r.POST("/login", ctl.Login)
	r.GET("/forget", ctl.Forget)
	r.GET("/logout", ctl.Logout)

	r.Use(middleware.AuthVerifition)
}

func (ctl *IndexController) RegisterRouter(r *gin.Engine) {
	r.GET("/", ctl.Index)
	gIndex := r.Group("/index")

	gIndex.GET("/Handinsert", ctl.Handinsert)
	gIndex.GET("/explodeExcel", ctl.ExplodeExcel)

	gIndex.POST("/saveDatas", ctl.SaveDatas)
	gIndex.POST("/saveImages", ctl.SaveImages)
	gIndex.POST("/importExcel", ctl.ImportExcel)

}

func (ctl *LogerController) RegisterRouter(r *gin.Engine) {
	gRoute := r.Group("/loger")

	gRoute.GET("/Querylog", ctl.Querylog)
	gRoute.GET("/viewImage", ctl.ViewImage)
	gRoute.GET("/exportDatas", ctl.ExportDatas)

	gRoute.POST("/searchData", ctl.SearchData)
	gRoute.POST("/searchResult", ctl.SearchResult)
	gRoute.POST("/deleteData", ctl.DeleteData)
	gRoute.POST("/saveLogs", ctl.SaveLogs)
	gRoute.POST("/checkImg", ctl.CheckImg)
	gRoute.POST("/openFileDir", ctl.OpenFileDir)
	gRoute.POST("/deleteLog", ctl.DeleteLog)
	gRoute.POST("/getData", ctl.GetData)
	gRoute.POST("/deleteImg", ctl.DeleteImg)
}

func (ctl *SystemController) RegisterRouter(r *gin.Engine) {
	gSystem := r.Group("/system")

	gSystem.GET("/SystemSettings", ctl.SystemSettings)

	gSystem.POST("/saveParamsDatas", ctl.SaveParamsDatas)
	gSystem.POST("/saveImgDirPath", ctl.SaveImgDirPath)
	gSystem.POST("/addResult", ctl.AddResult)
	gSystem.POST("/editReuslt", ctl.EditReuslt)
	gSystem.POST("/deleteResult", ctl.DeleteResult)
}
