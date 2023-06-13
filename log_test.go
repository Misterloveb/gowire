package gin_demo

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func TestLog(t *testing.T) {
	logger := logrus.New()
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "15:04:05.000",
	})
	filename := time.Now().Format("2006-01-02") + ".log"
	filepath := filepath.Join("internal/log", filename)
	hook := lumberjack.Logger{
		Filename:   filepath, // 日志文件路径
		MaxSize:    2,        // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 4,        // 日志文件最多保存多少个备份
		MaxAge:     1,        // 文件最多保存多少天
		Compress:   false,    // 是否压缩
	}
	logger.SetOutput(&hook)

	logger.SetLevel(logrus.DebugLevel)
	logg := logger.WithFields(logrus.Fields{
		"userid":    "",
		"requestid": "",
	})
	logg.Debug("调试代码", "234324", "222")
	logg.Info("测试代码")
	logg.Warning("警告")
	logg.Error("错误")
	logg.Fatal("致命错误")

}
