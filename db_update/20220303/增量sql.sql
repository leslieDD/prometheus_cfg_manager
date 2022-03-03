ALTER TABLE `line` DROP INDEX `uniq_label`,
  ADD UNIQUE INDEX `uniq_label` (`label`, `idc_id`) USING BTREE;

INSERT INTO `pro_cfg_manager`.`log_setting` (`level`, `label`, `selected`) VALUES ('error', '错误', '1');

INSERT INTO `pro_cfg_manager`.`page_function` (`page_name`, `page_nice_name`, `func_name`, `func_nice_name`) VALUES ('idc', 'IDC机房', 'get_idc', '获取机房列表');
INSERT INTO `pro_cfg_manager`.`page_function` (`page_name`, `page_nice_name`, `func_name`, `func_nice_name`) VALUES ('idc', 'IDC机房', 'add_idc', '新建机房');
INSERT INTO `pro_cfg_manager`.`page_function` (`page_name`, `page_nice_name`, `func_name`, `func_nice_name`) VALUES ('idc', 'IDC机房', 'update_idc', '更新机房');
INSERT INTO `pro_cfg_manager`.`page_function` (`page_name`, `page_nice_name`, `func_name`, `func_nice_name`) VALUES ('idc', 'IDC机房', 'del_idc', '删除机房');
INSERT INTO `pro_cfg_manager`.`page_function` (`page_name`, `page_nice_name`, `func_name`, `func_nice_name`) VALUES ('idc', 'IDC机房', 'get_idc_tree', '获取机房列表');
INSERT INTO `pro_cfg_manager`.`page_function` (`page_name`, `page_nice_name`, `func_name`, `func_nice_name`) VALUES ('idc', 'IDC机房', 'get_line', '获取线路信息');
INSERT INTO `pro_cfg_manager`.`page_function` (`page_name`, `page_nice_name`, `func_name`, `func_nice_name`) VALUES ('idc', 'IDC机房', 'add_line', '新建线路');
INSERT INTO `pro_cfg_manager`.`page_function` (`page_name`, `page_nice_name`, `func_name`, `func_nice_name`) VALUES ('idc', 'IDC机房', 'update_line', '更新线路');
INSERT INTO `pro_cfg_manager`.`page_function` (`page_name`, `page_nice_name`, `func_name`, `func_nice_name`) VALUES ('idc', 'IDC机房', 'del_line', '删除线路');
INSERT INTO `pro_cfg_manager`.`page_function` (`page_name`, `page_nice_name`, `func_name`, `func_nice_name`) VALUES ('idc', 'IDC机房', 'get_line_ipaddr', '获取线路地址池');
INSERT INTO `pro_cfg_manager`.`page_function` (`page_name`, `page_nice_name`, `func_name`, `func_nice_name`) VALUES ('idc', 'IDC机房', 'update_line_ipaddr', '更新线路地址池');
INSERT INTO `pro_cfg_manager`.`page_function` (`page_name`, `page_nice_name`, `func_name`, `func_nice_name`) VALUES ('idc', 'IDC机房', 'update_label_part', '只更新未设置IP');
INSERT INTO `pro_cfg_manager`.`page_function` (`page_name`, `page_nice_name`, `func_name`, `func_nice_name`) VALUES ('idc', 'IDC机房', 'update_label_all', '更新所有IP');
INSERT INTO `pro_cfg_manager`.`page_function` (`page_name`, `page_nice_name`, `func_name`, `func_nice_name`) VALUES ('idc', 'IDC机房', 'create_label_for_all_job', '在JOB组中生成标签');