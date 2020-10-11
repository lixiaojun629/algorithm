package dp

import (
	"testing"
)

func TestCutRope(t *testing.T) {
	t.Log(cuttingRope(5))
}

func TestCoinChange(t *testing.T) {
	//t.Log(CoinChangeRecursive([]int{2}, 3))
	//t.Log(CoinChangeRecursive([]int{1, 2, 5}, 11))
	//t.Log(CoinChangeDP([]int{2}, 3))
	//t.Log(CoinChangeDP([]int{1, 2, 5}, 11))
	t.Log(CoinChangeDP([]int{2, 5, 10, 1}, 27))
}

func TestLengthOfLIS(t *testing.T) {
	t.Log(LengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	//t.Log(LengthOfLISDP([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}

func TestMaxProduct(t *testing.T) {
	t.Log(MaxProduct([]int{2, 3, -2, 4}))
	t.Log(MaxProduct([]int{-2}))
}

func TestLongestCommonSubsequence(t *testing.T) {
	t.Log(LongestCommonSubsequence("a", "abac"))
}

func TestStockMaxProfit(t *testing.T) {
	//t.Log(StockMaxProfit([]int{3,3,5,0,0,3,1,4}))
	//t.Log(StockMaxProfit([]int{1,2,3,4,5}))
	//t.Log(StockMaxProfitK(2, []int{3, 3, 5, 0, 0, 3, 1, 4}))
	//t.Log(StockMaxProfitK(2, []int{1, 2, 3, 4, 5}))
	t.Log(greedy([]int{2, 4, 1}))
}

func TestLengthOfLongestSubstring(t *testing.T) {
	// res := LengthOfLongestSubstring("abcabcbb")
	// res := LengthOfLongestSubstring("pwwkew")
	// res := LengthOfLongestSubstring("aab")
	// fmt.Println(res)
	// t.Log("max length", res)
}

func TestPackageWeight(t *testing.T) {
	t.Log(calcWeight([]int{2, 2, 4, 6, 3}, 9))
	t.Log(calcWeight2([]int{2, 2, 4, 6, 3}, 9))
}
