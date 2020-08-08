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
