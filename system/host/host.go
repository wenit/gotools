package host

import (
	"encoding/json"
	"time"

	syshost "github.com/shirou/gopsutil/host"
)

// HostInfo 主机信息
type HostInfo struct {
	Hostname        string `json:"hostname"`        //主机名称
	Uptime          uint64 `json:"uptime"`          //已运行时间，单位秒
	BootTime        string `json:"bootTime"`        //启动时间，格式：2006-01-02 15:04:05
	Procs           uint64 `json:"procs"`           // 进程数量
	OS              string `json:"os"`              // ex: freebsd, linux
	Platform        string `json:"platform"`        // ex: ubuntu, linuxmint
	PlatformFamily  string `json:"platformFamily"`  // ex: debian, rhel
	PlatformVersion string `json:"platformVersion"` // version of the complete OS
	KernelVersion   string `json:"kernelVersion"`   // version of the OS kernel (if available)
	KernelArch      string `json:"kernelArch"`      // native cpu architecture queried at runtime, as returned by `uname -m` or empty string in case of error
	HostID          string `json:"hostId"`          // ex: uuid
	LastTime        string `json:"lastTime"`        //采集时间
}

func (o HostInfo) String() string {
	s, _ := json.Marshal(o)
	return string(s)
}

var currentHostInfo = NewHostInfo()

// NewHostInfo New Host info
func NewHostInfo() *HostInfo {
	return new(HostInfo)
}

// GetHostInfo 获取主机信息
func GetHostInfo() *HostInfo {
	h, err := syshost.Info()

	hostInfo := currentHostInfo
	if err == nil {
		hostInfo.BootTime = time.Unix(int64(h.BootTime), 0).Format("2006-01-02 15:04:05")
		hostInfo.Uptime = h.Uptime
		hostInfo.HostID = h.HostID
		hostInfo.Hostname = h.Hostname
		hostInfo.KernelArch = h.KernelArch
		hostInfo.KernelVersion = h.KernelVersion
		hostInfo.OS = h.OS
		hostInfo.Platform = h.Platform
		hostInfo.PlatformFamily = h.PlatformFamily
		hostInfo.PlatformVersion = h.PlatformVersion
		hostInfo.Procs = h.Procs
		hostInfo.LastTime = time.Now().Format("2006-01-02 15:04:05")
	}
	return hostInfo
}
