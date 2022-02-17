-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        10.7.3-MariaDB - mariadb.org binary distribution
-- 服务器操作系统:                      Win64
-- HeidiSQL 版本:                  11.3.0.6295
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

-- 正在导出表  pro_cfg_manager.annotations 的数据：~0 rows (大约)
DELETE FROM `annotations`;
/*!40000 ALTER TABLE `annotations` DISABLE KEYS */;
/*!40000 ALTER TABLE `annotations` ENABLE KEYS */;

-- 正在导出表  pro_cfg_manager.group_labels 的数据：~0 rows (大约)
DELETE FROM `group_labels`;
/*!40000 ALTER TABLE `group_labels` DISABLE KEYS */;
/*!40000 ALTER TABLE `group_labels` ENABLE KEYS */;

-- 正在导出表  pro_cfg_manager.group_machines 的数据：~0 rows (大约)
DELETE FROM `group_machines`;
/*!40000 ALTER TABLE `group_machines` DISABLE KEYS */;
/*!40000 ALTER TABLE `group_machines` ENABLE KEYS */;

-- 正在导出表  pro_cfg_manager.group_priv 的数据：~98 rows (大约)
DELETE FROM `group_priv`;
/*!40000 ALTER TABLE `group_priv` DISABLE KEYS */;
INSERT INTO `group_priv` (`id`, `group_id`, `func_id`) VALUES
	(399, 1, 74),
	(400, 1, 77),
	(401, 1, 76),
	(402, 1, 75),
	(403, 1, 78),
	(404, 1, 98),
	(405, 1, 100),
	(406, 1, 99),
	(407, 1, 80),
	(408, 1, 79),
	(409, 1, 86),
	(410, 1, 87),
	(411, 1, 85),
	(412, 1, 70),
	(413, 1, 72),
	(414, 1, 73),
	(415, 1, 69),
	(416, 1, 71),
	(417, 1, 95),
	(418, 1, 96),
	(419, 1, 93),
	(420, 1, 97),
	(421, 1, 94),
	(422, 1, 37),
	(423, 1, 33),
	(424, 1, 36),
	(425, 1, 35),
	(426, 1, 34),
	(427, 1, 50),
	(428, 1, 49),
	(429, 1, 45),
	(430, 1, 48),
	(431, 1, 44),
	(432, 1, 47),
	(433, 1, 46),
	(434, 1, 51),
	(435, 1, 52),
	(436, 1, 53),
	(437, 1, 31),
	(438, 1, 30),
	(439, 1, 90),
	(440, 1, 40),
	(441, 1, 43),
	(442, 1, 39),
	(443, 1, 42),
	(444, 1, 38),
	(445, 1, 41),
	(446, 1, 101),
	(447, 1, 103),
	(448, 1, 102),
	(449, 1, 83),
	(450, 1, 67),
	(451, 1, 8),
	(452, 1, 5),
	(453, 1, 105),
	(454, 1, 9),
	(455, 1, 6),
	(456, 1, 88),
	(457, 1, 7),
	(458, 1, 12),
	(459, 1, 81),
	(460, 1, 13),
	(461, 1, 14),
	(462, 1, 11),
	(463, 1, 15),
	(464, 1, 16),
	(465, 1, 26),
	(466, 1, 23),
	(467, 1, 20),
	(468, 1, 29),
	(469, 1, 27),
	(470, 1, 24),
	(471, 1, 21),
	(472, 1, 19),
	(473, 1, 28),
	(474, 1, 25),
	(475, 1, 22),
	(476, 1, 89),
	(477, 1, 104),
	(478, 1, 56),
	(479, 1, 59),
	(480, 1, 62),
	(481, 1, 65),
	(482, 1, 55),
	(483, 1, 58),
	(484, 1, 61),
	(485, 1, 64),
	(486, 1, 54),
	(487, 1, 57),
	(488, 1, 60),
	(489, 1, 63),
	(490, 1, 1),
	(491, 1, 2),
	(492, 1, 3),
	(493, 1, 66),
	(494, 1, 4),
	(495, 1, 68),
	(496, 1, 84);
/*!40000 ALTER TABLE `group_priv` ENABLE KEYS */;

-- 正在导出表  pro_cfg_manager.jobs 的数据：~0 rows (大约)
DELETE FROM `jobs`;
/*!40000 ALTER TABLE `jobs` DISABLE KEYS */;
/*!40000 ALTER TABLE `jobs` ENABLE KEYS */;

-- 正在导出表  pro_cfg_manager.job_group 的数据：~0 rows (大约)
DELETE FROM `job_group`;
/*!40000 ALTER TABLE `job_group` DISABLE KEYS */;
/*!40000 ALTER TABLE `job_group` ENABLE KEYS */;

-- 正在导出表  pro_cfg_manager.job_machines 的数据：~0 rows (大约)
DELETE FROM `job_machines`;
/*!40000 ALTER TABLE `job_machines` DISABLE KEYS */;
/*!40000 ALTER TABLE `job_machines` ENABLE KEYS */;

-- 正在导出表  pro_cfg_manager.labels 的数据：~0 rows (大约)
DELETE FROM `labels`;
/*!40000 ALTER TABLE `labels` DISABLE KEYS */;
/*!40000 ALTER TABLE `labels` ENABLE KEYS */;

-- 正在导出表  pro_cfg_manager.log_setting 的数据：~8 rows (大约)
DELETE FROM `log_setting`;
/*!40000 ALTER TABLE `log_setting` DISABLE KEYS */;
INSERT INTO `log_setting` (`id`, `level`, `label`, `selected`) VALUES
	(1, 'search', '查询', 0),
	(2, 'add', '增加', 1),
	(3, 'del', '删除', 1),
	(4, 'update', '更新', 1),
	(5, 'running', '运行', 1),
	(6, 'publish', '发布', 1),
	(7, 'login', '登录', 1),
	(8, 'reset', '重置', 1);
/*!40000 ALTER TABLE `log_setting` ENABLE KEYS */;

-- 正在导出表  pro_cfg_manager.machines 的数据：~0 rows (大约)
DELETE FROM `machines`;
/*!40000 ALTER TABLE `machines` DISABLE KEYS */;
/*!40000 ALTER TABLE `machines` ENABLE KEYS */;

-- 正在导出表  pro_cfg_manager.manager_group 的数据：~0 rows (大约)
DELETE FROM `manager_group`;
/*!40000 ALTER TABLE `manager_group` DISABLE KEYS */;
INSERT INTO `manager_group` (`id`, `name`, `enabled`, `update_at`, `update_by`) VALUES
	(1, 'administrator', 1, '2022-02-17 11:51:20', 'admin');
/*!40000 ALTER TABLE `manager_group` ENABLE KEYS */;

-- 正在导出表  pro_cfg_manager.manager_set 的数据：~0 rows (大约)
DELETE FROM `manager_set`;
/*!40000 ALTER TABLE `manager_set` DISABLE KEYS */;
INSERT INTO `manager_set` (`param_name`, `param_value`) VALUES
	(_binary 0x64656661756c745f67726f7570, _binary '');
/*!40000 ALTER TABLE `manager_set` ENABLE KEYS */;

-- 正在导出表  pro_cfg_manager.manager_user 的数据：~0 rows (大约)
DELETE FROM `manager_user`;
/*!40000 ALTER TABLE `manager_user` DISABLE KEYS */;
INSERT INTO `manager_user` (`id`, `username`, `nice_name`, `password`, `phone`, `salt`, `group_id`, `update_at`, `enabled`, `create_at`, `update_by`) VALUES
	(1, 'admin', '管理员', 'cc7550fb9f1b75d84b3677fdcd9d4c9f', '10086', 'ec7ceb50-c6c0-43b2-994e-b79cdb365457', 1, '2021-08-10 17:33:24', 1, '2021-08-10 17:33:24', '');
/*!40000 ALTER TABLE `manager_user` ENABLE KEYS */;

-- 正在导出表  pro_cfg_manager.monitor_labels 的数据：~0 rows (大约)
DELETE FROM `monitor_labels`;
/*!40000 ALTER TABLE `monitor_labels` DISABLE KEYS */;
/*!40000 ALTER TABLE `monitor_labels` ENABLE KEYS */;

-- 正在导出表  pro_cfg_manager.monitor_rules 的数据：~0 rows (大约)
DELETE FROM `monitor_rules`;
/*!40000 ALTER TABLE `monitor_rules` DISABLE KEYS */;
/*!40000 ALTER TABLE `monitor_rules` ENABLE KEYS */;

-- 正在导出表  pro_cfg_manager.operation_log 的数据：~1 rows (大约)
DELETE FROM `operation_log`;
/*!40000 ALTER TABLE `operation_log` DISABLE KEYS */;
INSERT INTO `operation_log` (`id`, `username`, `operate_type`, `ipaddr`, `operate_content`, `operate_result`, `operate_at`, `operate_error`) VALUES
	(1, 'admin', '', '127.0.0.1:59524', 'reset prometheus config data', 1, '2021-12-06 16:23:01', '成功'),
	(2, 'admin', '', '127.0.0.1:51385', 'reset prometheus config data', 1, '2022-02-17 14:14:14', '成功');
/*!40000 ALTER TABLE `operation_log` ENABLE KEYS */;

-- 正在导出表  pro_cfg_manager.options 的数据：~6 rows (大约)
DELETE FROM `options`;
/*!40000 ALTER TABLE `options` DISABLE KEYS */;
INSERT INTO `options` (`id`, `opt_key`, `opt_value`) VALUES
	(1, 'publish_at_null_subgroup', 'true'),
	(2, 'publish_at_remain_subgroup', 'true'),
	(3, 'publish_at_empty_nocreate_file', 'true'),
	(4, 'publish_jobs_also_ips', 'true'),
	(5, 'publish_jobs_also_reload_srv', 'true'),
	(6, 'publish_ips_also_reload_srv', 'true');
/*!40000 ALTER TABLE `options` ENABLE KEYS */;

-- 正在导出表  pro_cfg_manager.page_function 的数据：~98 rows (大约)
DELETE FROM `page_function`;
/*!40000 ALTER TABLE `page_function` DISABLE KEYS */;
INSERT INTO `page_function` (`id`, `page_name`, `page_nice_name`, `sub_page_name`, `sub_page_nice_name`, `func_name`, `func_nice_name`) VALUES
	(1, 'person', '个人中心', '', '', 'update_password', '更新密码'),
	(2, 'person', '个人中心', '', '', 'show_menu_prometheus_cfg_manager', '显示Prometheus管理菜单项'),
	(3, 'person', '个人中心', '', '', 'show_menu_administrator_cfg_manager', '显示用户及权限管理菜单项'),
	(4, 'register', '注册', '', '', 'register', '注册'),
	(5, 'ipManager', 'IP管理', '', '', 'search', '查询'),
	(6, 'ipManager', 'IP管理', '', '', 'update', '编辑'),
	(7, 'ipManager', 'IP管理', '', '', 'dis.enable', '启用/禁用'),
	(8, 'ipManager', 'IP管理', '', '', 'delete', '删除/批量删除'),
	(9, 'ipManager', 'IP管理', '', '', 'add', '添加'),
	(11, 'jobs', 'JOB组管理', 'jobs', 'JOB组管理', 'job_search', '查询组'),
	(12, 'jobs', 'JOB组管理', 'jobs', 'JOB组管理', 'job_add', '添加组'),
	(13, 'jobs', 'JOB组管理', 'jobs', 'JOB组管理', 'job_update', '编辑组'),
	(14, 'jobs', 'JOB组管理', 'jobs', 'JOB组管理', 'dis.enable', '启用/禁用组'),
	(15, 'jobs', 'JOB组管理', 'jobs', 'JOB组管理', 'ip_pool', '更新组IP池'),
	(16, 'jobs', 'JOB组管理', 'jobs', 'JOB组管理', 'delete', '删除组'),
	(19, 'jobs', 'JOB组管理', 'labelsJobs', '分组标签', 'search_sub_group', '查询子组'),
	(20, 'jobs', 'JOB组管理', 'labelsJobs', '分组标签', 'add_sub_group', '添加子组'),
	(21, 'jobs', 'JOB组管理', 'labelsJobs', '分组标签', 'update_sub_group', '更新子组'),
	(22, 'jobs', 'JOB组管理', 'labelsJobs', '分组标签', 'dis.enable_sub_group', '启用/禁用子组'),
	(23, 'jobs', 'JOB组管理', 'labelsJobs', '分组标签', 'delete_sub_group', '删除子组'),
	(24, 'jobs', 'JOB组管理', 'labelsJobs', '分组标签', 'search_sub_group_ip_pool', '查询子组IP池'),
	(25, 'jobs', 'JOB组管理', 'labelsJobs', '分组标签', 'update_sub_group_ip_pool', '更新子组IP池'),
	(26, 'jobs', 'JOB组管理', 'labelsJobs', '分组标签', 'search_sub_group_label', '查询子组标签列表'),
	(27, 'jobs', 'JOB组管理', 'labelsJobs', '分组标签', 'update_sub_group_label', '更新子组标签'),
	(28, 'jobs', 'JOB组管理', 'labelsJobs', '分组标签', 'dis.enable_sub_group_label', '启用/禁用子组标签'),
	(29, 'jobs', 'JOB组管理', 'labelsJobs', '分组标签', 'delete_sub_group_label', '删除子组标签'),
	(30, 'baseConfig', '基本配置', 'options', '选项编辑', 'search_options', '查询选项'),
	(31, 'baseConfig', '基本配置', 'options', '选项编辑', 'update_options', '更新选项'),
	(33, 'baseConfig', '基本配置', 'baseLabels', '公共标签', 'search', '查询'),
	(34, 'baseConfig', '基本配置', 'baseLabels', '公共标签', 'add', '增加'),
	(35, 'baseConfig', '基本配置', 'baseLabels', '公共标签', 'update', '更新'),
	(36, 'baseConfig', '基本配置', 'baseLabels', '公共标签', 'dis.enable', '启用/禁用'),
	(37, 'baseConfig', '基本配置', 'baseLabels', '公共标签', 'delete', '删除'),
	(38, 'baseConfig', '基本配置', 'reLabels', '标签重写', 'search', '查询'),
	(39, 'baseConfig', '基本配置', 'reLabels', '标签重写', 'add', '增加'),
	(40, 'baseConfig', '基本配置', 'reLabels', '标签重写', 'edit_name', '编辑名称'),
	(41, 'baseConfig', '基本配置', 'reLabels', '标签重写', 'update_rule', '更新规则'),
	(42, 'baseConfig', '基本配置', 'reLabels', '标签重写', 'dis.enable', '启用/禁用'),
	(43, 'baseConfig', '基本配置', 'reLabels', '标签重写', 'delete', '删除'),
	(44, 'baseConfig', '基本配置', 'defaultJobs', '默认分组', 'search', '查询'),
	(45, 'baseConfig', '基本配置', 'defaultJobs', '默认分组', 'add', '增加'),
	(46, 'baseConfig', '基本配置', 'defaultJobs', '默认分组', 'update', '更新'),
	(47, 'baseConfig', '基本配置', 'defaultJobs', '默认分组', 'dis.enable', '启用/禁用'),
	(48, 'baseConfig', '基本配置', 'defaultJobs', '默认分组', 'delete', '删除'),
	(49, 'baseConfig', '基本配置', 'checkYml', '测试配置', 'restart', '重启Prometheus服务'),
	(50, 'baseConfig', '基本配置', 'checkYml', '测试配置', 'check', '测试配置文件'),
	(51, 'baseConfig', '基本配置', 'editPrometheusYml', '模板编辑', 'load_tmpl', '加载模板数据'),
	(52, 'baseConfig', '基本配置', 'editPrometheusYml', '模板编辑', 'save_tmpl', '保存模板数据'),
	(53, 'baseConfig', '基本配置', 'empty', '重置', 'reset', '重置所有数据'),
	(54, 'noticeManager', '告警管理', '', '', 'search_tree', '获取告警规则列表'),
	(55, 'noticeManager', '告警管理', '', '', 'add_file', '添加文件'),
	(56, 'noticeManager', '告警管理', '', '', 'dis.enable', '启用/禁用分支'),
	(57, 'noticeManager', '告警管理', '', '', 'delete_tree_node', '删除所有子节点'),
	(58, 'noticeManager', '告警管理', '', '', 'delete_node', '删除此节点'),
	(59, 'noticeManager', '告警管理', '', '', 'rename_node', '重命名此节点'),
	(60, 'noticeManager', '告警管理', '', '', 'add_group', '增加组'),
	(61, 'noticeManager', '告警管理', '', '', 'import_rule_from_file', '从文件导入规则'),
	(62, 'noticeManager', '告警管理', '', '', 'add_rule', '添加规则'),
	(63, 'noticeManager', '告警管理', '', '', 'search_rule', '查询规则内容'),
	(64, 'noticeManager', '告警管理', '', '', 'update_rule', '更新规则'),
	(65, 'noticeManager', '告警管理', '', '', 'publish_rule', '发布'),
	(66, 'preview', '配置预览', '', '', 'search', '查看配置'),
	(67, 'ftree', '分组预览', '', '', 'list_all_file', '列表分组IP文件'),
	(68, 'ruleView', '告警规则预览', '', '', 'list_all_file', '列表告警规则文件'),
	(69, 'admin', '管理中心', 'user', '用户管理', 'search', '查看用户列表'),
	(70, 'admin', '管理中心', 'user', '用户管理', 'add', '增加用户'),
	(71, 'admin', '管理中心', 'user', '用户管理', 'update', '更新用户'),
	(72, 'admin', '管理中心', 'user', '用户管理', 'dis.enable', '启用/禁用用户'),
	(73, 'admin', '管理中心', 'user', '用户管理', 'delete', '删除用户'),
	(74, 'admin', '管理中心', 'group', '组管理', 'search', '查看用户组列表'),
	(75, 'admin', '管理中心', 'group', '组管理', 'add', '添加组'),
	(76, 'admin', '管理中心', 'group', '组管理', 'update', '更新组'),
	(77, 'admin', '管理中心', 'group', '组管理', 'dis.enable', '启用/禁用组'),
	(78, 'admin', '管理中心', 'group', '组管理', 'delete', '删除组'),
	(79, 'admin', '管理中心', 'privileges', '权限管理', 'search', '查看权限列表'),
	(80, 'admin', '管理中心', 'privileges', '权限管理', 'update', '更新权限列表'),
	(81, 'jobs', 'JOB组管理', 'jobs', 'JOB组管理', 'swap', '调整顺序'),
	(83, 'ftree', '分组预览', '', '', 'view_file', '查看IP文件内容'),
	(84, 'ruleView', '告警规则预览', '', '', 'view_file', '查看告警规则文件内容'),
	(85, 'admin', '管理中心', 'setting', '设置', 'search', '获取设置'),
	(86, 'admin', '管理中心', 'setting', '设置', 'update', '更新设置'),
	(87, 'admin', '管理中心', 'setting', '设置', 'reset', '重置数据'),
	(88, 'ipManager', 'IP管理', '', '', 'import', '批量添加'),
	(89, 'jobs', 'JOB组管理', 'labelsJobs', '分组标签', 'add_sub_group_label', '添加子组标签'),
	(90, 'baseConfig', '基本配置', 'reLabels', '标签重写', 'view_code', '展示代码'),
	(93, 'baseConfig', '基本配置', 'baseFields', '模板字段', 'search', '查询'),
	(94, 'baseConfig', '基本配置', 'baseFields', '模板字段', 'add', '增加'),
	(95, 'baseConfig', '基本配置', 'baseFields', '模板字段', 'update', '更新'),
	(96, 'baseConfig', '基本配置', 'baseFields', '模板字段', 'dis.enable', '启用/禁用'),
	(97, 'baseConfig', '基本配置', 'baseFields', '模板字段', 'delete', '删除'),
	(98, 'admin', '管理中心', 'log', '日志', 'search', '查询'),
	(99, 'admin', '管理中心', 'log', '日志', 'update', '更新设置'),
	(100, 'admin', '管理中心', 'log', '日志', 'clear', '清空日志'),
	(101, 'control', '控制中心', '', '', 'create', '重新创建配置'),
	(102, 'control', '控制中心', '', '', 'reload', '重新加载配置'),
	(103, 'control', '控制中心', '', '', 'createAndreload', '创建配置并加载'),
	(104, 'noticeManager', '告警管理', '', '', 'publish_empty_rule', '发布空规则'),
	(105, 'ipManager', 'IP管理', '', '', 'position', '更新IP位置信息');
/*!40000 ALTER TABLE `page_function` ENABLE KEYS */;

-- 正在导出表  pro_cfg_manager.relabels 的数据：~0 rows (大约)
DELETE FROM `relabels`;
/*!40000 ALTER TABLE `relabels` DISABLE KEYS */;
INSERT INTO `relabels` (`id`, `name`, `code`, `enabled`, `update_at`, `update_by`) VALUES
	(1, '空规则', '', 1, '2021-07-07 14:57:50', '');
/*!40000 ALTER TABLE `relabels` ENABLE KEYS */;

-- 正在导出表  pro_cfg_manager.rules_groups 的数据：~0 rows (大约)
DELETE FROM `rules_groups`;
/*!40000 ALTER TABLE `rules_groups` DISABLE KEYS */;
/*!40000 ALTER TABLE `rules_groups` ENABLE KEYS */;

-- 正在导出表  pro_cfg_manager.session 的数据：~2 rows (大约)
DELETE FROM `session`;
/*!40000 ALTER TABLE `session` DISABLE KEYS */;
INSERT INTO `session` (`id`, `token`, `user_id`, `update_at`) VALUES
	(93, 'd553bf31-8d9e-4ed8-b6a0-d4b78150f9b8', 7, '2021-08-10 08:19:56'),
	(95, 'fd5e775d-368d-4131-8281-52f27635925b', 1, '2022-02-16 03:52:26');
/*!40000 ALTER TABLE `session` ENABLE KEYS */;

-- 正在导出表  pro_cfg_manager.sub_group 的数据：~0 rows (大约)
DELETE FROM `sub_group`;
/*!40000 ALTER TABLE `sub_group` DISABLE KEYS */;
/*!40000 ALTER TABLE `sub_group` ENABLE KEYS */;

-- 正在导出表  pro_cfg_manager.system_log 的数据：~0 rows (大约)
DELETE FROM `system_log`;
/*!40000 ALTER TABLE `system_log` DISABLE KEYS */;
/*!40000 ALTER TABLE `system_log` ENABLE KEYS */;

-- 正在导出表  pro_cfg_manager.tmpl 的数据：~1 rows (大约)
DELETE FROM `tmpl`;
/*!40000 ALTER TABLE `tmpl` DISABLE KEYS */;
INSERT INTO `tmpl` (`tmpl`, `update_at`, `update_by`) VALUES
	('# my global config\nglobal:\n  scrape_interval:     15s # Set the scrape interval to every 15 seconds. Default is every 1 minute.\n  evaluation_interval: 15s # Evaluate rules every 15 seconds. The default is every 1 minute.\n  # scrape_timeout is set to the global default (10s).\n  # scrape_timeout: 60s\n\n# Alertmanager configuration\nalerting:\n  alertmanagers:\n  - static_configs:\n    - targets:\n      - 127.0.0.1:9093\n\n# Load rules once and periodically evaluate them according to the global \'evaluation_interval\'.\nrule_files:\n   - "{{.RelRuleDir}}/*.yml"\n  # - "first_rules.yml"\n  # - "second_rules.yml"\n\n# A scrape configuration containing exactly one endpoint to scrape:\n# Here it\'s Prometheus itself.\nscrape_configs:\n  # The job name is added as a label `job=<job_name>` to any timeseries scraped from this config.\n  - job_name: \'监控服务本机\'\n\n    # metrics_path defaults to \'/metrics\'\n    # scheme defaults to \'http\'.\n\n    static_configs:\n    - targets: [\'localhost:9090\']\n{{ range .Jobs }}\n  - job_name: \'{{.Name}}\'\n    file_sd_configs:\n      - files:\n        - "{{$.AbsConfDir}}/{{.Name}}.json"\n        refresh_interval: {{$.Fields.refresh_interval}}\n{{.Code}}\n{{ end }}\n', '2022-02-17 14:14:14', 'administrator');
/*!40000 ALTER TABLE `tmpl` ENABLE KEYS */;

-- 正在导出表  pro_cfg_manager.tmpl_fields 的数据：~2 rows (大约)
DELETE FROM `tmpl_fields`;
/*!40000 ALTER TABLE `tmpl_fields` DISABLE KEYS */;
INSERT INTO `tmpl_fields` (`id`, `key`, `value`, `enabled`, `update_at`, `update_by`) VALUES
	(1, 'metrics', '/metrics', 1, '2022-02-17 14:14:14', ''),
	(2, 'refresh_interval', '15s', 1, '2022-02-17 14:14:14', '');
/*!40000 ALTER TABLE `tmpl_fields` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
