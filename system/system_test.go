package system

import "testing"

func TestCurrentDirectory(t *testing.T) {
	t.Log(CurrentDirectory())
}

func TestParentDirectory(t *testing.T) {
	t.Log(ParentDirectory("/app/logs"))
}

func TestUIDAndGID(t *testing.T) {
	t.Log(UIDAndGID())
}

func TestLookupGid(t *testing.T) {
	t.Log(LookupGid("root"))
}

func TestLookupUIDAndGID(t *testing.T) {
	t.Log(LookupUIDAndGID("root"))
}

func TestGetPid(t *testing.T) {
	t.Log(GetPid())
}

func TestGetWd(t *testing.T) {
	t.Log(GetWd())
}

func TestGetTime(t *testing.T) {
	t.Log(GetTime())
}

func TestGetTimestamp(t *testing.T) {
	t.Log(GetTimestamp())
}

func TestGetTimestampMilli(t *testing.T) {
	t.Log(GetTimestampMilli())
}

func TestNumGoroutine(t *testing.T) {
	t.Log(NumGoroutine())
}
