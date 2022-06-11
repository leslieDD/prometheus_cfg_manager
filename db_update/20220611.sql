CREATE TABLE `job_labels` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(200) NOT NULL COLLATE 'utf8mb3_general_ci',
  `value` VARCHAR(200) NOT NULL COLLATE 'utf8mb3_general_ci',
  `update_at` DATETIME NOT NULL,
  `update_by` VARCHAR(50) NOT NULL COLLATE 'utf8mb3_general_ci',
  `job_id` INT(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) COMMENT = 'job组的全局标签' COLLATE = 'utf8mb3_general_ci' ENGINE = InnoDB;
INSERT INTO `pro_cfg_manager`.`page_function` (
    `page_name`,
    `page_nice_name`,
    `sub_page_name`,
    `sub_page_nice_name`,
    `func_name`,
    `func_nice_name`
  )
VALUES (
    'jobs',
    'JOB组管理',
    'jobs',
    'JOB组管理',
    'get_job_label',
    '获取组标签'
  );
INSERT INTO `pro_cfg_manager`.`page_function` (
    `page_name`,
    `page_nice_name`,
    `sub_page_name`,
    `sub_page_nice_name`,
    `func_name`,
    `func_nice_name`
  )
VALUES (
    'jobs',
    'JOB组管理',
    'jobs',
    'JOB组管理',
    'update_job_label',
    '更新组标签'
  );