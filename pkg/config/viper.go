package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	//设置viper
	conf := viper.New()
	conf.SetEnvPrefix("GW")
	conf.AutomaticEnv()
	conf.AddConfigPath("config")
	conf.SetConfigName("config")
	if err := conf.ReadInConfig(); err != nil {
		panic(fmt.Errorf("加载日志文件失败: %w", err))
	}
	return conf
}
