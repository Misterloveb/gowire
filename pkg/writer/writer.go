package writer

import (
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Writer struct {
	*lumberjack.Logger
}

func NewWriter(conf *viper.Viper, key string) *Writer {
	hook := &lumberjack.Logger{
		Filename:   conf.GetString(key + ".filepath"), // 日志文件路径
		MaxSize:    conf.GetInt(key + ".max_size"),    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: conf.GetInt(key + ".max_backups"), // 日志文件最多保存多少个备份
		MaxAge:     conf.GetInt(key + ".max_age"),     // 文件最多保存多少天
		Compress:   conf.GetBool(key + ".compress"),   // 是否压缩
	}
	return &Writer{
		Logger: hook,
	}
}
