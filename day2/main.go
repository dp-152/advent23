package main

import (
	"fmt"
	"path"
	"strconv"
	"strings"

	"github.com/dp-152/advent23/common/fs"
	p "github.com/dp-152/advent23/common/path"
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

func (c *cubeset) inflate(src *cubeset) {
	if src.red > c.red {
		c.red = src.red
	}

	if src.green > c.green {
		c.green = src.green
	}

	if src.blue > c.blue {
		c.blue = src.blue
	}
}

func (c *cubeset) pow() int64 {
	return c.red * c.green * c.blue
}

func main() {
	idxSum, powSum := parseFile(path.Join(p.OwnPath(), "input.txt"))
	fmt.Printf("Sum of lines: %d\n", idxSum)
	fmt.Printf("Sum of powers: %d\n", powSum)
}

func parseFile(inputFile string) (idxSum, powSum int64) {
	cancelChan := make(chan struct{})
	fileChan := fs.ReadLines(inputFile, cancelChan)

	for line := range fileChan {
		index, overflow, pow := parseGame(line)
		if !overflow {
			idxSum += index
		}
		powSum += pow
	}
	return
}

func parseGame(line string) (index int64, overflow bool, pow int64) {
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

	overflow, pow = parseSets(game)

	return
}

func parseSets(game string) (overflow bool, pow int64) {
	minset := new(cubeset)
	for _, setstr := range strings.Split(game, cSET_SEPARATOR) {
		set, of := parseVals(setstr)
		if of {
			overflow = true
		}

		minset.inflate(set)
	}
	pow = minset.pow()

	return
}

func parseVals(setstr string) (set *cubeset, overflow bool) {
	set = new(cubeset)
	for _, valstr := range strings.Split(setstr, cVAL_SEPARATOR) {
		set.parseVal(valstr)
	}

	overflow = isSetOverflow(set)

	return
}

func isSetOverflow(set *cubeset) bool {
	return set.red > constraint.red || set.green > constraint.green || set.blue > constraint.blue
}
