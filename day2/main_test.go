package main

import (
	"path"
	"testing"

	p "github.com/dp-152/advent23/common/path"
)

const cEXAMPLE_FILE = "input.example.txt"

func TestExampleInput__Challenge1(t *testing.T) {
	expected := int64(8)
	result, _ := parseFile(path.Join(p.OwnPath(), cEXAMPLE_FILE))

	if result != expected {
		t.Errorf("Failed challenge 1: expected %d, got %d", expected, result)
	}
}

func TestExampleInput__Challenge2(t *testing.T) {
	expected := int64(2286)
	_, result := parseFile(path.Join(p.OwnPath(), cEXAMPLE_FILE))

	if result != expected {
		t.Errorf("Failed challenge 2: expected %d, got %d", expected, result)
	}
}
