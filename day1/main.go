package main

import (
	"fmt"

	"github.com/dp-152/advent23/common/fs"
	"github.com/dp-152/advent23/common/path"
)

func main() {
	sum := 0

	cancelChan := make(chan struct{})
	fileChan := fs.ReadLines(fmt.Sprintf("%s/input.txt", path.OwnPath()), cancelChan)

	for line := range fileChan {
		sum += lineValue(line)
	}

	fmt.Printf("Sum of lines: %d\n", sum)
}

func lineValue(line string) int {
	left, right := 0, 0

	for _, runeVal := range line {
		num := int(runeVal - 0x30)
		if num <= 9 {
			if left < 1 {
				left = num
			}
			right = num
		}
	}
	return (left * 10) + right
}
