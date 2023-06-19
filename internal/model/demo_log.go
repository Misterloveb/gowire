package model

/******sql******
CREATE TABLE `work_log` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `kid` char(64) NOT NULL,
  `zhuansu` varchar(20) DEFAULT NULL COMMENT '转速',
  `qingjiao` varchar(20) DEFAULT NULL COMMENT '倾角',
  `hxfxll` varchar(20) DEFAULT NULL COMMENT '滑靴副泄漏量',
  `zsfxll` varchar(20) DEFAULT NULL COMMENT '柱塞副泄漏量',
  `plfxll` varchar(20) DEFAULT NULL COMMENT '配流副泄漏量',
  `xlwd` varchar(20) DEFAULT NULL COMMENT '泄漏温度',
  `ltkcxl` varchar(20) DEFAULT NULL COMMENT '连体快冲洗量',
  `qtcxl` varchar(20) DEFAULT NULL COMMENT '壳体快冲洗量',
  `cxwd` varchar(20) DEFAULT NULL COMMENT '冲洗温度',
  `lcckbz` varchar(20) DEFAULT NULL COMMENT '流场出口布置',
  `qwsrgl` varchar(20) DEFAULT NULL COMMENT '球碗生热功率',
  `result_id` int(11) NOT NULL COMMENT '结果类型id',
  `result_type` tinyint(2) unsigned NOT NULL COMMENT '结果类型',
  `result_info` varchar(255) NOT NULL COMMENT '结果类型名称',
  `result_value` varchar(500) NOT NULL COMMENT '结果类型的值',
  `insert_type` tinyint(1) NOT NULL COMMENT '录入方式 1 批量录入 2 手动录入',
  `addtime` int(11) NOT NULL COMMENT '添加时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC
******sql******/
// DemoLog [...]
type DemoLog struct {
	ID          int    `gorm:"primaryKey;column:id" json:"id"`
	Kid         string `gorm:"column:kid" json:"kid"`
	Zhuansu     string `gorm:"column:zhuansu" json:"zhuansu"`           // 转速
	Qingjiao    string `gorm:"column:qingjiao" json:"qingjiao"`         // 倾角
	Hxfxll      string `gorm:"column:hxfxll" json:"hxfxll"`             // 滑靴副泄漏量
	Zsfxll      string `gorm:"column:zsfxll" json:"zsfxll"`             // 柱塞副泄漏量
	Plfxll      string `gorm:"column:plfxll" json:"plfxll"`             // 配流副泄漏量
	Xlwd        string `gorm:"column:xlwd" json:"xlwd"`                 // 泄漏温度
	Ltkcxl      string `gorm:"column:ltkcxl" json:"ltkcxl"`             // 连体快冲洗量
	Qtcxl       string `gorm:"column:qtcxl" json:"qtcxl"`               // 壳体快冲洗量
	Cxwd        string `gorm:"column:cxwd" json:"cxwd"`                 // 冲洗温度
	Lcckbz      string `gorm:"column:lcckbz" json:"lcckbz"`             // 流场出口布置
	Qwsrgl      string `gorm:"column:qwsrgl" json:"qwsrgl"`             // 球碗生热功率
	ResultID    int    `gorm:"column:result_id" json:"result_id"`       // 结果类型id
	ResultType  uint8  `gorm:"column:result_type" json:"result_type"`   // 结果类型
	ResultInfo  string `gorm:"column:result_info" json:"result_info"`   // 结果类型名称
	ResultValue string `gorm:"column:result_value" json:"result_value"` // 结果类型的值
	InsertType  int8   `gorm:"column:insert_type" json:"insert_type"`   // 录入方式 1 批量录入 2 手动录入
	Addtime     int    `gorm:"column:addtime" json:"addtime"`           // 添加时间
}

// TableName get sql table name.获取数据库表名
func (w *DemoLog) TableName() string {
	return "demo_log"
}
