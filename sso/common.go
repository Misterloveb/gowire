package sso

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	redis2 "github.com/redis/go-redis/v9"
	"log"
)

func Session(sessionId string, store sessions.Store) gin.HandlerFunc {
	return sessions.Sessions(sessionId, store)
}
func RedisStore() sessions.Store {
	store, err := redis.NewStoreWithDB(10, "tcp", "localhost:6379", "", "1", []byte("secretsecretqwer"))
	if err != nil {
		log.Fatal(err)
	}
	return store
}
func CreatSign(appid, secret string) string {
	//验证签名是否正确
	hah := sha256.New()
	hah.Write([]byte(appid))
	hah.Write([]byte(secret))
	return hex.EncodeToString(hah.Sum(nil))
}
func NewRedis() *redis2.Client {
	rdb := redis2.NewClient(&redis2.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       1,
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatal(err)
	}
	return rdb
}
