package t6

func convert(s string, numRows int) string {
	fullColumnCount := (numRows - 1) * 2
	if fullColumnCount == 0 {
		return s
	}

	n := len(s)
	res := make([]byte, 0, n)
	shift := numRows*2 - 2
	for row := 0; row < numRows; row++ {
		for i := row; i < n; i += shift {
			res = append(res, s[i])
			if row > 0 && row < numRows-1 {
				idx := i + shift - 2*row
				if idx < n {
					res = append(res, s[idx])
				}
			}
		}
	}

	return string(res)
}
