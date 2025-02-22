package t38

import "strconv"

func countAndSay(n int) string {
	res := "1"
	for i := 1; i < n; i++ {
		res = rle(res)
	}
	return res
}

func rle(s string) string {
	var result string
	var currentSymbol string
	var currentCount int

	for _, v := range s {
		symbol := string(v)
		if symbol == currentSymbol {
			currentCount++
		} else {
			if len(currentSymbol) > 0 {
				result += strconv.Itoa(currentCount) + currentSymbol
			}
			currentCount = 1
			currentSymbol = symbol
		}
	}

	result += strconv.Itoa(currentCount) + currentSymbol
	return result
}
