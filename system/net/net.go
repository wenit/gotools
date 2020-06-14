package net

import (
	"github.com/shirou/gopsutil/net"
	. "github.com/wenit/gotools/system"
)

func GetConnections() []net.ConnectionStat {
	c, _ := net.Connections("all")
	return c
}

func GetConnectionsPid(pid int32) []net.ConnectionStat {
	c, _ := net.ConnectionsPid("all", pid)
	return c
}

func GetConnectionsCurrentPid() []net.ConnectionStat {
	c, _ := net.ConnectionsPid("all", int32(GetPid()))
	return c
}

func GetNetInterfaces() []net.InterfaceStat {
	i, _ := net.Interfaces()
	return i
}
