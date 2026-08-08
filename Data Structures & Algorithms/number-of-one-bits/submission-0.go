func hammingWeight(n int) int {
	result := 0
	for i := 0; i < 31; i++ {
		if n & 1 != 0 {
			result++
		}
		n >>= 1
	}
	return result
}
