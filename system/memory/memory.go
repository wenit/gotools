package memory

import (
	"encoding/json"
	"time"

	"github.com/shirou/gopsutil/mem"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

type MemoryInfo struct {
	Total     uint64  `json:"total"`     //总共内存（MB）
	Available uint64  `json:"available"` //实际可用内存（MB）
	Used      uint64  `json:"used"`      //已用内存（MB）
	Free      uint64  `json:"free"`      //剩余内存（MB）
	Percent   float64 `json:"percent"`   //使用率
	LastTime  string  `json:"lastTime"`  //采集时间
}

func (o MemoryInfo) String() string {
	s, _ := json.Marshal(o)
	return string(s)
}

var currentMemoryInfo = NewMemoryInfo()

// 创建内存信息对象
func NewMemoryInfo() *MemoryInfo {
	return new(MemoryInfo)
}

func GetMemoryInfo() *MemoryInfo {
	memoryInfo := currentMemoryInfo
	m, err := mem.VirtualMemory()
	if err == nil {
		memoryInfo.Total = m.Total / MB
		memoryInfo.Available = m.Available / MB
		memoryInfo.Used = m.Used / MB
		memoryInfo.Free = m.Free / MB
		memoryInfo.Percent = m.UsedPercent
		memoryInfo.LastTime = time.Now().Format("2006-01-02 15:04:05")
	}
	return memoryInfo
}
