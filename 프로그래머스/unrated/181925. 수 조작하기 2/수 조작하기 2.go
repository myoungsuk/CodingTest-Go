func solution(numLog []int) string {
    result := ""
	for i := 0; i < len(numLog)-1; i++ {

		c := numLog[i+1] - numLog[i]
		switch c {
		case 1:
			result += "w"
		case -1:
			result += "s"
		case 10:
			result += "d"
		case -10:
			result += "a"
		}

	}
	return result
}