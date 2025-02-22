package t20

func isValid(s string) bool {
	opensParentheses := map[string]bool{"(": true, "[": true, "{": true}
	closesParentheses := map[string]bool{")": true, "]": true, "}": true}

	parentheses := map[string]string{")": "(", "}": "{", "]": "["}

	var stack []string

	for i := 0; i < len(s); i++ {
		symbol := string(s[i])
		if ok, _ := opensParentheses[symbol]; ok {
			stack = append(stack, symbol)
			continue
		}

		if ok, _ := closesParentheses[symbol]; ok {
			if len(stack)-1 < 0 {
				return false
			}

			lastSymbol := stack[len(stack)-1]
			correctLastSymbol := parentheses[symbol]

			if lastSymbol != correctLastSymbol {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
