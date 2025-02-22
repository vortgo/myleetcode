package t29

import (
	"log"
	"math"
	"strconv"
)

func divide(dividend int, divisor int) int {
	dividend, isMinusDividend := removeMinus(dividend)
	divisor, isMinusDivisor := removeMinus(divisor)

	dividendStr := strconv.Itoa(dividend)

	var rs []int
	var reminder int
	for {
		if len(dividendStr) < 1 && reminder < divisor {
			break
		}

		var aStr string
		if reminder >= divisor {
			aStr = strconv.Itoa(reminder)
			reminder = 0
		} else {
			if len(dividendStr) == 0 {
				aStr = strconv.Itoa(reminder)
				reminder = 0
			} else {
				aStr = strconv.Itoa(reminder) + string(dividendStr[0])
				dividendStr = dividendStr[1:]
			}
		}

		aNum, err := strconv.Atoi(aStr)
		if err != nil {
			log.Fatal(err)
		}

		var res int
		res, reminder = div(aNum, divisor)
		rs = append(rs, res)
	}

	if len(rs) == 0 {
		return 0
	}

	var resultStr string
	for _, v := range rs {
		resultStr = resultStr + strconv.Itoa(v)
	}

	result, err := strconv.Atoi(resultStr)
	if err != nil {
		log.Fatal(err)
	}

	if (isMinusDivisor && isMinusDividend) || (!isMinusDivisor && !isMinusDividend) {
		return returnMaxMinInt(result)
	}

	return returnMaxMinInt(addMinus(result))
}

func removeMinus(num int) (int, bool) {
	strNum := strconv.Itoa(num)
	if strNum[0] == '-' {
		strNum = strNum[1:]
		newInt, err := strconv.Atoi(strNum)
		if err != nil {
			log.Fatal(err)
		}
		return newInt, true
	}

	return num, false
}

func addMinus(num int) int {
	strNum := strconv.Itoa(num)

	newInt := []byte{'-'}
	newInt = append(newInt, strNum...)
	i, err := strconv.Atoi(string(newInt))
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func returnMaxMinInt(i int) int {
	if i > math.MaxInt32 {
		return math.MaxInt32
	}

	if i < math.MinInt32 {
		return math.MinInt32
	}

	return i
}

func div(a, b int) (int, int) {
	result := 0
	for {
		if a < b {
			break
		}

		a -= b
		result++
	}
	return result, a
}
