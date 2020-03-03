package main

import "testing"

func TestSubFive(t *testing.T) {
	x := subFive(26)
	if x != 21 {
		t.Error("Expected: 21 got:", x)
	}
}