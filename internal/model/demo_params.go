package model

import "encoding/json"

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
// DemoParams [...]
type DemoParams struct {
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

func (w *DemoParams) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, w)
}

func (w *DemoParams) MarshalBinary() (data []byte, err error) {
	return json.Marshal(w)
}

// TableName get sql table name.获取数据库表名
func (w *DemoParams) TableName() string {
	return "demo_params"
}
