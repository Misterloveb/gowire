package provider

import (
	"github.com/Misterloveb/gowire/internal/dao"
	"github.com/google/wire"
)

var (
	DataResultDao = wire.NewSet(dao.NewWorkDataResultDao)
	DatasV3Dao    = wire.NewSet(dao.NewWorkDatasV3Dao)
	LogDao        = wire.NewSet(dao.NewWorkLogDao)
	ParamsDao     = wire.NewSet(dao.NewWorkParamsDao)
	RecordV3Dao   = wire.NewSet(dao.NewWorkRecordV3Dao)
	ResultDao     = wire.NewSet(dao.NewWorkResultDao)
	UserDao       = wire.NewSet(dao.NewUserDao)
)

var CommonDao = wire.NewSet(
	NewDao,
	DataResultDao,
	DatasV3Dao,
	LogDao,
	ParamsDao,
	RecordV3Dao,
	ResultDao,
	UserDao,
)
