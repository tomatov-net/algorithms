package compareTriplets

//https://www.hackerrank.com/challenges/compare-the-triplets/problem

func Validate(array []int32) bool {
	if len(array) != 3 {
		return false
	}

	for _, value := range array {
		if value > 100 || value < 1 {
			return false
		}
	}

	return true
}

func Run(aliceInput []int32, bobInput []int32) []int32 {
	result := []int32{0, 0}

	if !Validate(aliceInput) || !Validate(bobInput) {
		return result
	}

	for index, _ := range aliceInput {
		if aliceInput[index] > bobInput[index] {
			result[0] += 1
		}
		if aliceInput[index] < bobInput[index] {
			result[1] += 1
		}
	}

	return result
}
