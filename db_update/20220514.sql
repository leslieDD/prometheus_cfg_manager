ALTER TABLE `line`
ADD COLUMN `expand` TINYINT NOT NULL DEFAULT 0
AFTER `idc_id`;
ALTER TABLE `pool`
ADD COLUMN `expand` TINYINT NOT NULL DEFAULT 0
AFTER `ipaddrs`;
ALTER TABLE `idc`
ADD COLUMN `expand` TINYINT NOT NULL DEFAULT '0'
AFTER `enabled`;
ALTER TABLE `idc`
ADD COLUMN `view` TINYINT NOT NULL DEFAULT 0
AFTER `expand`;
ALTER TABLE `line`
ADD COLUMN `view` TINYINT NOT NULL DEFAULT '0'
AFTER `expand`;
ALTER TABLE `pool`
ADD COLUMN `view` TINYINT(4) NOT NULL DEFAULT '0'
AFTER `expand`;
INSERT INTO `pro_cfg_manager`.`page_function` (
    `page_name`,
    `page_nice_name`,
    `func_name`,
    `func_nice_name`
  )
VALUES ('idc', 'IDC机房', 'get_idc_tree_xls', '导出');