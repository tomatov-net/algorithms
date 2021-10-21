package plusMinus

import (
	"testing"
)

func TestRun(t *testing.T) {
	testData := []int32{-4, 3, -9, 0, 4, 1}
	expectedResult := []string{"0.500000", "0.333333", "0.166667"}
	result := Run(testData)

	for key, value := range result {
		if value != expectedResult[key] {
			t.Fail()
		}
	}
}
