INSERT INTO `pro_cfg_manager`.`page_function` (`page_name`, `page_nice_name`, `sub_page_name`, `sub_page_nice_name`, `func_name`, `func_nice_name`) VALUES ('admin', '管理中心', 'session', '会话管理', 'update', '更新会话参数');
ALTER TABLE `number_options`
	ADD UNIQUE INDEX `uniq_key` (`key`);
