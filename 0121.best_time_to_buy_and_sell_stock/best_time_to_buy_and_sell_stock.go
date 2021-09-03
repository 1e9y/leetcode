package best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	max := 0
	min := prices[0]
	for _, n := range prices {
		if n < min {
			min = n
		}
		if n-min > max {
			max = n - min
		}
	}
	return max
}
