-- MySQL dump 10.13  Distrib 8.0.27, for Linux (x86_64)
--
-- Host: localhost    Database: pro_cfg_manager
-- ------------------------------------------------------
-- Server version	8.0.27-0ubuntu0.21.10.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Current Database: `pro_cfg_manager`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `pro_cfg_manager` /*!40100 DEFAULT CHARACTER SET utf8 */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `pro_cfg_manager`;

--
-- Table structure for table `annotations`
--

DROP TABLE IF EXISTS `annotations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `annotations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `monitor_rules_id` int NOT NULL,
  `key` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `value` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_monitor_rules_id_key` (`monitor_rules_id`,`key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='监控规则的注释';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `annotations`
--

LOCK TABLES `annotations` WRITE;
/*!40000 ALTER TABLE `annotations` DISABLE KEYS */;
/*!40000 ALTER TABLE `annotations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `group_labels`
--

DROP TABLE IF EXISTS `group_labels`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `group_labels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `job_group_id` int NOT NULL,
  `key` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `value` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `enabled` tinyint NOT NULL DEFAULT '1',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_by` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_groupid_key` (`key`,`job_group_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='子组标签列表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `group_labels`
--

LOCK TABLES `group_labels` WRITE;
/*!40000 ALTER TABLE `group_labels` DISABLE KEYS */;
/*!40000 ALTER TABLE `group_labels` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `group_machines`
--

DROP TABLE IF EXISTS `group_machines`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `group_machines` (
  `id` int NOT NULL AUTO_INCREMENT,
  `job_group_id` int NOT NULL,
  `machines_id` int NOT NULL,
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique` (`machines_id`,`job_group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='子组IP列表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `group_machines`
--

LOCK TABLES `group_machines` WRITE;
/*!40000 ALTER TABLE `group_machines` DISABLE KEYS */;
/*!40000 ALTER TABLE `group_machines` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `group_priv`
--

DROP TABLE IF EXISTS `group_priv`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `group_priv` (
  `id` int NOT NULL AUTO_INCREMENT,
  `group_id` int NOT NULL,
  `func_id` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=399 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `group_priv`
--

LOCK TABLES `group_priv` WRITE;
/*!40000 ALTER TABLE `group_priv` DISABLE KEYS */;
INSERT INTO `group_priv` VALUES (302,1,78),(303,1,77),(304,1,76),(305,1,75),(306,1,74),(307,1,100),(308,1,99),(309,1,98),(310,1,80),(311,1,79),(312,1,85),(313,1,86),(314,1,87),(315,1,71),(316,1,72),(317,1,73),(318,1,69),(319,1,70),(320,1,93),(321,1,94),(322,1,95),(323,1,96),(324,1,97),(325,1,37),(326,1,36),(327,1,35),(328,1,34),(329,1,33),(330,1,50),(331,1,49),(332,1,48),(333,1,47),(334,1,46),(335,1,45),(336,1,44),(337,1,52),(338,1,51),(339,1,53),(340,1,30),(341,1,31),(342,1,43),(343,1,42),(344,1,41),(345,1,40),(346,1,39),(347,1,38),(348,1,90),(349,1,103),(350,1,102),(351,1,101),(352,1,83),(353,1,67),(354,1,5),(355,1,6),(356,1,7),(357,1,8),(358,1,9),(359,1,88),(360,1,15),(361,1,81),(362,1,11),(363,1,12),(364,1,13),(365,1,14),(366,1,16),(367,1,29),(368,1,19),(369,1,89),(370,1,20),(371,1,21),(372,1,22),(373,1,23),(374,1,24),(375,1,25),(376,1,26),(377,1,27),(378,1,28),(379,1,104),(380,1,65),(381,1,64),(382,1,63),(383,1,62),(384,1,61),(385,1,60),(386,1,59),(387,1,58),(388,1,57),(389,1,56),(390,1,55),(391,1,54),(392,1,1),(393,1,2),(394,1,3),(395,1,66),(396,1,4),(397,1,68),(398,1,84);
/*!40000 ALTER TABLE `group_priv` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `job_group`
--

DROP TABLE IF EXISTS `job_group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `job_group` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `jobs_id` int NOT NULL,
  `enabled` tinyint NOT NULL DEFAULT '1',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_by` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='还可以为单个分组中的IP地址进行分子组，为每个子组设置相应的标签列表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `job_group`
--

LOCK TABLES `job_group` WRITE;
/*!40000 ALTER TABLE `job_group` DISABLE KEYS */;
/*!40000 ALTER TABLE `job_group` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `job_machines`
--

DROP TABLE IF EXISTS `job_machines`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `job_machines` (
  `machine_id` int NOT NULL,
  `job_id` int NOT NULL,
  UNIQUE KEY `unique_mj` (`job_id`,`machine_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `job_machines`
--

LOCK TABLES `job_machines` WRITE;
/*!40000 ALTER TABLE `job_machines` DISABLE KEYS */;
/*!40000 ALTER TABLE `job_machines` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `jobs`
--

DROP TABLE IF EXISTS `jobs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `jobs` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(100) NOT NULL COMMENT '任务名称',
  `port` int NOT NULL COMMENT '端口号，对应exporter的端口号',
  `cfg_name` varchar(100) DEFAULT NULL COMMENT '在prometheus下生成配置的文件名称',
  `is_common` tinyint NOT NULL,
  `relabel_id` int NOT NULL,
  `display_order` int unsigned NOT NULL,
  `enabled` tinyint NOT NULL DEFAULT '1',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_by` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_unique` (`name`),
  KEY `order_unique` (`display_order`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='prometheus任务列表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `jobs`
--

LOCK TABLES `jobs` WRITE;
/*!40000 ALTER TABLE `jobs` DISABLE KEYS */;
/*!40000 ALTER TABLE `jobs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `labels`
--

DROP TABLE IF EXISTS `labels`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `labels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `label` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `enabled` tinyint NOT NULL DEFAULT '1',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_by` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_labels` (`label`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='标签列表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `labels`
--

LOCK TABLES `labels` WRITE;
/*!40000 ALTER TABLE `labels` DISABLE KEYS */;
/*!40000 ALTER TABLE `labels` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `log_setting`
--

DROP TABLE IF EXISTS `log_setting`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `log_setting` (
  `id` int NOT NULL AUTO_INCREMENT,
  `level` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `label` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `selected` tinyint NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb3 COMMENT='日志设置';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `log_setting`
--

LOCK TABLES `log_setting` WRITE;
/*!40000 ALTER TABLE `log_setting` DISABLE KEYS */;
INSERT INTO `log_setting` VALUES (1,'search','查询',0),(2,'add','增加',1),(3,'del','删除',1),(4,'update','更新',1),(5,'running','运行',1),(6,'publish','发布',1),(7,'login','登录',1),(8,'reset','重置',1);
/*!40000 ALTER TABLE `log_setting` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `machines`
--

DROP TABLE IF EXISTS `machines`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `machines` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `ipaddr` varchar(100) NOT NULL COMMENT 'IP地址',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `enabled` tinyint NOT NULL DEFAULT '1',
  `update_by` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `ipaddr_unique` (`ipaddr`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='机器列表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `machines`
--

LOCK TABLES `machines` WRITE;
/*!40000 ALTER TABLE `machines` DISABLE KEYS */;
/*!40000 ALTER TABLE `machines` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `manager_group`
--

DROP TABLE IF EXISTS `manager_group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `manager_group` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `enabled` tinyint NOT NULL,
  `update_at` datetime NOT NULL,
  `update_by` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_unique` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3 COMMENT='用户组';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `manager_group`
--

LOCK TABLES `manager_group` WRITE;
/*!40000 ALTER TABLE `manager_group` DISABLE KEYS */;
INSERT INTO `manager_group` VALUES (1,'administrator',1,'2021-12-06 16:19:21','admin');
/*!40000 ALTER TABLE `manager_group` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `manager_set`
--

DROP TABLE IF EXISTS `manager_set`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `manager_set` (
  `param_name` varbinary(100) NOT NULL,
  `param_value` varbinary(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='权限管理默认的设置参数';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `manager_set`
--

LOCK TABLES `manager_set` WRITE;
/*!40000 ALTER TABLE `manager_set` DISABLE KEYS */;
INSERT INTO `manager_set` VALUES (_binary 'default_group','');
/*!40000 ALTER TABLE `manager_set` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `manager_user`
--

DROP TABLE IF EXISTS `manager_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `manager_user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(100) NOT NULL,
  `nice_name` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `phone` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `salt` varchar(100) NOT NULL,
  `group_id` int NOT NULL,
  `update_at` datetime NOT NULL,
  `enabled` tinyint NOT NULL,
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_by` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_name` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3 COMMENT='用户';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `manager_user`
--

LOCK TABLES `manager_user` WRITE;
/*!40000 ALTER TABLE `manager_user` DISABLE KEYS */;
INSERT INTO `manager_user` VALUES (1,'admin','管理员','cc7550fb9f1b75d84b3677fdcd9d4c9f','10086','ec7ceb50-c6c0-43b2-994e-b79cdb365457',1,'2021-08-10 17:33:24',1,'2021-08-10 17:33:24','');
/*!40000 ALTER TABLE `manager_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `monitor_labels`
--

DROP TABLE IF EXISTS `monitor_labels`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `monitor_labels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `monitor_rules_id` int NOT NULL,
  `key` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `value` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_mid_lid` (`key`,`monitor_rules_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='监控规则的标签';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `monitor_labels`
--

LOCK TABLES `monitor_labels` WRITE;
/*!40000 ALTER TABLE `monitor_labels` DISABLE KEYS */;
/*!40000 ALTER TABLE `monitor_labels` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `monitor_rules`
--

DROP TABLE IF EXISTS `monitor_rules`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `monitor_rules` (
  `id` int NOT NULL AUTO_INCREMENT,
  `alert` varchar(100) NOT NULL,
  `expr` varchar(5000) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `for` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `sub_group_id` int NOT NULL,
  `enabled` tinyint NOT NULL,
  `description` varchar(300) NOT NULL,
  `update_by` varchar(100) NOT NULL,
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='具体的监控规则';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `monitor_rules`
--

LOCK TABLES `monitor_rules` WRITE;
/*!40000 ALTER TABLE `monitor_rules` DISABLE KEYS */;
/*!40000 ALTER TABLE `monitor_rules` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `operation_log`
--

DROP TABLE IF EXISTS `operation_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `operation_log` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(100) NOT NULL,
  `operate_type` varchar(50) NOT NULL,
  `ipaddr` varchar(100) NOT NULL,
  `operate_content` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `operate_result` tinyint NOT NULL,
  `operate_at` datetime NOT NULL,
  `operate_error` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3 COMMENT='操作日志';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `operation_log`
--

LOCK TABLES `operation_log` WRITE;
/*!40000 ALTER TABLE `operation_log` DISABLE KEYS */;
INSERT INTO `operation_log` VALUES (1,'admin','','127.0.0.1:59524','reset prometheus config data',1,'2021-12-06 16:23:01','成功');
/*!40000 ALTER TABLE `operation_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `options`
--

DROP TABLE IF EXISTS `options`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `options` (
  `id` int NOT NULL AUTO_INCREMENT,
  `opt_key` varchar(100) NOT NULL,
  `opt_value` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb3 COMMENT='选项';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `options`
--

LOCK TABLES `options` WRITE;
/*!40000 ALTER TABLE `options` DISABLE KEYS */;
INSERT INTO `options` VALUES (1,'publish_at_null_subgroup','true'),(2,'publish_at_remain_subgroup','true'),(3,'publish_at_empty_nocreate_file','true'),(4,'publish_jobs_also_ips','true'),(5,'publish_jobs_also_reload_srv','true'),(6,'publish_ips_also_reload_srv','true');
/*!40000 ALTER TABLE `options` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `page_function`
--

DROP TABLE IF EXISTS `page_function`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `page_function` (
  `id` int NOT NULL AUTO_INCREMENT,
  `page_name` varchar(100) NOT NULL,
  `page_nice_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `sub_page_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '',
  `sub_page_nice_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '',
  `func_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `func_nice_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=105 DEFAULT CHARSET=utf8mb3 COMMENT='页面功能';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `page_function`
--

LOCK TABLES `page_function` WRITE;
/*!40000 ALTER TABLE `page_function` DISABLE KEYS */;
INSERT INTO `page_function` VALUES (1,'person','个人中心','','','update_password','更新密码'),(2,'person','个人中心','','','show_menu_prometheus_cfg_manager','显示Prometheus管理菜单项'),(3,'person','个人中心','','','show_menu_administrator_cfg_manager','显示用户及权限管理菜单项'),(4,'register','注册','','','register','注册'),(5,'ipManager','IP管理','','','search','查询'),(6,'ipManager','IP管理','','','update','编辑'),(7,'ipManager','IP管理','','','dis.enable','启用/禁用'),(8,'ipManager','IP管理','','','delete','删除/批量删除'),(9,'ipManager','IP管理','','','add','添加'),(11,'jobs','JOB组管理','jobs','JOB组管理','job_search','查询组'),(12,'jobs','JOB组管理','jobs','JOB组管理','job_add','添加组'),(13,'jobs','JOB组管理','jobs','JOB组管理','job_update','编辑组'),(14,'jobs','JOB组管理','jobs','JOB组管理','dis.enable','启用/禁用组'),(15,'jobs','JOB组管理','jobs','JOB组管理','ip_pool','更新组IP池'),(16,'jobs','JOB组管理','jobs','JOB组管理','delete','删除组'),(19,'jobs','JOB组管理','labelsJobs','分组标签','search_sub_group','查询子组'),(20,'jobs','JOB组管理','labelsJobs','分组标签','add_sub_group','添加子组'),(21,'jobs','JOB组管理','labelsJobs','分组标签','update_sub_group','更新子组'),(22,'jobs','JOB组管理','labelsJobs','分组标签','dis.enable_sub_group','启用/禁用子组'),(23,'jobs','JOB组管理','labelsJobs','分组标签','delete_sub_group','删除子组'),(24,'jobs','JOB组管理','labelsJobs','分组标签','search_sub_group_ip_pool','查询子组IP池'),(25,'jobs','JOB组管理','labelsJobs','分组标签','update_sub_group_ip_pool','更新子组IP池'),(26,'jobs','JOB组管理','labelsJobs','分组标签','search_sub_group_label','查询子组标签列表'),(27,'jobs','JOB组管理','labelsJobs','分组标签','update_sub_group_label','更新子组标签'),(28,'jobs','JOB组管理','labelsJobs','分组标签','dis.enable_sub_group_label','启用/禁用子组标签'),(29,'jobs','JOB组管理','labelsJobs','分组标签','delete_sub_group_label','删除子组标签'),(30,'baseConfig','基本配置','options','选项编辑','search_options','查询选项'),(31,'baseConfig','基本配置','options','选项编辑','update_options','更新选项'),(33,'baseConfig','基本配置','baseLabels','公共标签','search','查询'),(34,'baseConfig','基本配置','baseLabels','公共标签','add','增加'),(35,'baseConfig','基本配置','baseLabels','公共标签','update','更新'),(36,'baseConfig','基本配置','baseLabels','公共标签','dis.enable','启用/禁用'),(37,'baseConfig','基本配置','baseLabels','公共标签','delete','删除'),(38,'baseConfig','基本配置','reLabels','标签重写','search','查询'),(39,'baseConfig','基本配置','reLabels','标签重写','add','增加'),(40,'baseConfig','基本配置','reLabels','标签重写','edit_name','编辑名称'),(41,'baseConfig','基本配置','reLabels','标签重写','update_rule','更新规则'),(42,'baseConfig','基本配置','reLabels','标签重写','dis.enable','启用/禁用'),(43,'baseConfig','基本配置','reLabels','标签重写','delete','删除'),(44,'baseConfig','基本配置','defaultJobs','默认分组','search','查询'),(45,'baseConfig','基本配置','defaultJobs','默认分组','add','增加'),(46,'baseConfig','基本配置','defaultJobs','默认分组','update','更新'),(47,'baseConfig','基本配置','defaultJobs','默认分组','dis.enable','启用/禁用'),(48,'baseConfig','基本配置','defaultJobs','默认分组','delete','删除'),(49,'baseConfig','基本配置','checkYml','测试配置','restart','重启Prometheus服务'),(50,'baseConfig','基本配置','checkYml','测试配置','check','测试配置文件'),(51,'baseConfig','基本配置','editPrometheusYml','模板编辑','load_tmpl','加载模板数据'),(52,'baseConfig','基本配置','editPrometheusYml','模板编辑','save_tmpl','保存模板数据'),(53,'baseConfig','基本配置','empty','重置','reset','重置所有数据'),(54,'noticeManager','告警管理','','','search_tree','获取告警规则列表'),(55,'noticeManager','告警管理','','','add_file','添加文件'),(56,'noticeManager','告警管理','','','dis.enable','启用/禁用分支'),(57,'noticeManager','告警管理','','','delete_tree_node','删除所有子节点'),(58,'noticeManager','告警管理','','','delete_node','删除此节点'),(59,'noticeManager','告警管理','','','rename_node','重命名此节点'),(60,'noticeManager','告警管理','','','add_group','增加组'),(61,'noticeManager','告警管理','','','import_rule_from_file','从文件导入规则'),(62,'noticeManager','告警管理','','','add_rule','添加规则'),(63,'noticeManager','告警管理','','','search_rule','查询规则内容'),(64,'noticeManager','告警管理','','','update_rule','更新规则'),(65,'noticeManager','告警管理','','','publish_rule','发布'),(66,'preview','配置预览','','','search','查看配置'),(67,'ftree','分组预览','','','list_all_file','列表分组IP文件'),(68,'ruleView','告警规则预览','','','list_all_file','列表告警规则文件'),(69,'admin','管理中心','user','用户管理','search','查看用户列表'),(70,'admin','管理中心','user','用户管理','add','增加用户'),(71,'admin','管理中心','user','用户管理','update','更新用户'),(72,'admin','管理中心','user','用户管理','dis.enable','启用/禁用用户'),(73,'admin','管理中心','user','用户管理','delete','删除用户'),(74,'admin','管理中心','group','组管理','search','查看用户组列表'),(75,'admin','管理中心','group','组管理','add','添加组'),(76,'admin','管理中心','group','组管理','update','更新组'),(77,'admin','管理中心','group','组管理','dis.enable','启用/禁用组'),(78,'admin','管理中心','group','组管理','delete','删除组'),(79,'admin','管理中心','privileges','权限管理','search','查看权限列表'),(80,'admin','管理中心','privileges','权限管理','update','更新权限列表'),(81,'jobs','JOB组管理','jobs','JOB组管理','swap','调整顺序'),(83,'ftree','分组预览','','','view_file','查看IP文件内容'),(84,'ruleView','告警规则预览','','','view_file','查看告警规则文件内容'),(85,'admin','管理中心','setting','设置','search','获取设置'),(86,'admin','管理中心','setting','设置','update','更新设置'),(87,'admin','管理中心','setting','设置','reset','重置数据'),(88,'ipManager','IP管理','','','import','批量添加'),(89,'jobs','JOB组管理','labelsJobs','分组标签','add_sub_group_label','添加子组标签'),(90,'baseConfig','基本配置','reLabels','标签重写','view_code','展示代码'),(93,'baseConfig','基本配置','baseFields','模板字段','search','查询'),(94,'baseConfig','基本配置','baseFields','模板字段','add','增加'),(95,'baseConfig','基本配置','baseFields','模板字段','update','更新'),(96,'baseConfig','基本配置','baseFields','模板字段','dis.enable','启用/禁用'),(97,'baseConfig','基本配置','baseFields','模板字段','delete','删除'),(98,'admin','管理中心','log','日志','search','查询'),(99,'admin','管理中心','log','日志','update','更新设置'),(100,'admin','管理中心','log','日志','clear','清空日志'),(101,'control','控制中心','','','create','重新创建配置'),(102,'control','控制中心','','','reload','重新加载配置'),(103,'control','控制中心','','','createAndreload','创建配置并加载'),(104,'noticeManager','告警管理','','','publish_empty_rule','发布空规则');
/*!40000 ALTER TABLE `page_function` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `relabels`
--

DROP TABLE IF EXISTS `relabels`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `relabels` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(200) NOT NULL,
  `code` text CHARACTER SET utf8 COLLATE utf8_general_ci,
  `enabled` tinyint NOT NULL DEFAULT '1',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_by` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb3 COMMENT='重写标签的配置';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `relabels`
--

LOCK TABLES `relabels` WRITE;
/*!40000 ALTER TABLE `relabels` DISABLE KEYS */;
INSERT INTO `relabels` VALUES (1,'空规则','',1,'2021-07-07 14:57:50','');
/*!40000 ALTER TABLE `relabels` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `rules_groups`
--

DROP TABLE IF EXISTS `rules_groups`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `rules_groups` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_by` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='监控规则组';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rules_groups`
--

LOCK TABLES `rules_groups` WRITE;
/*!40000 ALTER TABLE `rules_groups` DISABLE KEYS */;
/*!40000 ALTER TABLE `rules_groups` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `session`
--

DROP TABLE IF EXISTS `session`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `session` (
  `id` int NOT NULL AUTO_INCREMENT,
  `token` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `user_id` int NOT NULL,
  `update_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=97 DEFAULT CHARSET=utf8mb3 COMMENT='用户登录会话';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `session`
--

LOCK TABLES `session` WRITE;
/*!40000 ALTER TABLE `session` DISABLE KEYS */;
INSERT INTO `session` VALUES (93,'d553bf31-8d9e-4ed8-b6a0-d4b78150f9b8',7,'2021-08-10 08:19:56'),(95,'a1c0d5bd-e135-4f7d-ad9e-16020b594a0a',1,'2021-12-03 07:53:48');
/*!40000 ALTER TABLE `session` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sub_group`
--

DROP TABLE IF EXISTS `sub_group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sub_group` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `rules_groups_id` int NOT NULL,
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_by` varchar(50) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_name_rules_groups_id` (`name`,`rules_groups_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='监控规则组中的下级组，子组';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sub_group`
--

LOCK TABLES `sub_group` WRITE;
/*!40000 ALTER TABLE `sub_group` DISABLE KEYS */;
/*!40000 ALTER TABLE `sub_group` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_log`
--

DROP TABLE IF EXISTS `system_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `system_log` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `operate_type` varchar(50) NOT NULL,
  `ipaddr` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `operate_content` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `operate_result` tinyint NOT NULL,
  `operate_at` datetime NOT NULL,
  `operate_error` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=454 DEFAULT CHARSET=utf8mb3 COMMENT='系统日志';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_log`
--

LOCK TABLES `system_log` WRITE;
/*!40000 ALTER TABLE `system_log` DISABLE KEYS */;
INSERT INTO `system_log` VALUES (386,'admin','删除','110.88.106.120:3617','clear system log',1,'2021-08-10 17:33:01','成功'),(387,'admin','增加','110.88.106.120:3617','create manager reset secret',1,'2021-08-10 17:33:06','成功'),(388,'admin','重置','110.88.106.120:3689','reset manager data',1,'2021-08-10 17:33:24','成功'),(389,'admin','重置','110.88.106.120:3779','create reset prometheus config data key',1,'2021-08-10 17:33:39','成功'),(390,'admin','重置','110.88.106.120:1368','create reset prometheus config data key',1,'2021-08-10 17:39:26','成功'),(391,'admin','重置','110.88.106.120:2242','create reset prometheus config data key',1,'2021-08-10 17:42:45','成功'),(392,'admin','更新','110.88.106.120:3389','update sub group',1,'2021-08-10 17:46:18','成功'),(393,'admin','更新','110.88.106.120:3389','update sub group',1,'2021-08-10 17:46:20','成功'),(394,'admin','登录','110.88.106.120:3788','logout',1,'2021-08-10 17:48:51','成功'),(395,'admin','登录','110.88.106.120:3790','login',1,'2021-08-10 17:48:57','成功'),(396,'admin','登录','127.0.0.1:56112','login',0,'2021-12-03 15:53:38','密码错误'),(397,'admin','登录','127.0.0.1:56114','login',0,'2021-12-03 15:53:46','密码错误'),(398,'admin','登录','127.0.0.1:56116','login',1,'2021-12-03 15:53:49','成功'),(399,'admin','发布','127.0.0.1:56162','publish job',0,'2021-12-03 15:54:43','写入数据到磁盘失败'),(400,'admin','发布','127.0.0.1:56164','publish job',0,'2021-12-03 15:56:11','调用监控服务的API出现错误'),(401,'admin','发布','127.0.0.1:56188','publish job',0,'2021-12-03 15:56:23','调用监控服务的API出现错误'),(402,'admin','增加','127.0.0.1:56474','create websokcet',1,'2021-12-03 16:10:22','成功'),(403,'admin','重置','127.0.0.1:56480','create reset prometheus config data key',1,'2021-12-03 16:11:13','成功'),(404,'admin','运行','127.0.0.1:58510','publish sd_file',0,'2021-12-06 13:45:51','调用监控服务的API出现错误'),(405,'admin','发布','127.0.0.1:58542','publish job',0,'2021-12-06 14:47:02','调用监控服务的API出现错误'),(406,'admin','运行','127.0.0.1:58544','reload prometheus service',0,'2021-12-06 14:49:04','调用监控服务的API出现错误'),(407,'admin','增加','127.0.0.1:58708','create job',1,'2021-12-06 15:10:34','成功'),(408,'admin','增加','127.0.0.1:58718','create ip',1,'2021-12-06 15:10:57','成功'),(409,'admin','更新','127.0.0.1:58736','update manager user priv',1,'2021-12-06 15:11:18','成功'),(410,'admin','发布','127.0.0.1:58754','create prometheus config',0,'2021-12-06 15:11:32','调用监控服务的API出现错误'),(411,'admin','增加','127.0.0.1:58766','create ip',1,'2021-12-06 15:11:52','成功'),(412,'admin','发布','127.0.0.1:58784','create prometheus config',0,'2021-12-06 15:12:01','调用监控服务的API出现错误'),(413,'admin','更新','127.0.0.1:58794','get job sub group machines',1,'2021-12-06 15:12:16','成功'),(414,'admin','更新','127.0.0.1:58798','update sub group ip pool',1,'2021-12-06 15:12:18','成功'),(415,'admin','增加','127.0.0.1:58812','create ip',1,'2021-12-06 15:13:01','成功'),(416,'admin','发布','127.0.0.1:58820','create prometheus config',0,'2021-12-06 15:15:28','调用监控服务的API出现错误'),(417,'admin','发布','127.0.0.1:58856','create prometheus config',0,'2021-12-06 15:16:28','调用监控服务的API出现错误'),(418,'admin','更新','127.0.0.1:58886','get job sub group machines',1,'2021-12-06 15:19:04','成功'),(419,'admin','发布','127.0.0.1:58892','create prometheus config',0,'2021-12-06 15:23:13','执行事务出错'),(420,'admin','发布','127.0.0.1:58896','create prometheus config',0,'2021-12-06 15:33:29','执行事务出错'),(421,'admin','发布','127.0.0.1:58898','create prometheus config',0,'2021-12-06 15:34:30','调用监控服务的API出现错误'),(422,'admin','增加','127.0.0.1:58914','create ip',1,'2021-12-06 15:34:52','成功'),(423,'admin','发布','127.0.0.1:58920','create prometheus config',0,'2021-12-06 15:34:56','调用监控服务的API出现错误'),(424,'admin','运行','127.0.0.1:58960','publish sd_file',0,'2021-12-06 15:38:29','调用监控服务的API出现错误'),(425,'admin','运行','127.0.0.1:58962','publish sd_file',0,'2021-12-06 15:38:30','调用监控服务的API出现错误'),(426,'admin','运行','127.0.0.1:58972','reload prometheus service',0,'2021-12-06 15:40:18','调用监控服务的API出现错误'),(427,'admin','发布','127.0.0.1:59038','publish rule',1,'2021-12-06 15:44:39','成功'),(428,'admin','发布','127.0.0.1:59040','publish rule',1,'2021-12-06 15:44:42','成功'),(429,'admin','发布','127.0.0.1:59042','publish rule',1,'2021-12-06 15:44:42','成功'),(430,'admin','发布','127.0.0.1:59044','publish rule',1,'2021-12-06 15:44:43','成功'),(431,'admin','发布','127.0.0.1:59046','publish rule',1,'2021-12-06 15:44:44','成功'),(432,'admin','发布','127.0.0.1:59066','publish rule',1,'2021-12-06 15:44:55','成功'),(433,'admin','发布','127.0.0.1:59090','publish rule',1,'2021-12-06 15:45:29','成功'),(434,'admin','发布','127.0.0.1:59154','create prometheus config',0,'2021-12-06 15:48:11','调用监控服务的API出现错误'),(435,'admin','发布','127.0.0.1:59156','reload prometheus config',0,'2021-12-06 15:48:16','调用监控服务的API出现错误'),(436,'admin','发布','127.0.0.1:59158','create prometheus config',0,'2021-12-06 15:48:17','调用监控服务的API出现错误'),(437,'admin','发布','127.0.0.1:59160','publish rule',1,'2021-12-06 15:48:51','成功'),(438,'admin','发布','127.0.0.1:59170','publish rule',1,'2021-12-06 15:49:29','成功'),(439,'admin','发布','127.0.0.1:59180','publish rule',1,'2021-12-06 15:49:54','成功'),(440,'admin','发布','127.0.0.1:59182','publish rule',1,'2021-12-06 15:49:56','成功'),(441,'admin','发布','127.0.0.1:59196','publish rule',1,'2021-12-06 15:51:42','成功'),(442,'admin','发布','127.0.0.1:59198','publish rule',1,'2021-12-06 15:51:43','成功'),(443,'admin','发布','127.0.0.1:59200','publish rule',1,'2021-12-06 15:51:46','成功'),(444,'admin','发布','127.0.0.1:59202','publish rule',1,'2021-12-06 15:51:47','成功'),(445,'admin','发布','127.0.0.1:59250','publish rule',1,'2021-12-06 15:53:28','成功'),(446,'admin','更新','127.0.0.1:59296','update manager user priv',1,'2021-12-06 16:17:38','成功'),(447,'admin','更新','127.0.0.1:59306','update manager user priv',1,'2021-12-06 16:19:21','成功'),(448,'admin','增加','127.0.0.1:59414','create ip',1,'2021-12-06 16:21:43','成功'),(449,'admin','更新','127.0.0.1:59422','update ip',1,'2021-12-06 16:21:46','成功'),(450,'admin','发布','127.0.0.1:59430','create prometheus config',0,'2021-12-06 16:21:51','调用监控服务的API出现错误'),(451,'admin','增加','127.0.0.1:59476','create websokcet',1,'2021-12-06 16:22:28','成功'),(452,'admin','增加','127.0.0.1:59502','create websokcet',1,'2021-12-06 16:22:37','成功'),(453,'admin','重置','127.0.0.1:59522','create reset prometheus config data key',1,'2021-12-06 16:22:47','成功');
/*!40000 ALTER TABLE `system_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tmpl`
--

DROP TABLE IF EXISTS `tmpl`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tmpl` (
  `tmpl` longtext NOT NULL,
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_by` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='prometheus.yml';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tmpl`
--

LOCK TABLES `tmpl` WRITE;
/*!40000 ALTER TABLE `tmpl` DISABLE KEYS */;
INSERT INTO `tmpl` VALUES ('# my global config\nglobal:\n  scrape_interval:     15s # Set the scrape interval to every 15 seconds. Default is every 1 minute.\n  evaluation_interval: 15s # Evaluate rules every 15 seconds. The default is every 1 minute.\n  # scrape_timeout is set to the global default (10s).\n  # scrape_timeout: 60s\n\n# Alertmanager configuration\nalerting:\n  alertmanagers:\n  - static_configs:\n    - targets:\n      - 127.0.0.1:9093\n\n# Load rules once and periodically evaluate them according to the global \'evaluation_interval\'.\nrule_files:\n   - \"{{.RelRuleDir}}/*.yml\"\n  # - \"first_rules.yml\"\n  # - \"second_rules.yml\"\n\n# A scrape configuration containing exactly one endpoint to scrape:\n# Here it\'s Prometheus itself.\nscrape_configs:\n  # The job name is added as a label `job=<job_name>` to any timeseries scraped from this config.\n  - job_name: \'监控服务本机\'\n\n    # metrics_path defaults to \'/metrics\'\n    # scheme defaults to \'http\'.\n\n    static_configs:\n    - targets: [\'localhost:9090\']\n{{ range .Jobs }}\n  - job_name: \'{{.Name}}\'\n    file_sd_configs:\n      - files:\n        - \"{{$.AbsConfDir}}/{{.Name}}.json\"\n        refresh_interval: {{$.Fields.refresh_interval}}\n{{.Code}}\n{{ end }}\n','2021-12-06 16:23:01','administrator');
/*!40000 ALTER TABLE `tmpl` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tmpl_fields`
--

DROP TABLE IF EXISTS `tmpl_fields`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tmpl_fields` (
  `id` int NOT NULL AUTO_INCREMENT,
  `key` varchar(100) NOT NULL,
  `value` varchar(500) NOT NULL,
  `enabled` tinyint NOT NULL,
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_by` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `key_unique` (`key`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb3 COMMENT='模板字段';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tmpl_fields`
--

LOCK TABLES `tmpl_fields` WRITE;
/*!40000 ALTER TABLE `tmpl_fields` DISABLE KEYS */;
INSERT INTO `tmpl_fields` VALUES (1,'metrics','/metrics',1,'2021-12-06 16:23:01',''),(2,'refresh_interval','15s',1,'2021-12-06 16:23:01','');
/*!40000 ALTER TABLE `tmpl_fields` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-12-06 16:26:13
