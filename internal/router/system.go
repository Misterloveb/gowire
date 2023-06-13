package router

import (
	"gindemo/internal/controller"

	"github.com/gin-gonic/gin"
)

func systemRoute(r *gin.Engine) {
	system_route := controller.SystemController{}
	g_system := r.Group("/system")

	g_system.GET("/SystemSettings", system_route.SystemSettings)

	g_system.POST("/saveParamsDatas", system_route.SaveParamsDatas)
	g_system.POST("/saveImgDirPath", system_route.SaveImgDirPath)
	g_system.POST("/addResult", system_route.AddResult)
	g_system.POST("/editReuslt", system_route.EditReuslt)
	g_system.POST("/deleteResult", system_route.DeleteResult)
}
