package process

import (
	"testing"

	"github.com/shirou/gopsutil/process"
	. "github.com/wenit/gotools/system"
)

func TestProcess(t *testing.T) {
	t.Log(GetPid())
	//t.Log(process.Pids())
	p, _ := process.NewProcess(int32(GetPid()))
	t.Log(p.Cmdline())
	t.Log(p.Status())
	t.Log(p.MemoryInfo())
	t.Log(p.Threads())
	t.Log(p.Uids())
	t.Log(p.Name())
	t.Log(p.CreateTime())
	t.Log(p.Exe())
	t.Log(p.Cwd())
}

func TestGetProcessInfo(t *testing.T) {
	t.Log(GetProcessInfo())
}
