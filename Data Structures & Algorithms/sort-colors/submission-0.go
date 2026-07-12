func sortColors(nums []int) {
    low := 0
	mid := 0
	high := len(nums) - 1

	for mid <= high {
		if nums[mid] == 0 {
			temp := nums[low]
			nums[low] = nums[mid]
			nums[mid] = temp
			low++
			mid++
		}else if nums[mid] == 1 {
			mid++
		}else{
			temp := nums[high]
			nums[high] = nums[mid]
			nums[mid] = temp
			high--
		}
	}
}
