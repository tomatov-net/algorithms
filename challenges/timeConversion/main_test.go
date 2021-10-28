package timeConversion

import (
	"testing"
)

func TestRun(t *testing.T) {
	testData1 := "12:01:00PM"
	testData2 := "12:01:00AM"
	expectedResult1 := "12:01:00"
	expectedResult2 := "00:01:00"
	result1 := Run(testData1)
	result2 := Run(testData2)
	result3 := Run2(testData2)

	if result1 != expectedResult1 {
		t.Fail()
	}

	if result2 != expectedResult2 {
		t.Fail()
	}

	if result3 != expectedResult2 {
		t.Fail()
	}
}
