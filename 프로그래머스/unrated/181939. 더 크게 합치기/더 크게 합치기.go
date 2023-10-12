import "strconv"

func solution(a int, b int) int {
    A, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	B, _ := strconv.Atoi(strconv.Itoa(b) + strconv.Itoa(a))

	if A > B {
		return A
	} else {
		return B
	}
}