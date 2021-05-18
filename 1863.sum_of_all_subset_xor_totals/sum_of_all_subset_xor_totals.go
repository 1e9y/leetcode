package sum_of_all_subset_xor_totals

func subsetXORSum(nums []int) int {
	var totalSum func(int, int) int
	totalSum = func(i, sum int) int {
		if i == len(nums) {
			return sum
		}
		return totalSum(i+1, sum^nums[i]) + totalSum(i+1, sum)
	}
	return totalSum(0, 0)
}
