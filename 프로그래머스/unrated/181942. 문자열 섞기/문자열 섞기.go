func solution(str1 string, str2 string) string {
    var answer string

	for i := 0; i < len(str1); i++ {
		answer += string(str1[i]) + string(str2[i])
	}

	return answer
}