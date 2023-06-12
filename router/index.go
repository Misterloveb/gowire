package router

import (
	"gindemo/internal/controller"

	"github.com/gin-gonic/gin"
)

func indexRoute(r *gin.Engine) {
	index_route := controller.IndexController{}
	r.GET("/", index_route.Index)
	g_index := r.Group("/index")

	g_index.GET("/Handinsert", index_route.Handinsert)
	g_index.GET("/explodeExcel", index_route.ExplodeExcel)

	g_index.POST("/saveDatas", index_route.SaveDatas)
	g_index.POST("/saveImages", index_route.SaveImages)
	g_index.POST("/importExcel", index_route.ImportExcel)

}
