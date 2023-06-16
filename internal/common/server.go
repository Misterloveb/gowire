package common

import (
	"github.com/Misterloveb/gowire/internal/dao"
	"github.com/Misterloveb/gowire/pkg/log"
	"github.com/spf13/viper"
)

type Server struct {
	*viper.Viper
	*log.Logger
	*CommonDao
}
type CommonDao struct {
	*dao.WorkParamsDao
	*dao.WorkResultDao
	*dao.WorkDatasV3Dao
	*dao.WorkDataResultDao
	*dao.WorkRecordV3Dao
	*dao.UserDao
	*dao.WorkLogDao
}

func NewServer(log *log.Logger, ctl *CommonDao, conf *viper.Viper) *Server {
	return &Server{
		Logger:    log,
		CommonDao: ctl,
		Viper:     conf,
	}
}
