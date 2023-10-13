func solution(a int, d int, included []bool) int {
    result := 0

	for idx, item := range included {
		if item {
			result += a + (d * idx)
		}
	}
	return result
}