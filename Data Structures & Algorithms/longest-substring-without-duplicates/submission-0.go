func lengthOfLongestSubstring(s string) int {
	n := len(s)
	freq := make(map[byte]int)
	left := 0
	ans := 0

	for right := 0; right < n; right++ {
		freq[s[right]]++

		//Shrink the window from left, if we duplicates
		for freq[s[right]] > 1 {
			freq[s[left]]--
			left++
		}

		//len of the window is right - len + 1
		if ans < right - left + 1 {
			ans = right - left + 1
		}
	}
	return ans
}
