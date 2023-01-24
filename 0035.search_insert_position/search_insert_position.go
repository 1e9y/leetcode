package search_insert_position

func searchInsert(nums []int, target int) int {
	i := 0
	j := len(nums) - 1
	for i <= j {
		m := (i + j) / 2
		if nums[m] == target {
			return m
		}
		if nums[m] < target {
			i = m + 1
		} else {
			j = m - 1
		}
	}
	return i
}
