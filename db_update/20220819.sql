CREATE TABLE `crontab_api` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`name` VARCHAR(300) NOT NULL COLLATE 'utf8mb3_general_ci',
	`api` VARCHAR(300) NOT NULL COLLATE 'utf8mb3_general_ci',
	`update_at` DATETIME NOT NULL,
	`update_by` VARCHAR(200) NOT NULL DEFAULT '' COLLATE 'utf8mb3_general_ci',
	PRIMARY KEY (`id`) USING BTREE
)
COMMENT='定义任务调用的API列表'
COLLATE='utf8mb3_general_ci'
ENGINE=InnoDB
;

INSERT INTO `pro_cfg_manager`.`page_function` (`page_name`, `page_nice_name`, `sub_page_name`, `sub_page_nice_name`, `func_name`, `func_nice_name`) VALUES ('baseConfig', '基本配置', 'cronapi', '定义任务API', 'search', '查询');
INSERT INTO `pro_cfg_manager`.`page_function` (`page_name`, `page_nice_name`, `sub_page_name`, `sub_page_nice_name`, `func_name`, `func_nice_name`) VALUES ('baseConfig', '基本配置', 'cronapi', '定时任务API', 'add', '增加');
INSERT INTO `pro_cfg_manager`.`page_function` (`page_name`, `page_nice_name`, `sub_page_name`, `sub_page_nice_name`, `func_name`, `func_nice_name`) VALUES ('baseConfig', '基本配置', 'cronapi', '定时任务API', 'update', '更新');
INSERT INTO `pro_cfg_manager`.`page_function` (`page_name`, `page_nice_name`, `sub_page_name`, `sub_page_nice_name`, `func_name`, `func_nice_name`) VALUES ('baseConfig', '基本配置', 'cronapi', '定时任务API', 'dis.enable', '启用/禁用');
INSERT INTO `pro_cfg_manager`.`page_function` (`page_name`, `page_nice_name`, `sub_page_name`, `sub_page_nice_name`, `func_name`, `func_nice_name`) VALUES ('baseConfig', '基本配置', 'cronapi', '定时任务API', 'delete', '删除');

CREATE TABLE `crontab` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`name` VARCHAR(300) NOT NULL DEFAULT '' COLLATE 'utf8mb3_general_ci',
	`rule` TEXT NOT NULL COLLATE 'utf8mb3_general_ci',
	`api_id` INT(11) NOT NULL,
	`enabled` TINYINT(4) NOT NULL DEFAULT '1',
	`update_at` DATETIME NOT NULL,
	`update_by` VARCHAR(300) NOT NULL DEFAULT '' COLLATE 'utf8mb3_general_ci',
	PRIMARY KEY (`id`) USING BTREE
)
COMMENT='定时任务'
COLLATE='utf8mb3_general_ci'
ENGINE=InnoDB
;


INSERT INTO `pro_cfg_manager`.`page_function` (`page_name`, `page_nice_name`, `func_name`, `func_nice_name`) VALUES ('crontab', '定时任务', 'search', '查询');
INSERT INTO `pro_cfg_manager`.`page_function` (`page_name`, `page_nice_name`, `func_name`, `func_nice_name`) VALUES ('crontab', '定时任务', 'add', '增加');
INSERT INTO `pro_cfg_manager`.`page_function` (`page_name`, `page_nice_name`, `func_name`, `func_nice_name`) VALUES ('crontab', '定时任务', 'update', '更新');
INSERT INTO `pro_cfg_manager`.`page_function` (`page_name`, `page_nice_name`, `func_name`, `func_nice_name`) VALUES ('crontab', '定时任务', 'delete', '删除');
INSERT INTO `pro_cfg_manager`.`page_function` (`page_name`, `page_nice_name`, `func_name`, `func_nice_name`) VALUES ('crontab', '定时任务', 'dis.enable', '启用/禁用');
ALTER TABLE `crontab`
	ADD COLUMN `exec_cycle` INT NOT NULL DEFAULT '5' AFTER `api_id`;