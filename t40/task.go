package t40

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var result [][]int

	find(candidates, 0, 0, []int{}, target, &result)

	return result
}

func find(candidates []int, startFrom int, currentSum int, currentCombinations []int, target int, result *[][]int) {
	for i := startFrom; i < len(candidates); i++ {

		if i > startFrom && candidates[i] == candidates[i-1] {
			continue
		}

		if currentSum+candidates[i] == target {
			rCombinations := make([]int, len(currentCombinations))
			copy(rCombinations, currentCombinations)
			rCombinations = append(rCombinations, candidates[i])
			sort.Ints(rCombinations)
			*result = append(*result, rCombinations)

			return
		}

		if currentSum+candidates[i] > target {
			return
		}

		find(candidates, i+1, currentSum+candidates[i], append(currentCombinations, candidates[i]), target, result)
	}
}
