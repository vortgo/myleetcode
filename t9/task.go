package t9

func isPalindrome(x int) bool {
	s := x
	if x < 0 {
		return false
	}

	reverseNum := 0
	for x > 0 {
		rNum := x % 10
		x = (x - rNum) / 10
		reverseNum = reverseNum*10 + rNum
	}

	return s == reverseNum
}
