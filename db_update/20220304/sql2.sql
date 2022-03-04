ALTER TABLE `idc`
	ADD COLUMN `remark` TEXT NULL COMMENT '备注' AFTER `label`;
ALTER TABLE `pool`
	ADD COLUMN `remark` TEXT NULL AFTER `line_id`;


ALTER TABLE `job_group`
	ADD UNIQUE INDEX `uniq_name_jobs_id` (`jobs_id`, `name`);
ALTER TABLE `group_machines`
	DROP COLUMN `id`;
