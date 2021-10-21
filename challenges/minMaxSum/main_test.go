package minMaxSum

import (
	"testing"
)

func TestRun(t *testing.T) {
	testData := []int{256741038, 623958417, 467905213, 714532089, 938071625}
	expectedResult := []int{2063136757, 2744467344}
	result := Run(testData)

	for key, value := range result {
		if value != expectedResult[key] {
			t.Fail()
		}
	}
}
