package minMaxSum

import (
	"fmt"
	"math/rand"
)

//https://www.hackerrank.com/challenges/mini-max-sum/problem

func Validate(data []int) bool {
	return len(data) >= 5
}

func Run(dataArray []int) []int {
	var result []int
	var minSum int
	var maxSum int
	arrayLength := len(dataArray)

	if !Validate(dataArray) {
		return result
	}

	arraySorted := quickSort(dataArray)

	for i := 0; i < 4; i++ {
		minSum = minSum + arraySorted[i]
		maxSum = maxSum + arraySorted[arrayLength-i-1]
		fmt.Println(maxSum)
	}
	fmt.Println(dataArray)
	fmt.Println(arraySorted)
	fmt.Println(minSum)
	fmt.Println(maxSum)

	return []int{minSum, maxSum}
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
