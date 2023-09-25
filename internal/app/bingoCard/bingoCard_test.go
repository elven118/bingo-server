package bingoCard

import (
	"testing"
)

func TestIsExist(t *testing.T) {
	numbers := []int{1, 2, 9, 10, 13, 21, 48, 74}

	// test case 1
	result1 := isExist(2, numbers)
	expectedResult1 := true
	if expectedResult1 != result1 {
		t.Errorf("Test case 1 failed: Expected %v, got %v", expectedResult1, result1)
	}

	// test case 2
	result2 := isExist(9, numbers)
	expectedResult2 := true
	if expectedResult2 != result2 {
		t.Errorf("Test case 2 failed: Expected %v, got %v", expectedResult2, result2)
	}

	// test case 2
	result3 := isExist(20, numbers)
	expectedResult3 := false
	if expectedResult3 != result3 {
		t.Errorf("Test case 3 failed: Expected %v, got %v", expectedResult3, result3)
	}
}

func EqualSlices(slice1, slice2 []int) bool {
	// Check if the slices have the same length
	if len(slice1) != len(slice2) {
		return false
	}

	// Iterate over both slices and compare elements
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}

func IntPointersToInts(intPointers []*int) []int {
	ints := make([]int, len(intPointers))
	for i, ptr := range intPointers {
		ints[i] = *ptr
	}
	return ints
}

func Test1DigBingoCard(t *testing.T) {
	ResetLottery()
	AddLotteryNumber(1)
	AddLotteryNumber(2)
	AddLotteryNumber(9)
	AddLotteryNumber(10)
	AddLotteryNumber(12)
	AddLotteryNumber(13)
	AddLotteryNumber(21)
	AddLotteryNumber(24)
	AddLotteryNumber(48)
	AddLotteryNumber(56)
	AddLotteryNumber(61)
	AddLotteryNumber(74)

	// Test case 1 Valid diagonal
	cardNumbers := []int{9, 2, 13, 10, 12, 28, 21, 22, 24, 23, 39, 41, -1, 35, 40, 47, 56, 46, 48, 54, 61, 73, 67, 70, 74}
	expectedResult := []int{0, 1}
	result := ValidateBingoCard(cardNumbers)
	diagonal := IntPointersToInts(result.Diagonal)
	if !EqualSlices(diagonal, expectedResult) {
		t.Errorf("Test case 1 diagonal failed: Expected %v, got %v", expectedResult, diagonal)
	}
}

func Test2ColBingoCard(t *testing.T) {
	ResetLottery()
	AddLotteryNumber(9)
	AddLotteryNumber(28)
	AddLotteryNumber(39)
	AddLotteryNumber(47)
	AddLotteryNumber(61)

	// Test case 2: Valid col
	cardNumbers := []int{9, 2, 13, 10, 12, 28, 21, 22, 24, 23, 39, 41, -1, 35, 40, 47, 56, 46, 48, 54, 61, 73, 67, 70, 74}
	expectedResult := []int{0}
	result := ValidateBingoCard(cardNumbers)
	col := IntPointersToInts(result.Col)
	if !EqualSlices(col, expectedResult) {
		t.Errorf("Test case 2 col failed: Expected %v, got %v", expectedResult, col)
	}
}

func Test3RowBingoCard(t *testing.T) {
	ResetLottery()
	AddLotteryNumber(2)
	AddLotteryNumber(9)
	AddLotteryNumber(10)
	AddLotteryNumber(12)
	AddLotteryNumber(13)

	// Test case 3: Valid row
	cardNumbers := []int{9, 2, 13, 10, 12, 28, 21, 22, 24, 23, 39, 41, -1, 35, 40, 47, 56, 46, 48, 54, 61, 73, 67, 70, 74}
	expectedResult := []int{0}
	result := ValidateBingoCard(cardNumbers)
	row := IntPointersToInts(result.Row)
	if !EqualSlices(row, expectedResult) {
		t.Errorf("Test case 3 row failed: Expected %v, got %v", expectedResult, row)
	}
}

func Test4RowsBingoCard(t *testing.T) {
	ResetLottery()
	AddLotteryNumber(21)
	AddLotteryNumber(22)
	AddLotteryNumber(23)
	AddLotteryNumber(24)
	AddLotteryNumber(28)
	AddLotteryNumber(35)
	AddLotteryNumber(39)
	AddLotteryNumber(40)
	AddLotteryNumber(41)
	AddLotteryNumber(46)
	AddLotteryNumber(47)
	AddLotteryNumber(48)
	AddLotteryNumber(54)
	AddLotteryNumber(56)

	// Test case 4: Valid rows
	cardNumbers := []int{9, 2, 13, 10, 12, 28, 21, 22, 24, 23, 39, 41, -1, 35, 40, 47, 56, 46, 48, 54, 61, 73, 67, 70, 74}
	expectedResult := []int{1, 2, 3}
	result := ValidateBingoCard(cardNumbers)
	row := IntPointersToInts(result.Row)
	if !EqualSlices(row, expectedResult) {
		t.Errorf("Test case 4 rows failed: Expected %v, got %v", expectedResult, row)
	}
}

func Test5BingoCard(t *testing.T) {
	ResetLottery()
	AddLotteryNumber(1)
	AddLotteryNumber(3)
	AddLotteryNumber(4)
	AddLotteryNumber(5)
	AddLotteryNumber(7)
	AddLotteryNumber(8)
	AddLotteryNumber(21)
	AddLotteryNumber(22)
	AddLotteryNumber(23)
	AddLotteryNumber(24)
	AddLotteryNumber(28)
	AddLotteryNumber(35)
	AddLotteryNumber(39)
	AddLotteryNumber(40)
	AddLotteryNumber(41)
	AddLotteryNumber(46)
	AddLotteryNumber(47)
	AddLotteryNumber(48)
	AddLotteryNumber(54)
	AddLotteryNumber(55)
	AddLotteryNumber(56)
	AddLotteryNumber(64)
	AddLotteryNumber(66)
	AddLotteryNumber(70)
	AddLotteryNumber(75)

	// Test case 5: Invalid
	cardNumbers := []int{14, 1, 4, 2, 10, 27, 22, 21, 29, 17, 40, 37, -1, 45, 39, 50, 54, 56, 58, 46, 71, 63, 72, 74, 62}
	expectedResult := false
	result := ValidateBingoCard(cardNumbers)
	if expectedResult != result.IsValid {
		t.Errorf("Test case 5 failed: Expected %v, got %v", expectedResult, result.IsValid)
	}
}

func Test6BingoCard(t *testing.T) {
	ResetLottery()
	AddLotteryNumber(1)
	AddLotteryNumber(3)
	AddLotteryNumber(4)
	AddLotteryNumber(5)
	AddLotteryNumber(7)
	AddLotteryNumber(8)
	AddLotteryNumber(21)
	AddLotteryNumber(22)
	AddLotteryNumber(23)
	AddLotteryNumber(24)
	AddLotteryNumber(28)
	AddLotteryNumber(35)
	AddLotteryNumber(39)
	AddLotteryNumber(40)
	AddLotteryNumber(41)
	AddLotteryNumber(46)
	AddLotteryNumber(47)
	AddLotteryNumber(48)
	AddLotteryNumber(54)
	AddLotteryNumber(55)
	AddLotteryNumber(56)
	AddLotteryNumber(64)
	AddLotteryNumber(66)
	AddLotteryNumber(70)
	AddLotteryNumber(75)

	// Test case 5: Invalid
	cardNumbers := []int{14, 1, 4, 2, 10, 27, 22, 21, 29, 17, 40, 37, -1, 45, 39, 50, 54, 56, 58, 46, 71, 63, 72, 74, 62}
	result := ValidateBingoCard(cardNumbers)
	if EqualSlices(cardNumbers, result.Numbers) {
		t.Errorf("Test case numbers failed: Expected %v, got %v", cardNumbers, result.Numbers)
	}
}
