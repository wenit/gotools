package cpu

import (
	"encoding/json"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
)

var currentCpuInfo = NewCpuInfo()

var hasInt = false

func init() {
	c := time.Tick(1000 * time.Millisecond)
	go func() {
		for {
			<-c
			getCpuInfo()

		}
	}()
	hasInt = true
}

type CpuInfo struct {
	ModelName    string  `json:"modelName"`    //CPU型号
	Cores        int     `json:"cores"`        //CPU核心数
	LogicalCores int     `json:"logicalCores"` //CPU逻辑核心数
	Percent      float64 `json:"percent"`      //CPU使用率
	Load1        float64 `json:"load1"`        //1分钟平均负载
	Load5        float64 `json:"load5"`        //5分钟平均负载
	Load15       float64 `json:"load15"`       //15分钟平均负载
	LastTime     string  `json:"lastTime"`     //采集时间
}

func (o CpuInfo) String() string {
	s, _ := json.Marshal(o)
	return string(s)
}

func NewCpuInfo() *CpuInfo {
	return new(CpuInfo)
}

func getCpuInfo() *CpuInfo {
	cores, _ := cpu.Counts(false)
	logicalCores, _ := cpu.Counts(true)
	a, _ := load.Avg()
	l1 := a.Load1
	l5 := a.Load5
	l15 := a.Load15

	c, _ := cpu.Info()

	percents, _ := cpu.Percent(time.Millisecond*1000, false)

	cpuInfo := currentCpuInfo
	cpuInfo.Cores = cores
	cpuInfo.LogicalCores = logicalCores

	if len(percents) > 0 {
		cpuInfo.Percent = percents[0]
	}

	if len(c) > 0 {
		cpuInfo.ModelName = c[0].ModelName
	}

	cpuInfo.Load1 = l1
	cpuInfo.Load5 = l5
	cpuInfo.Load15 = l15

	cpuInfo.LastTime = time.Now().Format("2006-01-02 15:04:05")

	//log.Println(cpuInfo.Percent)

	return cpuInfo
}

func GetCpuInfo() *CpuInfo {
	return currentCpuInfo
}
