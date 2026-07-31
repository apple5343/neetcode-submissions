func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		prev := result[len(result) - 1]
		cur := intervals[i]
		if max(prev[0], cur[0]) <= min(prev[1], cur[1]) {
			prev[1] = max(prev[1], cur[1])
		} else {
			result = append(result, []int{cur[0], cur[1]})
		}
	}
	return result
}
