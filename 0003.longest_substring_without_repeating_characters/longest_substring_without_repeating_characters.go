package longest_substring_without_repeating_characters

func max(a, b int) int {
	if a > b {
		return a

	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	answer := 0
	var pos [128]int

	start := 0
	for end, char := range s {
		start = max(start, pos[char])
		answer = max(answer, end-start+1)
		pos[char] = end + 1
	}

	return answer
}
