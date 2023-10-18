package main

import (
	"fmt"
)

func main() {
	arr := []int{0, 1, 2, 4, 3}
	queries := [][]int{{0, 4, 2}, {0, 3, 2}, {0, 2, 2}}

	fmt.Print(solution(arr, queries))
}

func solution(arr []int, queries [][]int) []int {
	var result = []int{}

	for _, query := range queries {
		k := query[2]
		min := 1000001

		for _, n := range arr[query[0] : query[1]+1] {
			if n > k && n < min {
				min = n
			}
		}

		if min == 1000001 {
			result = append(result, -1)
		} else {
			result = append(result, min)
		}
	}
	return result
}
