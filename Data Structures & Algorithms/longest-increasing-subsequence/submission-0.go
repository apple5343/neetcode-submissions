func lengthOfLIS(nums []int) int {
    seq := []int{nums[0]}
	for _, num := range nums[1:] {
		if seq[len(seq) - 1] < num {
			seq = append(seq, num)
		} else {
			seq[bs(seq, num)] = num
		}
	}
	return len(seq)
}

func bs(nums []int, x int) int {
	l := -1
	r := len(nums)
	
	for l + 1 < r {
		m := l + (r-l)/2
		if nums[m] >= x {
			r = m
		} else {
			l = m
		}
	}
	return r
}
