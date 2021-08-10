#!/bin/bash
go build
rm -rf /usr/local/prometheus_cfg_manager/pro_cfg_manager
cp pro_cfg_manager /usr/local/prometheus_cfg_manager
if [ -f /usr/local/prometheus_cfg_manager/config.yml ];then
    mv /usr/local/prometheus_cfg_manager/config.yml /usr/local/prometheus_cfg_manager/config.yml.$(date +"%s")
fi
cp -f ./config.yml /usr/local/prometheus_cfg_manager/config.yml
rm -rf /usr/local/prometheus_cfg_manager/static
cp -rf ./static /usr/local/prometheus_cfg_manager
systemctl restart pro_cfg_manager