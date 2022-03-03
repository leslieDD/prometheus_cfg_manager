ALTER TABLE `line` DROP INDEX `uniq_label`,
  ADD UNIQUE INDEX `uniq_label` (`label`, `idc_id`) USING BTREE;
  
INSERT INTO `pro_cfg_manager`.`log_setting` (`level`, `label`, `selected`) VALUES ('error', '错误', '1');

