INSERT INTO `pro_cfg_manager`.`page_function` (
    `page_name`,
    `page_nice_name`,
    `sub_page_name`,
    `sub_page_nice_name`,
    `func_name`,
    `func_nice_name`
  )
VALUES (
    'idc',
    'IDC机房',
    '',
    '',
    'expand_ipaddr',
    '扩展IP地址'
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
    'instance',
    '实例导入',
    '',
    '',
    'get_data',
    '获取实例数据'
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
    'instance',
    '实例导入',
    '',
    '',
    'put_data',
    '导入实例数据'
  );
insert options(opt_key, opt_value)
values('position_ipaddr', 'false');
insert options(opt_key, opt_value)
values('sync_prometheus_status', false);