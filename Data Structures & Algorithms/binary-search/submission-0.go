func search(nums []int, target int) int {
	idx := -1
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right - left) / 2;
		if nums[mid] == target {
			idx = mid
			return idx
		}else if nums[mid] < target {
			left = mid + 1
		}else {
			right = mid - 1
		}
	}
	return idx
}
