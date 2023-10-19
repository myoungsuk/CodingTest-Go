import ("strconv")

func check(num int) bool {
	str := strconv.Itoa(num)
	for _, r := range str {
		if r != '5' && r != '0' {
			return false
		}
	}
	return true
}

func solution(l int, r int) []int {
	var result []int

	for i := l / 5 * 5; i <= r; i++ {
		if check(i) {
			result = append(result, i)
		}
	}

	if len(result) == 0 {
		result = append(result, -1)
	}

	return result
}