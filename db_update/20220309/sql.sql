INSERT INTO `pro_cfg_manager`.`page_function` (
    `page_name`,
    `page_nice_name`,
    `sub_page_name`,
    `sub_page_nice_name`,
    `func_name`,
    `func_nice_name`
  )
VALUES (
    'admin',
    '管理中心',
    'session',
    '会话管理',
    'search',
    '获取用户会话'
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
    'admin',
    '管理中心',
    'session',
    '会话管理',
    'delete',
    '删除用户会话'
  );
ALTER TABLE `session`
ADD COLUMN `ipaddr` VARCHAR(120) NULL DEFAULT NULL
AFTER `token`;