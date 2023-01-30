package non_decreasing_subsequences

const (
	MinNum = -100
	MaxNum = 100
)

func findSubsequences(nums []int) [][]int {
	result := make([][]int, 0)
	sequence := make([]int, 0, len(nums))
	backtrack(nums, 0, sequence, &result)
	return result
}

func backtrack(nums []int, start int, sequence []int, result *[][]int) {
	if len(sequence) > 1 {
		temp := make([]int, len(sequence))
		copy(temp, sequence)
		*result = append(*result, temp)
	}

	visited := make(map[int]bool, MaxNum-MinNum+1)
	for i := start; i < len(nums); i++ {
		idx := nums[i] - MinNum
		if start > 0 && nums[i] < nums[start-1] {
			continue
		}
		if visited[idx] {
			continue
		}
		visited[idx] = true
		sequence = append(sequence, nums[i])
		backtrack(nums, i+1, sequence, result)
		sequence = sequence[:len(sequence)-1]
	}
}
