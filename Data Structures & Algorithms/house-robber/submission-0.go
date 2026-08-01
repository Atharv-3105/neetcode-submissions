func rob(nums []int) int {
	memo := make([]int, len(nums) + 1)
	for i := 0; i < len(nums); i++{
		memo[i] = -1
	}

	var dfs func(i int) int
    dfs = func(idx int) int {
		if idx >= len(nums) {
			return 0
		}
		if memo[idx] != -1{
			return memo[idx]
		}

		memo[idx] = max(nums[idx] + dfs(idx + 2), dfs(idx + 1))
		return memo[idx]
	}

	return dfs(0)
}
func max(a, b int) int{
	if a > b {
		return a
	}
	return  b
}
