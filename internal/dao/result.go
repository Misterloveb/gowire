package dao

import "gindemo/internal/model"

type WorkResultDao struct {
	*Dao
}

func NewWorkResultDao(dao *Dao) *WorkResultDao {
	return &WorkResultDao{
		Dao: dao,
	}
}

func (w *WorkResultDao) GetData() []model.WorkResult {
	resdata := make([]model.WorkResult, 0, 30)
	w.Db.Find(&resdata)
	return resdata
}
func (w *WorkResultDao) Insert(data []*model.WorkResult) error {
	return w.Db.Create(data).Statement.Error
}
func (w *WorkResultDao) Update(data *model.WorkResult) error {
	return w.Db.Save(data).Statement.Error
}
func (w *WorkResultDao) Delete(data *model.WorkResult) error {
	return w.Db.Delete(data).Statement.Error
}
