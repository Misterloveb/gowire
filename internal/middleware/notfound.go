package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotFound404() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		if ctx.Writer.Status() == http.StatusNotFound {
			obj := map[string]string{
				"time":  "5",
				"title": "",
			}
			ctx.HTML(http.StatusOK, "error404.html", obj)
		}
	}
}
