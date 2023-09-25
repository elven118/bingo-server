package bingoCard

import (
	"bingo/graph/model"
	"math/rand"
	"sort"
	"time"
)

var lotteryNumbers []int

func randomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min+1)
}

func sliceContains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func GenBingoCard() []int {
	var (
		rangeFrom = 1
		rangeTo   = 15
	)
	var randomArray []int
	for c := 0; c < 5; c++ {
		for i := 0; i < 5; i++ {
			if c == 2 && i == 2 {
				randomArray = append(randomArray, -1)
			} else {
				randomNubmer := randomInt(rangeFrom, rangeTo)
				for sliceContains(randomArray, randomNubmer) {
					randomNubmer = randomInt(rangeFrom, rangeTo)
				}
				randomArray = append(randomArray, randomNubmer)
			}
		}
		rangeFrom += 15
		rangeTo += 15
	}

	return randomArray
}

// InsertSorted inserts an element into a sorted slice while maintaining the sorted order.
func insertSorted(element int, numbers []int) []int {
	i := sort.Search(len(numbers), func(i int) bool { return numbers[i] >= element })
	if i < 0 || i > len(numbers) {
		return numbers
	}
	numbers = append(numbers[:i], append([]int{element}, numbers[i:]...)...)
	return numbers
}

func isExist(target int, numbers []int) bool {
	i := sort.Search(len(numbers), func(i int) bool { return numbers[i] >= target })
	if i < len(numbers) && numbers[i] == target {
		return true
	} else {
		return false
	}
}

func AddLotteryNumber(element int) int {
	lotteryNumbers = insertSorted(element, lotteryNumbers)
	return element
}

func ResetLottery() bool {
	lotteryNumbers = make([]int, 0)
	return true
}

func ValidateBingoCard(cardNumbers []int) *model.ValidateResult {
	result := &model.ValidateResult{
		Row:      []*int{},
		Col:      []*int{},
		Diagonal: []*int{},
		IsValid:  false,
		Numbers:  append([]int{}, cardNumbers...),
	}
	exists := make([]bool, len(cardNumbers))
	exists[12] = true

	// check row
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			index := r*5 + c
			if cardNumbers[index] == -1 {
				if !exists[index] {
					break
				}
			} else {
				// binary search
				exist := isExist(cardNumbers[index], lotteryNumbers)
				exists[index] = exist
				cardNumbers[index] = -1
				if !exist {
					break
				}
			}
			if c == 4 {
				row := r
				result.Row = append(result.Row, &row)
				result.IsValid = true
			}
		}
	}

	// check col
	for c := 0; c < 5; c++ {
		for r := 0; r < 5; r++ {
			index := c + r*5
			if cardNumbers[index] == -1 {
				if !exists[index] {
					break
				}
			} else {
				// binary search
				exist := isExist(cardNumbers[index], lotteryNumbers)
				exists[index] = exist
				cardNumbers[index] = -1
				if !exist {
					break
				}
			}
			if r == 4 {
				col := c
				result.Col = append(result.Col, &col)
				result.IsValid = true
			}
		}
	}

	// check diagonal
	for i := 0; i <= 24; i += 6 {
		if cardNumbers[i] == -1 {
			if !exists[i] {
				break
			}
		} else {
			// binary search
			exist := isExist(cardNumbers[i], lotteryNumbers)
			exists[i] = exist
			cardNumbers[i] = -1
			if !exist {
				break
			}
		}
		if i == 24 {
			d := 0
			result.Diagonal = append(result.Diagonal, &d)
			result.IsValid = true
		}
	}

	for i := 4; i <= 20; i += 4 {
		if cardNumbers[i] == -1 {
			if !exists[i] {
				break
			}
		} else {
			// binary search
			exist := isExist(cardNumbers[i], lotteryNumbers)
			exists[i] = exist
			cardNumbers[i] = -1
			if !exist {
				break
			}
		}
		if i == 20 {
			d := 1
			result.Diagonal = append(result.Diagonal, &d)
			result.IsValid = true
		}
	}
	return result
}
