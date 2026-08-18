func groupAnagrams(strs []string) [][]string {
	freqMap := make(map[[26]int][]string)

	for _, s := range strs{
		var currFreq [26]int
		for _, c := range s{
			currFreq[c - 'a']++
		}

		freqMap[currFreq] = append(freqMap[currFreq], s)
	}

	var result [][]string 
	for _, grp := range freqMap{
		result = append(result, grp)
	}

	return result
}
