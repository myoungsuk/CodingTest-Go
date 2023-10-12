func solution(arr []string) string {
    answer := ""
	for i := 0; i < len(arr); i++ {
		answer += arr[i]
	}
	return answer
}