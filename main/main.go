package main

import (
	"fmt"
	"sort"
)

func main() {
	a := 2
	b := 6
	c := 1

	fmt.Print(solution(a, b, c))
}

func solution(a int, b int, c int) int {

	original := []int{a, b, c}
	arr := []int{a, b, c}
	count := GetMaxSameCount([]int{a, b, c})
	ret := 1

	for i := 0; i < count; i++ {
		ret *= Sum(arr)
		arr = Multiply(arr, original)
	}

	return ret
}

func Multiply(a []int, b []int) []int {
	var ret []int
	for i := range a {
		ret = append(ret, a[i]*b[i])
	}
	return ret
}

func GetMaxSameCount(arr []int) int {
	sort.Ints(arr)
	prev := arr[0]
	ret := 1

	for i := 1; i < len(arr); i++ {
		if prev == arr[i] {
			ret++
		}
		prev = arr[i]
	}
	return ret
}

func Sum(arr []int) int {
	ret := 0
	for _, item := range arr {
		ret += item
	}
	return ret
}
