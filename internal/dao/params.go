package dao

import (
	"context"
	"encoding/json"
	"github.com/Misterloveb/gowire/internal/model"
	"log"
	"net/url"
	"strings"
	"time"
)

type WorkParamsDao struct {
	*Dao
}

func NewWorkParamsDao(dao *Dao) *WorkParamsDao {
	return &WorkParamsDao{
		Dao: dao,
	}
}

func (w *WorkParamsDao) GetData() []*model.DemoParams {
	resdata := make([]*model.DemoParams, 0, 30)
	if w.Rdb == nil {
		w.Db.Order("`order`").Find(&resdata)
		return resdata
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stat := w.Rdb.Get(ctx, "wpm")
	if stat.Err() != nil {
		w.Db.Order("`order`").Find(&resdata)
		jsRes, err := json.Marshal(resdata)
		if err != nil {
			log.Println(err.Error())
			return nil
		}
		w.Rdb.Set(ctx, "wpm", jsRes, 3600*time.Second)
	} else {
		json.Unmarshal([]byte(stat.Val()), &resdata)
	}
	return resdata
}
func (w *WorkParamsDao) Update(data url.Values) error {
	dst := &model.DemoParams{}
	str_buff := strings.Builder{}
	args_arr := make([]any, 0, len(data)*2)
	str_buff.Grow(1024)
	str_buff.WriteString("UPDATE ")
	str_buff.WriteString(dst.TableName())
	str_buff.WriteString(" SET `childrens` = case `name` ")
	for name, childrens := range data {
		str_buff.WriteString(" WHEN ? THEN ?")
		args_arr = append(args_arr, name, childrens[0])
	}
	str_buff.WriteString(" END")
	res := w.Db.Exec(str_buff.String(), args_arr...)
	if res.Error != nil {
		log.Println(res.Error.Error())
		return res.Error
	}
	return nil
}
