INSERT INTO `pro_cfg_manager`.`page_function` (`page_name`, `page_nice_name`, `sub_page_name`, `sub_page_nice_name`, `func_name`, `func_nice_name`) VALUES ('jobs', 'JOB组管理', 'jobs', 'JOB组管理', 'get_job_mirror', '获取镜像');
INSERT INTO `pro_cfg_manager`.`page_function` (`page_name`, `page_nice_name`, `sub_page_name`, `sub_page_nice_name`, `func_name`, `func_nice_name`) VALUES ('jobs', 'JOB组管理', 'jobs', 'JOB组管理', 'update_job_mirror', '更新镜像');
CREATE TABLE `job_mirror` (
	`mirrors` TEXT NULL DEFAULT NULL COLLATE 'utf8mb3_general_ci',
	`update_at` DATETIME NOT NULL,
	`update_by` VARCHAR(50) NOT NULL COLLATE 'utf8mb3_general_ci',
	`job_id` INT(11) NOT NULL,
	PRIMARY KEY (`job_id`) USING BTREE
)
COMMENT='job组镜像'
COLLATE='utf8mb3_general_ci'
ENGINE=InnoDB
;
