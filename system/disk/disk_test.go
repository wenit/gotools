package disk

import (
	"testing"

	"github.com/shirou/gopsutil/disk"
)

func TestGetDiskInfo(t *testing.T) {
	t.Log(GetDiskInfo())
}

func TestGetDiskInfoByPath(t *testing.T) {
	t.Log(GetDiskInfoByPath("/"))
}

func TestGetAllDiskInfo(t *testing.T) {
	t.Log(GetAllDiskInfo(true))
}

func TestGetAllPhysicalDiskInfo(t *testing.T) {
	t.Log(GetAllPhysicalDiskInfo())
}

func TestGetDiskInfo2(t *testing.T) {
	t.Log(disk.IOCounters())
	t.Log(disk.Partitions(false))
}
