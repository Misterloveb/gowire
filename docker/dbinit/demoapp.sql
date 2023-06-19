SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
CREATE DATABASE IF NOT EXISTS `demoapp` CHARACTER SET 'utf8' COLLATE 'utf8_general_ci';
USE `demoapp`;
-- ----------------------------
-- Table structure for demo_dataresult
-- ----------------------------
DROP TABLE IF EXISTS `demo_dataresult`;
CREATE TABLE `demo_dataresult`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `pkid` char(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '所属参数表kid',
  `result_id` int(11) NOT NULL COMMENT '所属结果表id',
  `result_type` tinyint(2) UNSIGNED NOT NULL COMMENT '结果类型',
  `result_info` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '结果名称',
  `result_value` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '结果值',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `kidandresid`(`pkid`, `result_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of demo_dataresult
-- ----------------------------

-- ----------------------------
-- Table structure for demo_datas_v3
-- ----------------------------
DROP TABLE IF EXISTS `demo_datas_v3`;
CREATE TABLE `demo_datas_v3`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `kid` char(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `zhuansu` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '速度',
  `qingjiao` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '角度',
  `hxfxll` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '容量1',
  `zsfxll` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '容量2',
  `plfxll` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '容量3',
  `xlwd` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '环境温度',
  `ltkcxl` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '容量4',
  `qtcxl` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '容量5',
  `cxwd` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '平均温度',
  `lcckbz` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '规格',
  `qwsrgl` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '功率',
  `insert_type` tinyint(1) NOT NULL COMMENT '1 批量录入 2 手动录入',
  `insert_time` datetime NOT NULL COMMENT '录入时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `datakid`(`kid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of demo_datas_v3
-- ----------------------------

-- ----------------------------
-- Table structure for demo_log
-- ----------------------------
DROP TABLE IF EXISTS `demo_log`;
CREATE TABLE `demo_log`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `kid` char(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `zhuansu` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '转速',
  `qingjiao` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '倾角',
  `hxfxll` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '滑靴副泄漏量',
  `zsfxll` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '柱塞副泄漏量',
  `plfxll` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '配流副泄漏量',
  `xlwd` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '泄漏温度',
  `ltkcxl` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '连体快冲洗量',
  `qtcxl` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '壳体快冲洗量',
  `cxwd` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '冲洗温度',
  `lcckbz` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '流场出口布置',
  `qwsrgl` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '球碗生热功率',
  `result_id` int(11) NOT NULL COMMENT '结果类型id',
  `result_type` tinyint(2) UNSIGNED NOT NULL COMMENT '结果类型',
  `result_info` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '结果类型名称',
  `result_value` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '结果类型的值',
  `insert_type` tinyint(1) NOT NULL COMMENT '录入方式 1 批量录入 2 手动录入',
  `addtime` int(11) NOT NULL COMMENT '添加时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of demo_log
-- ----------------------------

-- ----------------------------
-- Table structure for demo_params
-- ----------------------------
DROP TABLE IF EXISTS `demo_params`;
CREATE TABLE `demo_params`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `info` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '参数中文名',
  `name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '参数英文名',
  `unit` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '参数单位',
  `type` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '参数类型',
  `htmltype` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '参数值的类型',
  `length` int(11) NOT NULL COMMENT '参数值的长度',
  `order` int(11) NOT NULL DEFAULT 1 COMMENT '参数顺序',
  `childrens` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '参数的值',
  `html_width` smallint(3) UNSIGNED NOT NULL COMMENT '参数的宽度',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of demo_params
-- ----------------------------
INSERT INTO `demo_params` VALUES (1, '速度', 'zhuansu', 'rpm', 'varchar', 'select', 20, 1, '1000|1500|2000|3000|3500|3850|4000|5000|6000', 110);
INSERT INTO `demo_params` VALUES (2, '角度', 'qingjiao', '°', 'varchar', 'select', 20, 2, '0|17.5', 90);
INSERT INTO `demo_params` VALUES (3, '容量1', 'hxfxll', 'L/min', 'varchar', 'select', 20, 3, '1|4|8|9|18|27|31.5|34.65|36|45|54', 180);
INSERT INTO `demo_params` VALUES (4, '容量2', 'zsfxll', 'L/min', 'varchar', 'select', 20, 4, '0.5|2|4|4.5|9|13.5|15.75|17.325|18|22.5|27', 180);
INSERT INTO `demo_params` VALUES (5, '容量3', 'plfxll', 'L/min', 'varchar', 'select', 20, 5, '1|4|8|9|18|27|31.5|34.65|36|45|54', 180);
INSERT INTO `demo_params` VALUES (6, '环境温度', 'xlwd', '℃', 'varchar', 'select', 20, 6, '60|100|130', 150);
INSERT INTO `demo_params` VALUES (7, '容量4', 'ltkcxl', 'L/min', 'varchar', 'select', 20, 7, '0|5|10|20', 180);
INSERT INTO `demo_params` VALUES (8, '容量5', 'qtcxl', 'L/min', 'varchar', 'select', 20, 8, '0|5|10|20', 180);
INSERT INTO `demo_params` VALUES (9, '平均温度', 'cxwd', '℃', 'varchar', 'select', 20, 9, '20|80|90', 150);
INSERT INTO `demo_params` VALUES (10, '规格', 'lcckbz', '', 'varchar', 'select', 20, 10, '(a)|(b)|(c)|(d)|(e)', 150);
INSERT INTO `demo_params` VALUES (11, '功率', 'qwsrgl', 'W', 'varchar', 'select', 20, 11, '0|10|15|20|50|90', 160);

-- ----------------------------
-- Table structure for demo_record_v3
-- ----------------------------
DROP TABLE IF EXISTS `demo_record_v3`;
CREATE TABLE `demo_record_v3`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `pkid` char(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '数据表kid',
  `result_id` int(11) UNSIGNED NOT NULL COMMENT '结果类型id',
  `filesize` int(11) NOT NULL COMMENT '附件大小，单位kb',
  `filetype` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '附件扩展名',
  `filename` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '附件名称',
  `filepath` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '附件地址',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '附件状态',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `kidfilename`(`pkid`, `result_id`, `filename`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of demo_record_v3
-- ----------------------------

-- ----------------------------
-- Table structure for demo_result
-- ----------------------------
DROP TABLE IF EXISTS `demo_result`;
CREATE TABLE `demo_result`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `info` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '结果类型中文名',
  `type` tinyint(2) NOT NULL COMMENT '1:图片;2:数值;3:文件夹地址',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of demo_result
-- ----------------------------
INSERT INTO `demo_result` VALUES (1, '热成像分布图', 1);
INSERT INTO `demo_result` VALUES (2, '温度走势分布图', 1);
INSERT INTO `demo_result` VALUES (3, '天气走势分布图', 1);
INSERT INTO `demo_result` VALUES (4, '地势走势分布图', 1);
INSERT INTO `demo_result` VALUES (5, '流沙分布图', 1);
INSERT INTO `demo_result` VALUES (6, '地壳分布图', 1);
INSERT INTO `demo_result` VALUES (7, '湿地分布图', 1);
INSERT INTO `demo_result` VALUES (8, '海洋分布图', 1);
INSERT INTO `demo_result` VALUES (9, '液面高度', 2);
INSERT INTO `demo_result` VALUES (10, '温度最高', 2);
INSERT INTO `demo_result` VALUES (11, '温度最低', 2);
INSERT INTO `demo_result` VALUES (12, '详情文件', 3);
INSERT INTO `demo_result` VALUES (13, '实例名称', 2);
INSERT INTO `demo_result` VALUES (14, '实例说明', 2);

-- ----------------------------
-- Table structure for demo_user
-- ----------------------------
DROP TABLE IF EXISTS `demo_user`;
CREATE TABLE `demo_user`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `username` varchar(60) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `password` char(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `salt` char(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username`) USING BTREE,
  INDEX `idx_work_user_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of demo_user
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
