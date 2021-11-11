package fraudulentActivityNotifications

import (
	"fmt"
	"math"
	"sort"
)

//https://www.hackerrank.com/challenges/fraudulent-activity-notifications/problem

func Run(expenditure []int, d int) int {
	var result int

	length := len(expenditure)

	for i := 0; i < length; i++ {

		if d == length {
			break
		}

		fmt.Println(i, length)

		sliceToAnalyze := expenditure[i:d]
		median := getMedian(sliceToAnalyze)

		//fmt.Println(sliceToAnalyze, d, median)

		if float64(expenditure[d]) >= 2*median {
			result++
		}

		d++
	}

	fmt.Println(result)

	return result
}

func getMedian(arr []int) float64 {
	if !sort.IntsAreSorted(arr) {
		sort.Ints(arr)
	}

	length := float64(len(arr))
	centerValueKey := int64(math.Floor(length / 2))

	if int(length)%2 != 0 {
		return float64(arr[centerValueKey])
	}

	median := (arr[centerValueKey] + arr[centerValueKey+1]) / 2

	return float64(median)
}
