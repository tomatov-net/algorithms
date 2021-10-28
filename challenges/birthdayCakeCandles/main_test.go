package birthdayCakeCandles

import (
	"testing"
)

func TestRun(t *testing.T) {
	testData := []int{3, 2, 1, 3}
	expectedResult := 2
	result := Run(testData)

	if result != expectedResult {
		t.Fail()
	}
}
