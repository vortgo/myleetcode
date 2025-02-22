package t13

func romanToInt(s string) int {
	mComplex := map[string]int{
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	mSimple := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	num := 0

	for len(s) > 0 {
		if len(s) >= 2 {
			complexPath := s[:2]
			if v, ok := mComplex[complexPath]; ok {
				num += v
				s = s[2:]
				continue
			}
		}

		num += mSimple[s[:1]]
		s = s[1:]

	}
	return num
}
