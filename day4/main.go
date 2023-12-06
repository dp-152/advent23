package main

import (
	"fmt"
	"math"

	"github.com/dp-152/advent23/common/fs"
)

type lotteryCard struct {
	id      int
	winNums map[int]bool
	ownNums map[int]bool
	matchCt *int
}

const cCARD_SUFFIX = ':'
const cNUMGROUP_SEPARATOR = '|'

func (c *lotteryCard) score() int {
	if c.matchCt == nil {
		c.matches()
	}

	if *c.matchCt == 0 {
		return 0
	}

	return int(math.Pow(2, float64(*c.matchCt-1)))
}

func (c *lotteryCard) matches() int {
	if c.matchCt == nil {
		matchCount := 0
		for num := range c.ownNums {
			if _, ok := c.winNums[num]; ok {
				matchCount += 1
			}
		}
		c.matchCt = &matchCount
	}

	return *c.matchCt
}

func main() {
	pointsSum := sumPoints("input.txt")

	fmt.Printf("Sum of lottery card points: %d\n", pointsSum)

	copies := countCopies("input.txt")

	fmt.Printf("Sum of all cards: %d\n", copies)
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

func countCopies(inputFile string) (count int) {
	cards := make([]*lotteryCard, 0, 100)
	cardCounts := make(map[int]int)

	cancelChan := make(chan struct{})
	fileChan := fs.ReadLines(inputFile, cancelChan)

	for line := range fileChan {
		card := parseLine(line)
		cards = append(cards, card)
		cardCounts[card.id] = 1
	}

	for _, card := range cards {
		count += cardCounts[card.id]
		matches := card.matches()
		for copyOffset := 1; copyOffset <= matches; copyOffset += 1 {
			cardCounts[card.id+copyOffset] += cardCounts[card.id]
		}
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
