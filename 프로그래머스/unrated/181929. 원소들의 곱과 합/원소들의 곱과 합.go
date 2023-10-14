func solution(num_list []int) int {
    sq := 1
	sum := 0
	
	for _, num := range num_list {
		sq *= num
		sum += num
	}
	
	if sq < sum * sum {
		return 1
	}
	
	return 0
}