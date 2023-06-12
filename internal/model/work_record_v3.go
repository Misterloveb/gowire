package model

import "gorm.io/gorm"

/******sql******
CREATE TABLE `work_record_v3` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `pkid` char(64) NOT NULL COMMENT '数据表kid',
  `result_id` int(11) unsigned NOT NULL COMMENT '结果类型id',
  `filesize` int(11) NOT NULL COMMENT '附件大小，单位kb',
  `filetype` varchar(10) NOT NULL COMMENT '附件扩展名',
  `filename` varchar(255) NOT NULL COMMENT '附件名称',
  `filepath` varchar(255) NOT NULL COMMENT '附件地址',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '附件状态',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `kidfilename` (`pkid`,`result_id`,`filename`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC
******sql******/
// WorkRecordV3 [...]
type WorkRecordV3 struct {
	ID       int    `gorm:"primaryKey;column:id" json:"-"`
	Pkid     string `gorm:"column:pkid" json:"pkid"`               // 数据表kid
	ResultID int    `gorm:"column:result_id" json:"resultId"`      // 结果类型id
	Filesize int    `gorm:"column:filesize" json:"filesize"`       // 附件大小，单位kb
	Filetype string `gorm:"column:filetype" json:"filetype"`       // 附件扩展名
	Filename string `gorm:"column:filename" json:"filename"`       // 附件名称
	Filepath string `gorm:"column:filepath" json:"filepath"`       // 附件地址
	Status   bool   `gorm:"column:status;default:1" json:"status"` // 附件状态
}

func (w *WorkRecordV3) Insert(dbobj *gorm.DB, wdata ...*WorkRecordV3) error {
	thisdb := db
	if dbobj != nil {
		thisdb = dbobj
	}
	lennum := 1
	if n := len(wdata); n != 0 {
		lennum = n
	}
	insert_data := make([]*WorkRecordV3, lennum)
	if lennum > 1 {
		for k, v := range wdata {
			insert_data[k] = v
		}
	} else {
		insert_data[0] = w
	}
	if err := thisdb.Create(insert_data); err.Error != nil {
		return err.Error
	}
	return nil
}
func (w *WorkRecordV3) GetDataByWhere(col string, query any, arg ...any) []*WorkRecordV3 {
	res := make([]*WorkRecordV3, 20)
	db.Select(col).Where(query, arg...).Find(&res)
	return res
}
func (w *WorkRecordV3) Delete(dbobj *gorm.DB, query any, arg ...any) {
	tmpdb := db
	if dbobj != nil {
		tmpdb = dbobj
	}
	tmpdb.Where(query, arg...).Delete(w)
}

// TableName get sql table name.获取数据库表名
func (w *WorkRecordV3) TableName() string {
	return "work_record_v3"
}
