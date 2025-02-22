package t39

import (
	"sort"
	"strconv"
)

func combinationSum(candidates []int, target int) [][]int {
	m := make(map[string][]int)
	find(candidates, []int{}, 0, target, m)

	var result [][]int

	for _, v := range m {
		result = append(result, v)
	}

	return result
}

func find(combinations []int, currentCombinations []int, currentSum, t int, result map[string][]int) {
	for l := 0; l < len(combinations); l++ {
		if combinations[l]+currentSum == t {
			combinationCopy := make([]int, len(currentCombinations))
			copy(combinationCopy, currentCombinations)
			combinationCopy = append(combinationCopy, combinations[l])
			sort.Ints(combinationCopy)
			result[getKey(combinationCopy)] = combinationCopy
			continue
		}

		if combinations[l]+currentSum > t {
			continue
		}

		//combinations[l]+num < t
		find(combinations, append(currentCombinations, combinations[l]), combinations[l]+currentSum, t, result)
	}
}

func getKey(i []int) string {
	var key []byte
	for _, v := range i {
		key = append(key, strconv.Itoa(v)...)
	}

	return string(key)
}
