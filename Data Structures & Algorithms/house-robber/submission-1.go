func rob(nums []int) int {
    //Bottom-up solution will be that for at each step we have a choice of either taking it and jumping
	//2 indices or skipping it and moving to next index
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++{
		dp[i] = max(nums[i] + dp[i - 2], dp[i - 1])
	}

	return dp[n - 1]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
