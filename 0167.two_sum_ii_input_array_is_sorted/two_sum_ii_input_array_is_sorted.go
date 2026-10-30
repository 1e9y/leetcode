package two_sum_ii_input_array_is_sorted

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1

	for l < r {
		sum := numbers[l] + numbers[r]

		switch {
		case sum > target:
			r--
		case sum < target:
			l++
		default:
			return []int{l + 1, r + 1}
		}
	}

	panic("no solution")
}
