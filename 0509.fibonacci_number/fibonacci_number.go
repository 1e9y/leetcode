package fibonacci_number

var dp = map[int]int{
	0: 0,
	1: 1,
	2: 1,
}

func fib(n int) (value int) {
	value, ok := dp[n]
	if ok {
		return
	}

	value = fib(n-1) + fib(n-2)
	dp[n] = value

	return
}
