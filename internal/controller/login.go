package controller

import (
	"crypto/sha256"
	"fmt"
	"gindemo/internal/model"
	"gindemo/pkg/verificationcode"
	"log"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mojocn/base64Captcha"
)

type LoginController struct {
	VerifiCode *verificationcode.VerifiCode
}

func (l *LoginController) Login(ctx *gin.Context) {
	// captcha_id := ctx.PostForm("captid")
	session := sessions.Default(ctx)
	// vercode := ctx.PostForm("vercode")
	if l.VerifiCode == nil {
		ctx.JSON(200, gin.H{"status": 0, "msg": "验证码错误"})
		return
	}
	// if !l.VerifiCode.Verify(captcha_id, vercode) {
	// 	ctx.JSON(200, gin.H{"status": 0, "msg": "验证码错误"})
	// 	return
	// }
	username := ctx.PostForm("username")
	if !validate(username, ctx.PostForm("password")) {
		ctx.JSON(200, gin.H{"status": 0, "msg": "用户不存在或秘密错误"})
		return
	}
	session.Set("userinfo", username)
	session.Save()
	ctx.JSON(200, gin.H{"status": 1})
}
func validate(username string, password string) bool {
	user := &model.User{}
	u := user.GetData("username=?", username)
	if len(u) == 0 {
		return false
	}
	if strings.Compare(u[0].Password, hashPassword([]byte(password), u[0].Salt)) != 0 {
		return false
	}
	return true
}
func (l *LoginController) Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Delete("userinfo")
	session.Save()
	ctx.HTML(http.StatusOK, "login.html", "")
}
func (l *LoginController) Forget(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "forget.html", "")
}
func (l *LoginController) Register(ctx *gin.Context) {
	user := &model.User{}
	pwd := []byte(strings.Trim(ctx.PostForm("password"), " "))
	salt := fmt.Sprintf("%x", sha256.Sum256([]byte(uuid.NewString())))
	user.Salt = salt
	user.UserName = strings.Trim(ctx.PostForm("nickname"), " ")
	user.Password = hashPassword(pwd, salt)
	if err := user.Insert(); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"status": 0})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": 1})
}
func hashPassword(pwd []byte, salt string) string {
	pwd_arr := make([]byte, 0, len(salt)+len(pwd))
	pwd_arr = append(pwd_arr, salt...)
	pwd_arr = append(pwd_arr, pwd...)
	return fmt.Sprintf("%x", sha256.Sum256(pwd_arr))
}

func (l *LoginController) GetCaptcha(ctx *gin.Context) {
	driver := base64Captcha.DefaultDriverDigit
	l.VerifiCode = verificationcode.NewVerifiCode(driver)
	id, b64s, err := l.VerifiCode.Generate()
	if err != nil {
		log.Println(err.Error())
		return
	}
	ctx.JSON(200, gin.H{
		"id":   id,
		"b64s": b64s,
	})
}
