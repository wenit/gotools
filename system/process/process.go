package process

import (
	"encoding/json"
	"time"

	"github.com/shirou/gopsutil/process"
	. "github.com/wenit/gotools/system"
)

type ProcessInfo struct {
	Pid        int32  `json:"pid"`
	Name       string `json:"name"`
	Cwd        string `json:"cwd"`
	Status     string `json:"status"`
	Parent     int32  `json:"parent"`
	NumThreads int32  `json:"numThreads"`
	Cmdline    string `json:"cmdline"`
	CreateTime string `json:"createTime"`
	LastTime   string `json:"lastTime"` //采集时间
}

func NewProcessInfo() *ProcessInfo {
	return new(ProcessInfo)
}

func (o ProcessInfo) String() string {
	s, _ := json.Marshal(o)
	return string(s)
}

func GetProcessInfo() *ProcessInfo {
	p, _ := process.NewProcess(int32(GetPid()))

	processInfo := NewProcessInfo()

	processInfo.Pid = p.Pid
	name, _ := p.Name()
	processInfo.Name = name
	cmdline, _ := p.Cmdline()
	processInfo.Cmdline = cmdline

	cwd, _ := p.Cwd()

	processInfo.Cwd = cwd

	createTime, _ := p.CreateTime()
	processInfo.CreateTime = time.Unix(int64(createTime/1000), 0).Format("2006-01-02 15:04:05")
	processInfo.LastTime = GetTime()
	return processInfo
}
