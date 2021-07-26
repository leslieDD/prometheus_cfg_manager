package utils

import (
	"net"

	"github.com/shirou/gopsutil/host"
)

func CheckIPAddr(ipaddr string) bool {
	return net.ParseIP(ipaddr) == nil
}

func HostInfo() *host.InfoStat {
	info, err := host.Info()
	if err != nil {
		return nil
	}
	return info
}
