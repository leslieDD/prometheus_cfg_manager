ALTER TABLE `crontab`
    ADD COLUMN `cid` VARCHAR(200) DEFAULT NULL AFTER `name`;

ALTER TABLE `crontab`
    ADD COLUMN `line_title` VARCHAR(200) DEFAULT NULL AFTER `show_title`;

INSERT INTO `pro_cfg_manager`.`page_function` (
    `page_name`,
    `page_nice_name`,
    `sub_page_name`,
    `sub_page_nice_name`,
    `func_name`,
    `func_nice_name`
)
VALUES (
   'crontab',
   '定时任务',
   '',
   '',
   'load_title',
   '加载线标题'
);