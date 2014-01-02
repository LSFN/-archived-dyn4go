package dyn4go

import (
	"testing"
)

func TestEpsilon(t *testing.T) {
	if Epsilon == 0 {
		t.Error("Epsilon is equal to 0")
	}
	if 1 + Epsilon != 1 {
		t.Error("Epsilon computed incorrectly: 1 + Epsilon != 1")
	}
}