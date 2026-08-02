func maxArea(heights []int) int {
	mx := 0
	l := 0
	r := len(heights) - 1
	for l < r {
		mx = max(mx, (r-l)*min(heights[l], heights[r]))
		if heights[l] < heights[r] {
			l++
		} else {
			r--
		}
	}
	return mx
}
