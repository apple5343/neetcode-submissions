func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	result := [][]int{}
	n := len(nums)
	for i := 0; i < n; i++ {
		if i != 0 && nums[i-1] == nums[i] {
			continue
		}
		l := i + 1
		r := n - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				l++
				for l < r && nums[l-1] == nums[l] {
					l++
				}
			} else if sum > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return result
}
