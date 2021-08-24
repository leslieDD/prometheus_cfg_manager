#!/bin/bash
git pull
go build
systemctl stop pro_cfg_manager
cp pro_cfg_manager /usr/local/prometheus_cfg_manager
systemctl restart pro_cfg_manager