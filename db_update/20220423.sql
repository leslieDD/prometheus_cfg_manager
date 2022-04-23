ALTER TABLE `job_machines`
ADD COLUMN `blacked` INT NULL DEFAULT '0' COMMENT '是否禁用，在本组中'
AFTER `job_id`;
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
    'job_search_black',
    '查询组黑名单'
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
    'job_update_black',
    '更新组黑名单'
  );