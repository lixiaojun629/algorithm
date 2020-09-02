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

//CombinationSum4 换硬币组合数量，硬币顺序不同算不同的组合
//377. https://leetcode-cn.com/problems/combination-sum-iv/
func CombinationSum4(nums []int, target int) int {
	//sum[i] 组成i的最大组合数
	sum := make([]int, target+1)
	sum[0] = 1
	for i := 0; i <= target; i++ {
		for _, num := range nums {
			if i >= num {
				sum[i] += sum[i-num]
			}
		}
	}
	return sum[target]
}

//CoinChange2 换硬币组合数量，硬币顺序不同算一个组合
//518 https://leetcode-cn.com/problems/coin-change-2/
func CoinChange2(amount int, coins []int) int {
	sum := make([]int, amount+1)
	sum[0] = 1

	for _, coin := range coins { //硬币循环在外层，表示每次加一个硬币进来，确保组合不重复
		for i := coin; i <= amount; i++ {
			sum[i] += sum[i-coin]
		}
	}
	return sum[amount]
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
