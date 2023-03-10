package climbing_stairs

func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}
	prev, curr, next := 1, 2, 0
	for i := 2; i < n; i++ {
		next = prev + curr
		prev = curr
		curr = next
	}
	return curr
}
