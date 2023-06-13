package http

import (
	"context"
	"fmt"
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

func NewHTTP(handler http.Handler, conf *viper.Viper) *HTTP {
	return &HTTP{
		addr:    conf.GetString("app.addr"),
		handler: handler,
	}
}

func (h *HTTP) Start() {
	server := http.Server{
		Addr:    "",
		Handler: h.handler,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(fmt.Sprintln("服务器启动失败", err.Error()))
		}
	}()
	close_signal := make(chan os.Signal, 1)
	signal.Notify(close_signal, syscall.SIGINT, syscall.SIGTERM)
	<-close_signal
	log.Println("系统开始关闭。。。")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("系统关闭失败:", err.Error())
	}
	log.Println("系统已成功关闭！将再5秒后自动重启")
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*10)
	restart := time.AfterFunc(time.Second*5, func() {
		h.Start()
	})
	defer cancel()
	select {
	case <-ctx.Done():
		log.Fatal("系统重启超时")
	case <-restart.C:
		restart.Stop()
		log.Println("系统重启成功")
	}
}
