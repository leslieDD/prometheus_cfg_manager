ALTER TABLE `line` DROP INDEX `uniq_label`,
  ADD UNIQUE INDEX `uniq_label` (`label`, `idc_id`) USING BTREE;