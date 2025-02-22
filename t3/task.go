package t3

func lengthOfLongestSubstring(s string) int {
	runeStr := []rune(s)
	var existSymbols = make(map[rune]bool)
	var longestSlice []rune
	var maxLen int
	for i := 0; i < len(runeStr); i++ {
		r := runeStr[i]
		if _, ok := existSymbols[r]; !ok {
			existSymbols[r] = true
			longestSlice = append(longestSlice, r)
		} else {
			if len(longestSlice) > maxLen {
				maxLen = len(longestSlice)
			}
			var symbolIndexInSlice int
			for k, v := range longestSlice {
				delete(existSymbols, v)
				if v == r {
					symbolIndexInSlice = k
					break
				}
			}
			longestSlice = longestSlice[symbolIndexInSlice+1:]
			longestSlice = append(longestSlice, r)
			existSymbols[r] = true
		}
	}
	if len(longestSlice) > maxLen {
		maxLen = len(longestSlice)
	}

	return maxLen
}

/*
func lengthOfLongestSubstring(s string) int {
	runeStr := []rune(s)
	var maxLen int
	for len(runeStr) > maxLen {
		l := getLenWithoutRepeating(runeStr)
		if l > maxLen {
			maxLen = l
		}
		runeStr = runeStr[1:]
	}

	return maxLen
}

func getLenWithoutRepeating(s []rune) int {
	var existSymbols = make(map[rune]int)
	var maxLen int
	for i := 0; i < len(s); i++ {
		symbol := s[i]
		if _, ok := existSymbols[symbol]; !ok {
			maxLen++
			existSymbols[symbol] = i
		} else {
			return maxLen
		}
	}

	return maxLen
}
*/
