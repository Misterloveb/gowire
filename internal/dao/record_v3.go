package dao

import (
	"github.com/Misterloveb/gowire/internal/model"
	"gorm.io/gorm"
)

type WorkRecordV3Dao struct {
	*Dao
}

func NewWorkRecordV3Dao(dao *Dao) *WorkRecordV3Dao {
	return &WorkRecordV3Dao{
		Dao: dao,
	}
}

func (w *WorkRecordV3Dao) Insert(dbobj *gorm.DB, wdata []*model.RecordV3) error {
	thisdb := w.Db
	if dbobj != nil {
		thisdb = dbobj
	}
	if err := thisdb.Create(wdata); err.Error != nil {
		return err.Error
	}
	return nil
}
func (w *WorkRecordV3Dao) GetDataByWhere(col string, query any, arg ...any) []*model.RecordV3 {
	res := make([]*model.RecordV3, 20)
	w.Db.Select(col).Where(query, arg...).Find(&res)
	return res
}
func (w *WorkRecordV3Dao) Delete(dbobj *gorm.DB, query any, arg ...any) {
	tmpdb := w.Db
	if dbobj != nil {
		tmpdb = dbobj
	}
	tmpdb.Where(query, arg...).Delete(&model.RecordV3{})
}
