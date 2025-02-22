package t1

func twoSum(nums []int, target int) []int {
	var existNumber = make(map[int]int)

	for i := 0; i < len(nums); i++ {
		searchNumber := target - nums[i]

		_, ok := existNumber[searchNumber]
		if !ok {
			existNumber[nums[i]] = i
		} else {
			return []int{existNumber[searchNumber], i}
		}
	}

	return []int{}
}
