package simpleArraySum

//https://www.hackerrank.com/challenges/simple-array-sum/problem

func Validate(array []int) bool {
	if len(array) <= 0 || len(array) > 1000 {
		return false
	}

	return true
}

func Run(array []int) int {
	result := 0

	if !Validate(array) {
		return result
	}

	for _, value := range array {
		result += value
	}

	return result
}
