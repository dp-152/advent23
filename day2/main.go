package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dp-152/advent23/common/fs"
	"github.com/dp-152/advent23/common/path"
)

const cGAME_PREFIX = "Game "
const cRED_SUFFIX = "red"
const cGREEN_SUFFIX = "green"
const cBLUE_SUFFIX = "blue"
const cSET_SEPARATOR = "; "
const cVAL_SEPARATOR = ", "
const cGAME_SEPARATOR = ": "

var constraint = &cubeset{
	red:   12,
	green: 13,
	blue:  14,
}

type cubeset struct {
	red,
	green,
	blue int64
}

func (c *cubeset) parseVal(valstr string) {
	split := strings.Split(valstr, " ")
	ct, err := strconv.ParseInt(split[0], 10, 32)

	if err != nil {
		panic(err)
	}

	switch split[1] {
	case cRED_SUFFIX:
		c.red = ct
	case cGREEN_SUFFIX:
		c.green = ct
	case cBLUE_SUFFIX:
		c.blue = ct
	}
}

func main() {
	var sum int64 = 0

	cancelChan := make(chan struct{})
	fileChan := fs.ReadLines(fmt.Sprintf("%s/input.txt", path.OwnPath()), cancelChan)

	for line := range fileChan {
		index, overflow := parseGame(line)
		if !overflow {
			sum += index
		}
	}

	fmt.Printf("Sum of lines: %d\n", sum)
}

func parseGame(line string) (index int64, overflow bool) {
	game, ok := strings.CutPrefix(line, cGAME_PREFIX)
	if !ok {
		panic(fmt.Errorf("cannot parse game: prefix \"%s\" not found in line %s[...]", cGAME_PREFIX, line[:20]))
	}

	idx, game, ok := strings.Cut(game, cGAME_SEPARATOR)
	if !ok {
		panic(fmt.Errorf("cannot parse game: cannot get index: separator \"%s\" not found in line %s[...]", cGAME_SEPARATOR, game[:20]))
	}

	index, err := strconv.ParseInt(idx, 10, 64)
	if err != nil {
		panic(err)
	}

	overflow = parseSets(game)

	return
}

func parseSets(game string) (overflow bool) {
	for _, setstr := range strings.Split(game, cSET_SEPARATOR) {
		if parseVals(setstr) {
			overflow = true
			break
		}
	}

	return
}

func parseVals(setstr string) (overflow bool) {
	set := new(cubeset)
	for _, valstr := range strings.Split(setstr, cVAL_SEPARATOR) {
		set.parseVal(valstr)
	}

	return isSetOverflow(set)
}

func isSetOverflow(set *cubeset) bool {
	return set.red > constraint.red || set.green > constraint.green || set.blue > constraint.blue
}
