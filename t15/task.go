package t15

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	seen := make(map[string]bool)
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		left := i + 1
		right := len(nums) - 1

		if left >= len(nums) {
			break
		}

		for left < right {
			if nums[i]+nums[left]+nums[right] == 0 {
				tempResult := []int{nums[i], nums[left], nums[right]}

				key := fmt.Sprintf("%d,%d,%d", tempResult[0], tempResult[1], tempResult[2])
				if !seen[key] {
					seen[key] = true
					result = append(result, tempResult)
				}

				left++
				right--

				if left == len(nums) {
					break
				}
			}

			if nums[i]+nums[left]+nums[right] > 0 {
				right--
			}
			if nums[i]+nums[left]+nums[right] < 0 {
				left++
			}
		}
	}

	return result
}
