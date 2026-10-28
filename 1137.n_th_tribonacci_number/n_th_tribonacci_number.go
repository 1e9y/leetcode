package n_th_tribonacci_number

var dp = map[int]int{
	0: 0,
	1: 1,
	2: 1,
	3: 2,
}

func tribonacci(n int) int {
	v, ok := dp[n]
	if ok {
		return v
	}

	v = tribonacci(n-1) + tribonacci(n-2) + tribonacci(n-3)
	dp[n] = v

	return v
}
