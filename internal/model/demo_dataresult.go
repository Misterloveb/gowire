package model

/******sql******
CREATE TABLE `work_dataresult` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `pkid` char(64) NOT NULL COMMENT '所属参数表kid',
  `result_id` int(11) NOT NULL COMMENT '所属结果表id',
  `result_type` tinyint(2) unsigned NOT NULL COMMENT '结果类型',
  `result_info` varchar(255) NOT NULL COMMENT '结果名称',
  `result_value` varchar(500) NOT NULL DEFAULT '' COMMENT '结果值',
  PRIMARY KEY (`id`),
  UNIQUE KEY `kidandresid` (`pkid`,`result_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8
******sql******/
// Dataresult [...]
type Dataresult struct {
	ID          int    `gorm:"primaryKey;column:id" json:"-"`
	Pkid        string `gorm:"column:pkid" json:"pkid"`                 // 所属参数表kid
	ResultID    string `gorm:"column:result_id" json:"result_id"`       // 所属结果表id
	ResultType  uint8  `gorm:"column:result_type" json:"result_type"`   // 结果类型
	ResultInfo  string `gorm:"column:result_info" json:"result_info"`   // 结果名称
	ResultValue string `gorm:"column:result_value" json:"result_value"` // 结果值
}

// TableName get sql table name.获取数据库表名
func (w *Dataresult) TableName() string {
	return "demo_dataresult"
}
