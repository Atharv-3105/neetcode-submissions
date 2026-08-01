func minCostClimbingStairs(cost []int) int {
    n := len(cost)
	dp := make([]int, len(cost) + 1)
	for i := 2; i <= n;  i++{
		dp[i] = min(cost[i - 1] + dp[i - 1], cost[i - 2] + dp[i - 2])
	}

	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
