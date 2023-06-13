package router

import (
	"gindemo/internal/controller"

	"github.com/gin-gonic/gin"
)

func logerRoute(r *gin.Engine) {
	loger_route := controller.LogerController{}
	g_route := r.Group("/loger")

	g_route.GET("/Querylog", loger_route.Querylog)
	g_route.GET("/viewImage", loger_route.ViewImage)
	g_route.GET("/exportDatas", loger_route.ExportDatas)

	g_route.POST("/searchData", loger_route.SearchData)
	g_route.POST("/searchResult", loger_route.SearchResult)
	g_route.POST("/deleteData", loger_route.DeleteData)
	g_route.POST("/saveLogs", loger_route.SaveLogs)
	g_route.POST("/checkImg", loger_route.CheckImg)
	g_route.POST("/openFileDir", loger_route.OpenFileDir)
	g_route.POST("/deleteLog", loger_route.DeleteLog)
	g_route.POST("/getData", loger_route.GetData)
	g_route.POST("/deleteImg", loger_route.DeleteImg)
}
