package longestcommonprefix

func longestCommonPrefix(strs []string) (result string) {
  common := ""
	// var bucket 
	index := 0
	if len(strs) == 1 {
		return strs[0]
	}

	char := ""
	for {
		if index == len(strs) {
			index = 0
			char = ""
			continue
		}

		if len(strs[index]) == 0 {
			break
		}

		curr := strs[index][0:1]
		if index == 0 {
			char += curr
		}

		if char != curr {
			break
		}

		str := strs[index][1:]
		strs[index] = str
		if index == len(strs) - 1 {
			common += char
		}
		index++
	}

	return common
}

func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	
	template := strs[0]
	if len(strs) == 1 {
		return template
	}
	
	for i, char := range template {
		for _, str := range strs {
			if i >= len(str) || char != rune(str[i]) {
				return template[0:i]
			}
		}
	}

	return template
}