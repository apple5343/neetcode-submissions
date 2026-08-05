func combinationSum(nums []int, target int) [][]int {
	var backtracking func(path []int, sum, i int) [][]int
	backtracking = func(path []int, sum, i int) [][]int {
		if sum == target {
			cpy := make([]int, len(path))
			copy(cpy, path)
			return [][]int{cpy}
		}
		result := [][]int{}
		for j, val := range nums[i:] {
			if sum+val <= target {
				result = append(result, backtracking(append(path, val), sum+val, i+j)...)
			}
		}
		return result
	}
	return backtracking([]int{}, 0, 0)
}
