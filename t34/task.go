package t34

func searchRange(nums []int, target int) []int {
	index := s(nums, target, 0)

	if index == -1 {
		return []int{-1, -1}
	}

	startIndex := index
	endIndex := index
	for {
		if startIndex == 0 {
			if nums[startIndex] != target {
				startIndex++
				break
			}
			break
		}

		if nums[startIndex] != target {
			startIndex++
			break
		}
		startIndex--
	}

	for {
		if endIndex > len(nums)-1 {
			break
		}

		if nums[endIndex] != target {
			break
		}
		endIndex++
	}

	return []int{startIndex, endIndex - 1}
}

func s(nums []int, target, startIndex int) int {
	if len(nums) == 1 && nums[0] != target || len(nums) == 0 {
		return -1
	}

	center := (len(nums) - 1) / 2

	if nums[center] > target {
		return s(nums[:center], target, startIndex)
	}

	if nums[center] < target {
		startIndex += center + 1
		return s(nums[center+1:], target, startIndex)
	}

	return startIndex + center
}
