package log

import (
	"github.com/golang-module/carbon/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
)

type Logger struct {
	*logrus.Logger
}

func NewLogger(conf *viper.Viper, writer io.Writer) *Logger {
	logger := logrus.New()
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: carbon.DateTimeMilliLayout,
	})
	logger.SetLevel(logrus.Level(conf.GetUint32("serverlog.level")))
	logger.SetOutput(writer)
	return &Logger{logger}
}
