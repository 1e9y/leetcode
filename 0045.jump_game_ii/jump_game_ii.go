package jump_game_ii

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	var min, start, end int
	for i, n := range nums {
		end = max(end, i+n)
		if end >= len(nums)-1 {
			return min + 1
		}
		if i == start {
			min++
			start = end
		}
	}

	return min
}
