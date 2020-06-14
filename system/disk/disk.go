package disk

import (
	"encoding/json"
	"time"

	"github.com/shirou/gopsutil/disk"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

type DiskInfo struct {
	Path              string  `json:"path"`              //挂着点
	Device            string  `json:"device"`            //设备名称
	Fstype            string  `json:"fstype"`            //文件系统类型
	Total             uint64  `json:"total"`             //总计大小（MB）
	Free              uint64  `json:"free"`              //剩余大小（MB）
	Used              uint64  `json:"used"`              //已使用大小（MB）
	UsedPercent       float64 `json:"usedPercent"`       //使用率
	InodesTotal       uint64  `json:"inodesTotal"`       //inode统计大小
	InodesUsed        uint64  `json:"inodesUsed"`        //inode已使用大小
	InodesFree        uint64  `json:"inodesFree"`        //inode剩余大小
	InodesUsedPercent float64 `json:"inodesUsedPercent"` //inode使用率
	LastTime          string  `json:"lastTime"`          //采集时间
}

func (o DiskInfo) String() string {
	s, _ := json.Marshal(o)
	return string(s)
}

var currentDiskInfo = NewDiskInfo()

func NewDiskInfo() *DiskInfo {
	return new(DiskInfo)
}

// 获取根路径磁盘信息
func GetDiskInfo() *DiskInfo {
	return GetDiskInfoByPath("/")
}

// 获取磁盘信息
func GetDiskInfoByPath(path string) *DiskInfo {
	d, err := disk.Usage(path)
	diskInfo := NewDiskInfo()
	if err == nil {
		diskInfo.Path = d.Path
		diskInfo.Fstype = d.Fstype
		diskInfo.Total = d.Total / MB
		diskInfo.Used = d.Used / MB
		diskInfo.Free = d.Free / MB
		diskInfo.UsedPercent = d.UsedPercent
		diskInfo.InodesTotal = d.InodesTotal
		diskInfo.InodesUsed = d.InodesUsed
		diskInfo.InodesFree = d.InodesFree
		diskInfo.InodesUsedPercent = d.InodesUsedPercent
		diskInfo.LastTime = time.Now().Format("2006-01-02 15:04:05")
	}
	currentDiskInfo = diskInfo
	return diskInfo
}

// GetAllDiskInfo 获取所有物理磁盘信息
// Partitions returns disk partitions. If all is false, returns
// physical devices only (e.g. hard disks, cd-rom drives, USB keys)
// and ignore all others (e.g. memory partitions such as /dev/shm)
func GetAllPhysicalDiskInfo() []*DiskInfo {
	return GetAllDiskInfo(false)
}

// GetAllDiskInfo Partitions returns disk partitions. If all is false, returns
// physical devices only (e.g. hard disks, cd-rom drives, USB keys)
// and ignore all others (e.g. memory partitions such as /dev/shm)
func GetAllDiskInfo(all bool) []*DiskInfo {
	s, _ := disk.Partitions(all)
	allDisk := make([]*DiskInfo, 0, 5)
	if s != nil && len(s) > 0 {
		for _, v := range s {
			path := v.Mountpoint
			info := GetDiskInfoByPath(path)
			info.Device = v.Device
			allDisk = append(allDisk, info)
		}
	}
	return allDisk
}

func ge() {
	disk.IOCounters()
}
