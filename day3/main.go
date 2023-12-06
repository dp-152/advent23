package main

import (
	"fmt"
	"math"

	"github.com/dp-152/advent23/common/fs"
)

type eSchematicPartType int

const (
	eUNKNOWN eSchematicPartType = iota
	eSYMBOL
	eNUMBER
)

type partAddr struct {
	x,
	y int
}

type engineSchematicItem struct {
	addr      partAddr
	addrRight partAddr
	itemType  eSchematicPartType
	itemValue int
}

type engineSchematic struct {
	separator string
	addrMap   map[string]*engineSchematicItem
	valMap    map[string][]*engineSchematicItem
	typeMap   map[eSchematicPartType][]*engineSchematicItem
}

func (m *engineSchematic) Add(item *engineSchematicItem) {
	addrLeft := fmt.Sprintf("%d%s%d", item.addr.x, m.separator, item.addr.y)
	addrRight := fmt.Sprintf("%d%s%d", item.addrRight.x, m.separator, item.addrRight.y)
	m.addrMap[addrLeft] = item
	m.addrMap[addrRight] = item

	addrVal := fmt.Sprintf("%d%s%d", item.itemType, m.separator, item.itemValue)
	m.valMap[addrVal] = append(m.valMap[addrVal], item)

	m.typeMap[item.itemType] = append(m.typeMap[item.itemType], item)
}

func (m *engineSchematic) GetAtAddr(x, y int) (item *engineSchematicItem, ok bool) {
	addr := fmt.Sprintf("%d%s%d", x, m.separator, y)
	item, ok = m.addrMap[addr]
	return
}

func (m *engineSchematic) GetBySymbol(symbol rune) (items []*engineSchematicItem) {
	items = m.valMap[fmt.Sprintf("%d%s%d", eSYMBOL, m.separator, symbol)]
	return
}

func (m *engineSchematic) GetByType(itemType eSchematicPartType) (items []*engineSchematicItem) {
	items = m.typeMap[itemType]
	return
}

func (m *engineSchematic) GetAdjacent(origin *engineSchematicItem) []*engineSchematicItem {
	uniq := make(map[*engineSchematicItem]bool)

	for xoffs := -1; xoffs <= 1; xoffs += 1 {
		xaddr := origin.addr.x + xoffs
		for yoffs := -1; yoffs <= 1; yoffs += 1 {
			yaddr := origin.addr.y + yoffs
			if match, ok := m.GetAtAddr(xaddr, yaddr); ok && match != origin {
				uniq[match] = true
			}
		}
	}

	found := make([]*engineSchematicItem, len(uniq))
	idx := 0
	for key := range uniq {
		found[idx] = key
		idx += 1
	}
	return found
}

func main() {
	engine := loadEngineSchematic("input.txt")

	partsSum := sumPartNumbers(engine)

	fmt.Printf("Sum of all parts: %d\n", partsSum)
}

func sumPartNumbers(engine *engineSchematic) (sum int) {
	seen := make(map[*engineSchematicItem]bool)
	for _, source := range engine.GetByType(eSYMBOL) {
		adjItems := engine.GetAdjacent(source)
		for _, adjItem := range adjItems {
			if adjItem.itemType == eNUMBER {
				if _, ok := seen[adjItem]; !ok {
					seen[adjItem] = true
					sum += adjItem.itemValue
				}
			}
		}
	}
	return
}

func loadEngineSchematic(inputFile string) *engineSchematic {
	engine := &engineSchematic{
		separator: "#",
		addrMap:   make(map[string]*engineSchematicItem),
		valMap:    make(map[string][]*engineSchematicItem),
		typeMap:   make(map[eSchematicPartType][]*engineSchematicItem),
	}

	cancelChan := make(chan struct{})
	fileChan := fs.ReadLines(inputFile, cancelChan)

	var lineIdx int = 0
	for line := range fileChan {
		scanLine(engine, line, lineIdx)
		lineIdx += 1
	}

	return engine
}

func scanLine(engine *engineSchematic, line string, lineIdx int) {
	numbuf := make([]int8, 0, 5)

	var numlx int
	var numrx int
	for pos, runeVal := range line {
		if runeVal >= '0' && runeVal <= '9' {
			if len(numbuf) == 0 {
				numlx = pos
			}
			numrx = pos

			numbuf = append(numbuf, int8(runeVal-'0'))
		} else {
			if len(numbuf) > 0 {
				mapNum(engine, numbuf, numlx, numrx, lineIdx)
				numbuf = numbuf[:0]
				numlx = -1
			}

			if runeVal != '.' {
				mapSym(engine, runeVal, pos, lineIdx)
			}
		}
	}
	if len(numbuf) > 0 {
		mapNum(engine, numbuf, numlx, numrx, lineIdx)
	}
}

func mapNum(engine *engineSchematic, numbuf []int8, lx, rx, y int) {
	var numval int
	for pos, val := range numbuf {
		numval += int(val) * int(math.Pow10(len(numbuf)-pos-1))
	}

	engine.Add(&engineSchematicItem{
		addr: partAddr{
			x: lx,
			y: y,
		},
		addrRight: partAddr{
			x: rx,
			y: y,
		},
		itemType:  eNUMBER,
		itemValue: numval,
	})
}

func mapSym(engine *engineSchematic, symbol rune, x, y int) {
	engine.Add(&engineSchematicItem{
		addr: partAddr{
			x: x,
			y: y,
		},
		addrRight: partAddr{
			x: x,
			y: y,
		},
		itemType:  eSYMBOL,
		itemValue: int(symbol),
	})
}
