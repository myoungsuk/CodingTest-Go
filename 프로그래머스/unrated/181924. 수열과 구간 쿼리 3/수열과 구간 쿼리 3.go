func solution(arr []int, queries [][]int) []int {
    
    for i := 0; i < len(queries); i++ {
		start := queries[i][0]
		end := queries[i][1]

		temp := arr[start]
		arr[start] = arr[end]
		arr[end] = temp
	}

	return arr
}