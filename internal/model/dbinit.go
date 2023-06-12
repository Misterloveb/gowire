package model

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var once sync.Once

func stdoutloger() logger.Interface {
	return logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		LogLevel: logger.Info,
	})
}
func DB() *gorm.DB {
	if db == nil {
		once.Do(func() {
			dsn := fmt.Sprintf(
				"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
				viper.GetString("database.user"),
				viper.GetString("database.pwd"),
				viper.GetString("database.ip"),
				viper.GetUint("database.port"),
				viper.GetString("database.db"),
				viper.GetString("database.charset"),
			)
			dbobj, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
				Logger: stdoutloger(),
			})
			if err != nil {
				panic("数据库启动失败！")
			}
			sqldb, err := dbobj.DB()
			if err != nil {
				panic("数据库获取失败！")
			}
			//最大空闲连接时间
			sqldb.SetConnMaxIdleTime(time.Second * viper.GetDuration("database.MaxIdleTime"))
			//空闲连接池最大数量
			sqldb.SetMaxIdleConns(viper.GetInt("database.MaxIdleConns"))
			//最大打开的连接数
			sqldb.SetMaxOpenConns(viper.GetInt("database.MaxOpenConns"))
			//连接可复用的最长时间
			sqldb.SetConnMaxLifetime(time.Hour * viper.GetDuration("database.MaxLifetime"))
			db = dbobj
		})
	}
	return db
}
