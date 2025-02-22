package t35

func searchInsert(nums []int, target int) int {
	return find(nums, target, 0)
}

func find(nums []int, target int, fromStart int) int {
	centerIndex := (len(nums) - 1) / 2
	if nums[centerIndex] == target {
		return fromStart + centerIndex
	}

	if len(nums) == 1 {
		if nums[0] < target {
			return fromStart + 1
		} else {
			return fromStart
		}
	}

	if nums[centerIndex] > target {
		nums = nums[:centerIndex]
	} else {
		nums = nums[centerIndex+1:]
		fromStart += centerIndex + 1
	}
	return find(nums, target, fromStart)
}
