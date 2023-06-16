package provider

import (
	"github.com/Misterloveb/gowire/internal/common"
	"github.com/Misterloveb/gowire/internal/dao"
	"github.com/Misterloveb/gowire/internal/server"
	"github.com/Misterloveb/gowire/pkg/config"
	"github.com/Misterloveb/gowire/pkg/http"
	"github.com/Misterloveb/gowire/pkg/log"
	"github.com/Misterloveb/gowire/pkg/writer"
	"github.com/google/wire"
)

var Config = wire.NewSet(config.NewViper)

var (
	daoStrcut = wire.Struct(new(dao.Dao), "Db")
	NewDao    = wire.NewSet(daoStrcut, NewDB, NewRedis)
	NewDB     = wire.NewSet(dao.NewDB)
	NewRedis  = wire.NewSet(dao.NewRedis)
)
var (
	DataResultDao = wire.NewSet(dao.NewWorkDataResultDao)
	DatasV3Dao    = wire.NewSet(dao.NewWorkDatasV3Dao)
	LogDao        = wire.NewSet(dao.NewWorkLogDao)
	ParamsDao     = wire.NewSet(dao.NewWorkParamsDao)
	RecordV3Dao   = wire.NewSet(dao.NewWorkRecordV3Dao)
	ResultDao     = wire.NewSet(dao.NewWorkResultDao)
	UserDao       = wire.NewSet(dao.NewUserDao)
	CommonDao     = wire.NewSet(
		NewDao,
		DataResultDao,
		DatasV3Dao,
		LogDao,
		ParamsDao,
		RecordV3Dao,
		ResultDao,
		UserDao,
	)
)

var (
	CtlStuct  = wire.Struct(new(server.RegisterController), "*")
	GinServer = wire.NewSet(server.NewGinEngine, CtlStuct)
)

var (
	serverLogKey = wire.Value("serverlog")
	ServerWriter = wire.NewSet(writer.NewWriter, serverLogKey)
	ServerLog    = wire.NewSet(log.NewLogger, ServerWriter)
)
var (
	CommonCtl    = wire.Struct(new(common.CommonDao), "*")
	CommonServer = wire.NewSet(common.NewServer, ServerLog, CommonCtl)
)

var HTTPServer = wire.NewSet(http.NewHTTP)
