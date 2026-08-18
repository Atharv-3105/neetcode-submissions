type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0{
		return ""
	}

	//Get the len of each string and append size + special char + string for all the strings
	//Ex: 5$hello7$helloww3$hel
	var sizes []string 
	for _, str := range strs{
		sizes = append(sizes, strconv.Itoa(len(str)))
	}

	return strings.Join(sizes, ",") + "$" + strings.Join(strs, "")
}

func (s *Solution) Decode(encoded string) []string {
	if encoded == "" {
		return []string{}
	}

	parts := strings.SplitN(encoded, "$", 2)
	sizes := strings.Split(parts[0], ",")
	var result []string 
	i := 0
	for _, s := range sizes {
		if s == ""{
			continue
		}

		lengthStr, _ := strconv.Atoi(s)
		result = append(result, parts[1][i : i + lengthStr])
		i += lengthStr
	}

	return result
}
