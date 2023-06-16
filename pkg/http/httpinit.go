package http

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type HTTP struct {
	addr    string
	handler http.Handler
}

func NewHTTP(handler *gin.Engine, conf *viper.Viper) *HTTP {
	return &HTTP{
		addr:    conf.GetString("app.addr"),
		handler: handler,
	}
}

func (h *HTTP) Start() {
	ser := http.Server{
		Addr:    h.addr,
		Handler: h.handler,
	}
	go func() {
		if err := ser.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(fmt.Sprintln("服务器启动失败", err.Error()))
		}
	}()
	close_signal := make(chan os.Signal, 1)
	signal.Notify(close_signal, syscall.SIGINT, syscall.SIGTERM)
	<-close_signal
	log.Println("系统开始关闭。。。")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := ser.Shutdown(ctx); err != nil {
		log.Fatal("系统关闭失败:", err.Error())
	}
	log.Println("系统已成功关闭！")
}
