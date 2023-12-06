package main

import (
	"testing"
)

const cEXAMPLE_FILE = "input.example.txt"

func TestExampleInput__Challenge1(t *testing.T) {
	expected := 13

	result := sumPoints(cEXAMPLE_FILE)

	if result != expected {
		t.Errorf("Failed challenge 1: expected %d, got %d", expected, result)
	}
}

func TestExampleInput__Challenge2(t *testing.T) {
	expected := 30

	result := countCopies(cEXAMPLE_FILE)

	if result != expected {
		t.Errorf("Failed challenge 2: expected: %d, got %d", expected, result)
	}
}
