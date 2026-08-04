func topKFrequent(nums []int, k int) []int {
	d := 1000
	maxFreq := 0
	freqs := make([]int, d*2+1)
	for _, num := range nums {
		num += d
		freqs[num]++
		maxFreq = max(maxFreq, freqs[num])
	}
	ordered := make([][]int, maxFreq + 1)
	for num, freq := range freqs {
		ordered[freq] = append(ordered[freq], num-d)
	}
	result := make([]int, 0, k)
	for len(result) < k {
		if len(ordered[maxFreq]) != 0 {
			result = append(result, ordered[maxFreq]...)
		}
		maxFreq--
	}
	return result
}
