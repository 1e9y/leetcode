package top_k_frequent_elements

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int, len(nums))
	for _, n := range nums {
		freq[n]++
	}

	buckets := make([][]int, len(nums)+1)
	for num, count := range freq {
		buckets[count] = append(buckets[count], num)
	}

	answer := make([]int, 0, k)
	for i := len(buckets) - 1; i > 0 && k > 0; i-- {
		answer = append(answer, buckets[i]...)
		k -= len(buckets[i])
	}

	return answer
}
