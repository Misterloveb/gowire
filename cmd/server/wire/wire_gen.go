// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/Misterloveb/gowire/internal/common"
	"github.com/Misterloveb/gowire/internal/controller"
	"github.com/Misterloveb/gowire/internal/dao"
	"github.com/Misterloveb/gowire/internal/server"
	"github.com/Misterloveb/gowire/pkg/config"
	"github.com/Misterloveb/gowire/pkg/http"
	"github.com/Misterloveb/gowire/pkg/log"
	"github.com/Misterloveb/gowire/pkg/writer"
)

// Injectors from wire.go:

func App() *http.HTTP {
	viper := config.NewViper()
	string2 := _wireStringValue
	writerWriter := writer.NewWriter(viper, string2)
	logger := log.NewLogger(viper, writerWriter)
	db := dao.NewDB(viper)
	client := dao.NewRedis(viper)
	daoDao := &dao.Dao{
		Db:  db,
		Rdb: client,
	}
	workParamsDao := dao.NewWorkParamsDao(daoDao)
	workResultDao := dao.NewWorkResultDao(daoDao)
	workDatasV3Dao := dao.NewWorkDatasV3Dao(daoDao)
	workDataResultDao := dao.NewWorkDataResultDao(daoDao)
	workRecordV3Dao := dao.NewWorkRecordV3Dao(daoDao)
	userDao := dao.NewUserDao(daoDao)
	workLogDao := dao.NewWorkLogDao(daoDao)
	commonDao := &common.CommonDao{
		WorkParamsDao:     workParamsDao,
		WorkResultDao:     workResultDao,
		WorkDatasV3Dao:    workDatasV3Dao,
		WorkDataResultDao: workDataResultDao,
		WorkRecordV3Dao:   workRecordV3Dao,
		UserDao:           userDao,
		WorkLogDao:        workLogDao,
	}
	commonServer := common.NewServer(logger, commonDao, viper)
	loginController := &controller.LoginController{
		Server: commonServer,
	}
	logerController := &controller.LogerController{
		Server: commonServer,
	}
	indexController := &controller.IndexController{
		Server: commonServer,
	}
	systemController := &controller.SystemController{
		Server: commonServer,
	}
	registerController := &controller.RegisterController{
		Login:  loginController,
		Loger:  logerController,
		Index:  indexController,
		System: systemController,
	}
	engine := server.NewGinEngine(viper, registerController)
	httpHTTP := http.NewHTTP(engine, viper)
	return httpHTTP
}

var (
	_wireStringValue = "serverlog"
)
