package dao

import (
	"github.com/Misterloveb/gowire/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WorkDatasV3Dao struct {
	*Dao
}

func NewWorkDatasV3Dao(dao *Dao) *WorkDatasV3Dao {
	return &WorkDatasV3Dao{
		Dao: dao,
	}
}

func (w *WorkDatasV3Dao) GetData() []*model.DatasV3 {
	res := make([]*model.DatasV3, 0, 30)
	w.Db.Find(&res)
	return res
}

func (w *WorkDatasV3Dao) GetCount(data *model.DatasV3) int64 {
	var count int64
	w.Db.Model(data).Select("COUNT(*) AS count").Count(&count)
	return count
}
func (w *WorkDatasV3Dao) GetDataByWhere(data *model.DatasV3, offset, limit int) []*model.DatasV3 {
	res := make([]*model.DatasV3, 0, 30)
	w.Db.Where(data).Limit(limit).Offset(offset).Find(&res)
	return res
}

func (w *WorkDatasV3Dao) Delete(dbobj *gorm.DB, query any, arg ...any) {
	tmpdb := w.Db
	if dbobj != nil {
		tmpdb = dbobj
	}
	tmpdb.Where(query, arg...).Delete(w)
}
func (w *WorkDatasV3Dao) Insert(dbobj *gorm.DB, data []*model.DatasV3) error {
	thisdb := w.Db
	if dbobj != nil {
		thisdb = dbobj
	}
	res := thisdb.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "kid"}},
		DoUpdates: clause.AssignmentColumns([]string{"insert_type", "insert_time"}),
	}).Create(data)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
