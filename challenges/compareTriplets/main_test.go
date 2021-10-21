package compareTriplets

import (
	"testing"
)

func TestRun(t *testing.T) {
	testDataAlice := []int32{5, 6, 7}
	testDataBob := []int32{3, 6, 10}
	expectedResult := []int32{1, 1}
	result := Run(testDataAlice, testDataBob)

	for i, v := range result {
		if v != expectedResult[i] {
			t.Fail()
		}
	}
}
