package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	//设置viper
	conf := viper.New()
	conf.AddConfigPath("config")
	conf.SetConfigName("config")
	if err := conf.ReadInConfig(); err != nil {
		panic(fmt.Errorf("加载日志文件失败: %w", err))
	}
	//监听文件变化事件
	// conf.OnConfigChange(func(e fsnotify.Event) {
	// 	if conf.IsSet("webapp.static_path") {
	// 		fmt.Println("Config file changed:", e.Name, "webapp.static_path:", viper.GetString("webapp.static_path"))
	// 	}
	// })
	return conf
}
