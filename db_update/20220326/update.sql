ALTER TABLE `session`
	ADD COLUMN `login_at` DATETIME NOT NULL AFTER `update_at`;