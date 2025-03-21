package main

import (
	"math"
	"testing"
)

func TestDummy(t *testing.T) {
	got := int(math.Abs(-1))
	want := 1
	if got != want {
		t.Errorf("Abs(-1) = %d, want %d", got, want)
	}
}
