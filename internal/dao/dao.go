package dao

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/Misterloveb/gowire/pkg/writer"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

type Dao struct {
	Db  *gorm.DB
	Rdb *redis.Client
}

//go:embed dbinit/demoapp.sql
var initSql string

func NewDB(conf *viper.Viper) *gorm.DB {
	initDataBase(conf)
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		conf.GetString("mysql.user"),
		conf.GetString("mysql.pwd"),
		conf.GetString("mysql.ip"),
		conf.GetInt("mysql.port"),
		conf.GetString("mysql.db"),
		conf.GetString("mysql.charset"),
	)
	newLogger := logger.New(
		log.New(writer.NewWriter(conf, "dblog"), "\r\n", log.LstdFlags), // io newWriter
		logger.Config{
			SlowThreshold:             time.Second * conf.GetDuration("dblog.SlowThreshold"), // Slow SQL threshold
			LogLevel:                  logger.LogLevel(conf.GetInt("dblog.level")),           // Log level
			IgnoreRecordNotFoundError: true,                                                  // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      conf.GetBool("dblog.ParameterizedQueries"),            // Don't include params in the SQL log
			Colorful:                  false,                                                 // Disable color
		},
	)
	dbobj, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Println("ip=", conf.GetString("mysql.ip"))
		log.Println("port=", conf.GetString("mysql.port"))
		log.Fatal("数据库连接失败", err.Error())
	}
	sqldb, err := dbobj.DB()
	if err != nil {
		log.Fatal("数据库获取失败", err.Error())
	}
	//最大空闲连接时间
	sqldb.SetConnMaxIdleTime(time.Second * conf.GetDuration("mysql.MaxIdleTime"))
	//空闲连接池最大数量
	sqldb.SetMaxIdleConns(conf.GetInt("mysql.MaxIdleConns"))
	//最大打开的连接数
	sqldb.SetMaxOpenConns(conf.GetInt("mysql.MaxOpenConns"))
	//连接可复用的最长时间
	sqldb.SetConnMaxLifetime(time.Hour * conf.GetDuration("mysql.MaxLifetime"))
	return dbobj
}
func NewRedis(conf *viper.Viper) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.GetString("redis.ip") + ":" + conf.GetString("redis.port"),
		Password: conf.GetString("redis.password"),
		DB:       conf.GetInt("redis.db"),
	})
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Println(fmt.Sprintf("redis 连接失败: %s", err.Error()))
		return nil
	}
	return rdb
}
func initDataBase(conf *viper.Viper) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local&multiStatements=true",
		conf.GetString("mysql.user"),
		conf.GetString("mysql.pwd"),
		conf.GetString("mysql.ip"),
		conf.GetInt("mysql.port"),
		"",
		conf.GetString("mysql.charset"),
	)
	dbname := conf.GetString("mysql.db")
	dbobj, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败：", err.Error())
	}
	res, err := dbobj.Raw("show databases like '" + dbname + "'").Rows()
	if !res.Next() {
		//不存在数据库则创建
		sta := dbobj.Exec(initSql).Statement
		if sta.Error != nil {
			log.Fatal("数据库创建失败！", sta.Error.Error())
		}
	}
}
