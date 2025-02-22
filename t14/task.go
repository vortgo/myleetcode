package t14

func longestCommonPrefix(strs []string) string {
	var maxPrefix string
	firstWord := strs[0]
main:
	for i := 0; i < len(firstWord); i++ {
		for _, word := range strs[1:] {
			if len(word) == i || word[i] != firstWord[i] {
				break main
			}
		}
		maxPrefix = firstWord[:i+1]
	}

	return maxPrefix
}
