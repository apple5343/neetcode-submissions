func hasDuplicate(nums []int) bool {
    visited := make(map[int]bool)
	for _, num := range nums {
		if visited[num] {
			return true
		}
		visited[num] = true
	}
	return false
}
