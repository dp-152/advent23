package main

import (
	"fmt"
	"path"
	"strconv"
	"strings"

	"github.com/dp-152/advent23/common/fs"
	p "github.com/dp-152/advent23/common/path"
)

type eAlmanacElementType int

const (
	eSEED eAlmanacElementType = iota
	eSOIL
	eFERT
	eWATER
	eLIGHT
	eTEMP
	eHUM
	eLOC
)

type rangeMapEntry struct {
	sourceStart uint64
	targetStart uint64
	rang        uint64
}

type rangeMap struct {
	sourceType eAlmanacElementType
	targetType eAlmanacElementType
	entries    []*rangeMapEntry
	_isSorted  bool
}

type almanac struct {
	seeds    []uint64
	bySource map[eAlmanacElementType]*rangeMap
	byTarget map[eAlmanacElementType]*rangeMap
}

var elementTypeMap = map[string]eAlmanacElementType{
	"seed":        eSEED,
	"soil":        eSOIL,
	"fertilizer":  eFERT,
	"water":       eWATER,
	"light":       eLIGHT,
	"temperature": eTEMP,
	"humidity":    eHUM,
	"location":    eLOC,
}

func (m *rangeMap) ensureSorted() {
	if m._isSorted {
		return
	}

	quicksort(m.entries, 0, len(m.entries)-1)
	m._isSorted = true
}

func (m *rangeMap) AddRange(rang *rangeMapEntry) {
	m._isSorted = false
	m.entries = append(m.entries, rang)
}

func (m *rangeMap) Lookup(val uint64) (entry *rangeMapEntry, found bool) {
	m.ensureSorted()

	left, right := 0, len(m.entries)-1

L:
	for left <= right {
		index := ((right - left) / 2) + left
		entry = m.entries[index]
		rangeStart := entry.sourceStart

		rangeEnd := rangeStart + entry.rang

		switch {
		case val >= rangeStart && val <= rangeEnd:
			found = true
			break L
		case val > rangeEnd:
			left = index + 1
		case val < rangeStart:
			right = index - 1
		}
	}

	return
}

func main() {
	location := lowestLoc(path.Join(p.OwnPath(), "input.txt"))

	fmt.Printf("Lowest location: %d", location)
}

func lowestLoc(inputFile string) uint64 {
	cancelChan := make(chan struct{})
	fileChan := fs.ReadLines(inputFile, cancelChan)

	alman := &almanac{
		seeds:    make([]uint64, 0, 10),
		bySource: make(map[eAlmanacElementType]*rangeMap),
		byTarget: make(map[eAlmanacElementType]*rangeMap),
	}

	var currMap *rangeMap = nil
	for line := range fileChan {
		currMap = parseLine(alman, currMap, line)
	}

	locs := make(map[uint64]uint64)

	for _, seed := range alman.seeds {
		var lookup uint64 = seed
		for _, elType := range []eAlmanacElementType{eSEED, eSOIL, eFERT, eWATER, eLIGHT, eTEMP, eHUM} {
			entry := getRange(lookup, alman, elType)
			offset := lookup - entry.sourceStart
			lookup = entry.targetStart + offset
		}
		locs[seed] = lookup
	}

	var result uint64 = 0xffff_ffff_ffff_ffff
	for loc := range locs {
		if locs[loc] < result {
			result = locs[loc]
		}
	}

	return result
}

func getRange(lookup uint64, alman *almanac, elType eAlmanacElementType) (entry *rangeMapEntry) {
	var found bool
	entry, found = alman.bySource[elType].Lookup(lookup)

	if !found {
		entry = &rangeMapEntry{
			sourceStart: lookup,
			targetStart: lookup,
			rang:        1,
		}
	}

	return
}

func parseLine(dest *almanac, currMap *rangeMap, line string) *rangeMap {
	if len(line) > 0 {
		switch {
		case strings.HasPrefix(line, "seeds:"):
			dest.seeds = parseSeeds(line)
		case line[0] >= 'a' && line[0] <= 'z':
			currMap = parseMap(line)
			dest.bySource[currMap.sourceType] = currMap
			dest.byTarget[currMap.targetType] = currMap
		case line[0] >= '0' && line[0] <= '9':
			parseRange(currMap, line)
		}
	}

	return currMap
}

func parseSeeds(line string) []uint64 {
	line, ok := strings.CutPrefix(line, "seeds: ")

	if !ok {
		panic(fmt.Errorf("invalid string format for seeds: expected 'seeds: 00 00 00 00', got %s", line))
	}

	nums := strings.Split(line, " ")

	result := make([]uint64, 0, 20)

	for _, num := range nums {
		val, err := strconv.ParseUint(num, 10, 64)

		if err != nil {
			panic(err)
		}
		result = append(result, val)
	}

	return result
}

func parseMap(line string) *rangeMap {
	line, ok := strings.CutSuffix(line, " map:")

	if !ok {
		panic(fmt.Errorf("invalid string format for map: expected 'y-to-z map:', got %s", line))
	}

	source, target, ok := strings.Cut(line, "-to-")

	if !ok {
		panic(fmt.Errorf("invalid string format for map: expected 'y-to-z map:', got %s", line))
	}

	return &rangeMap{
		sourceType: elementTypeMap[source],
		targetType: elementTypeMap[target],
		entries:    make([]*rangeMapEntry, 0, 20),
	}
}

func parseRange(rngMap *rangeMap, line string) {
	nums := strings.Split(line, " ")

	dest, err := strconv.ParseUint(nums[0], 10, 64)
	if err != nil {
		panic(err)
	}
	src, err := strconv.ParseUint(nums[1], 10, 64)
	if err != nil {
		panic(err)
	}
	rng, err := strconv.ParseUint(nums[2], 10, 64)
	if err != nil {
		panic(err)
	}

	rngMap.AddRange(&rangeMapEntry{
		targetStart: dest,
		sourceStart: src,
		rang:        rng,
	})
}

func quicksort(list []*rangeMapEntry, leftIndex, rightIndex int) {
	if leftIndex >= rightIndex {
		return
	}

	pivot := list[rightIndex]
	partIdx := leftIndex

	for i := leftIndex; i < rightIndex; i += 1 {
		if list[i].sourceStart < pivot.sourceStart {
			temp := list[i]
			list[i] = list[partIdx]
			list[partIdx] = temp
			partIdx += 1
		}
	}

	list[rightIndex] = list[partIdx]
	list[partIdx] = pivot

	quicksort(list, leftIndex, partIdx-1)
	quicksort(list, partIdx+1, rightIndex)
}
