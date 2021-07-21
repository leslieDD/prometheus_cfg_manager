-- --------------------------------------------------------
-- 主机:                           27.151.29.204
-- 服务器版本:                        8.0.21 - Source distribution
-- 服务器操作系统:                      Linux
-- HeidiSQL 版本:                  11.3.0.6295
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- 导出 pro_cfg_manager 的数据库结构
DROP DATABASE IF EXISTS `pro_cfg_manager`;
CREATE DATABASE IF NOT EXISTS `pro_cfg_manager` /*!40100 DEFAULT CHARACTER SET utf8 */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `pro_cfg_manager`;

-- 导出  表 pro_cfg_manager.annotations 结构
DROP TABLE IF EXISTS `annotations`;
CREATE TABLE IF NOT EXISTS `annotations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `monitor_rules_id` int NOT NULL,
  `key` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `value` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_monitor_rules_id_key` (`monitor_rules_id`,`key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='监控规则的注释';

-- 正在导出表  pro_cfg_manager.annotations 的数据：~368 rows (大约)
DELETE FROM `annotations`;
/*!40000 ALTER TABLE `annotations` DISABLE KEYS */;
/*!40000 ALTER TABLE `annotations` ENABLE KEYS */;

-- 导出  表 pro_cfg_manager.group_labels 结构
DROP TABLE IF EXISTS `group_labels`;
CREATE TABLE IF NOT EXISTS `group_labels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `job_group_id` int NOT NULL,
  `key` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `value` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `enabled` tinyint NOT NULL DEFAULT '1',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_groupid_key` (`key`,`job_group_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='子组标签列表';

-- 正在导出表  pro_cfg_manager.group_labels 的数据：~2 rows (大约)
DELETE FROM `group_labels`;
/*!40000 ALTER TABLE `group_labels` DISABLE KEYS */;
/*!40000 ALTER TABLE `group_labels` ENABLE KEYS */;

-- 导出  表 pro_cfg_manager.group_machines 结构
DROP TABLE IF EXISTS `group_machines`;
CREATE TABLE IF NOT EXISTS `group_machines` (
  `id` int NOT NULL AUTO_INCREMENT,
  `job_group_id` int NOT NULL,
  `machines_id` int NOT NULL,
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique` (`machines_id`,`job_group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='子组IP列表';

-- 正在导出表  pro_cfg_manager.group_machines 的数据：~218 rows (大约)
DELETE FROM `group_machines`;
/*!40000 ALTER TABLE `group_machines` DISABLE KEYS */;
/*!40000 ALTER TABLE `group_machines` ENABLE KEYS */;

-- 导出  表 pro_cfg_manager.jobs 结构
DROP TABLE IF EXISTS `jobs`;
CREATE TABLE IF NOT EXISTS `jobs` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(100) NOT NULL COMMENT '任务名称',
  `port` int NOT NULL COMMENT '端口号，对应exporter的端口号',
  `cfg_name` varchar(100) DEFAULT NULL COMMENT '在prometheus下生成配置的文件名称',
  `is_common` tinyint NOT NULL,
  `relabel_id` int NOT NULL,
  `display_order` int unsigned NOT NULL,
  `enabled` tinyint NOT NULL DEFAULT '1',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_unique` (`name`),
  KEY `order_unique` (`display_order`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='prometheus任务列表';

-- 正在导出表  pro_cfg_manager.jobs 的数据：~28 rows (大约)
DELETE FROM `jobs`;
/*!40000 ALTER TABLE `jobs` DISABLE KEYS */;
/*!40000 ALTER TABLE `jobs` ENABLE KEYS */;

-- 导出  表 pro_cfg_manager.job_group 结构
DROP TABLE IF EXISTS `job_group`;
CREATE TABLE IF NOT EXISTS `job_group` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `jobs_id` int NOT NULL,
  `enabled` tinyint NOT NULL DEFAULT '1',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='还可以为单个分组中的IP地址进行分子组，为每个子组设置相应的标签列表';

-- 正在导出表  pro_cfg_manager.job_group 的数据：~32 rows (大约)
DELETE FROM `job_group`;
/*!40000 ALTER TABLE `job_group` DISABLE KEYS */;
/*!40000 ALTER TABLE `job_group` ENABLE KEYS */;

-- 导出  表 pro_cfg_manager.job_machines 结构
DROP TABLE IF EXISTS `job_machines`;
CREATE TABLE IF NOT EXISTS `job_machines` (
  `machine_id` int NOT NULL,
  `job_id` int NOT NULL,
  UNIQUE KEY `unique_mj` (`job_id`,`machine_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 正在导出表  pro_cfg_manager.job_machines 的数据：~231 rows (大约)
DELETE FROM `job_machines`;
/*!40000 ALTER TABLE `job_machines` DISABLE KEYS */;
/*!40000 ALTER TABLE `job_machines` ENABLE KEYS */;

-- 导出  表 pro_cfg_manager.labels 结构
DROP TABLE IF EXISTS `labels`;
CREATE TABLE IF NOT EXISTS `labels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `label` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `enabled` tinyint NOT NULL DEFAULT '1',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_labels` (`label`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='标签列表';

-- 正在导出表  pro_cfg_manager.labels 的数据：~3 rows (大约)
DELETE FROM `labels`;
/*!40000 ALTER TABLE `labels` DISABLE KEYS */;
/*!40000 ALTER TABLE `labels` ENABLE KEYS */;

-- 导出  表 pro_cfg_manager.machines 结构
DROP TABLE IF EXISTS `machines`;
CREATE TABLE IF NOT EXISTS `machines` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `ipaddr` varchar(100) NOT NULL COMMENT 'IP地址',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `enabled` tinyint NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `ipaddr_unique` (`ipaddr`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='机器列表';

-- 正在导出表  pro_cfg_manager.machines 的数据：~208 rows (大约)
DELETE FROM `machines`;
/*!40000 ALTER TABLE `machines` DISABLE KEYS */;
/*!40000 ALTER TABLE `machines` ENABLE KEYS */;

-- 导出  表 pro_cfg_manager.monitor_labels 结构
DROP TABLE IF EXISTS `monitor_labels`;
CREATE TABLE IF NOT EXISTS `monitor_labels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `monitor_rules_id` int NOT NULL,
  `key` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `value` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_mid_lid` (`key`,`monitor_rules_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='监控规则的标签';

-- 正在导出表  pro_cfg_manager.monitor_labels 的数据：~184 rows (大约)
DELETE FROM `monitor_labels`;
/*!40000 ALTER TABLE `monitor_labels` DISABLE KEYS */;
/*!40000 ALTER TABLE `monitor_labels` ENABLE KEYS */;

-- 导出  表 pro_cfg_manager.monitor_rules 结构
DROP TABLE IF EXISTS `monitor_rules`;
CREATE TABLE IF NOT EXISTS `monitor_rules` (
  `id` int NOT NULL AUTO_INCREMENT,
  `alert` varchar(100) NOT NULL,
  `expr` varchar(5000) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `for` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `sub_group_id` int NOT NULL,
  `enabled` tinyint NOT NULL,
  `description` varchar(300) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='具体的监控规则';

-- 正在导出表  pro_cfg_manager.monitor_rules 的数据：~184 rows (大约)
DELETE FROM `monitor_rules`;
/*!40000 ALTER TABLE `monitor_rules` DISABLE KEYS */;
/*!40000 ALTER TABLE `monitor_rules` ENABLE KEYS */;

-- 导出  表 pro_cfg_manager.operation_log 结构
DROP TABLE IF EXISTS `operation_log`;
CREATE TABLE IF NOT EXISTS `operation_log` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(100) NOT NULL,
  `ipaddr` varchar(100) NOT NULL,
  `operate_name` varchar(100) NOT NULL,
  `operate_result` tinyint NOT NULL,
  `operate_at` datetime NOT NULL,
  `operate_error` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=136 DEFAULT CHARSET=utf8 COMMENT='操作日志';

-- 正在导出表  pro_cfg_manager.operation_log 的数据：~0 rows (大约)
DELETE FROM `operation_log`;
/*!40000 ALTER TABLE `operation_log` DISABLE KEYS */;
INSERT INTO `operation_log` (`id`, `username`, `ipaddr`, `operate_name`, `operate_result`, `operate_at`, `operate_error`) VALUES
	(7, '', '127.0.0.1:51164', 'reset_system', 1, '2021-07-21 16:00:58', ''),
	(8, '', '127.0.0.1:63202', 'reset_system', 1, '2021-07-21 16:00:59', ''),
	(9, '', '127.0.0.1:49816', 'reset_system', 1, '2021-07-21 16:04:34', ''),
	(10, '', '127.0.0.1:52119', 'reset_system', 1, '2021-07-21 16:04:34', ''),
	(11, '', '127.0.0.1:52252', 'reset_system', 1, '2021-07-21 16:04:34', ''),
	(12, '', '127.0.0.1:61637', 'reset_system', 1, '2021-07-21 16:04:34', ''),
	(13, '', '127.0.0.1:51046', 'reset_system', 1, '2021-07-21 16:04:34', ''),
	(14, '', '127.0.0.1:58614', 'reset_system', 1, '2021-07-21 16:04:34', ''),
	(15, '', '127.0.0.1:57452', 'reset_system', 1, '2021-07-21 16:04:34', ''),
	(16, '', '127.0.0.1:51434', 'reset_system', 1, '2021-07-21 16:04:35', ''),
	(17, '', '127.0.0.1:61305', 'reset_system', 1, '2021-07-21 16:04:35', ''),
	(18, '', '127.0.0.1:54271', 'reset_system', 1, '2021-07-21 16:07:00', ''),
	(19, '', '127.0.0.1:65424', 'reset_system', 1, '2021-07-21 16:07:00', ''),
	(20, '', '127.0.0.1:57313', 'reset_system', 1, '2021-07-21 16:07:00', ''),
	(21, '', '127.0.0.1:49678', 'reset_system', 1, '2021-07-21 16:07:00', ''),
	(22, '', '127.0.0.1:54662', 'reset_system', 1, '2021-07-21 16:07:00', ''),
	(23, '', '127.0.0.1:63308', 'reset_system', 1, '2021-07-21 16:07:01', ''),
	(24, '', '127.0.0.1:61646', 'reset_system', 1, '2021-07-21 16:07:01', ''),
	(25, '', '127.0.0.1:55009', 'reset_system', 1, '2021-07-21 16:07:01', ''),
	(26, '', '127.0.0.1:60500', 'reset_system', 1, '2021-07-21 16:07:01', ''),
	(27, '', '127.0.0.1:64514', 'reset_system', 1, '2021-07-21 16:07:01', ''),
	(28, '', '127.0.0.1:58225', 'reset_system', 1, '2021-07-21 16:07:02', ''),
	(29, '', '127.0.0.1:56942', 'reset_system', 1, '2021-07-21 16:07:02', ''),
	(30, '', '127.0.0.1:59924', 'reset_system', 1, '2021-07-21 16:07:02', ''),
	(31, '', '127.0.0.1:60048', 'reset_system', 1, '2021-07-21 16:07:02', ''),
	(32, '', '127.0.0.1:65291', 'reset_system', 1, '2021-07-21 16:07:02', ''),
	(33, '', '127.0.0.1:55444', 'reset_system', 1, '2021-07-21 16:07:02', ''),
	(34, '', '127.0.0.1:58005', 'reset_system', 1, '2021-07-21 16:07:03', ''),
	(35, '', '127.0.0.1:56234', 'reset_system', 1, '2021-07-21 16:07:03', ''),
	(36, '', '127.0.0.1:59942', 'reset_system', 1, '2021-07-21 16:11:29', ''),
	(37, '', '127.0.0.1:61165', 'reset_system', 1, '2021-07-21 16:11:29', ''),
	(38, '', '127.0.0.1:52554', 'reset_system', 1, '2021-07-21 16:11:29', ''),
	(39, '', '127.0.0.1:56160', 'reset_system', 1, '2021-07-21 16:11:30', ''),
	(40, '', '127.0.0.1:50005', 'reset_system', 1, '2021-07-21 16:11:30', ''),
	(41, '', '127.0.0.1:62027', 'reset_system', 1, '2021-07-21 16:11:30', ''),
	(42, '', '127.0.0.1:60412', 'reset_system', 1, '2021-07-21 16:11:30', ''),
	(43, '', '127.0.0.1:52935', 'reset_system', 1, '2021-07-21 16:11:30', ''),
	(44, '', '127.0.0.1:62708', 'reset_system', 1, '2021-07-21 16:11:30', ''),
	(45, '', '127.0.0.1:49388', 'reset_system', 1, '2021-07-21 16:11:31', ''),
	(46, '', '127.0.0.1:58063', 'reset_system', 1, '2021-07-21 16:11:31', ''),
	(47, '', '127.0.0.1:64095', 'reset_system', 1, '2021-07-21 16:11:31', ''),
	(48, '', '127.0.0.1:55255', 'reset_system', 1, '2021-07-21 16:11:31', ''),
	(49, '', '127.0.0.1:50480', 'reset_system', 1, '2021-07-21 16:11:31', ''),
	(50, '', '127.0.0.1:60387', 'reset_system', 1, '2021-07-21 16:11:31', ''),
	(51, '', '127.0.0.1:50788', 'reset_system', 1, '2021-07-21 16:11:32', ''),
	(52, '', '127.0.0.1:49595', 'reset_system', 1, '2021-07-21 16:11:32', ''),
	(53, '', '127.0.0.1:57385', 'reset_system', 1, '2021-07-21 16:11:32', ''),
	(54, '', '127.0.0.1:62381', 'reset_system', 1, '2021-07-21 16:11:32', ''),
	(55, '', '127.0.0.1:59900', 'reset_system', 1, '2021-07-21 16:11:32', ''),
	(56, '', '127.0.0.1:53911', 'reset_system', 1, '2021-07-21 16:11:32', ''),
	(57, '', '127.0.0.1:61607', 'reset_system', 1, '2021-07-21 16:11:33', ''),
	(58, '', '127.0.0.1:51798', 'reset_system', 1, '2021-07-21 16:11:33', ''),
	(59, '', '127.0.0.1:65398', 'reset_system', 1, '2021-07-21 16:11:33', ''),
	(60, '', '127.0.0.1:57436', 'reset_system', 1, '2021-07-21 16:11:33', ''),
	(61, '', '127.0.0.1:57679', 'reset_system', 1, '2021-07-21 16:11:33', ''),
	(62, '', '127.0.0.1:59502', 'reset_system', 1, '2021-07-21 16:11:33', ''),
	(63, '', '127.0.0.1:55761', 'reset_system', 1, '2021-07-21 16:11:34', ''),
	(64, '', '127.0.0.1:50679', 'reset_system', 1, '2021-07-21 16:11:34', ''),
	(65, '', '127.0.0.1:62706', 'reset_system', 1, '2021-07-21 16:11:34', ''),
	(66, '', '127.0.0.1:58766', 'reset_system', 1, '2021-07-21 16:11:34', ''),
	(67, '', '127.0.0.1:59672', 'reset_system', 1, '2021-07-21 16:11:35', ''),
	(68, '', '127.0.0.1:53665', 'reset_system', 1, '2021-07-21 16:11:35', ''),
	(69, '', '127.0.0.1:64638', 'reset_system', 1, '2021-07-21 16:11:35', ''),
	(70, '', '127.0.0.1:63532', 'reset_system', 1, '2021-07-21 16:11:35', ''),
	(71, '', '127.0.0.1:64227', 'reset_system', 1, '2021-07-21 16:11:35', ''),
	(72, '', '127.0.0.1:63731', 'reset_system', 1, '2021-07-21 16:11:35', ''),
	(73, '', '127.0.0.1:65050', 'reset_system', 1, '2021-07-21 16:11:35', ''),
	(74, '', '127.0.0.1:52339', 'reset_system', 1, '2021-07-21 16:11:36', ''),
	(75, '', '127.0.0.1:49175', 'reset_system', 1, '2021-07-21 16:11:36', ''),
	(76, '', '127.0.0.1:55903', 'reset_system', 1, '2021-07-21 16:11:36', ''),
	(77, '', '127.0.0.1:59779', 'reset_system', 1, '2021-07-21 16:11:36', ''),
	(78, '', '127.0.0.1:54194', 'reset_system', 1, '2021-07-21 16:11:36', ''),
	(79, '', '127.0.0.1:64674', 'reset_system', 1, '2021-07-21 16:11:36', ''),
	(80, '', '127.0.0.1:62719', 'reset_system', 1, '2021-07-21 16:11:37', ''),
	(81, '', '127.0.0.1:53036', 'reset_system', 1, '2021-07-21 16:11:37', ''),
	(82, '', '127.0.0.1:53018', 'reset_system', 1, '2021-07-21 16:11:37', ''),
	(83, '', '127.0.0.1:52102', 'reset_system', 1, '2021-07-21 16:11:37', ''),
	(84, '', '127.0.0.1:63306', 'reset_system', 1, '2021-07-21 16:11:37', ''),
	(85, '', '127.0.0.1:63459', 'reset_system', 1, '2021-07-21 16:11:38', ''),
	(86, '', '127.0.0.1:53981', 'reset_system', 1, '2021-07-21 16:11:38', ''),
	(87, '', '127.0.0.1:58889', 'reset_system', 1, '2021-07-21 16:11:38', ''),
	(88, '', '127.0.0.1:57234', 'reset_system', 1, '2021-07-21 16:11:38', ''),
	(89, '', '127.0.0.1:52622', 'reset_system', 1, '2021-07-21 16:11:38', ''),
	(90, '', '127.0.0.1:49981', 'reset_system', 1, '2021-07-21 16:11:38', ''),
	(91, '', '127.0.0.1:56984', 'reset_system', 1, '2021-07-21 16:11:39', ''),
	(92, '', '127.0.0.1:59366', 'reset_system', 1, '2021-07-21 16:11:39', ''),
	(93, '', '127.0.0.1:57074', 'reset_system', 1, '2021-07-21 16:11:39', ''),
	(94, '', '127.0.0.1:63202', 'reset_system', 0, '2021-07-21 16:21:43', '重置KEY不匹配'),
	(95, '', '127.0.0.1:64762', 'reset_system', 0, '2021-07-21 16:21:43', '重置KEY不匹配'),
	(96, '', '127.0.0.1:55503', 'reset_system', 0, '2021-07-21 16:21:43', '重置KEY不匹配'),
	(97, '', '127.0.0.1:58455', 'reset_system', 0, '2021-07-21 16:21:43', '重置KEY不匹配'),
	(98, '', '127.0.0.1:59602', 'reset_system', 0, '2021-07-21 16:21:44', '重置KEY不匹配'),
	(99, '', '127.0.0.1:59841', 'reset_system', 0, '2021-07-21 16:38:35', '重置KEY不匹配'),
	(100, '', '127.0.0.1:55936', 'reset_system', 0, '2021-07-21 16:38:51', '重置KEY不匹配'),
	(101, '', '127.0.0.1:56121', 'reset_system', 0, '2021-07-21 16:40:35', '重置KEY不匹配'),
	(102, '', '127.0.0.1:58610', 'reset_system', 0, '2021-07-21 16:40:38', '重置KEY不匹配'),
	(103, '', '127.0.0.1:59429', 'reset_system', 0, '2021-07-21 16:42:03', '重置KEY不匹配'),
	(104, '', '127.0.0.1:56131', 'reset_system', 0, '2021-07-21 16:46:04', 'load key error'),
	(105, '', '127.0.0.1:50641', 'reset_system', 0, '2021-07-21 16:46:07', 'load key error'),
	(106, '', '127.0.0.1:60932', 'reset_system', 0, '2021-07-21 16:48:32', 'WHERE conditions required'),
	(107, '', '127.0.0.1:52161', 'reset_system', 1, '2021-07-21 16:50:03', ''),
	(108, '', '127.0.0.1:49339', 'reset_system', 1, '2021-07-21 16:55:59', '操作成功'),
	(109, '', '127.0.0.1:54465', 'reset_system', 1, '2021-07-21 17:04:21', '操作成功'),
	(110, '', '127.0.0.1:54876', 'reset_system', 1, '2021-07-21 17:04:57', '操作成功'),
	(111, '', '127.0.0.1:57030', 'reset_system', 1, '2021-07-21 17:31:12', '操作成功'),
	(112, '', '127.0.0.1:50537', 'reset_system', 0, '2021-07-21 17:31:49', 'running, try again later.'),
	(113, '', '127.0.0.1:59721', 'reset_system', 1, '2021-07-21 17:32:08', '操作成功'),
	(114, '', '127.0.0.1:64326', 'reset_system', 1, '2021-07-21 17:32:08', '操作成功'),
	(115, '', '127.0.0.1:64906', 'reset_system', 0, '2021-07-21 17:32:17', 'running, try again later.'),
	(116, '', '127.0.0.1:63874', 'reset_system', 0, '2021-07-21 17:32:21', 'running, try again later.'),
	(117, '', '127.0.0.1:59897', 'reset_system', 0, '2021-07-21 17:32:24', 'running, try again later.'),
	(118, '', '127.0.0.1:55557', 'reset_system', 0, '2021-07-21 17:45:10', 'running, try again later.'),
	(119, '', '127.0.0.1:61723', 'reset_system', 0, '2021-07-21 17:46:28', 'running, try again later.'),
	(120, '', '127.0.0.1:61218', 'reset_system', 1, '2021-07-21 17:46:31', '操作成功'),
	(121, '', '127.0.0.1:56601', 'reset_system', 0, '2021-07-21 17:46:33', 'running, try again later.'),
	(122, '', '127.0.0.1:64371', 'reset_system', 0, '2021-07-21 17:46:35', 'running, try again later.'),
	(123, '', '127.0.0.1:58749', 'reset_system', 0, '2021-07-21 17:46:37', 'running, try again later.'),
	(124, '', '127.0.0.1:64771', 'reset_system', 0, '2021-07-21 17:46:39', 'running, try again later.'),
	(125, '', '127.0.0.1:54655', 'reset_system', 0, '2021-07-21 17:47:21', 'running, try again later.'),
	(126, '', '127.0.0.1:64240', 'reset_system', 0, '2021-07-21 17:47:22', 'running, try again later.'),
	(127, '', '127.0.0.1:50846', 'reset_system', 0, '2021-07-21 17:47:23', 'running, try again later.'),
	(128, '', '127.0.0.1:56911', 'reset_system', 1, '2021-07-21 17:47:23', '操作成功'),
	(129, '', '127.0.0.1:62747', 'reset_system', 1, '2021-07-21 17:47:26', '操作成功'),
	(130, '', '127.0.0.1:56761', 'reset_system', 1, '2021-07-21 17:47:35', '操作成功'),
	(131, '', '127.0.0.1:49499', 'reset_system', 1, '2021-07-21 17:47:36', '操作成功'),
	(132, '', '127.0.0.1:57048', 'reset_system', 1, '2021-07-21 17:47:43', '操作成功'),
	(133, '', '127.0.0.1:56037', 'reset_system', 1, '2021-07-21 17:47:44', '操作成功'),
	(134, '', '127.0.0.1:63334', 'reset_system', 1, '2021-07-21 17:49:27', '操作成功'),
	(135, '', '127.0.0.1:60555', 'reset_system', 1, '2021-07-21 17:49:30', '操作成功');
/*!40000 ALTER TABLE `operation_log` ENABLE KEYS */;

-- 导出  表 pro_cfg_manager.options 结构
DROP TABLE IF EXISTS `options`;
CREATE TABLE IF NOT EXISTS `options` (
  `id` int NOT NULL AUTO_INCREMENT,
  `opt_key` varchar(100) NOT NULL,
  `opt_value` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COMMENT='选项';

-- 正在导出表  pro_cfg_manager.options 的数据：~6 rows (大约)
DELETE FROM `options`;
/*!40000 ALTER TABLE `options` DISABLE KEYS */;
INSERT INTO `options` (`id`, `opt_key`, `opt_value`) VALUES
	(1, 'publish_at_null_subgroup', 'true'),
	(2, 'publish_at_remain_subgroup', 'true'),
	(3, 'publish_at_empty_nocreate_file', 'true'),
	(4, 'publish_jobs_also_ips', 'true'),
	(5, 'publish_jobs_also_reload_srv', 'true'),
	(6, 'publish_ips_also_reload_srv', 'true');
/*!40000 ALTER TABLE `options` ENABLE KEYS */;

-- 导出  表 pro_cfg_manager.relabels 结构
DROP TABLE IF EXISTS `relabels`;
CREATE TABLE IF NOT EXISTS `relabels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(200) NOT NULL,
  `code` text CHARACTER SET utf8 COLLATE utf8_general_ci,
  `enabled` tinyint NOT NULL DEFAULT '1',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8 COMMENT='重写标签的配置';

-- 正在导出表  pro_cfg_manager.relabels 的数据：~4 rows (大约)
DELETE FROM `relabels`;
/*!40000 ALTER TABLE `relabels` DISABLE KEYS */;
INSERT INTO `relabels` (`id`, `name`, `code`, `enabled`, `update_at`) VALUES
	(1, '空规则', '', 1, '2021-07-07 14:57:50');
/*!40000 ALTER TABLE `relabels` ENABLE KEYS */;

-- 导出  表 pro_cfg_manager.rules_groups 结构
DROP TABLE IF EXISTS `rules_groups`;
CREATE TABLE IF NOT EXISTS `rules_groups` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='监控规则组';

-- 正在导出表  pro_cfg_manager.rules_groups 的数据：~7 rows (大约)
DELETE FROM `rules_groups`;
/*!40000 ALTER TABLE `rules_groups` DISABLE KEYS */;
/*!40000 ALTER TABLE `rules_groups` ENABLE KEYS */;

-- 导出  表 pro_cfg_manager.sub_group 结构
DROP TABLE IF EXISTS `sub_group`;
CREATE TABLE IF NOT EXISTS `sub_group` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `rules_groups_id` int NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_name_rules_groups_id` (`name`,`rules_groups_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='监控规则组中的下级组，子组';

-- 正在导出表  pro_cfg_manager.sub_group 的数据：~18 rows (大约)
DELETE FROM `sub_group`;
/*!40000 ALTER TABLE `sub_group` DISABLE KEYS */;
/*!40000 ALTER TABLE `sub_group` ENABLE KEYS */;

-- 导出  表 pro_cfg_manager.test 结构
DROP TABLE IF EXISTS `test`;
CREATE TABLE IF NOT EXISTS `test` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;

-- 正在导出表  pro_cfg_manager.test 的数据：~4 rows (大约)
DELETE FROM `test`;
/*!40000 ALTER TABLE `test` DISABLE KEYS */;
INSERT INTO `test` (`id`, `name`) VALUES
	(1, '0'),
	(2, '1'),
	(3, '2'),
	(4, '3');
/*!40000 ALTER TABLE `test` ENABLE KEYS */;

-- 导出  表 pro_cfg_manager.tmpl 结构
DROP TABLE IF EXISTS `tmpl`;
CREATE TABLE IF NOT EXISTS `tmpl` (
  `tmpl` longtext NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='prometheus.yml';

-- 正在导出表  pro_cfg_manager.tmpl 的数据：~1 rows (大约)
DELETE FROM `tmpl`;
/*!40000 ALTER TABLE `tmpl` DISABLE KEYS */;
INSERT INTO `tmpl` (`tmpl`) VALUES
	('# my global config\nglobal:\n  scrape_interval:     15s # Set the scrape interval to every 15 seconds. Default is every 1 minute.\n  evaluation_interval: 15s # Evaluate rules every 15 seconds. The default is every 1 minute.\n  # scrape_timeout is set to the global default (10s).\n  # scrape_timeout: 60s\n\n# Alertmanager configuration\nalerting:\n  alertmanagers:\n  - static_configs:\n    - targets:\n      - 127.0.0.1:9093\n\n# Load rules once and periodically evaluate them according to the global \'evaluation_interval\'.\nrule_files:\n   - "{{.RelRuleDir}}/*.yml"\n  # - "first_rules.yml"\n  # - "second_rules.yml"\n\n# A scrape configuration containing exactly one endpoint to scrape:\n# Here it\'s Prometheus itself.\nscrape_configs:\n  # The job name is added as a label `job=<job_name>` to any timeseries scraped from this config.\n  - job_name: \'监控服务本机\'\n\n    # metrics_path defaults to \'/metrics\'\n    # scheme defaults to \'http\'.\n\n    static_configs:\n    - targets: [\'localhost:9090\']\n{{ range .Jobs }}\n  - job_name: \'{{.Name}}\'\n    file_sd_configs:\n      - files:\n        - "{{$.AbsConfDir}}/{{.Name}}.json"\n        refresh_interval: 15s\n{{.Code}}\n{{ end }}\n'),
	('# my global config\nglobal:\n  scrape_interval:     15s # Set the scrape interval to every 15 seconds. Default is every 1 minute.\n  evaluation_interval: 15s # Evaluate rules every 15 seconds. The default is every 1 minute.\n  # scrape_timeout is set to the global default (10s).\n  # scrape_timeout: 60s\n\n# Alertmanager configuration\nalerting:\n  alertmanagers:\n  - static_configs:\n    - targets:\n      - 127.0.0.1:9093\n\n# Load rules once and periodically evaluate them according to the global \'evaluation_interval\'.\nrule_files:\n   - "{{.RelRuleDir}}/*.yml"\n  # - "first_rules.yml"\n  # - "second_rules.yml"\n\n# A scrape configuration containing exactly one endpoint to scrape:\n# Here it\'s Prometheus itself.\nscrape_configs:\n  # The job name is added as a label `job=<job_name>` to any timeseries scraped from this config.\n  - job_name: \'监控服务本机\'\n\n    # metrics_path defaults to \'/metrics\'\n    # scheme defaults to \'http\'.\n\n    static_configs:\n    - targets: [\'localhost:9090\']\n{{ range .Jobs }}\n  - job_name: \'{{.Name}}\'\n    file_sd_configs:\n      - files:\n        - "{{$.AbsConfDir}}/{{.Name}}.json"\n        refresh_interval: 15s\n{{.Code}}\n{{ end }}\n');
/*!40000 ALTER TABLE `tmpl` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
