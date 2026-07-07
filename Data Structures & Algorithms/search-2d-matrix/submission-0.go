func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m; i++ {
		firstVal := matrix[i][0]
		lastVal := matrix[i][n - 1]
		if firstVal <= target && lastVal >= target {
			return binarySearch(target, matrix[i])
		}
	}
	return false
}
func binarySearch(target int, nums []int) bool {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right - left) / 2
		if nums[mid] == target {
			return true
		}else if nums[mid] < target {
			left = mid + 1
		}else{
			right = mid - 1
		}
	}
	return false
}
