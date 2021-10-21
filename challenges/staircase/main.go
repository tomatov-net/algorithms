package staircase

import (
	"fmt"
	"strings"
)

//https://www.hackerrank.com/challenges/staircase/problem

func Validate(data int32) bool {
	if data == 0 || data > 100 {
		return false
	}

	return true
}

func Run(size int32) []string {
	type figure struct {
		symbol string
		count  int32
	}
	var figures []figure
	var result []string

	if !Validate(size) {
		return result
	}

	figures = append(figures, figure{symbol: " ", count: size - 1})
	figures = append(figures, figure{symbol: "#", count: 1})

	for i := 0; i < int(size); i++ {
		var row []string
		for j := 0; j < int(figures[0].count); j++ {
			row = append(row, figures[0].symbol)
		}

		for j := 0; j < int(figures[1].count); j++ {
			row = append(row, figures[1].symbol)
		}
		figures[0].count--
		figures[1].count++
		result = append(result, strings.Join(row, ""))
	}

	for _, value := range result {
		fmt.Println(value)
	}

	return result
}
