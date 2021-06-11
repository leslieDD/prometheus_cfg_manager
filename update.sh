#!/bin/bash
go build
rm -rf /usr/local/prometheus_cfg_manager/pro_cfg_manager
cp pro_cfg_manager /usr/local/prometheus_cfg_manager
rm -rf /usr/local/prometheus_cfg_manager/static
cp -rf ./static /usr/local/prometheus_cfg_manager/
systemctl restart pro_cfg_manager