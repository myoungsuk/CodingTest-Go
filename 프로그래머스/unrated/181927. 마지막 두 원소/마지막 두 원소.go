func solution(num_list []int) []int {
    last := num_list[len(num_list)-1]
	front := num_list[len(num_list)-2]

	if last > front {
		return append(num_list, last-front)
	} else {
		return append(num_list, last*2)
	}
}