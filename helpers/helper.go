package helpers

import "testing"

const (
	PartA = iota
	PartB
)

func Check (t testing.TB, expect, got int) {
	t.Helper()
	if got != expect {
		t.Errorf("Expected %d, got %d", expect, got)
	}
}