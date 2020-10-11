package dp

import "strconv"

//https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/
func translateNum(num int) int {
	str := strconv.Itoa(num)
	//dp[i] 从1到第i个数字，可以翻译的方法总数
	dp := make([]int, len(str)+1)
	dp[0] = 1
	for i := 1; i <= len(str); i++ {
		dp[i] = dp[i-1]
		if i == 1 {
			continue
		}
		num := str[i-2 : i]
		if num >= "10" && num <= "25" {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(str)]
}

func translateNum1(num int) int {
	a1, a2, a3 := 0, 0, 1
	str := strconv.Itoa(num)
	for i := 0; i < len(str); i++ {
		a1, a2 = a2, a3
		if i == 0 {
			continue
		}
		num := str[i-1 : i+1]
		if num >= "10" && num <= "25" {
			a3 += a1
		}
	}
	return a3
}
