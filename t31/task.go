package t31

import "sort"

func nextPermutation(nums []int) {
	var endIndex int
	for i := len(nums) - 1; i >= 1; i-- {
		if nums[i] > nums[i-1] {
			endIndex = i
			break
		}
	}

	replIndex := len(nums) - 1
	if endIndex == 0 {
		for i := 0; i < replIndex; i++ {
			replIndex = len(nums) - 1 - i
			nums[i], nums[replIndex] = nums[replIndex], nums[i]
		}
		return
	}

	preEndIndex := endIndex - 1
	endNums := nums[endIndex:]
	sort.Ints(endNums)

	preEndNum := nums[preEndIndex]

	replIndex = len(endNums) - 1
	for i := 0; i < len(endNums); i++ {
		if preEndNum < endNums[i] {
			replIndex = i
			break
		}
	}

	nums[preEndIndex] = endNums[replIndex]
	endNums[replIndex] = preEndNum
}
