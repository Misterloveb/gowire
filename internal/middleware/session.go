package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

func Session(conf *viper.Viper, store sessions.Store) gin.HandlerFunc {
	return sessions.Sessions(conf.GetString("session.sessionId"), store)
}

func NewCookieStore(conf *viper.Viper) sessions.Store {
	store := cookie.NewStore(
		[]byte(conf.GetString("session.authentication")),
		[]byte(conf.GetString("session.encryption")),
	)
	store.Options(sessions.Options{
		Path:     conf.GetString("cookie.Path"),
		Domain:   conf.GetString("cookie.Domain"),
		MaxAge:   conf.GetInt("cookie.MaxAge"),
		Secure:   conf.GetBool("cookie.Secure"),
		HttpOnly: conf.GetBool("cookie.HttpOnly"),
		SameSite: http.SameSite(conf.GetInt("cookie.SameSite")),
	})

	return store
}
