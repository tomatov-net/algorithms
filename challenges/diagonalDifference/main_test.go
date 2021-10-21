package diagonalDifference

import (
	"testing"
)

func TestRun(t *testing.T) {
	testData := [][]int32{{1, 2, 3}, {4, 5, 6}, {9, 8, 9}}
	expectedResult := 2
	result := Run(testData)

	if expectedResult != int(result) {
		t.Fail()
	}
}
