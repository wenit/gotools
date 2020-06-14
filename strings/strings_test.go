package strings

import "testing"

func TestUUID(t *testing.T) {
	t.Log(UUID())
}

func TestShortUUID(t *testing.T) {
	t.Log(ShortUUID())
}
