package dao

import (
	"gindemo/internal/model"
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

func (w *WorkDatasV3Dao) GetData() []*model.WorkDatasV3 {
	res := make([]*model.WorkDatasV3, 0, 30)
	w.db.Find(&res)
	return res
}

func (w *WorkDatasV3Dao) GetCount() int64 {
	var count int64
	w.db.Model(w).Select("COUNT(*) AS count").Count(&count)
	return count
}
func (w *WorkDatasV3Dao) GetDataByWhere(offset, limit int) []*model.WorkDatasV3 {
	res := make([]*model.WorkDatasV3, 0, 30)
	w.db.Where(w).Limit(limit).Offset(offset).Find(&res)
	return res
}

func (w *WorkDatasV3Dao) Delete(dbobj *gorm.DB, query any, arg ...any) {
	tmpdb := w.db
	if dbobj != nil {
		tmpdb = dbobj
	}
	tmpdb.Where(query, arg...).Delete(w)
}
func (w *WorkDatasV3Dao) Insert(dbobj *gorm.DB, data []*model.WorkDatasV3) error {
	thisdb := w.db
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
