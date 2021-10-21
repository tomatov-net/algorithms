package diagonalDifference

import "fmt"

//https://www.hackerrank.com/challenges/diagonal-difference/problem

func Run(array [][]int32) int32 {
	var leftDiagonal []int32
	var rightDiagonal []int32
	var diagonalIndex int

	for rowNumber, rowValue := range array {
		leftDiagonal = append(leftDiagonal, array[rowNumber][diagonalIndex])
		rightDiagonal = append(rightDiagonal, array[rowNumber][(len(rowValue)-diagonalIndex-1)])
		diagonalIndex++
	}

	fmt.Println(leftDiagonal)
	fmt.Println(rightDiagonal)
	fmt.Println(sum(leftDiagonal))
	fmt.Println(sum(rightDiagonal))

	return abs(sum(leftDiagonal) - sum(rightDiagonal))
}

func sum(numbers []int32) int32 {
	var result int32
	for _, value := range numbers {
		result += value
	}
	return result
}

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}
