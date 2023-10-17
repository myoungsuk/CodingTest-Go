func solution(n int, control string) int {
    sum := n

	for _, v := range control {
		switch string(v) {
		case "w":
			sum += 1
		case "s":
			sum -= 1
		case "d":
			sum += 10
		case "a":
			sum -= 10
		}
	}
	return sum
}