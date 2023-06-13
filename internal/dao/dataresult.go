package dao

import (
	"gindemo/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WorkDataResultDao struct {
	*Dao
}

func NewWorkDataResultDao(dao *Dao) *WorkDataResultDao {
	return &WorkDataResultDao{
		Dao: dao,
	}
}

func (w *WorkDataResultDao) Insert(data []*model.WorkDataresult) error {
	res := w.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "pkid"}, {Name: "result_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"result_id", "result_type", "result_info", "result_value"}),
	}).Create(data)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
func (w *WorkDataResultDao) Delete(dbobj *gorm.DB, query any, arg ...any) {
	tmpdb := w.db
	if dbobj != nil {
		tmpdb = dbobj
	}
	tmpdb.Where(query, arg...).Delete(&model.WorkDataresult{})
}
func (w *WorkDataResultDao) GetData() []model.WorkDataresult {
	resdata := make([]model.WorkDataresult, 0, 30)
	w.db.Find(&resdata)
	return resdata
}
func (w *WorkDataResultDao) GetDataByWhere(pkid, result_id string) []model.WorkDataresult {
	resdata := make([]model.WorkDataresult, 0, 30)
	w.db.Where("`pkid` = ? AND `result_id` = ?", pkid, result_id).Last(&resdata)
	return resdata
}
