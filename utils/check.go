package utils

import "net"

func CheckIPAddr(ipaddr string) bool {
	return net.ParseIP(ipaddr) == nil
}
