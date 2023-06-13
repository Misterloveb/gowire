package dao

import (
	"gindemo/internal/model"
	"log"
	"net/url"
	"strings"
)

type WorkParamsDao struct {
	*Dao
}

func NewWorkParamsDao(dao *Dao) *WorkParamsDao {
	return &WorkParamsDao{
		Dao: dao,
	}
}

func (w *WorkParamsDao) GetData() []*model.WorkParams {
	resdata := make([]*model.WorkParams, 0, 30)
	w.db.Order("`order`").Find(&resdata)
	return resdata
}
func (w *WorkParamsDao) Update(data url.Values) error {
	dst := &model.WorkParams{}
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
	res := w.db.Exec(str_buff.String(), args_arr...)
	if res.Error != nil {
		log.Println(res.Error.Error())
		return res.Error
	}
	return nil
}
