package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthVerifition(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("userinfo")
	if user == nil {
		obj := map[string]string{
			"time":  "5",
			"title": "未登录系统！",
		}
		c.HTML(http.StatusUnauthorized, "error401.html", obj)
		c.Abort()
	}
	c.Next()
}
