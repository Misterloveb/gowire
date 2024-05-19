package ssomain

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Misterloveb/gowire/sso"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"
)

type appServer struct {
	domain   string
	secret   string
	redirect string
	logout   string
}

var ssoRedirect = map[string]appServer{
	"app1": {
		domain:   "www.app1.com:8081",
		secret:   "app1_123456",
		redirect: "http://www.app1.com:8081/ssologin",
		logout:   "http://www.app1.com:8081/synclogout",
	},
	"app2": {
		domain:   "www.app2.com:8082",
		secret:   "app2_123456",
		redirect: "http://www.app2.com:8082/ssologin",
		logout:   "http://www.app2.com:8082/synclogout",
	},
}

type appView struct {
	Appid     string `json:"appid"`
	Sign      string `json:"sign"`
	Token     string `json:"token"`
	SessionId string `json:"session_id"` //每个客户端对应的sso上的sessionid
}

var memcache = cache.New(5*time.Minute, 10*time.Minute)

func TestSSO(t *testing.T) {
	rdb := sso.NewRedis()
	r := gin.Default()
	r.Static("/static", "../../static")
	r.LoadHTMLFiles("./login.html")
	r.Use(sso.Session("sso", sso.RedisStore()))
	//登录接口
	g := r.Group("/sso", MiddlewareAuth)
	g.GET("/login", func(ctx *gin.Context) {
		appid := ctx.Query("appid")
		res, _ := ssoRedirect[appid]
		//检查是否有应用登录
		sobj := sessions.Default(ctx)
		loginapp := sobj.Get("loginapp")
		if loginapp == nil || loginapp.(string) == "" {
			//没有客户端登录过
			ctx.HTML(http.StatusOK, "login.html", gin.H{
				"appid":    appid,
				"redirect": res.redirect,
			})
			return
		}
		//已登录过,生成并保存token以用作验证
		token := creatToken(appid)
		memcache.Set(token, res.secret, time.Minute)
		ctx.Redirect(http.StatusFound, res.redirect+"?token="+token+"&sid="+sobj.ID())
	})

	g.POST("/login", func(ctx *gin.Context) {
		appid := ctx.PostForm("appid")
		redirect := ctx.PostForm("redirect")
		username := ctx.PostForm("username")
		data, _ := ssoRedirect[appid]
		if strings.Compare(data.redirect, redirect) != 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": 0, "msg": "登录失败！回调地址不合法"})
			return
		}
		//生成token,以及写入cookie，记录登录信息
		sessobj := sessions.Default(ctx)
		sessobj.Set("loginapp", appid) //用于get /login验证
		sessobj.Save()
		sessid := sessobj.ID()
		token := creatToken(appid) //临时token
		setMemCache(sessid, "", "", username)
		memcache.Set(token, appid, time.Minute)
		ctx.JSON(http.StatusOK, gin.H{"status": 1, "url": data.redirect + "?token=" + token + "&sid=" + sessid})
	})
	g.POST("/checktoken", func(ctx *gin.Context) {
		//需要传入 appid,sign(appid+appsecret),sessionid
		var appdata appView
		if err := ctx.ShouldBindJSON(&appdata); err != nil {
			ctx.JSON(http.StatusOK, gin.H{"status": "fail", "msg": "json格式错误"})
			return
		}
		data, _ := ssoRedirect[appdata.Appid]

		//验证签名是否正确
		if !chekSign(appdata.Appid, data.secret, appdata.Sign) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "msg": "签名验证失败"})
			return
		}
		//验证token是否存在
		serverobj, ok := memcache.Get("serverList")
		if !ok {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "msg": "不合法请求"})
			return
		}
		serv, _ := serverobj.(*serverLogined)
		_, ok = memcache.Get(appdata.Token)
		if !ok {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "msg": "token不合法"})
			return
		}
		memcache.Delete(appdata.Token)
		userobj := serv.Get(appdata.SessionId)
		if userobj == nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "msg": "sessionid不合法"})
			return
		}
		userinfo := userobj.userinfo
		token := creatToken(appdata.Appid) //正式的token
		serv.AddToken(appdata.SessionId, token)
		serv.AddAppid(appdata.SessionId, appdata.Appid)
		memcache.Set("serverList", serv, cache.NoExpiration)
		ctx.JSON(http.StatusOK, gin.H{"status": "ok", "msg": "", "userinfo": userinfo, "token": token})
	})
	g.GET("/logout", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		appid := ctx.Query("appid")
		sign := ctx.Query("sign")
		data, _ := ssoRedirect[appid]
		//验证签名是否正确
		if !chekSign(appid, data.secret, sign) {
			ctx.String(http.StatusOK, "签名验证失败")
			return
		}
		//删除sso自身session
		session.Clear()
		session.Save()
		//通知其他已登录客户端退出，http
		//if err := logoutWithHttp(session.ID()); err != nil {
		//	ctx.String(200, "通知失败：", err.Error())
		//	return
		//}
		//通知其他已登录客户端退出，redis
		if err := logoutWithRedis(rdb, session.ID()); err != nil {
			ctx.String(200, "通知失败：", err.Error())
			return
		}
		//成功后跳转到发起退出请求的app首页
		ctx.Redirect(http.StatusFound, "http://"+data.domain)
	})
	r.Run("www.sso.com:8083")
}

//redis发布订阅退出
func logoutWithRedis(rdb *redis.Client, sessionid string) (err error) {
	serverlist, ok := memcache.Get("serverList")
	if !ok {
		return errors.New("非法请求")
	}
	serv, _ := serverlist.(*serverLogined)
	sobj := serv.Get(sessionid)
	serv.Del(sessionid) //删除对应的会话信息
	memcache.Set("serverList", serv, cache.NoExpiration)
	if sobj == nil {
		return nil
	}
	tokenlist := sobj.token
	response := make(map[string]string, len(tokenlist))
	for k, appid := range sobj.appid {
		response[appid] = tokenlist[k]
	}
	reponsebyte, _ := json.Marshal(response)
	//用sso的sessionid区分不同的客户端
	//重试机制
	for i := 0; i < 3; i++ {
		res := rdb.Publish(context.Background(), sessionid, reponsebyte)
		if err = res.Err(); err != nil {
			log.Printf("%s-%d-发布失败：%s", sessionid, i, err.Error())
			continue
		}
		break
	}
	return
}

//http方式通知退出
func logoutWithHttp(sessionid string) error {
	serverlist, ok := memcache.Get("serverList")
	if !ok {
		return errors.New("非法请求")
	}
	serv, _ := serverlist.(*serverLogined)
	sobj := serv.Get(sessionid)
	serv.Del(sessionid) //删除对应的会话信息
	memcache.Set("serverList", serv, cache.NoExpiration)
	if sobj == nil {
		return nil
	}
	tokenlist := sobj.token
	appidlist := sobj.appid
	var httperr error
	for k, appid := range appidlist {
		rdata := "{\"token\":\"" + tokenlist[k] + "\"}"
		_, httperr = http.Post(ssoRedirect[appid].logout, "application/json", bytes.NewReader([]byte(rdata)))
	}
	return httperr
}
func setMemCache(sessid, appid, token, userinfo string) {
	serverobj, ok := memcache.Get("serverList")
	if ok {
		server, _ := serverobj.(*serverLogined)
		if appid != "" {
			server.AddAppid(sessid, appid)
		}
		if token != "" {
			server.AddToken(sessid, token)
		}
		if userinfo != "" {
			server.AddUser(sessid, userinfo)
		}
		memcache.Set("serverList", server, cache.NoExpiration)
	} else {
		serverToken := &serverLogined{}
		if appid != "" {
			serverToken.AddAppid(sessid, appid)
		}
		if token != "" {
			serverToken.AddToken(sessid, token)
		}
		if userinfo != "" {
			serverToken.AddUser(sessid, userinfo)
		}
		memcache.Set("serverList", serverToken, cache.NoExpiration)
	}
}
func chekSign(appid, secret, sign string) bool {
	//验证签名是否正确
	hah := sha256.New()
	hah.Write([]byte(appid))
	hah.Write([]byte(secret))
	if strings.Compare(sign, hex.EncodeToString(hah.Sum(nil))) != 0 {
		return false
	}
	return true
}
func creatToken(appid string) string {
	hsh := sha256.New()
	hsh.Write([]byte(appid))
	hsh.Write([]byte(uuid.NewString()))
	return base64.URLEncoding.EncodeToString([]byte(fmt.Sprintf("%x", hsh.Sum(nil))))
}
func MiddlewareAuth(ctx *gin.Context) {
	//校验appid合法性
	appid := ""
	if ctx.Request.Method == http.MethodGet {
		appid = ctx.Query("appid")
	}
	if ctx.Request.Method == http.MethodPost {
		if ctx.Request.Header.Get("Content-Type") == "application/json" {
			ctx.Next()
			return
		}
		appid = ctx.PostForm("appid")
	}
	_, ok := ssoRedirect[appid]
	if !ok {
		ctx.String(http.StatusOK, "非法请求")
		ctx.Abort()
		return
	}
	ctx.Next()
}
