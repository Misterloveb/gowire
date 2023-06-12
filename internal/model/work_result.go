package model

/******sql******
CREATE TABLE `work_result` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `info` varchar(255) NOT NULL COMMENT '结果类型中文名',
  `modtype` tinyint(2) NOT NULL COMMENT '1:图片;2:数值;3:文件夹地址',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8
******sql******/
// WorkResult [...]
type WorkResult struct {
	ID   int    `gorm:"primaryKey;column:id" json:"-" form:"id"`
	Info string `gorm:"column:info" json:"info" form:"info"`    // 结果类型中文名
	Type int8   `gorm:"column:type" json:"modtype" form:"type"` // 1:图片;2:数值;3:文件夹地址
}

func (w *WorkResult) GetData() []WorkResult {
	resdata := make([]WorkResult, 0, 30)
	db.Find(&resdata)
	return resdata
}
func (w *WorkResult) Insert() error {
	stam := db.Create(w).Statement
	return stam.Error
}
func (w *WorkResult) Update() error {
	stam := db.Save(w)
	return stam.Error
}
func (w *WorkResult) Delete() error {
	stam := db.Delete(w)
	return stam.Error
}

// TableName get sql table name.获取数据库表名
func (w *WorkResult) TableName() string {
	return "work_result"
}
