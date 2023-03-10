package longest_increasing_subsequence

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	result := 1
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
				result = max(result, dp[i])
			}
		}
	}
	return result
}

/*
	F(S{1,i}): {
		1: n = 1
		max(F(S{0,j}))+1: s_i > max(s_j)
		max(F(S{0,j})): else
	}

	j
i		0	1	0	3	2	3
	0	1
	1	1	2
	0	1	2	2
	3	1	2	2	3
	2	1	2	2	3	3
	3	1	2	2	3	3	4

*/
