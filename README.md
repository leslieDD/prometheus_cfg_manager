#### Prometheus Config Manager

##### 介绍

`prometheus_cfg_manager`是用来管理`prometheus`的相关配置，包括`prometheus.yml`，告警规则，及服务发现等

##### 安装

* 在本项目库中有一SQL文件：`prometheus_cfg_manager.sql`，导入到`mysql`中即可
* 配置本项目库中`config.yml`配置文件，根据里面的提示进行配置
* 运行`./install.sh`脚本，即可安装完成

##### 使用

在浏览器中打开`http://x.x.x.x:8200`(默认)即可使用，其中：`x.x.x.x`是服务器IP地址，`8200`为服务的监听地址，如果在`config.yml`中配置了其它端口，请使用相应的端口号。默认账号密码是：admin/admin

