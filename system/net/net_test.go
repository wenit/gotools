package net

import (
	"testing"

	. "github.com/wenit/gotools/system"
)

func TestGetConnections(t *testing.T) {
	t.Log(GetConnections())
}

func TestGetConnectionsPid(t *testing.T) {
	t.Log(GetConnectionsPid(int32(GetPid())))
}

func TestGetNetInterfaces(t *testing.T) {
	t.Log(GetNetInterfaces())
}

func TestGetConnectionsCurrentPid(t *testing.T) {
	t.Log(GetConnectionsCurrentPid())
}
