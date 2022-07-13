package sol

func maxCoins(nums []int) int {
	numsLen := len(nums)
	bufferNums := make([]int, numsLen+2)
	bufferNums[0], bufferNums[numsLen+1] = 1, 1
	copy(bufferNums[1:], nums)
	dp := make([][]int, numsLen+2)
	for row := range dp {
		dp[row] = make([]int, numsLen+2)
	}
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for left := numsLen; left >= 1; left-- {
		for right := left; right <= numsLen; right++ {
			dp[left][right] = 0
			for i := left; i <= right; i++ {
				// extract from i
				coins := bufferNums[left-1]*bufferNums[i]*bufferNums[right+1] + dp[left][i-1] + dp[i+1][right]
				dp[left][right] = max(dp[left][right], coins)
			}
		}
	}
	return dp[1][numsLen]
}
