package t26

func removeDuplicates(nums []int) int {
	var unic []int
	unic1 := make([]int, 0)

	unic = append(unic, nums[0])
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			unic = append(unic, nums[i])
		}
	}

	for i := 0; i < len(unic); i++ {
		nums[i] = unic[i]
	}

	return len(unic)
}

type A struct {
}

var a1 A
var a2 = &A{}
var a3 = new(A)
