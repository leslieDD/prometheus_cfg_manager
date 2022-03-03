#!/bin/bash
# 编译前端
cd admin
npm install
npm run build
cd ..
# 编译后端
go build
# 清理和备份旧文件
rm -rf /usr/local/prometheus_cfg_manager/pro_cfg_manager
mkdir -p /usr/local/prometheus_cfg_manager
rm -rf /usr/local/prometheus_cfg_manager/static
# 拷贝新文件
cp pro_cfg_manager /usr/local/prometheus_cfg_manager
if [ ! -f /usr/local/prometheus_cfg_manager/config.yml ];then
    # mv /usr/local/prometheus_cfg_manager/config.yml /usr/local/prometheus_cfg_manager/config.yml.$(date +"%s") 
    cp -f ./config.yml /usr/local/prometheus_cfg_manager/config.yml
fi
cp -rf ./admin/dist /usr/local/prometheus_cfg_manager/static
# 启动服务，需要把pro_cfg_manager.service注册到系统中
systemctl restart pro_cfg_manager
