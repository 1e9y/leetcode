package min_cost_climbing_stairs

func minCostClimbingStairs(cost []int) int {
	minCost := map[int]int{
		0: 0,
		1: 0,
	}

	for i := 2; i <= len(cost); i++ {
		minCost[i] = min(cost[i-1]+minCost[i-1], cost[i-2]+minCost[i-2])
	}

	return minCost[len(cost)]
}
