func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	last := 0
	n := 1
    for i := 1; i < len(intervals); i++ {
		if intervals[last][1] <= intervals[i][0] {
			n++
			last = i
		}
	}
	return len(intervals) - n
}
