func twoSum(nums []int, target int) []int {
    seenMap := make(map[int]int)
	for i, num := range nums {
		el := target - num
		if idx, ok := seenMap[el]; ok {
			return []int{idx, i}
		}

		seenMap[num] = i
	}

	return []int{}
}
