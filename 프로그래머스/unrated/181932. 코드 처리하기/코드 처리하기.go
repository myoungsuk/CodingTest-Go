func solution(code string) string {
   mode := 0
	answer := ""

	for i, c := range code {
		if c == '1' {
			mode = 1 - mode
			continue
		}

		if mode == 0 {
			if i%2 == 0 {
				answer += string(c)
			}
		} else {
			if i%2 != 0 {
				answer += string(c)
			}
		}
	}

	if answer == "" {
		return "EMPTY"
	}

	return answer
}