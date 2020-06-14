package cpu

import (
	"os"
	"runtime"
	"testing"
	"time"

	. "github.com/shirou/gopsutil/cpu"
)

func TestGetCpuInfo(t *testing.T) {
	for i := 0; i < 10; i++ {
		//GetCpuInfo()
		t.Log(GetCpuInfo())
		time.Sleep(time.Second)
	}

}

func testCPUPercent(t *testing.T, percpu bool) {
	numcpu := runtime.NumCPU()
	testCount := 3

	if runtime.GOOS != "windows" {
		testCount = 100
		v, err := Percent(time.Millisecond, percpu)
		if err != nil {
			t.Errorf("error %v", err)
		}
		// Skip CircleCI which CPU num is different
		if os.Getenv("CIRCLECI") != "true" {
			if (percpu && len(v) != numcpu) || (!percpu && len(v) != 1) {
				t.Fatalf("wrong number of entries from CPUPercent: %v", v)
			}
		}
	}
	for i := 0; i < testCount; i++ {
		duration := time.Duration(10) * time.Microsecond
		v, err := Percent(duration, percpu)
		if err != nil {
			t.Errorf("error %v", err)
		}
		for _, percent := range v {
			// Check for slightly greater then 100% to account for any rounding issues.
			if percent < 0.0 || percent > 100.0001*float64(numcpu) {
				t.Fatalf("CPUPercent value is invalid: %f", percent)
			}
		}
	}
}

func TestCPUPercent(t *testing.T) {
	testCPUPercent(t, false)
}

func TestCPUPercentPerCpu(t *testing.T) {
	testCPUPercent(t, true)
}
