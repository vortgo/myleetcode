package t5

func longestPalindrome(s string) string {

	start, end := 0, 0
	for currentCenter := 0; currentCenter < len(s); currentCenter++ {
		leven := expand(s, currentCenter, currentCenter)
		lodd := expand(s, currentCenter, currentCenter+1)

		l := max(leven, lodd)

		if l > end-start {
			start = currentCenter - (l-1)/2
			end = currentCenter + l/2
		}
	}

	return s[start : end+1]
}

func expand(s string, li, ri int) int {
	for li >= 0 && ri < len(s) && s[li] == s[ri] {
		li--
		ri++
	}
	return ri - li - 1
}
