package main

import (
	"fmt"
	"strconv"
)

func main() {
	num_list := []int{3, 4, 5, 2, 1}
	fmt.Print(solution(num_list))
}

func solution(num_list []int) int {
	var odd string
	var even string

	for i := 0; i < len(num_list); i++ {
		if num_list[i]%2 == 0 {
			odd += strconv.Itoa(num_list[i])
		} else {
			even += strconv.Itoa(num_list[i])
		}
	}

	oddInt, _ := strconv.Atoi(odd)
	evenInt, _ := strconv.Atoi(even)

	return oddInt + evenInt
}
