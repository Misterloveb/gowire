package controller

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/Misterloveb/gowire/internal/common"
	"github.com/Misterloveb/gowire/internal/model"
	"github.com/Misterloveb/gowire/pkg/verificationcode"
	"log"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mojocn/base64Captcha"
)

type LoginController struct {
	VerifyCode *verificationcode.VerifiCode
	*common.Server
}

func (ctl *LoginController) Login(ctx *gin.Context) {
	if err := ctl.validate(ctx); err != nil {
		ctx.JSON(200, gin.H{"status": 0, "msg": err.Error()})
		return
	}
	session := sessions.Default(ctx)
	session.Set("userinfo", ctx.PostForm("username"))
	session.Save()
	ctx.JSON(200, gin.H{"status": 1})
}
func (ctl *LoginController) validate(ctx *gin.Context) error {
	captcha_id := ctx.PostForm("captid")
	vercode := ctx.PostForm("vercode")
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	if ctl.VerifyCode == nil || !ctl.VerifyCode.Verify(captcha_id, vercode) {
		return errors.New("验证码错误")
	}
	u := ctl.UserDao.GetData("username=?", username)
	if len(u) == 0 {
		return errors.New("用户不存在或密码错误")
	}
	if strings.Compare(u[0].Password, hashPassword([]byte(password), u[0].Salt)) != 0 {
		return errors.New("用户不存在或密码错误")
	}
	return nil
}
func (ctl *LoginController) Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Delete("userinfo")
	session.Save()
	ctx.HTML(http.StatusOK, "login.html", "")
}
func (ctl *LoginController) Register(ctx *gin.Context) {
	user := &model.User{}
	pwd := []byte(strings.Trim(ctx.PostForm("password"), " "))
	salt := fmt.Sprintf("%x", sha256.Sum256([]byte(uuid.NewString())))
	user.Salt = salt
	user.UserName = strings.Trim(ctx.PostForm("nickname"), " ")
	user.Password = hashPassword(pwd, salt)
	if err := ctl.UserDao.Insert(user); err != nil {
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

func (ctl *LoginController) GetCaptcha(ctx *gin.Context) {
	driver := base64Captcha.DefaultDriverDigit
	ctl.VerifyCode = verificationcode.NewVerifiCode(driver)
	id, b64s, err := ctl.VerifyCode.Generate()
	if err != nil {
		log.Println(err.Error())
		return
	}
	ctx.JSON(200, gin.H{
		"id":   id,
		"b64s": b64s,
	})
}
