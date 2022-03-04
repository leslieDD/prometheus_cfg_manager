ALTER TABLE `idc`
	ADD COLUMN `remark` TEXT NULL COMMENT '备注' AFTER `label`;
ALTER TABLE `pool`
	ADD COLUMN `remark` TEXT NULL AFTER `line_id`;