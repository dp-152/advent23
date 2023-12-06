package main

import (
	"fmt"

	"github.com/dp-152/advent23/common/fs"
)

type lotteryCard struct {
	id      int
	winNums map[int]bool
	ownNums map[int]bool
}

const cCARD_SUFFIX = ':'
const cNUMGROUP_SEPARATOR = '|'

func (c *lotteryCard) score() (score int) {
	for num := range c.ownNums {
		if _, ok := c.winNums[num]; ok {
			if score == 0 {
				score = 1
			} else {
				score *= 2
			}
		}
	}
	return
}

func main() {
	pointsSum := sumPoints("input.txt")

	fmt.Printf("Sum of lottery cards: %d", pointsSum)
}

func sumPoints(inputFile string) (sum int) {
	cancelChan := make(chan struct{})
	fileChan := fs.ReadLines(inputFile, cancelChan)

	for line := range fileChan {
		card := parseLine(line)
		sum += card.score()
	}

	return
}

func parseLine(line string) *lotteryCard {
	card := &lotteryCard{
		winNums: make(map[int]bool),
		ownNums: make(map[int]bool),
	}

	var target *map[int]bool

	currNum := -1
	for _, runeVal := range line {
		switch {
		case runeVal >= '0' && runeVal <= '9':
			if currNum == -1 {
				currNum = int(runeVal) - '0'
			} else {
				currNum = currNum*10 + (int(runeVal) - '0')
			}
		case runeVal == cCARD_SUFFIX:
			card.id = currNum
			currNum = -1
			target = &card.winNums
		case runeVal == cNUMGROUP_SEPARATOR:
			if currNum > -1 {
				(*target)[currNum] = true
				currNum = -1
			}
			target = &card.ownNums
		case runeVal == ' ' && currNum > -1:
			(*target)[currNum] = true
			currNum = -1
		}
	}

	if currNum > -1 {
		(*target)[currNum] = true
	}

	return card
}
