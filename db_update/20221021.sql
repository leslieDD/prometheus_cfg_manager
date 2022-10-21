ALTER TABLE `crontab`
    ADD COLUMN `status` TINYINT(4) NOT NULL DEFAULT '0' AFTER `unit`;

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
           'check_status',
           '检索状态'
       );