package dp

import (
	"fmt"
)

//https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/

func StockMaxProfit(prices []int) int {
	maxTransCount := 2
	length := len(prices)
	fmt.Println(prices)
	//mp[i][k][s]表示第i天交易k次后，拥有股票(s=1表示有，s=0表示没有)状态下的已获取的最大利润；
	//买和卖算一次交易，买的时候增加k;
	mp := make([][][]int, length+1)
	for i := 0; i < length+1; i++ {
		mp[i] = make([][]int, maxTransCount+1)
		for j := 0; j <= maxTransCount; j++ {
			mp[i][j] = make([]int, 2)
		}
	}

	//bad case
	minNum := -1 << 31
	for i := 0; i < length; i++ {
		mp[i][0][1] = minNum //每天，总交易次数是0情况下，不可能持有股票
	}
	for k := 0; k <= maxTransCount; k++ {
		mp[0][k][1] = minNum //第0天，还没开始交易，不可能存在持有股票的情况
	}
	//state transfer
	for i := 1; i <= length; i++ {
		for k := 1; k <= maxTransCount; k++ {
			fmt.Printf("i=%d,k=%d\n", i, k)
			fmt.Printf("mp[%d][%d][0]:%d , ", i-1, k, mp[i-1][k][0])
			fmt.Printf("mp[%d][%d][1]:%d , ", i-1, k, mp[i-1][k][1])
			fmt.Printf("mp[%d][%d][0]:%d , ", i-1, k-1, mp[i-1][k-1][0])
			fmt.Printf("mp[%d][%d][1]:%d\n", i-1, k-1, mp[i-1][k-1][1])
			mp[i][k][1] = max(mp[i-1][k][1], mp[i-1][k-1][0]-prices[i-1]) //buy
			mp[i][k][0] = max(mp[i-1][k][0], mp[i-1][k][1]+prices[i-1])   //sell
			fmt.Printf("mp[%d][%d][0]:%d , ", i, k, mp[i][k][0])
			fmt.Printf("mp[%d][%d][1]:%d\n", i, k, mp[i][k][1])
		}
	}

	return mp[length][maxTransCount][0]
}

func max(a, b int) int {
	if a < b {
		a = b
	}
	return a
}

func StockMaxProfitK(k int, prices []int) int {
	n := len(prices)
	if k > n/2 {
		return greedy(prices)
	}
	//mp[nn][kk][s]表示第nn天，第kk次交易，股票是否持有(s==0||s==1)状态下的最大利润
	mp := make([][][]int, 2)
	for nn := 0; nn <= 1; nn++ {
		mp[nn] = make([][]int, k+1)
		for kk := 0; kk <= k; kk++ {
			mp[nn][kk] = make([]int, 2)
		}
	}

	//bad case, 特殊处理持有状态下的最大利润，以下两种持有状态都不可能发生，设置为最小值
	//如果不处理，默认值为0，导致状态转移时卖出不可能持有的股票，结果不准。
	minInt := -1 << 31
	for nn := 0; nn <= 1; nn++ {
		mp[nn][0][1] = minInt
	}
	for kk := 0; kk <= k; kk++ {
		mp[0][kk][1] = minInt
	}

	//状态转移
	for nn := 1; nn <= n; nn++ {
		xn, yn := nn%2, (nn-1)%2
		for kk := 1; kk <= k; kk++ {
			//买入和卖出加一起算一次交易，买入时kk+1
			mp[xn][kk][0] = max(mp[yn][kk][0], mp[yn][kk][1]+prices[nn-1])
			mp[xn][kk][1] = max(mp[yn][kk][1], mp[yn][kk-1][0]-prices[nn-1])
		}
	}
	return mp[n%2][k][0]
}

func greedy(prices []int) int {
	n := len(prices)
	fmt.Println(prices)
	profit := 0
	for nn := 1; nn < n; nn++ {
		if prices[nn] > prices[nn-1] {
			profit += prices[nn] - prices[nn-1]
		}
	}
	return profit
}
