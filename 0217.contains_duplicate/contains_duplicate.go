package contains_duplicate

import "slices"

func containsDuplicate(nums []int) bool {
	slices.Sort(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			return true
		}
	}

	return false
}
