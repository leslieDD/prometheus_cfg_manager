#### Prometheus Config Manager

##### 介绍

`prometheus_cfg_manager`是用来管理`prometheus`的相关配置，包括`prometheus.yml`，告警规则，及服务发现等

##### 安装

* 在本项目库中有一SQL文件：`prometheus_cfg_manager.sql`，导入到`mysql`中即可
* 配置本项目库中`config.yml`配置文件，根据里面的提示进行配置
* 运行`./install.sh`脚本，即可安装完成

##### 使用

在浏览器中打开`http://x.x.x.x:8200`(默认)即可使用，其中：`x.x.x.x`是服务器IP地址，`8200`为服务的监听地址，如果在`config.yml`中配置了其它端口，请使用相应的端口号。默认账号密码是：admin/admin

##### 附

告警规则可以参考：[https://awesome-prometheus-alerts.grep.to/rules.html](https://awesome-prometheus-alerts.grep.to/rules.html)，支持此网站规则导入到本服务中，在`告警管理`页面中操作。`告警管理`中支持右键下拉菜单进行操作



##### 以下是部分截图

![](https://github.com/leslieDD/prometheus_cfg_manager/blob/master/images/person.png)

![](https://github.com/leslieDD/prometheus_cfg_manager/blob/master/images/ipManager.png)

![](https://github.com/leslieDD/prometheus_cfg_manager/blob/master/images/group.png)

![](https://github.com/leslieDD/prometheus_cfg_manager/blob/master/images/ymlView.png)

![](https://github.com/leslieDD/prometheus_cfg_manager/blob/master/images/admin.png)
