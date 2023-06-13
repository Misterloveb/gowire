package dao

import "gindemo/internal/model"

type WorkLogDao struct {
	*Dao
}

func NewWorkLogDao(dao *Dao) *WorkLogDao {
	return &WorkLogDao{
		Dao: dao,
	}
}

func (w *WorkLogDao) Insert(data []*model.WorkLog) error {
	if res := w.db.Create(data); res.Error != nil {
		return res.Error
	}
	return nil
}
func (w *WorkLogDao) GetData(offset, limit int) []*model.WorkLog {
	res := make([]*model.WorkLog, 0, 32)
	w.db.Order("`addtime` desc").Offset(offset).Limit(limit).Find(&res)
	return res
}
func (w *WorkLogDao) Count() int64 {
	var count int64
	w.db.Model(&model.WorkLog{}).Select("COUNT(*) AS count").Count(&count)
	return count
}
func (w *WorkLogDao) Delete(query any, args ...any) error {
	if res := w.db.Where(query, args...).Delete(&model.WorkLog{}); res.Error != nil {
		return res.Error
	}
	return nil
}
