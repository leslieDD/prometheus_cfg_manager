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
CREATE DATABASE IF NOT EXISTS `pro_cfg_manager` /*!40100 DEFAULT CHARACTER SET utf8 */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `pro_cfg_manager`;

-- 导出  表 pro_cfg_manager.annotations 结构
CREATE TABLE IF NOT EXISTS `annotations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `monitor_rules_id` int NOT NULL,
  `key` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `value` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_monitor_rules_id_key` (`monitor_rules_id`,`key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='监控规则的注释';

-- 数据导出被取消选择。

-- 导出  表 pro_cfg_manager.group_labels 结构
CREATE TABLE IF NOT EXISTS `group_labels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `job_group_id` int NOT NULL,
  `key` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `value` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `enabled` tinyint NOT NULL DEFAULT '1',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_by` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_groupid_key` (`key`,`job_group_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='子组标签列表';

-- 数据导出被取消选择。

-- 导出  表 pro_cfg_manager.group_machines 结构
CREATE TABLE IF NOT EXISTS `group_machines` (
  `id` int NOT NULL AUTO_INCREMENT,
  `job_group_id` int NOT NULL,
  `machines_id` int NOT NULL,
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique` (`machines_id`,`job_group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='子组IP列表';

-- 数据导出被取消选择。

-- 导出  表 pro_cfg_manager.group_priv 结构
CREATE TABLE IF NOT EXISTS `group_priv` (
  `id` int NOT NULL AUTO_INCREMENT,
  `group_id` int NOT NULL,
  `func_id` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=99 DEFAULT CHARSET=utf8;

-- 数据导出被取消选择。

-- 导出  表 pro_cfg_manager.jobs 结构
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
  `update_by` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_unique` (`name`),
  KEY `order_unique` (`display_order`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='prometheus任务列表';

-- 数据导出被取消选择。

-- 导出  表 pro_cfg_manager.job_group 结构
CREATE TABLE IF NOT EXISTS `job_group` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `jobs_id` int NOT NULL,
  `enabled` tinyint NOT NULL DEFAULT '1',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_by` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='还可以为单个分组中的IP地址进行分子组，为每个子组设置相应的标签列表';

-- 数据导出被取消选择。

-- 导出  表 pro_cfg_manager.job_machines 结构
CREATE TABLE IF NOT EXISTS `job_machines` (
  `machine_id` int NOT NULL,
  `job_id` int NOT NULL,
  UNIQUE KEY `unique_mj` (`job_id`,`machine_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 数据导出被取消选择。

-- 导出  表 pro_cfg_manager.labels 结构
CREATE TABLE IF NOT EXISTS `labels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `label` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `enabled` tinyint NOT NULL DEFAULT '1',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_by` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_labels` (`label`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='标签列表';

-- 数据导出被取消选择。

-- 导出  表 pro_cfg_manager.log_setting 结构
CREATE TABLE IF NOT EXISTS `log_setting` (
  `id` int NOT NULL AUTO_INCREMENT,
  `level` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `label` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `selected` tinyint NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8 COMMENT='日志设置';

-- 数据导出被取消选择。

-- 导出  表 pro_cfg_manager.machines 结构
CREATE TABLE IF NOT EXISTS `machines` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `ipaddr` varchar(100) NOT NULL COMMENT 'IP地址',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `enabled` tinyint NOT NULL DEFAULT '1',
  `update_by` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `ipaddr_unique` (`ipaddr`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='机器列表';

-- 数据导出被取消选择。

-- 导出  表 pro_cfg_manager.manager_group 结构
CREATE TABLE IF NOT EXISTS `manager_group` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `enabled` tinyint NOT NULL,
  `update_at` datetime NOT NULL,
  `update_by` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_unique` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='用户组';

-- 数据导出被取消选择。

-- 导出  表 pro_cfg_manager.manager_set 结构
CREATE TABLE IF NOT EXISTS `manager_set` (
  `param_name` varbinary(100) NOT NULL,
  `param_value` varbinary(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='权限管理默认的设置参数';

-- 数据导出被取消选择。

-- 导出  表 pro_cfg_manager.manager_user 结构
CREATE TABLE IF NOT EXISTS `manager_user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(100) NOT NULL,
  `nice_name` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `phone` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `salt` varchar(100) NOT NULL,
  `group_id` int NOT NULL,
  `update_at` datetime NOT NULL,
  `enabled` tinyint NOT NULL,
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_by` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_name` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='用户';

-- 数据导出被取消选择。

-- 导出  表 pro_cfg_manager.monitor_labels 结构
CREATE TABLE IF NOT EXISTS `monitor_labels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `monitor_rules_id` int NOT NULL,
  `key` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `value` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_mid_lid` (`key`,`monitor_rules_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='监控规则的标签';

-- 数据导出被取消选择。

-- 导出  表 pro_cfg_manager.monitor_rules 结构
CREATE TABLE IF NOT EXISTS `monitor_rules` (
  `id` int NOT NULL AUTO_INCREMENT,
  `alert` varchar(100) NOT NULL,
  `expr` varchar(5000) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `for` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `sub_group_id` int NOT NULL,
  `enabled` tinyint NOT NULL,
  `description` varchar(300) NOT NULL,
  `update_by` varchar(100) NOT NULL,
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='具体的监控规则';

-- 数据导出被取消选择。

-- 导出  表 pro_cfg_manager.operation_log 结构
CREATE TABLE IF NOT EXISTS `operation_log` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(100) NOT NULL,
  `operate_type` varchar(50) NOT NULL,
  `ipaddr` varchar(100) NOT NULL,
  `operate_content` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `operate_result` tinyint NOT NULL,
  `operate_at` datetime NOT NULL,
  `operate_error` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=149 DEFAULT CHARSET=utf8 COMMENT='操作日志';

-- 数据导出被取消选择。

-- 导出  表 pro_cfg_manager.options 结构
CREATE TABLE IF NOT EXISTS `options` (
  `id` int NOT NULL AUTO_INCREMENT,
  `opt_key` varchar(100) NOT NULL,
  `opt_value` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COMMENT='选项';

-- 数据导出被取消选择。

-- 导出  表 pro_cfg_manager.page_function 结构
CREATE TABLE IF NOT EXISTS `page_function` (
  `id` int NOT NULL AUTO_INCREMENT,
  `page_name` varchar(100) NOT NULL,
  `page_nice_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `sub_page_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '',
  `sub_page_nice_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '',
  `func_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `func_nice_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=101 DEFAULT CHARSET=utf8 COMMENT='页面功能';

-- 数据导出被取消选择。

-- 导出  表 pro_cfg_manager.relabels 结构
CREATE TABLE IF NOT EXISTS `relabels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(200) NOT NULL,
  `code` text CHARACTER SET utf8 COLLATE utf8_general_ci,
  `enabled` tinyint NOT NULL DEFAULT '1',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_by` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8 COMMENT='重写标签的配置';

-- 数据导出被取消选择。

-- 导出  表 pro_cfg_manager.rules_groups 结构
CREATE TABLE IF NOT EXISTS `rules_groups` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_by` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='监控规则组';

-- 数据导出被取消选择。

-- 导出  表 pro_cfg_manager.session 结构
CREATE TABLE IF NOT EXISTS `session` (
  `id` int NOT NULL AUTO_INCREMENT,
  `token` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `user_id` int NOT NULL,
  `update_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=95 DEFAULT CHARSET=utf8 COMMENT='用户登录会话';

-- 数据导出被取消选择。

-- 导出  表 pro_cfg_manager.sub_group 结构
CREATE TABLE IF NOT EXISTS `sub_group` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `rules_groups_id` int NOT NULL,
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_by` varchar(50) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_name_rules_groups_id` (`name`,`rules_groups_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='监控规则组中的下级组，子组';

-- 数据导出被取消选择。

-- 导出  表 pro_cfg_manager.system_log 结构
CREATE TABLE IF NOT EXISTS `system_log` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `operate_type` varchar(50) NOT NULL,
  `ipaddr` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `operate_content` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `operate_result` tinyint NOT NULL,
  `operate_at` datetime NOT NULL,
  `operate_error` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=392 DEFAULT CHARSET=utf8 COMMENT='系统日志';

-- 数据导出被取消选择。

-- 导出  表 pro_cfg_manager.tmpl 结构
CREATE TABLE IF NOT EXISTS `tmpl` (
  `tmpl` longtext NOT NULL,
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_by` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='prometheus.yml';

-- 数据导出被取消选择。

-- 导出  表 pro_cfg_manager.tmpl_fields 结构
CREATE TABLE IF NOT EXISTS `tmpl_fields` (
  `id` int NOT NULL AUTO_INCREMENT,
  `key` varchar(100) NOT NULL,
  `value` varchar(500) NOT NULL,
  `enabled` tinyint NOT NULL,
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_by` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `key_unique` (`key`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='模板字段';

-- 数据导出被取消选择。

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
