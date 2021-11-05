package encryptionProblem

import (
	"fmt"
	"math"
	"strings"
)

//https://www.hackerrank.com/challenges/encryption/problem

func Run(str string) string {
	strFiltered := strings.Replace(str, " ", "", -1)
	strLength := len(strFiltered)
	var rows, columns int
	sqrRoot := math.Sqrt(float64(strLength))
	_, floatPart := math.Modf(sqrRoot)
	sqrRootInt := int(sqrRoot)

	columns = int(sqrRoot)
	if floatPart == 0 {
		rows = int(sqrRoot)
		columns = rows
	} else {
		rows, columns = sqrRootInt, sqrRootInt+1
	}

	chunkedArray := chunk(strFiltered, columns)
	rows = len(chunkedArray)
	var resultArr []string
	var tempResult []string
	fmt.Println(chunkedArray)

	for column := 0; column < columns; column++ {
		for row := 0; row < rows; row++ {
			if isset(chunkedArray[row], column) {
				tempResult = append(tempResult, chunkedArray[row][column])
			}
		}
		resultArr = append(resultArr, strings.Join(tempResult, ""))
		tempResult = []string{}
	}
	result := strings.Join(resultArr, " ")

	fmt.Println(resultArr)
	return result
}

func isset(arr []string, column int) bool {
	return len(arr)-1 >= column
}

func chunk(str string, size int) [][]string {
	var result [][]string
	charsArr := strings.Split(str, "")

	for i := 0; i < len(charsArr); i += size {
		end := i + size

		if end > len(charsArr) {
			end = len(charsArr)
		}

		result = append(result, charsArr[i:end])
	}
	return result
}
