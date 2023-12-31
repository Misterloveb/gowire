package model

import "encoding/json"

/******sql******
CREATE TABLE `work_result` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `info` varchar(255) NOT NULL COMMENT '结果类型中文名',
  `modtype` tinyint(2) NOT NULL COMMENT '1:图片;2:数值;3:文件夹地址',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8
******sql******/
// DemoResult [...]
type DemoResult struct {
	ID   int    `gorm:"primaryKey;column:id" json:"-" form:"id"`
	Info string `gorm:"column:info" json:"info" form:"info"`    // 结果类型中文名
	Type int8   `gorm:"column:type" json:"modtype" form:"type"` // 1:图片;2:数值;3:文件夹地址
}

// TableName get sql table name.获取数据库表名
func (w *DemoResult) TableName() string {
	return "demo_result"
}
func (w *DemoResult) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, w)
}

func (w *DemoResult) MarshalBinary() (data []byte, err error) {
	return json.Marshal(w)
}
