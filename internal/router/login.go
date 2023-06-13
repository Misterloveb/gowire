package router

import (
	"gindemo/internal/controller"
	"gindemo/internal/middleware"
	"github.com/gin-gonic/gin"
)

func loginRoute(r *gin.Engine) {
	login_route := &controller.LoginController{}
	r.StaticFile("/login", "./internal/view/login/login.html")
	r.StaticFile("/register", "./internal/view/login/reg.html")
	r.POST("/register", login_route.Register)
	r.POST("/captcha", login_route.GetCaptcha)
	r.POST("/login", login_route.Login)
	r.GET("/forget", login_route.Forget)
	r.GET("/logout", login_route.Logout)

	r.Use(middleware.AuthVerifition)
}
