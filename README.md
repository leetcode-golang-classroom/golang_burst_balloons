# golang_burst_balloons

You are given `n` balloons, indexed from `0` to `n - 1`. Each balloon is painted with a number on it represented by an array `nums`. You are asked to burst all the balloons.

If you burst the `ith` balloon, you will get `nums[i - 1] * nums[i] * nums[i + 1]` coins. If `i - 1` or `i + 1` goes out of bounds of the array, then treat it as if there is a balloon with a `1` painted on it.

Return *the maximum coins you can collect by bursting the balloons wisely*.

## Examples

**Example 1:**

```
Input: nums = [3,1,5,8]
Output: 167
Explanation:
nums = [3,1,5,8] --> [3,5,8] --> [3,8] --> [8] --> []
coins =  3*1*5    +   3*5*8   +  1*3*8  + 1*8*1 = 167
```

**Example 2:**

```
Input: nums = [1,5]
Output: 10

```

**Constraints:**

- `n == nums.length`
- `1 <= n <= 300`
- `0 <= nums[i] <= 100`

## 解析

給定一個正整數陣列 nums, 代表會放在氣球上的 label

定義每次可以 burst 一個氣球 然後把該氣球與鄰近的氣球 label 做相乘 作為可以拿到的 coins

當 burst 其中一個氣球後， 其左右位置的氣球就變成鄰近的氣球

舉例 [3,1,5,8] 當 burst label 是 1 的氣球 取得 3*1*5 coin 然後 nums 變成 [3,5,8]

當 burst 的左右有超出範圍 的狀況會把該位置當作 label 是 1

要求寫一個演算法算出能夠burst 出做多的 coins 數量

思考如何去窮舉每個可能

由於左右都會被當成 1

為了方便思考 可以在 nums 左右新增一個元素 設定值為 1

如下圖：

![](https://i.imgur.com/89Saguh.png)

則可以觀察發現

要窮舉可以透過 針對一個 L,R 範圍遍歷 所有 i = L.. R

找出最大的值會有以下的關係式



![](https://i.imgur.com/v8uUzm6.png)

對每個 i 需要遍歷 L 到 R 大約 是 O(n)

每個 L, R 的可能有 O($n^2$) 的可能

透過 memorization 避免走過重複的路徑

所以總共的時間複雜度是 O($n^3$)

空間複雜度是 O($n^2$)

如果透過 DFS 遞迴方式 還會有 call stack 的 空間複雜度

Tabulation Dynamic Programming 方式的話

可以透過以下方式

透過 left 起始位置，找出所有可能的 right

找出所有的 i

定義 dp[left][right] =  從 bufferNums[left:right+1] 能夠找到最大的 coins 數值

所求就是 dp[1][len(nums)]

時間複雜度O($n^3$)

空間複雜度是 O($n^2$)

## 程式碼
```go
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

```
## 困難點

1. 要思考如何窮舉出所有可能性
2. 要看出每個乘機運算的遞迴關係

## Solve Point

- [x]  建立 二維矩陣 dp 用來存中間運算的結果
- [x]  逐步依照遞迴關係式去運算結果