-- --------------------------------------------------------
-- 主机:                           192.168.1.7
-- 服务器版本:                        5.7.31-34-log - Percona Server (GPL), Release 34, Revision 2e68637
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
CREATE DATABASE IF NOT EXISTS `pro_cfg_manager` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `pro_cfg_manager`;

-- 导出  表 pro_cfg_manager.jobs 结构
CREATE TABLE IF NOT EXISTS `jobs` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(100) NOT NULL COMMENT '任务名称',
  `port` int(11) NOT NULL COMMENT '端口号，对应exporter的端口号',
  `cfg_name` varchar(100) DEFAULT NULL COMMENT '在prometheus下生成配置的文件名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='prometheus任务列表';

-- 正在导出表  pro_cfg_manager.jobs 的数据：~3 rows (大约)
DELETE FROM `jobs`;
/*!40000 ALTER TABLE `jobs` DISABLE KEYS */;
INSERT INTO `jobs` (`id`, `name`, `port`, `cfg_name`) VALUES
	(1, '云存储节点', 0, NULL),
	(2, '云空闲节点', 0, NULL),
	(3, '云计算节点', 0, NULL);
/*!40000 ALTER TABLE `jobs` ENABLE KEYS */;

-- 导出  表 pro_cfg_manager.machines 结构
CREATE TABLE IF NOT EXISTS `machines` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `ipaddr` varchar(100) NOT NULL COMMENT 'IP地址',
  `job_id` json NOT NULL COMMENT '所属任务列表',
  PRIMARY KEY (`id`),
  UNIQUE KEY `ipaddr_unique` (`ipaddr`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='机器列表';

-- 正在导出表  pro_cfg_manager.machines 的数据：~1 rows (大约)
DELETE FROM `machines`;
/*!40000 ALTER TABLE `machines` DISABLE KEYS */;
INSERT INTO `machines` (`id`, `ipaddr`, `job_id`) VALUES
	(1, '117.27.240.126', '[1]');
/*!40000 ALTER TABLE `machines` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
