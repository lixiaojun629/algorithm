package bit

//338 https://leetcode-cn.com/problems/counting-bits/
func countBits(num int) []int {
	counts := make([]int, num+1)
	for i := 1; i <= num; i++ {
		counts[i] = counts[i&(i-1)] + 1
	}
	return counts
}
