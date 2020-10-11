//Package dp is about dynamic programming
//淘宝双十一购物节有各种促销活动，比如满200减50，假设女朋友的购物车有n个（n>100）想买的商品，她希望从里面选几个
//在凑够满减的条件的前提下，让选出来的商品价格总和最大程度地接近满减条件（200元），这样可以最大限度的薅羊毛。
//作为程序员的你，能不能编个代码来帮她搞定
package dp

import "fmt"

//https://zh.wikipedia.org/wiki/%E8%83%8C%E5%8C%85%E9%97%AE%E9%A2%98
//背包问题，f(W) = max{f(W-1),{f(W-w[j])+p[j]|w[j]<=W}}(背包总重量，w[j]物品j的重量,p[j]物品j的价格
//ShoppingCart 动态规划购物车里面的商品, items 商品价格， n 商品个数，w 满减条件
func ShoppingCart(items []int, n, w int) {
	//states[i] 第i次选择后，物品总价数组
	//states[i][j] 第i次选择后，物品总价为j
	states := make([][]bool, n)
	for i := range states {
		states[i] = make([]bool, 3*w)
	}

	states[0][0] = true
	if items[0] <= 3*w {
		states[0][items[0]] = true
	}

	for i := 1; i < n; i++ {
		for j := 0; j < 3*w; j++ {
			if states[i-1][j] == true {
				states[i][j] = true //不购买第i个商品
			}
		}

		for j := 0; j < 3*w-items[i]; j++ {
			if states[i-1][j] == true {
				states[i][j+items[i]] = true //购买第i个商品
			}
		}
	}

	j := w
	for ; j < 3*w; j++ {
		if states[n-1][j] {
			break
		}
	}
	fmt.Println(j)
	if j == 3*w-1 {
		return //没有可行解
	}

	for i := n - 1; i >= 1; i-- {
		if j-items[i] >= 0 && states[i-1][j-items[i]] == true {
			fmt.Println(i, items[i])
			j -= items[i]
		}
	}

	if j != 0 {
		fmt.Println(0, items[0])
	}
}

//0-1背包问题
//有n个不可分割的物品，每个物品重量可能不同，这些物品的重量组成数组weights = [2,2,4,6,3]
//要把这些物品装到一个背包中，背包能承受的最大重量为 maxWeight, 在满足背包最大重量限制的前提下，
//背包中物品的总重量是多少？

func calcWeight(weights []int, maxWeight int) int {
	n := len(weights)
	states := make([][]bool, n)
	for i := 0; i < n; i++ {
		states[i] = make([]bool, maxWeight+1)
	}
	//states[i][w] 第i次把物品放到背包后，背包的重量是否可能为w
	states[0][0] = true
	if weights[0] <= maxWeight {
		states[0][weights[0]] = true
	}
	//状态转移 states[i] => states[i+1]
	for i := 1; i < n; i++ {
		for w := 0; w <= maxWeight; w++ {
			states[i][w] = states[i-1][w]
		}
		for w := 0; w+weights[i] <= maxWeight; w++ {
			states[i][w+weights[i]] = states[i][w+weights[i]] || states[i-1][w]
		}
	}
	ans := 0
	for w := maxWeight; w >= 0; w-- {
		if states[n-1][w] {
			ans = w
			break
		}
	}
	for w, i := ans, n-1; i >= 0 && w-weights[i] >= 0; i-- {
		if states[i][w-weights[i]] {
			fmt.Println("add", weights[i])
			w -= weights[i]
		}
	}
	return ans
}

func calcWeight2(weights []int, maxWeight int) int {
	n := len(weights)
	states := make([]bool, maxWeight+1)
	//states[i][w] 第i次把物品放到背包后，背包的重量是否可能为w
	states[0] = true
	if weights[0] <= maxWeight {
		states[weights[0]] = true
	}
	//状态转移 states[i] => states[i+1]
	for i := 1; i < n; i++ {
		for w := maxWeight; w >= weights[i]; w-- {
			states[w] = states[w] || states[w-weights[i]]
		}
	}
	for w := maxWeight; w >= 0; w-- {
		if states[w] {
			return w
		}
	}
	return 0
}
