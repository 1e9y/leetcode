package unique_paths

var memo = make(map[int]map[int]int)

func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	if v, ok := memo[m][n]; ok {
		return v
	}
	paths := uniquePaths(m-1, n) + uniquePaths(m, n-1)
	if _, ok := memo[m]; !ok {
		memo[m] = make(map[int]int)
	}
	memo[m][n] = paths
	return paths
}
