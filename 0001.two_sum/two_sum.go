package two_sum

func twoSum(nums []int, target int) []int {
	store := make(map[int]int)
	for i, n := range nums {
		d := target - n
		if j, ok := store[d]; ok {
			return []int{j, i}
		}
		store[n] = i
	}
	return nil
}
