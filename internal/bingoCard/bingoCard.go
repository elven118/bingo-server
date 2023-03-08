package bingoCard

import (
	"math/rand"
	"time"
)

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
