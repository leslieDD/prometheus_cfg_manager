INSERT INTO `pro_cfg_manager`.`page_function` (`page_name`, `page_nice_name`, `sub_page_name`, `sub_page_nice_name`, `func_name`, `func_nice_name`) VALUES ('admin', '管理中心', 'session', '会话管理', 'update', '更新会话参数');
CREATE TABLE `number_options` (
	`key` VARCHAR(50) NOT NULL COLLATE 'utf8mb3_general_ci',
	`value` INT(11) NOT NULL DEFAULT '0',
	`commit` VARCHAR(100) NULL DEFAULT NULL COLLATE 'utf8mb3_general_ci',
	UNIQUE INDEX `uniq_key` (`key`) USING BTREE
)
COMMENT='值为数据的选项'
COLLATE='utf8mb3_general_ci'
ENGINE=InnoDB
;
