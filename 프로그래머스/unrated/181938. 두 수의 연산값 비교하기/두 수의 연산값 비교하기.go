import "strconv"

func solution(a int, b int) int {
    
    A, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	B := 2 * a * b

	if A >= B {
		return A
	} else {
		return B
	}
}