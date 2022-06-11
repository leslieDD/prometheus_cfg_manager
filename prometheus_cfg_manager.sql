-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        10.7.3-MariaDB - mariadb.org binary distribution
-- 服务器操作系统:                      Win64
-- HeidiSQL 版本:                  11.3.0.6295
-- --------------------------------------------------------
/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */
;
/*!40101 SET NAMES utf8 */
;
/*!50503 SET NAMES utf8mb4 */
;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */
;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */
;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */
;
-- 导出 pro_cfg_manager 的数据库结构
DROP DATABASE IF EXISTS `pro_cfg_manager`;
CREATE DATABASE IF NOT EXISTS `pro_cfg_manager`
/*!40100 DEFAULT CHARACTER SET utf8mb3 */
;
USE `pro_cfg_manager`;
-- 导出  表 pro_cfg_manager.annotations 结构
DROP TABLE IF EXISTS `annotations`;
CREATE TABLE IF NOT EXISTS `annotations` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `monitor_rules_id` int(11) NOT NULL,
  `key` varchar(500) NOT NULL,
  `value` varchar(500) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_monitor_rules_id_key` (`monitor_rules_id`, `key`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb3 COMMENT = '监控规则的注释';
-- 正在导出表  pro_cfg_manager.annotations 的数据：~0 rows (大约)
DELETE FROM `annotations`;
/*!40000 ALTER TABLE `annotations` DISABLE KEYS */
;
/*!40000 ALTER TABLE `annotations` ENABLE KEYS */
;
-- 导出  表 pro_cfg_manager.group_labels 结构
DROP TABLE IF EXISTS `group_labels`;
CREATE TABLE IF NOT EXISTS `group_labels` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `job_group_id` int(11) NOT NULL,
  `key` varchar(100) NOT NULL,
  `value` varchar(100) NOT NULL,
  `enabled` tinyint(4) NOT NULL DEFAULT 1,
  `update_at` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `update_by` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_groupid_key` (`key`, `job_group_id`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb3 COMMENT = '子组标签列表';
-- 正在导出表  pro_cfg_manager.group_labels 的数据：~3 rows (大约)
DELETE FROM `group_labels`;
/*!40000 ALTER TABLE `group_labels` DISABLE KEYS */
;
/*!40000 ALTER TABLE `group_labels` ENABLE KEYS */
;
-- 导出  表 pro_cfg_manager.group_machines 结构
DROP TABLE IF EXISTS `group_machines`;
CREATE TABLE IF NOT EXISTS `group_machines` (
  `job_group_id` int(11) NOT NULL,
  `machines_id` int(11) NOT NULL,
  UNIQUE KEY `unique` (`machines_id`, `job_group_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb3 COMMENT = '子组IP列表';
-- 正在导出表  pro_cfg_manager.group_machines 的数据：~0 rows (大约)
DELETE FROM `group_machines`;
/*!40000 ALTER TABLE `group_machines` DISABLE KEYS */
;
/*!40000 ALTER TABLE `group_machines` ENABLE KEYS */
;
-- 导出  表 pro_cfg_manager.group_priv 结构
DROP TABLE IF EXISTS `group_priv`;
CREATE TABLE IF NOT EXISTS `group_priv` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `group_id` int(11) NOT NULL,
  `func_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 125 DEFAULT CHARSET = utf8mb3;
-- 正在导出表  pro_cfg_manager.group_priv 的数据：~124 rows (大约)
DELETE FROM `group_priv`;
/*!40000 ALTER TABLE `group_priv` DISABLE KEYS */
;
INSERT INTO `group_priv` (`id`, `group_id`, `func_id`)
VALUES (1, 1, 1),
  (2, 1, 2),
  (3, 1, 3),
  (4, 1, 4),
  (5, 1, 5),
  (6, 1, 6),
  (7, 1, 7),
  (8, 1, 8),
  (9, 1, 9),
  (10, 1, 11),
  (11, 1, 12),
  (12, 1, 13),
  (13, 1, 14),
  (14, 1, 15),
  (15, 1, 16),
  (16, 1, 19),
  (17, 1, 20),
  (18, 1, 21),
  (19, 1, 22),
  (20, 1, 23),
  (21, 1, 24),
  (22, 1, 25),
  (23, 1, 26),
  (24, 1, 27),
  (25, 1, 28),
  (26, 1, 29),
  (27, 1, 30),
  (28, 1, 31),
  (29, 1, 33),
  (30, 1, 34),
  (31, 1, 35),
  (32, 1, 36),
  (33, 1, 37),
  (34, 1, 38),
  (35, 1, 39),
  (36, 1, 40),
  (37, 1, 41),
  (38, 1, 42),
  (39, 1, 43),
  (40, 1, 44),
  (41, 1, 45),
  (42, 1, 46),
  (43, 1, 47),
  (44, 1, 48),
  (45, 1, 49),
  (46, 1, 50),
  (47, 1, 51),
  (48, 1, 52),
  (49, 1, 53),
  (50, 1, 54),
  (51, 1, 55),
  (52, 1, 56),
  (53, 1, 57),
  (54, 1, 58),
  (55, 1, 59),
  (56, 1, 60),
  (57, 1, 61),
  (58, 1, 62),
  (59, 1, 63),
  (60, 1, 64),
  (61, 1, 65),
  (62, 1, 66),
  (63, 1, 67),
  (64, 1, 68),
  (65, 1, 69),
  (66, 1, 70),
  (67, 1, 71),
  (68, 1, 72),
  (69, 1, 73),
  (70, 1, 74),
  (71, 1, 75),
  (72, 1, 76),
  (73, 1, 77),
  (74, 1, 78),
  (75, 1, 79),
  (76, 1, 80),
  (77, 1, 81),
  (78, 1, 83),
  (79, 1, 84),
  (80, 1, 85),
  (81, 1, 86),
  (82, 1, 87),
  (83, 1, 88),
  (84, 1, 89),
  (85, 1, 90),
  (86, 1, 93),
  (87, 1, 94),
  (88, 1, 95),
  (89, 1, 96),
  (90, 1, 97),
  (91, 1, 98),
  (92, 1, 99),
  (93, 1, 100),
  (94, 1, 101),
  (95, 1, 102),
  (96, 1, 103),
  (97, 1, 104),
  (98, 1, 105),
  (99, 1, 106),
  (100, 1, 107),
  (101, 1, 108),
  (102, 1, 109),
  (103, 1, 110),
  (104, 1, 111),
  (105, 1, 112),
  (106, 1, 113),
  (107, 1, 114),
  (108, 1, 115),
  (109, 1, 116),
  (110, 1, 117),
  (111, 1, 118),
  (112, 1, 119),
  (113, 1, 120),
  (114, 1, 121),
  (115, 1, 122),
  (116, 1, 123),
  (117, 1, 124),
  (118, 1, 125),
  (119, 1, 126),
  (120, 1, 127),
  (121, 1, 128),
  (122, 1, 129),
  (123, 1, 130),
  (124, 1, 131);
/*!40000 ALTER TABLE `group_priv` ENABLE KEYS */
;
-- 导出  表 pro_cfg_manager.idc 结构
DROP TABLE IF EXISTS `idc`;
CREATE TABLE IF NOT EXISTS `idc` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `label` varchar(200) NOT NULL,
  `remark` text DEFAULT NULL COMMENT '备注',
  `enabled` tinyint(4) NOT NULL DEFAULT 1,
  `update_at` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `update_by` varchar(200) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_label` (`label`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb3 COMMENT = '机房';
-- 正在导出表  pro_cfg_manager.idc 的数据：~0 rows (大约)
DELETE FROM `idc`;
/*!40000 ALTER TABLE `idc` DISABLE KEYS */
;
/*!40000 ALTER TABLE `idc` ENABLE KEYS */
;
-- 导出  表 pro_cfg_manager.jobs 结构
DROP TABLE IF EXISTS `jobs`;
CREATE TABLE IF NOT EXISTS `jobs` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(100) NOT NULL COMMENT '任务名称',
  `port` int(11) NOT NULL COMMENT '端口号，对应exporter的端口号',
  `cfg_name` varchar(100) DEFAULT NULL COMMENT '在prometheus下生成配置的文件名称',
  `is_common` tinyint(4) NOT NULL,
  `relabel_id` int(11) NOT NULL,
  `display_order` int(10) unsigned NOT NULL,
  `enabled` tinyint(4) NOT NULL DEFAULT 1,
  `update_at` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `update_by` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_unique` (`name`),
  KEY `order_unique` (`display_order`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb3 COMMENT = 'prometheus任务列表';
-- 正在导出表  pro_cfg_manager.jobs 的数据：~0 rows (大约)
DELETE FROM `jobs`;
/*!40000 ALTER TABLE `jobs` DISABLE KEYS */
;
/*!40000 ALTER TABLE `jobs` ENABLE KEYS */
;
-- 导出  表 pro_cfg_manager.job_group 结构
DROP TABLE IF EXISTS `job_group`;
CREATE TABLE IF NOT EXISTS `job_group` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `jobs_id` int(11) NOT NULL,
  `enabled` tinyint(4) NOT NULL DEFAULT 1,
  `update_at` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `update_by` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_name_jobs_id` (`jobs_id`, `name`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb3 COMMENT = '还可以为单个分组中的IP地址进行分子组，为每个子组设置相应的标签列表';
-- 正在导出表  pro_cfg_manager.job_group 的数据：~0 rows (大约)
DELETE FROM `job_group`;
/*!40000 ALTER TABLE `job_group` DISABLE KEYS */
;
/*!40000 ALTER TABLE `job_group` ENABLE KEYS */
;
-- 导出  表 pro_cfg_manager.job_machines 结构
DROP TABLE IF EXISTS `job_machines`;
CREATE TABLE IF NOT EXISTS `job_machines` (
  `machine_id` int(11) NOT NULL,
  `job_id` int(11) NOT NULL,
  `blacked` int(11) DEFAULT 0 COMMENT '是否禁用，在本组中',
  UNIQUE KEY `unique_mj` (`job_id`, `machine_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb3;
-- 正在导出表  pro_cfg_manager.job_machines 的数据：~0 rows (大约)
DELETE FROM `job_machines`;
/*!40000 ALTER TABLE `job_machines` DISABLE KEYS */
;
/*!40000 ALTER TABLE `job_machines` ENABLE KEYS */
;
-- 导出  表 pro_cfg_manager.labels 结构
DROP TABLE IF EXISTS `labels`;
CREATE TABLE IF NOT EXISTS `labels` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `label` varchar(100) NOT NULL,
  `enabled` tinyint(4) NOT NULL DEFAULT 1,
  `update_at` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `update_by` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_labels` (`label`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb3 COMMENT = '标签列表';
-- 正在导出表  pro_cfg_manager.labels 的数据：~0 rows (大约)
DELETE FROM `labels`;
/*!40000 ALTER TABLE `labels` DISABLE KEYS */
;
/*!40000 ALTER TABLE `labels` ENABLE KEYS */
;
-- 导出  表 pro_cfg_manager.line 结构
DROP TABLE IF EXISTS `line`;
CREATE TABLE IF NOT EXISTS `line` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `label` varchar(200) NOT NULL,
  `enabled` tinyint(4) NOT NULL DEFAULT 1,
  `update_at` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `update_by` varchar(200) NOT NULL,
  `idc_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb3 COMMENT = '线路，从属于机房(idc)';
-- 正在导出表  pro_cfg_manager.line 的数据：~0 rows (大约)
DELETE FROM `line`;
/*!40000 ALTER TABLE `line` DISABLE KEYS */
;
/*!40000 ALTER TABLE `line` ENABLE KEYS */
;
-- 导出  表 pro_cfg_manager.log_setting 结构
DROP TABLE IF EXISTS `log_setting`;
CREATE TABLE IF NOT EXISTS `log_setting` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `level` varchar(50) NOT NULL,
  `label` varchar(50) NOT NULL,
  `selected` tinyint(4) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 13 DEFAULT CHARSET = utf8mb3 COMMENT = '日志设置';
-- 正在导出表  pro_cfg_manager.log_setting 的数据：~8 rows (大约)
DELETE FROM `log_setting`;
/*!40000 ALTER TABLE `log_setting` DISABLE KEYS */
;
INSERT INTO `log_setting` (`id`, `level`, `label`, `selected`)
VALUES (1, 'search', '查询', 0),
  (2, 'add', '增加', 1),
  (3, 'del', '删除', 1),
  (4, 'update', '更新', 1),
  (5, 'running', '运行', 1),
  (6, 'publish', '发布', 1),
  (7, 'login', '登录', 1),
  (8, 'reset', '重置', 1);
/*!40000 ALTER TABLE `log_setting` ENABLE KEYS */
;
-- 导出  表 pro_cfg_manager.machines 结构
DROP TABLE IF EXISTS `machines`;
CREATE TABLE IF NOT EXISTS `machines` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `ipaddr` varchar(100) NOT NULL COMMENT 'IP地址',
  `idc_name` varchar(200) DEFAULT NULL COMMENT '所在机房',
  `line_name` varchar(200) DEFAULT NULL COMMENT '所在线路',
  `position` text DEFAULT NULL,
  `update_at` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `enabled` tinyint(4) NOT NULL DEFAULT 1,
  `update_by` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `ipaddr_unique` (`ipaddr`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb3 COMMENT = '机器列表';
-- 正在导出表  pro_cfg_manager.machines 的数据：~0 rows (大约)
DELETE FROM `machines`;
/*!40000 ALTER TABLE `machines` DISABLE KEYS */
;
/*!40000 ALTER TABLE `machines` ENABLE KEYS */
;
-- 导出  表 pro_cfg_manager.manager_group 结构
DROP TABLE IF EXISTS `manager_group`;
CREATE TABLE IF NOT EXISTS `manager_group` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `enabled` tinyint(4) NOT NULL,
  `update_at` datetime NOT NULL,
  `update_by` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_unique` (`name`)
) ENGINE = InnoDB AUTO_INCREMENT = 2 DEFAULT CHARSET = utf8mb3 COMMENT = '用户组';
-- 正在导出表  pro_cfg_manager.manager_group 的数据：~1 rows (大约)
DELETE FROM `manager_group`;
/*!40000 ALTER TABLE `manager_group` DISABLE KEYS */
;
INSERT INTO `manager_group` (
    `id`,
    `name`,
    `enabled`,
    `update_at`,
    `update_by`
  )
VALUES (1, 'administrator', 1, '2022-05-08 13:20:51', '');
/*!40000 ALTER TABLE `manager_group` ENABLE KEYS */
;
-- 导出  表 pro_cfg_manager.manager_set 结构
DROP TABLE IF EXISTS `manager_set`;
CREATE TABLE IF NOT EXISTS `manager_set` (
  `param_name` varbinary(100) NOT NULL,
  `param_value` varbinary(100) NOT NULL
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb3 COMMENT = '权限管理默认的设置参数';
-- 正在导出表  pro_cfg_manager.manager_set 的数据：~1 rows (大约)
DELETE FROM `manager_set`;
/*!40000 ALTER TABLE `manager_set` DISABLE KEYS */
;
INSERT INTO `manager_set` (`param_name`, `param_value`)
VALUES (_binary 0x64656661756c745f67726f7570, _binary '');
/*!40000 ALTER TABLE `manager_set` ENABLE KEYS */
;
-- 导出  表 pro_cfg_manager.manager_user 结构
DROP TABLE IF EXISTS `manager_user`;
CREATE TABLE IF NOT EXISTS `manager_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(100) NOT NULL,
  `nice_name` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `phone` varchar(50) NOT NULL DEFAULT '',
  `salt` varchar(100) NOT NULL,
  `group_id` int(11) NOT NULL,
  `update_at` datetime NOT NULL,
  `enabled` tinyint(4) NOT NULL,
  `create_at` datetime NOT NULL DEFAULT current_timestamp(),
  `update_by` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_name` (`username`)
) ENGINE = InnoDB AUTO_INCREMENT = 2 DEFAULT CHARSET = utf8mb3 COMMENT = '用户';
-- 正在导出表  pro_cfg_manager.manager_user 的数据：~1 rows (大约)
DELETE FROM `manager_user`;
/*!40000 ALTER TABLE `manager_user` DISABLE KEYS */
;
INSERT INTO `manager_user` (
    `id`,
    `username`,
    `nice_name`,
    `password`,
    `phone`,
    `salt`,
    `group_id`,
    `update_at`,
    `enabled`,
    `create_at`,
    `update_by`
  )
VALUES (
    1,
    'admin',
    '管理员',
    'd3af7695b6c3c977aafe904e99b789e7',
    '10086',
    'da456512-0967-479d-9a26-05c1c98ad28e',
    1,
    '2022-05-08 13:20:51',
    1,
    '2022-05-08 13:20:51',
    ''
  );
/*!40000 ALTER TABLE `manager_user` ENABLE KEYS */
;
-- 导出  表 pro_cfg_manager.monitor_labels 结构
DROP TABLE IF EXISTS `monitor_labels`;
CREATE TABLE IF NOT EXISTS `monitor_labels` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `monitor_rules_id` int(11) NOT NULL,
  `key` varchar(100) NOT NULL,
  `value` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_mid_lid` (`key`, `monitor_rules_id`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb3 COMMENT = '监控规则的标签';
-- 正在导出表  pro_cfg_manager.monitor_labels 的数据：~0 rows (大约)
DELETE FROM `monitor_labels`;
/*!40000 ALTER TABLE `monitor_labels` DISABLE KEYS */
;
/*!40000 ALTER TABLE `monitor_labels` ENABLE KEYS */
;
-- 导出  表 pro_cfg_manager.monitor_rules 结构
DROP TABLE IF EXISTS `monitor_rules`;
CREATE TABLE IF NOT EXISTS `monitor_rules` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `alert` varchar(100) NOT NULL,
  `expr` varchar(5000) NOT NULL,
  `for` varchar(50) NOT NULL,
  `sub_group_id` int(11) NOT NULL,
  `enabled` tinyint(4) NOT NULL,
  `description` varchar(300) NOT NULL,
  `update_by` varchar(100) NOT NULL,
  `update_at` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb3 COMMENT = '具体的监控规则';
-- 正在导出表  pro_cfg_manager.monitor_rules 的数据：~0 rows (大约)
DELETE FROM `monitor_rules`;
/*!40000 ALTER TABLE `monitor_rules` DISABLE KEYS */
;
/*!40000 ALTER TABLE `monitor_rules` ENABLE KEYS */
;
-- 导出  表 pro_cfg_manager.number_options 结构
DROP TABLE IF EXISTS `number_options`;
CREATE TABLE IF NOT EXISTS `number_options` (
  `key` varchar(50) NOT NULL,
  `value` int(11) NOT NULL DEFAULT 0,
  `commit` varchar(100) DEFAULT NULL,
  UNIQUE KEY `uniq_key` (`key`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb3 COMMENT = '值为数据的选项';
-- 正在导出表  pro_cfg_manager.number_options 的数据：~1 rows (大约)
DELETE FROM `number_options`;
/*!40000 ALTER TABLE `number_options` DISABLE KEYS */
;
INSERT INTO `number_options` (`key`, `value`, `commit`)
VALUES ('session_expire', 3600, NULL);
/*!40000 ALTER TABLE `number_options` ENABLE KEYS */
;
-- 导出  表 pro_cfg_manager.operation_log 结构
DROP TABLE IF EXISTS `operation_log`;
CREATE TABLE IF NOT EXISTS `operation_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(100) NOT NULL,
  `operate_type` varchar(50) NOT NULL,
  `ipaddr` varchar(100) NOT NULL,
  `operate_content` varchar(100) NOT NULL,
  `operate_result` tinyint(4) NOT NULL,
  `operate_at` datetime NOT NULL,
  `operate_error` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 8 DEFAULT CHARSET = utf8mb3 COMMENT = '操作日志';
-- 正在导出表  pro_cfg_manager.operation_log 的数据：~1 rows (大约)
DELETE FROM `operation_log`;
/*!40000 ALTER TABLE `operation_log` DISABLE KEYS */
;
INSERT INTO `operation_log` (
    `id`,
    `username`,
    `operate_type`,
    `ipaddr`,
    `operate_content`,
    `operate_result`,
    `operate_at`,
    `operate_error`
  )
VALUES (
    6,
    'admin',
    '',
    '127.0.0.1:55273',
    'reset prometheus config data',
    1,
    '2022-04-01 02:50:17',
    '成功'
  ),
  (
    7,
    'admin',
    '',
    '127.0.0.1:62196',
    'reset prometheus config data',
    1,
    '2022-05-08 13:20:24',
    '成功'
  );
/*!40000 ALTER TABLE `operation_log` ENABLE KEYS */
;
-- 导出  表 pro_cfg_manager.options 结构
DROP TABLE IF EXISTS `options`;
CREATE TABLE IF NOT EXISTS `options` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `opt_key` varchar(100) NOT NULL,
  `opt_value` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 10 DEFAULT CHARSET = utf8mb3 COMMENT = '选项';
-- 正在导出表  pro_cfg_manager.options 的数据：~9 rows (大约)
DELETE FROM `options`;
/*!40000 ALTER TABLE `options` DISABLE KEYS */
;
INSERT INTO `options` (`id`, `opt_key`, `opt_value`)
VALUES (1, 'publish_at_null_subgroup', 'true'),
  (2, 'publish_at_remain_subgroup', 'true'),
  (3, 'publish_at_empty_nocreate_file', 'true'),
  (4, 'publish_jobs_also_ips', 'true'),
  (5, 'publish_jobs_also_reload_srv', 'true'),
  (6, 'publish_ips_also_reload_srv', 'true'),
  (7, 'position_ipaddr', 'false'),
  (8, 'sync_prometheus_status', 'false'),
  (9, 'expand_skip_ipv6', 'false'),
  (10, 'group_by_idc_line_label', 'true'),
  (11, 'labeled_ip_no_label', 'true'),
  (12, 'labeled_all_ip', 'false');
;
/*!40000 ALTER TABLE `options` ENABLE KEYS */
;
-- 导出  表 pro_cfg_manager.page_function 结构
DROP TABLE IF EXISTS `page_function`;
CREATE TABLE IF NOT EXISTS `page_function` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `page_name` varchar(100) NOT NULL,
  `page_nice_name` varchar(100) NOT NULL,
  `sub_page_name` varchar(100) DEFAULT '',
  `sub_page_nice_name` varchar(100) DEFAULT '',
  `func_name` varchar(100) NOT NULL,
  `func_nice_name` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 132 DEFAULT CHARSET = utf8mb3 COMMENT = '页面功能';
-- 正在导出表  pro_cfg_manager.page_function 的数据：~101 rows (大约)
DELETE FROM `page_function`;
/*!40000 ALTER TABLE `page_function` DISABLE KEYS */
;
INSERT INTO `page_function` (
    `id`,
    `page_name`,
    `page_nice_name`,
    `sub_page_name`,
    `sub_page_nice_name`,
    `func_name`,
    `func_nice_name`
  )
VALUES (
    1,
    'person',
    '个人中心',
    '',
    '',
    'update_password',
    '更新密码'
  ),
  (
    2,
    'person',
    '个人中心',
    '',
    '',
    'show_menu_prometheus_cfg_manager',
    '显示Prometheus管理菜单项'
  ),
  (
    3,
    'person',
    '个人中心',
    '',
    '',
    'show_menu_administrator_cfg_manager',
    '显示用户及权限管理菜单项'
  ),
  (4, 'register', '注册', '', '', 'register', '注册'),
  (5, 'ipManager', 'IP管理', '', '', 'search', '查询'),
  (6, 'ipManager', 'IP管理', '', '', 'update', '编辑'),
  (
    7,
    'ipManager',
    'IP管理',
    '',
    '',
    'dis.enable',
    '启用/禁用'
  ),
  (
    8,
    'ipManager',
    'IP管理',
    '',
    '',
    'delete',
    '删除/批量删除'
  ),
  (9, 'ipManager', 'IP管理', '', '', 'add', '添加'),
  (
    11,
    'jobs',
    'JOB组管理',
    'jobs',
    'JOB组管理',
    'job_search',
    '查询组'
  ),
  (
    12,
    'jobs',
    'JOB组管理',
    'jobs',
    'JOB组管理',
    'job_add',
    '添加组'
  ),
  (
    13,
    'jobs',
    'JOB组管理',
    'jobs',
    'JOB组管理',
    'job_update',
    '编辑组'
  ),
  (
    14,
    'jobs',
    'JOB组管理',
    'jobs',
    'JOB组管理',
    'dis.enable',
    '启用/禁用组'
  ),
  (
    15,
    'jobs',
    'JOB组管理',
    'jobs',
    'JOB组管理',
    'ip_pool',
    '更新组IP池'
  ),
  (
    16,
    'jobs',
    'JOB组管理',
    'jobs',
    'JOB组管理',
    'delete',
    '删除组'
  ),
  (
    19,
    'jobs',
    'JOB组管理',
    'labelsJobs',
    '分组标签',
    'search_sub_group',
    '查询子组'
  ),
  (
    20,
    'jobs',
    'JOB组管理',
    'labelsJobs',
    '分组标签',
    'add_sub_group',
    '添加子组'
  ),
  (
    21,
    'jobs',
    'JOB组管理',
    'labelsJobs',
    '分组标签',
    'update_sub_group',
    '更新子组'
  ),
  (
    22,
    'jobs',
    'JOB组管理',
    'labelsJobs',
    '分组标签',
    'dis.enable_sub_group',
    '启用/禁用子组'
  ),
  (
    23,
    'jobs',
    'JOB组管理',
    'labelsJobs',
    '分组标签',
    'delete_sub_group',
    '删除子组'
  ),
  (
    24,
    'jobs',
    'JOB组管理',
    'labelsJobs',
    '分组标签',
    'search_sub_group_ip_pool',
    '查询子组IP池'
  ),
  (
    25,
    'jobs',
    'JOB组管理',
    'labelsJobs',
    '分组标签',
    'update_sub_group_ip_pool',
    '更新子组IP池'
  ),
  (
    26,
    'jobs',
    'JOB组管理',
    'labelsJobs',
    '分组标签',
    'search_sub_group_label',
    '查询子组标签列表'
  ),
  (
    27,
    'jobs',
    'JOB组管理',
    'labelsJobs',
    '分组标签',
    'update_sub_group_label',
    '更新子组标签'
  ),
  (
    28,
    'jobs',
    'JOB组管理',
    'labelsJobs',
    '分组标签',
    'dis.enable_sub_group_label',
    '启用/禁用子组标签'
  ),
  (
    29,
    'jobs',
    'JOB组管理',
    'labelsJobs',
    '分组标签',
    'delete_sub_group_label',
    '删除子组标签'
  ),
  (
    30,
    'baseConfig',
    '基本配置',
    'options',
    '选项编辑',
    'search_options',
    '查询选项'
  ),
  (
    31,
    'baseConfig',
    '基本配置',
    'options',
    '选项编辑',
    'update_options',
    '更新选项'
  ),
  (
    33,
    'baseConfig',
    '基本配置',
    'baseLabels',
    '公共标签',
    'search',
    '查询'
  ),
  (
    34,
    'baseConfig',
    '基本配置',
    'baseLabels',
    '公共标签',
    'add',
    '增加'
  ),
  (
    35,
    'baseConfig',
    '基本配置',
    'baseLabels',
    '公共标签',
    'update',
    '更新'
  ),
  (
    36,
    'baseConfig',
    '基本配置',
    'baseLabels',
    '公共标签',
    'dis.enable',
    '启用/禁用'
  ),
  (
    37,
    'baseConfig',
    '基本配置',
    'baseLabels',
    '公共标签',
    'delete',
    '删除'
  ),
  (
    38,
    'baseConfig',
    '基本配置',
    'reLabels',
    '监控规则',
    'search',
    '查询'
  ),
  (
    39,
    'baseConfig',
    '基本配置',
    'reLabels',
    '监控规则',
    'add',
    '增加'
  ),
  (
    40,
    'baseConfig',
    '基本配置',
    'reLabels',
    '监控规则',
    'edit_name',
    '编辑名称'
  ),
  (
    41,
    'baseConfig',
    '基本配置',
    'reLabels',
    '监控规则',
    'update_rule',
    '更新规则'
  ),
  (
    42,
    'baseConfig',
    '基本配置',
    'reLabels',
    '监控规则',
    'dis.enable',
    '启用/禁用'
  ),
  (
    43,
    'baseConfig',
    '基本配置',
    'reLabels',
    '监控规则',
    'delete',
    '删除'
  ),
  (
    44,
    'baseConfig',
    '基本配置',
    'defaultJobs',
    '默认分组',
    'search',
    '查询'
  ),
  (
    45,
    'baseConfig',
    '基本配置',
    'defaultJobs',
    '默认分组',
    'add',
    '增加'
  ),
  (
    46,
    'baseConfig',
    '基本配置',
    'defaultJobs',
    '默认分组',
    'update',
    '更新'
  ),
  (
    47,
    'baseConfig',
    '基本配置',
    'defaultJobs',
    '默认分组',
    'dis.enable',
    '启用/禁用'
  ),
  (
    48,
    'baseConfig',
    '基本配置',
    'defaultJobs',
    '默认分组',
    'delete',
    '删除'
  ),
  (
    49,
    'baseConfig',
    '基本配置',
    'checkYml',
    '测试配置',
    'restart',
    '重启Prometheus服务'
  ),
  (
    50,
    'baseConfig',
    '基本配置',
    'checkYml',
    '测试配置',
    'check',
    '测试配置文件'
  ),
  (
    51,
    'baseConfig',
    '基本配置',
    'editPrometheusYml',
    '模板编辑',
    'load_tmpl',
    '加载模板数据'
  ),
  (
    52,
    'baseConfig',
    '基本配置',
    'editPrometheusYml',
    '模板编辑',
    'save_tmpl',
    '保存模板数据'
  ),
  (
    53,
    'baseConfig',
    '基本配置',
    'empty',
    '重置',
    'reset',
    '重置所有数据'
  ),
  (
    54,
    'noticeManager',
    '告警管理',
    '',
    '',
    'search_tree',
    '获取告警规则列表'
  ),
  (
    55,
    'noticeManager',
    '告警管理',
    '',
    '',
    'add_file',
    '添加文件'
  ),
  (
    56,
    'noticeManager',
    '告警管理',
    '',
    '',
    'dis.enable',
    '启用/禁用分支'
  ),
  (
    57,
    'noticeManager',
    '告警管理',
    '',
    '',
    'delete_tree_node',
    '删除所有子节点'
  ),
  (
    58,
    'noticeManager',
    '告警管理',
    '',
    '',
    'delete_node',
    '删除此节点'
  ),
  (
    59,
    'noticeManager',
    '告警管理',
    '',
    '',
    'rename_node',
    '重命名此节点'
  ),
  (
    60,
    'noticeManager',
    '告警管理',
    '',
    '',
    'add_group',
    '增加组'
  ),
  (
    61,
    'noticeManager',
    '告警管理',
    '',
    '',
    'import_rule_from_file',
    '从文件导入规则'
  ),
  (
    62,
    'noticeManager',
    '告警管理',
    '',
    '',
    'add_rule',
    '添加规则'
  ),
  (
    63,
    'noticeManager',
    '告警管理',
    '',
    '',
    'search_rule',
    '查询规则内容'
  ),
  (
    64,
    'noticeManager',
    '告警管理',
    '',
    '',
    'update_rule',
    '更新规则'
  ),
  (
    65,
    'noticeManager',
    '告警管理',
    '',
    '',
    'publish_rule',
    '发布'
  ),
  (66, 'preview', '配置预览', '', '', 'search', '查看配置'),
  (
    67,
    'ftree',
    '分组预览',
    '',
    '',
    'list_all_file',
    '列表分组IP文件'
  ),
  (
    68,
    'ruleView',
    '告警规则预览',
    '',
    '',
    'list_all_file',
    '列表告警规则文件'
  ),
  (
    69,
    'admin',
    '管理中心',
    'user',
    '用户管理',
    'search',
    '查看用户列表'
  ),
  (
    70,
    'admin',
    '管理中心',
    'user',
    '用户管理',
    'add',
    '增加用户'
  ),
  (
    71,
    'admin',
    '管理中心',
    'user',
    '用户管理',
    'update',
    '更新用户'
  ),
  (
    72,
    'admin',
    '管理中心',
    'user',
    '用户管理',
    'dis.enable',
    '启用/禁用用户'
  ),
  (
    73,
    'admin',
    '管理中心',
    'user',
    '用户管理',
    'delete',
    '删除用户'
  ),
  (
    74,
    'admin',
    '管理中心',
    'group',
    '组管理',
    'search',
    '查看用户组列表'
  ),
  (
    75,
    'admin',
    '管理中心',
    'group',
    '组管理',
    'add',
    '添加组'
  ),
  (
    76,
    'admin',
    '管理中心',
    'group',
    '组管理',
    'update',
    '更新组'
  ),
  (
    77,
    'admin',
    '管理中心',
    'group',
    '组管理',
    'dis.enable',
    '启用/禁用组'
  ),
  (
    78,
    'admin',
    '管理中心',
    'group',
    '组管理',
    'delete',
    '删除组'
  ),
  (
    79,
    'admin',
    '管理中心',
    'privileges',
    '权限管理',
    'search',
    '查看权限列表'
  ),
  (
    80,
    'admin',
    '管理中心',
    'privileges',
    '权限管理',
    'update',
    '更新权限列表'
  ),
  (
    81,
    'jobs',
    'JOB组管理',
    'jobs',
    'JOB组管理',
    'swap',
    '调整顺序'
  ),
  (
    83,
    'ftree',
    '分组预览',
    '',
    '',
    'view_file',
    '查看IP文件内容'
  ),
  (
    84,
    'ruleView',
    '告警规则预览',
    '',
    '',
    'view_file',
    '查看告警规则文件内容'
  ),
  (
    85,
    'admin',
    '管理中心',
    'setting',
    '设置',
    'search',
    '获取设置'
  ),
  (
    86,
    'admin',
    '管理中心',
    'setting',
    '设置',
    'update',
    '更新设置'
  ),
  (
    87,
    'admin',
    '管理中心',
    'setting',
    '设置',
    'reset',
    '重置数据'
  ),
  (
    88,
    'ipManager',
    'IP管理',
    '',
    '',
    'import',
    '批量添加'
  ),
  (
    89,
    'jobs',
    'JOB组管理',
    'labelsJobs',
    '分组标签',
    'add_sub_group_label',
    '添加子组标签'
  ),
  (
    90,
    'baseConfig',
    '基本配置',
    'reLabels',
    '监控规则',
    'view_code',
    '展示代码'
  ),
  (
    93,
    'baseConfig',
    '基本配置',
    'baseFields',
    '模板字段',
    'search',
    '查询'
  ),
  (
    94,
    'baseConfig',
    '基本配置',
    'baseFields',
    '模板字段',
    'add',
    '增加'
  ),
  (
    95,
    'baseConfig',
    '基本配置',
    'baseFields',
    '模板字段',
    'update',
    '更新'
  ),
  (
    96,
    'baseConfig',
    '基本配置',
    'baseFields',
    '模板字段',
    'dis.enable',
    '启用/禁用'
  ),
  (
    97,
    'baseConfig',
    '基本配置',
    'baseFields',
    '模板字段',
    'delete',
    '删除'
  ),
  (98, 'admin', '管理中心', 'log', '日志', 'search', '查询'),
  (
    99,
    'admin',
    '管理中心',
    'log',
    '日志',
    'update',
    '更新设置'
  ),
  (
    100,
    'admin',
    '管理中心',
    'log',
    '日志',
    'clear',
    '清空日志'
  ),
  (
    101,
    'control',
    '控制中心',
    '',
    '',
    'create',
    '重新创建配置'
  ),
  (
    102,
    'control',
    '控制中心',
    '',
    '',
    'reload',
    '重新加载配置'
  ),
  (
    103,
    'control',
    '控制中心',
    '',
    '',
    'createAndreload',
    '创建配置并加载'
  ),
  (
    104,
    'noticeManager',
    '告警管理',
    '',
    '',
    'publish_empty_rule',
    '发布空规则'
  ),
  (
    105,
    'ipManager',
    'IP管理',
    '',
    '',
    'position',
    '更新IP位置信息'
  ),
  (106, 'idc', 'IDC机房', '', '', 'get_idc', '获取机房信息'),
  (107, 'idc', 'IDC机房', '', '', 'add_idc', '新建机房'),
  (
    108,
    'idc',
    'IDC机房',
    '',
    '',
    'update_idc',
    '更新机房'
  ),
  (109, 'idc', 'IDC机房', '', '', 'del_idc', '删除机房'),
  (
    110,
    'idc',
    'IDC机房',
    '',
    '',
    'get_idc_tree',
    '获取机房列表'
  ),
  (
    111,
    'idc',
    'IDC机房',
    '',
    '',
    'get_line',
    '获取线路信息'
  ),
  (112, 'idc', 'IDC机房', '', '', 'add_line', '新建线路'),
  (
    113,
    'idc',
    'IDC机房',
    '',
    '',
    'update_line',
    '更新线路'
  ),
  (114, 'idc', 'IDC机房', '', '', 'del_line', '删除线路'),
  (
    115,
    'idc',
    'IDC机房',
    '',
    '',
    'get_line_ipaddr',
    '获取线路地址池'
  ),
  (
    116,
    'idc',
    'IDC机房',
    '',
    '',
    'update_line_ipaddr',
    '更新线路地址池'
  ),
  (
    117,
    'idc',
    'IDC机房',
    '',
    '',
    'update_label_part',
    '只更新未设置IP'
  ),
  (
    118,
    'idc',
    'IDC机房',
    '',
    '',
    'update_label_all',
    '更新所有IP'
  ),
  (
    119,
    'idc',
    'IDC机房',
    '',
    '',
    'create_label_for_all_job',
    '在JOB组中生成标签'
  ),
  (
    120,
    'ipManager',
    'IP管理',
    '',
    '',
    'download',
    '导出IP'
  ),
  (
    121,
    'admin',
    '管理中心',
    'session',
    '会话管理',
    'search',
    '获取用户会话'
  ),
  (
    122,
    'admin',
    '管理中心',
    'session',
    '会话管理',
    'delete',
    '删除用户会话'
  ),
  (
    123,
    'control',
    '控制中心',
    '',
    '',
    'get_prometheus_url',
    '打开Prometheus管理后台'
  ),
  (
    124,
    'jobs',
    'JOB组管理',
    'labelsJobs',
    '分组标签',
    'delete_empty_sub_group',
    '删除没有IP的空子组'
  ),
  (
    125,
    'admin',
    '管理中心',
    'session',
    '会话管理',
    'update',
    '更新会话参数'
  ),
  (
    126,
    'jobs',
    'JOB组管理',
    'jobs',
    'JOB组管理',
    'job_search_black',
    '查询组黑名单'
  ),
  (
    127,
    'jobs',
    'JOB组管理',
    'jobs',
    'JOB组管理',
    'job_update_black',
    '更新组黑名单'
  ),
  (
    128,
    'jobs',
    'JOB组管理',
    'jobs',
    'JOB组管理',
    'update_job_subgroup',
    '更新子组'
  ),
  (
    129,
    'idc',
    'IDC机房',
    '',
    '',
    'expand_ipaddr',
    '扩展IP地址'
  ),
  (
    130,
    'instance',
    '实例导入',
    '',
    '',
    'get_data',
    '获取实例数据'
  ),
  (
    131,
    'instance',
    '实例导入',
    '',
    '',
    'put_data',
    '导入实例数据'
  );
/*!40000 ALTER TABLE `page_function` ENABLE KEYS */
;
-- 导出  表 pro_cfg_manager.pool 结构
DROP TABLE IF EXISTS `pool`;
CREATE TABLE IF NOT EXISTS `pool` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `line_id` int(11) NOT NULL COMMENT '线路ID',
  `remark` text DEFAULT NULL,
  `ipaddrs` text DEFAULT NULL COMMENT '地址池，以英文分号隔开',
  `update_at` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `update_by` varchar(200) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_line_id` (`line_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb3 COMMENT = '线路的地址池，从属于line表';
-- 正在导出表  pro_cfg_manager.pool 的数据：~0 rows (大约)
DELETE FROM `pool`;
/*!40000 ALTER TABLE `pool` DISABLE KEYS */
;
/*!40000 ALTER TABLE `pool` ENABLE KEYS */
;
-- 导出  表 pro_cfg_manager.relabels 结构
DROP TABLE IF EXISTS `relabels`;
CREATE TABLE IF NOT EXISTS `relabels` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(200) NOT NULL,
  `code` text DEFAULT NULL,
  `enabled` tinyint(4) NOT NULL DEFAULT 1,
  `update_at` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `update_by` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_name` (`name`)
) ENGINE = InnoDB AUTO_INCREMENT = 41 DEFAULT CHARSET = utf8mb3 COMMENT = '重写标签的配置';
-- 正在导出表  pro_cfg_manager.relabels 的数据：~1 rows (大约)
DELETE FROM `relabels`;
/*!40000 ALTER TABLE `relabels` DISABLE KEYS */
;
INSERT INTO `relabels` (
    `id`,
    `name`,
    `code`,
    `enabled`,
    `update_at`,
    `update_by`
  )
VALUES (1, '空规则', '', 1, '2022-03-24 22:44:52', 'admin');
/*!40000 ALTER TABLE `relabels` ENABLE KEYS */
;
-- 导出  表 pro_cfg_manager.rules_groups 结构
DROP TABLE IF EXISTS `rules_groups`;
CREATE TABLE IF NOT EXISTS `rules_groups` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `update_at` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `update_by` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_name` (`name`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb3 COMMENT = '监控规则组';
-- 正在导出表  pro_cfg_manager.rules_groups 的数据：~0 rows (大约)
DELETE FROM `rules_groups`;
/*!40000 ALTER TABLE `rules_groups` DISABLE KEYS */
;
/*!40000 ALTER TABLE `rules_groups` ENABLE KEYS */
;
-- 导出  表 pro_cfg_manager.session 结构
DROP TABLE IF EXISTS `session`;
CREATE TABLE IF NOT EXISTS `session` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `token` varchar(100) NOT NULL,
  `ipaddr` varchar(120) DEFAULT NULL,
  `user_id` int(11) NOT NULL,
  `update_at` datetime NOT NULL,
  `login_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_user_id` (`user_id`)
) ENGINE = InnoDB AUTO_INCREMENT = 136 DEFAULT CHARSET = utf8mb3 COMMENT = '用户登录会话';
-- 正在导出表  pro_cfg_manager.session 的数据：~0 rows (大约)
DELETE FROM `session`;
/*!40000 ALTER TABLE `session` DISABLE KEYS */
;
INSERT INTO `session` (
    `id`,
    `token`,
    `ipaddr`,
    `user_id`,
    `update_at`,
    `login_at`
  )
VALUES (
    109,
    '62a2a974-caed-4ac1-9f8c-f35abe7d51e3',
    '127.0.0.1:57796',
    1,
    '2022-05-08 13:20:54',
    '0000-00-00 00:00:00'
  );
/*!40000 ALTER TABLE `session` ENABLE KEYS */
;
-- 导出  表 pro_cfg_manager.sub_group 结构
DROP TABLE IF EXISTS `sub_group`;
CREATE TABLE IF NOT EXISTS `sub_group` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `rules_groups_id` int(11) NOT NULL,
  `update_at` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `update_by` varchar(50) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_name_rules_groups_id` (`name`, `rules_groups_id`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb3 COMMENT = '监控规则组中的下级组，子组';
-- 正在导出表  pro_cfg_manager.sub_group 的数据：~0 rows (大约)
DELETE FROM `sub_group`;
/*!40000 ALTER TABLE `sub_group` DISABLE KEYS */
;
/*!40000 ALTER TABLE `sub_group` ENABLE KEYS */
;
-- 导出  表 pro_cfg_manager.system_log 结构
DROP TABLE IF EXISTS `system_log`;
CREATE TABLE IF NOT EXISTS `system_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(100) NOT NULL,
  `operate_type` varchar(50) NOT NULL,
  `ipaddr` varchar(100) NOT NULL,
  `operate_content` varchar(100) NOT NULL,
  `operate_result` tinyint(4) NOT NULL,
  `operate_at` datetime NOT NULL,
  `operate_error` text NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb3 COMMENT = '系统日志';
-- 正在导出表  pro_cfg_manager.system_log 的数据：~0 rows (大约)
DELETE FROM `system_log`;
/*!40000 ALTER TABLE `system_log` DISABLE KEYS */
;
/*!40000 ALTER TABLE `system_log` ENABLE KEYS */
;
-- 导出  表 pro_cfg_manager.tmpl 结构
DROP TABLE IF EXISTS `tmpl`;
CREATE TABLE IF NOT EXISTS `tmpl` (
  `tmpl` longtext NOT NULL,
  `update_at` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `update_by` varchar(100) NOT NULL
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb3 COMMENT = 'prometheus.yml';
-- 正在导出表  pro_cfg_manager.tmpl 的数据：~2 rows (大约)
DELETE FROM `tmpl`;
/*!40000 ALTER TABLE `tmpl` DISABLE KEYS */
;
INSERT INTO `tmpl` (`tmpl`, `update_at`, `update_by`)
VALUES (
    '# my global config\nglobal:\n  scrape_interval:     15s # Set the scrape interval to every 15 seconds. Default is every 1 minute.\n  evaluation_interval: 15s # Evaluate rules every 15 seconds. The default is every 1 minute.\n  # scrape_timeout is set to the global default (10s).\n  # scrape_timeout: 60s\n\n# Alertmanager configuration\nalerting:\n  alertmanagers:\n  - static_configs:\n    - targets:\n      - 127.0.0.1:9093\n\n# Load rules once and periodically evaluate them according to the global \'evaluation_interval\'.\nrule_files:\n   - "{{.RelRuleDir}}/*.yml"\n  # - "first_rules.yml"\n  # - "second_rules.yml"\n\n# A scrape configuration containing exactly one endpoint to scrape:\n# Here it\'s Prometheus itself.\nscrape_configs:\n  # The job name is added as a label `job=<job_name>` to any timeseries scraped from this config.\n  - job_name: \'监控服务本机\'\n\n    # metrics_path defaults to \'/metrics\'\n    # scheme defaults to \'http\'.\n\n    static_configs:\n    - targets: [\'localhost:9090\']\n{{ range .Jobs }}\n  - job_name: \'{{.Name}}\'\n    file_sd_configs:\n      - files:\n        - "{{$.AbsConfDir}}/{{.Name}}.json"\n        refresh_interval: {{$.Fields.refresh_interval}}\n{{.Code}}\n{{ end }}\n',
    '2022-05-08 13:20:24',
    'administrator'
  );
/*!40000 ALTER TABLE `tmpl` ENABLE KEYS */
;
-- 导出  表 pro_cfg_manager.tmpl_fields 结构
DROP TABLE IF EXISTS `tmpl_fields`;
CREATE TABLE IF NOT EXISTS `tmpl_fields` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `key` varchar(100) NOT NULL,
  `value` varchar(500) NOT NULL,
  `enabled` tinyint(4) NOT NULL,
  `update_at` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `update_by` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `key_unique` (`key`)
) ENGINE = InnoDB AUTO_INCREMENT = 3 DEFAULT CHARSET = utf8mb3 COMMENT = '模板字段';
-- 正在导出表  pro_cfg_manager.tmpl_fields 的数据：~2 rows (大约)
DELETE FROM `tmpl_fields`;
/*!40000 ALTER TABLE `tmpl_fields` DISABLE KEYS */
;
INSERT INTO `tmpl_fields` (
    `id`,
    `key`,
    `value`,
    `enabled`,
    `update_at`,
    `update_by`
  )
VALUES (
    1,
    'metrics',
    '/metrics',
    1,
    '2022-05-08 13:20:24',
    ''
  ),
  (
    2,
    'refresh_interval',
    '15s',
    1,
    '2022-05-08 13:20:24',
    ''
  );
/*!40000 ALTER TABLE `tmpl_fields` ENABLE KEYS */
;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */
;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */
;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */
;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */
;