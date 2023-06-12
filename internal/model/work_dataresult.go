package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

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
// WorkDataresult [...]
type WorkDataresult struct {
	ID          int    `gorm:"primaryKey;column:id" json:"-"`
	Pkid        string `gorm:"column:pkid" json:"pkid"`                 // 所属参数表kid
	ResultID    string `gorm:"column:result_id" json:"result_id"`       // 所属结果表id
	ResultType  uint8  `gorm:"column:result_type" json:"result_type"`   // 结果类型
	ResultInfo  string `gorm:"column:result_info" json:"result_info"`   // 结果名称
	ResultValue string `gorm:"column:result_value" json:"result_value"` // 结果值
}

func (w *WorkDataresult) Insert(dbobj *gorm.DB, wdata ...*WorkDataresult) error {
	thisdb := db
	if dbobj != nil {
		thisdb = dbobj
	}
	lennum := 1
	if n := len(wdata); n != 0 {
		lennum = n
	}
	insert_data := make([]*WorkDataresult, lennum)
	if lennum > 1 {
		for k, v := range wdata {
			insert_data[k] = v
		}
	} else {
		insert_data[0] = w
	}

	res := thisdb.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "pkid"}, {Name: "result_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"result_id", "result_type", "result_info", "result_value"}),
	}).Create(insert_data)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
func (w *WorkDataresult) Delete(dbobj *gorm.DB, query any, arg ...any) {
	tmpdb := db
	if dbobj != nil {
		tmpdb = dbobj
	}
	tmpdb.Where(query, arg...).Delete(w)
}
func (w *WorkDataresult) GetData() []WorkDataresult {
	resdata := make([]WorkDataresult, 0, 30)
	db.Find(&resdata)
	return resdata
}
func (w *WorkDataresult) GetDataByWhere(pkid, result_id string) []WorkDataresult {
	resdata := make([]WorkDataresult, 0, 30)
	db.Where("`pkid` = ? AND `result_id` = ?", pkid, result_id).Last(&resdata)
	return resdata
}

// TableName get sql table name.获取数据库表名
func (w *WorkDataresult) TableName() string {
	return "work_dataresult"
}
