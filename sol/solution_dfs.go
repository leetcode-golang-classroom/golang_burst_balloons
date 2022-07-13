package sol

func maxCoinsDFS(nums []int) int {
	numLen := len(nums)
	bufferNums := make([]int, numLen+2)
	bufferNums[0], bufferNums[numLen+1] = 1, 1
	copy(bufferNums[1:], nums)
	cache := make([][]int, numLen+2)
	for row := range cache {
		cache[row] = make([]int, numLen+2)
	}
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var dfs func(left, right int) int
	dfs = func(left, right int) int {
		if left > right {
			return 0
		}
		if cache[left][right] != 0 {
			return cache[left][right]
		}
		for pivot := left; pivot <= right; pivot++ {
			coins := bufferNums[left-1] * bufferNums[pivot] * bufferNums[right+1]
			coins += dfs(left, pivot-1) + dfs(pivot+1, right)
			cache[left][right] = max(cache[left][right], coins)
		}
		return cache[left][right]
	}
	return dfs(1, numLen)
}
