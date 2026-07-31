
func coinChange(coins []int, amount int) int {
    dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = 1 << 20
	}
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		for _, coin := range coins {
			if i - coin >= 0 {
				dp[i] = min(dp[i], dp[i-coin] + 1)
			}
		}
	}
	if dp[amount] == 1 << 20 {
		return -1
	}
	return dp[amount]
}
