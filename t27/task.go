package t27

func removeElement(nums []int, val int) int {
	var rightLink = len(nums) - 1
	for i := 0; i <= rightLink; {
		if nums[rightLink] == val {
			rightLink--
			continue
		}
		if nums[i] == val {
			nums[i], nums[rightLink] = nums[rightLink], nums[i]
			rightLink--
		}
		i++
	}

	return rightLink + 1
}
