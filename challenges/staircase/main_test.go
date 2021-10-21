package staircase

import (
	"testing"
)

func TestRun(t *testing.T) {
	var testData int32 = 6
	expectedResult := []string{
		"     #",
		"    ##",
		"   ###",
		"  ####",
		" #####",
		"######",
	}
	result := Run(testData)

	for key, value := range result {
		if value != expectedResult[key] {
			t.Fail()
		}
	}
}
