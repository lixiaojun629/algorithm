package dp

//https://leetcode-cn.com/problems/coin-change/

//dp[i]表示总额为i的零钱，使用硬币的最小数量
//dp[i] = min(dp[i-coins[0...j]+1)
func CoinChangeDP(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}

	dp[0] = 0
	for i := 0; i <= amount; i++ {
		for _, coin := range coins {
			if i >= coin {
				dp[i] = minInt(dp[i-coin]+1, dp[i])
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func minInt(a, b int) int {
	if a < b {
		b = a
	}
	return b
}

//递归暴力破解
func CoinChangeRecursive(coins []int, amount int) int {
	min = amount
	change(coins, 0, amount)
	return res
}

var res = -1
var min = 0

func change(coins []int, count, amount int) {
	if amount == 0 {
		if count < min {
			min = count
			res = min
		}
		return
	}
	for _, coin := range coins {
		if amount >= coin {
			change(coins, count+1, amount-coin)
		}
	}
}
