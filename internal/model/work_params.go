package model

import (
	"log"
	"net/url"
	"strings"
)

/******sql******
CREATE TABLE `work_params` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `info` varchar(255) NOT NULL COMMENT '参数中文名',
  `name` varchar(20) NOT NULL COMMENT '参数英文名',
  `unit` varchar(20) DEFAULT NULL COMMENT '参数单位',
  `modtype` varchar(20) NOT NULL COMMENT '参数类型',
  `htmltype` varchar(20) NOT NULL COMMENT '参数值的类型',
  `length` int(11) NOT NULL COMMENT '参数值的长度',
  `order` int(11) NOT NULL DEFAULT '1' COMMENT '参数顺序',
  `childrens` varchar(500) DEFAULT NULL COMMENT '参数的值',
  `html_width` smallint(3) unsigned NOT NULL COMMENT '参数的宽度',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8
******sql******/
// WorkParams [...]
type WorkParams struct {
	ID        int    `gorm:"primaryKey;column:id" json:"-"`
	Info      string `gorm:"column:info" json:"info"`            // 参数中文名
	Name      string `gorm:"column:name" json:"name"`            // 参数英文名
	Unit      string `gorm:"column:unit" json:"unit"`            // 参数单位
	Type      string `gorm:"column:modtype" json:"modtype"`      // 参数类型
	HTMLtype  string `gorm:"column:htmltype" json:"htmltype"`    // 参数值的类型
	Length    int    `gorm:"column:length" json:"length"`        // 参数值的长度
	Order     int    `gorm:"column:order" json:"order"`          // 参数顺序
	Childrens string `gorm:"column:childrens" json:"childrens"`  // 参数的值
	HTMLWidth uint16 `gorm:"column:html_width" json:"htmlWidth"` // 参数的宽度
}

func (w *WorkParams) GetData() []*WorkParams {
	resdata := make([]*WorkParams, 0, 30)
	db.Order("`order`").Find(&resdata)
	return resdata
}
func (w *WorkParams) Update(data url.Values) error {
	str_buff := strings.Builder{}
	args_arr := make([]any, 0, len(data)*2)
	str_buff.Grow(1024)
	str_buff.WriteString("UPDATE ")
	str_buff.WriteString(w.TableName())
	str_buff.WriteString(" SET `childrens` = case `name` ")
	for name, childrens := range data {
		str_buff.WriteString(" WHEN ? THEN ?")
		args_arr = append(args_arr, name, childrens[0])
	}
	str_buff.WriteString(" END")
	res := db.Exec(str_buff.String(), args_arr...)
	if res.Error != nil {
		log.Println(res.Error.Error())
		return res.Error
	}
	return nil
}

// TableName get sql table name.获取数据库表名
func (w *WorkParams) TableName() string {
	return "work_params"
}
