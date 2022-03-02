-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        10.7.3-MariaDB - mariadb.org binary distribution
-- 服务器操作系统:                      Win64
-- HeidiSQL 版本:                  11.3.0.6295
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- 导出 pro_cfg_manager 的数据库结构
CREATE DATABASE IF NOT EXISTS `pro_cfg_manager` /*!40100 DEFAULT CHARACTER SET utf8mb3 */;
USE `pro_cfg_manager`;

-- 导出  表 pro_cfg_manager.idc 结构
CREATE TABLE IF NOT EXISTS `idc` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `label` varchar(200) NOT NULL,
  `enabled` tinyint(4) NOT NULL DEFAULT 1,
  `update_at` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `update_by` varchar(200) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_label` (`label`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb3 COMMENT='机房';

-- 正在导出表  pro_cfg_manager.idc 的数据：~5 rows (大约)
/*!40000 ALTER TABLE `idc` DISABLE KEYS */;
REPLACE INTO `idc` (`id`, `label`, `enabled`, `update_at`, `update_by`) VALUES
	(1, '机房联通', 1, '2022-03-01 16:45:42', 'admin'),
	(4, 'ddd', 1, '2022-03-01 17:18:31', 'admin'),
	(5, 'ddd2vvv', 1, '2022-03-01 17:58:31', 'admin');
/*!40000 ALTER TABLE `idc` ENABLE KEYS */;

-- 导出  表 pro_cfg_manager.line 结构
CREATE TABLE IF NOT EXISTS `line` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `label` varchar(200) NOT NULL,
  `enabled` tinyint(4) NOT NULL DEFAULT 1,
  `update_at` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `update_by` varchar(200) NOT NULL,
  `idc_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_label` (`label`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb3 COMMENT='线路，从属于机房(idc)';

-- 正在导出表  pro_cfg_manager.line 的数据：~5 rows (大约)
/*!40000 ALTER TABLE `line` DISABLE KEYS */;
REPLACE INTO `line` (`id`, `label`, `enabled`, `update_at`, `update_by`, `idc_id`) VALUES
	(5, '1234', 1, '2022-03-01 17:26:25', 'admin', 1),
	(6, '12342323', 1, '2022-03-01 17:26:25', 'admin', 1);
/*!40000 ALTER TABLE `line` ENABLE KEYS */;

-- 导出  表 pro_cfg_manager.pool 结构
CREATE TABLE IF NOT EXISTS `pool` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `line_id` int(11) NOT NULL COMMENT '线路ID',
  `ipaddrs` text DEFAULT NULL COMMENT '地址池，以英文分号隔开',
  `update_at` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `update_by` varchar(200) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_line_id` (`line_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb3 COMMENT='线路的地址池，从属于line表';

-- 正在导出表  pro_cfg_manager.pool 的数据：~1 rows (大约)
/*!40000 ALTER TABLE `pool` DISABLE KEYS */;
REPLACE INTO `pool` (`id`, `line_id`, `ipaddrs`, `update_at`, `update_by`) VALUES
	(4, 5, '218.85.157.0/24', '2022-02-03 14:29:29', 'admin'),
	(5, 6, '8.8.0.0/16', '2022-02-03 17:58:48', 'admin');
/*!40000 ALTER TABLE `pool` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
