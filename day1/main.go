package main

import (
	"fmt"
	"strings"

	"github.com/dp-152/advent23/common/fs"
	"github.com/dp-152/advent23/common/path"
)

var nums = []string{"nil", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var maxchars int

func init() {
	for _, s := range nums {
		l := len(s)
		if l > maxchars {
			maxchars = l
		}
	}
}

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

	for pos, runeVal := range line {
		val, ok := valFromNum(runeVal)
		if ok {
			if left < 1 {
				left = val
			}
			right = val
			continue
		}

		val, ok = valFromStr(line[pos:toEnd(len(line), pos, 5)])
		if ok {
			if left < 1 {
				left = val
			}
			right = val
		}
	}
	return (left * 10) + right
}

func valFromNum(runeVal rune) (val int, ok bool) {
	num := int(runeVal - 0x30)
	if num <= 9 {
		ok = true
		val = num
	}

	return
}

func valFromStr(cmp string) (val int, ok bool) {
	for num, word := range nums {
		if strings.HasPrefix(cmp, word) {
			ok = true
			val = num
		}
	}

	return
}

func toEnd(length, offset, size int) int {
	index := offset + size
	overflow := index - length
	finalIndex := index - overflow

	return finalIndex
}
