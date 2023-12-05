package main

import (
	"fmt"
	"path"
	"strings"

	"github.com/dp-152/advent23/common/fs"
	p "github.com/dp-152/advent23/common/path"
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
	sum := parseFile(path.Join(p.OwnPath(), "input.txt"))

	fmt.Printf("Sum of lines: %d\n", sum)
}

func parseFile(inputFile string) (sum int) {
	cancelChan := make(chan struct{})
	fileChan := fs.ReadLines(inputFile, cancelChan)

	for line := range fileChan {
		sum += lineValue(line)
	}

	return
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
