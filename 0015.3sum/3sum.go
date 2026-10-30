package three_sum

import "slices"

func threeSum(nums []int) [][]int {
	var answer [][]int

	slices.Sort(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		if nums[i] > 0 {
			break
		}

		target := -nums[i]
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[l] + nums[r]
			switch {
			case sum > target:
				r--
			case sum < target:
				l++
			default:
				answer = append(answer, []int{nums[i], nums[l], nums[r]})
				l, r = l+1, r-1
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}

	return answer
}
