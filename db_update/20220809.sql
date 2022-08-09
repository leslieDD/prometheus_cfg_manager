ALTER TABLE `monitor_rules`
	ADD COLUMN `api_enabled` TINYINT NOT NULL DEFAULT '0' AFTER `enabled`,
	ADD COLUMN `api_content` VARCHAR(200) NOT NULL DEFAULT '' AFTER `api_enabled`;
