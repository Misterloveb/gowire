package main

import (
	"fmt"
	"gindemo/internal/model"
	"gindemo/router"

	"github.com/spf13/viper"
)

func appStart() {
	//设置viper
	viper.SetConfigFile("./config/config.ini")
	// viper.SetConfigName("config")
	// viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	// viper.OnConfigChange(func(e fsnotify.Event) {
	// 	if viper.IsSet("webapp.static_path") {
	// 		fmt.Println("Config file changed:", e.Name, "webapp.static_path:", viper.GetString("webapp.static_path"))
	// 	}

	// })
	// viper.WatchConfig()
	//初始化数据库
	_ = model.DB()
	//注册路由
	r := router.InitRoutes()
	r.Run(":8080")
}
