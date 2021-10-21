package plusMinus

import "fmt"

//https://www.hackerrank.com/challenges/plus-minus/problem

func Run(array []int32) []string {
	type ratio struct {
		positive float32
		negative float32
		zero     float32
	}

	var r ratio

	for _, value := range array {
		if value > 0 {
			r.positive++
		} else if value < 0 {
			r.negative++
		} else {
			r.zero++
		}
	}
	count := float32(len(array))

	result := []string{
		fmt.Sprintf("%.6f", r.positive/count),
		fmt.Sprintf("%.6f", r.negative/count),
		fmt.Sprintf("%.6f", r.zero/count),
	}
	fmt.Println(result)

	for _, val := range result {
		fmt.Println(val)
	}

	return result
}
