package birthdayCakeCandles

import (
	"fmt"
	"math/rand"
)

//https://www.hackerrank.com/challenges/birthday-cake-candles/problem

func Validate(data []int) bool {
	return len(data) <= 1000000
}

func Run(dataArray []int) int {
	result := 0
	var maxNumber int
	arrayLen := len(dataArray)

	if !Validate(dataArray) {
		return result
	}

	arraySorted := quickSort(dataArray)
	maxNumber = arraySorted[arrayLen-1]

	for i := len(arraySorted) - 1; i >= 0; i-- {
		if arraySorted[i] == maxNumber {
			result++
		} else {
			break
		}
	}
	fmt.Println(result)

	return result
}

func quickSort(data []int) []int {
	if len(data) < 2 {
		return data
	}
	left, right := 0, len(data)-1

	// Pick a pivot
	pivotIdx := rand.Int() % len(data)

	// Move the pivot to the right
	data[pivotIdx], data[right] = data[right], data[pivotIdx]

	// Pile elements smaller than the pivot on the left
	for i := range data {
		if data[i] < data[right] {
			data[i], data[left] = data[left], data[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	data[left], data[right] = data[right], data[left]

	// Go down the rabbit hole
	quickSort(data[:left])
	quickSort(data[left+1:])

	return data
}
