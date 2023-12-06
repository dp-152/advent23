package main

import (
	"testing"
)

const cEXAMPLE_FILE = "input.example.txt"

func TestExampleInput__Challenge1(t *testing.T) {
	expected := 4361
	engine := loadEngineSchematic(cEXAMPLE_FILE)

	result := sumPartNumbers(engine)

	if result != expected {
		t.Errorf("Failed challenge 1: expected %d, got %d", expected, result)
	}
}
