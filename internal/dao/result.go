package dao

import (
	"context"
	"encoding/json"
	"github.com/Misterloveb/gowire/internal/model"
	"log"
	"time"
)

type WorkResultDao struct {
	*Dao
}

func NewWorkResultDao(dao *Dao) *WorkResultDao {
	return &WorkResultDao{
		Dao: dao,
	}
}

func (w *WorkResultDao) GetData() []*model.DemoResult {
	resdata := make([]*model.DemoResult, 0, 30)
	if w.Rdb == nil {
		w.Db.Find(&resdata)
		return resdata
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stat := w.Rdb.Get(ctx, "wrs")
	if stat.Err() != nil {
		w.Db.Find(&resdata)
		jsRes, err := json.Marshal(resdata)
		if err != nil {
			log.Println(err.Error())
			return nil
		}
		w.Rdb.Set(ctx, "wrs", jsRes, 3600*time.Second)
	} else {
		json.Unmarshal([]byte(stat.Val()), &resdata)
	}
	return resdata
}
func (w *WorkResultDao) Insert(data []*model.DemoResult) error {
	return w.Db.Create(data).Statement.Error
}
func (w *WorkResultDao) Update(data *model.DemoResult) error {
	return w.Db.Save(data).Statement.Error
}
func (w *WorkResultDao) Delete(data *model.DemoResult) error {
	return w.Db.Delete(data).Statement.Error
}
