func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	left := 0
	ans := len(nums) + 1

	windowSum := 0
	
	//We maintain windowSum; and if windowSum >= target we try to shrink the window
	for right := 0; right < n; right++ {

		windowSum += nums[right]

		//Shrink window if it's sum is greater than or equal to target
		for windowSum >= target {
			ans = min(ans, right - left + 1)
			windowSum -= nums[left]
			left++
		}
	}

	if ans == len(nums) + 1 {
		return 0
	}
	return ans

	
		
	
}
