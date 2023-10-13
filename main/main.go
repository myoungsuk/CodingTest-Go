package main

import (
	"fmt"
)

func main() {
	n := 10
	fmt.Print(solution(n))
}

func solution(n int) int {
	answer := 0

	if n%2 == 1 {
		for i := 0; i <= n; i++ {
			if i%2 == 1 {
				answer += i
			}
		}
	} else {
		for i := 1; i <= n; i++ {
			if i%2 == 0 {
				answer += i * i
			}
		}
	}
	return answer
}
