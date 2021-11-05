package encryptionProblem

import (
	"testing"
)

func TestRun(t *testing.T) {
	testData := "feedthedog    "
	expectedResult := "fto ehg ee dd"
	result := Run(testData)

	if result != expectedResult {
		t.Fail()
	}
}
