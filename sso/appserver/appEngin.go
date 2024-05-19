package appserver

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Misterloveb/gowire/sso"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/redis/go-redis/v9"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	cookiename    = "PHPSESSIONID"
	ssoserver     = "http://www.sso.com:8083"
	ssoLogin      = ssoserver + "/sso/login?appid="
	ssoChecktoken = ssoserver + "/sso/checktoken"
	ssoLogout     = ssoserver + "/sso/logout"
)

type appEngin struct {
	cache     *cache.Cache
	r         *gin.Engine
	rdb       *redis.Client
	appId     string
	appSecret string
	addr      string
}
type userinfo struct {
	name string
}

func NewApp(appId, appSecret, addr string) *appEngin {
	app := &appEngin{}
	app.cache = cache.New(5*time.Minute, 10*time.Minute)
	app.r = gin.Default()
	app.rdb = sso.NewRedis()
	app.appId = appId
	app.appSecret = appSecret
	app.addr = addr
	return app
}

func (a *appEngin) init() {
	//注册session
	a.r.Use(sso.Session(cookiename, sso.RedisStore()))
	g := a.r.Group("/source", a.middleWareAuth())
	//受保护资源
	g.GET("/", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		ssotoken := session.Get("ssotoken")
		user, _ := a.cache.Get(ssotoken.(string))
		ctx.JSON(http.StatusOK, gin.H{
			"server":   a.appId,
			"username": user.(*userinfo).name,
		})
	})
	//退出登录
	g.GET("/logout", func(ctx *gin.Context) {
		urlstr := fmt.Sprintf(ssoLogout+"?appid=%s&sign=%s", a.appId, sso.CreatSign(a.appId, a.appSecret))
		ctx.Redirect(http.StatusFound, urlstr)
	})
	//首页index
	a.r.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusFound, "/source")
	})
	//异步退出，用于sso http通知退出方式(考虑安全，应该采用token验证)
	a.r.POST("/synclogout", func(ctx *gin.Context) {
		var appdata gin.H
		if err := ctx.ShouldBindJSON(&appdata); err != nil {
			log.Println("json 格式错误")
			return
		}
		ssotoken := appdata["token"]
		if ssotoken == nil {
			log.Println("token 为空")
			return
		}
		//删除自身session
		a.cache.Delete(ssotoken.(string))
	})
	//sso登录成功后回调
	a.r.GET("/ssologin", func(ctx *gin.Context) {
		token := ctx.Query("token")
		ssoSessionId := ctx.Query("sid")
		if token == "" {
			log.Println("token 获取失败")
			return
		}
		if ssoSessionId == "" {
			log.Println("ssoSessionId 获取失败")
			return
		}
		username, token, err := a.checkToken(token, ssoSessionId)
		if err != nil {
			log.Println(err)
			return
		}
		//sso登录成功
		sessionobj := sessions.Default(ctx)
		sessionobj.Set("ssotoken", token)
		sessionobj.Save()
		a.cache.Set(token, &userinfo{name: username}, cache.NoExpiration)
		//redis订阅退出,根据sso的sessionid区分不同客户端
		go a.subLogout(ssoSessionId)
		ctx.Redirect(http.StatusFound, "/")
	})
}

//订阅sso退出事件
func (a *appEngin) subLogout(ssoid string) {
	pubsub := a.rdb.Subscribe(context.Background(), ssoid)
	if err := pubsub.Ping(context.Background()); err != nil {
		log.Println(a.appId, "订阅失败")
		return
	}
	log.Println(a.appId, "订阅成功--", ssoid)
	defer pubsub.Close()
	tokenlist := make(map[string]string, 10)
	//开始监听
	for msg := range pubsub.Channel() {
		if msg.Payload != "" {
			//收到退出消息
			if err := json.Unmarshal([]byte(msg.Payload), &tokenlist); err != nil {
				log.Println("json解析失败", err.Error())
				return
			}
			//清除自身session
			a.cache.Delete(tokenlist[a.appId])
			log.Println(a.appId, "退出成功")
			return
		}
		log.Println("退出失败，收到消息为空")
		return
	}
}
func (a *appEngin) checkToken(token string, ssoid string) (string, string, error) {
	data := gin.H{"appid": a.appId, "sign": sso.CreatSign(a.appId, a.appSecret), "token": token, "session_id": ssoid}
	res, err := json.Marshal(data)
	if err != nil {
		return "", "", errors.New("json 格式化失败" + err.Error())
	}
	requertsso, err := http.Post(ssoChecktoken, "application/json", bytes.NewReader(res))
	var reqdata gin.H
	req, _ := io.ReadAll(requertsso.Body)
	defer requertsso.Body.Close()
	if err := json.Unmarshal(req, &reqdata); err != nil {
		return "", "", errors.New("json 解析失败" + err.Error())
	}
	if reqdata["status"].(string) != "ok" {
		return "", "", errors.New("校验token失败：" + reqdata["msg"].(string))
	}
	//校验成功,返回用户信息和新的token
	return reqdata["userinfo"].(string), reqdata["token"].(string), nil
}
func (a *appEngin) middleWareAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Cookie(cookiename)
		ssologinurl := ssoLogin + a.appId
		if err != nil || cookie == "" {
			ctx.Redirect(http.StatusFound, ssologinurl)
			ctx.Abort()
			return
		}
		ssotoken := sessions.Default(ctx).Get("ssotoken")
		if ssotoken == nil {
			ctx.Redirect(http.StatusFound, ssologinurl)
			ctx.Abort()
			return
		}
		_, ok := a.cache.Get(ssotoken.(string))
		if !ok {
			ctx.Redirect(http.StatusFound, ssologinurl)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
func (a *appEngin) Run() {
	a.init()
	a.r.Run(a.addr)
}
