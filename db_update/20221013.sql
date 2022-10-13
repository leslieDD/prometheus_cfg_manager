ALTER TABLE `crontab`
	ADD COLUMN `show_title` TINYINT(4) NOT NULL DEFAULT '0' AFTER `enabled`;