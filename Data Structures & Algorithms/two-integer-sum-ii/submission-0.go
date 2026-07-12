func twoSum(numbers []int, target int) []int {
	var indices []int
	left := 0
	right := len(numbers) - 1
	for left < right {
		currSum := numbers[left] + numbers[right]
		if currSum == target {
			indices = append(indices, left + 1)
			indices = append(indices, right + 1)
			return indices
		}else if currSum < target {
			left++
		}else{
			right--
		}
	}

	return indices

}
