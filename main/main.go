package main

import (
	"fmt"
)

func main() {
	num_list := []int{3, 4, 5, 2, 1}
	fmt.Print(solution(num_list))
}

func solution(num_list []int) int {

	sq := 1
	sum := 0

	for _, num := range num_list {
		sq *= num
		sum += num
	}

	if sq < sum*sum {
		return 1
	}

	return 0
}
