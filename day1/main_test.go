package main

import (
	"path"
	"testing"

	p "github.com/dp-152/advent23/common/path"
)

const cEXAMPLE_FILE_1 = "input.example1.txt"
const cEXAMPLE_FILE_2 = "input.example2.txt"

func TestExampleInput__Challenge1(t *testing.T) {
	expected := 142
	result := parseFile(path.Join(p.OwnPath(), cEXAMPLE_FILE_1))

	if result != expected {
		t.Errorf("Failed challenge 1: expected %d, got %d", expected, result)
	}
}

func TestExampleInput__Challenge2(t *testing.T) {
	expected := 281
	result := parseFile(path.Join(p.OwnPath(), cEXAMPLE_FILE_2))

	if result != expected {
		t.Errorf("Failed challenge 2: expected %d, got %d", expected, result)
	}
}
