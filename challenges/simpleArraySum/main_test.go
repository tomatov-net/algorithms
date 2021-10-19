package simpleArraySum

import (
	"testing"
)

func TestRun(t *testing.T) {
	testData := []int{1, 2, 3, 4, 10, 11}
	expectedResult := 31
	result := Run(testData)

	if expectedResult != result {
		t.Fail()
	}
}
