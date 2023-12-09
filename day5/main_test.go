package main

import (
	"testing"
)

const cEXAMPLE_FILE = "input.example.txt"

func TestExampleInput__Challenge1(t *testing.T) {
	var expected uint64 = 35

	result := lowestLoc(cEXAMPLE_FILE)

	if result != expected {
		t.Errorf("Failed challenge 1: expected %d, got %d", expected, result)
	}
}
