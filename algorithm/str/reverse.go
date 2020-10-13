package str

import (
	"math"
	"strconv"
)

//https://leetcode-cn.com/problems/reverse-integer/
func reverse(x int) int {
	if x <= -1<<31 || x > math.MaxInt32 {
		return 0
	}
	if x < 0 {
		ans := reverse(-x)
		return -ans
	}
	str := strconv.Itoa(x)
	target := make([]byte, 0)
	leadZero := true
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] != '0' {
			leadZero = false
		}
		if !leadZero {
			target = append(target, str[i])
		}
	}

	num, _ := strconv.Atoi(string(target))
	if num > math.MaxInt32 {
		return 0
	}
	return num
}
