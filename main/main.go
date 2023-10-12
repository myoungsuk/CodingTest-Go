package main

import (
	"fmt"
)

func main() {
	a := 98
	b := 2
	fmt.Print(solution(a, b))
}

func solution(num int, n int) int {

	if num%n == 0 {
		return 1
	} else {
		return 0
	}

}
