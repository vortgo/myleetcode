package t7

import (
	"math"
	"strconv"
)

func reverse(x int) int {
	maxInt32S := []rune(strconv.Itoa(math.MaxInt32))
	//minInt32S := []rune(strconv.Itoa(math.MinInt32))[1:]
	negative := false
	s := []rune(strconv.Itoa(x))
	firstChar := s[0]
	if firstChar == '-' {
		negative = true
		s = s[1:]
	}

	reversed := make([]rune, len(s))
	ri := len(reversed) - 1
	for _, r := range s {
		reversed[ri] = r
		ri--
	}

	if len(reversed) > len(maxInt32S) {
		return 0
	}

	if len(reversed) <= len(maxInt32S) {

		if len(reversed) == len(maxInt32S) {
			if string(reversed) > string(maxInt32S) {
				return 0
			}
		}

		result := reversed
		if negative {
			result = []rune{'-'}
			result = append(result, reversed...)
		}
		revNum, err := strconv.Atoi(string(result))
		if err != nil {
			return 0
		}

		return revNum
	}

	return 0
}
