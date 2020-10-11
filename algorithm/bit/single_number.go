package bit

//https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/
func singleNumbers(nums []int) []int {
	var ret int = 0
	for _, num := range nums {
		ret ^= num
	}
	var div int = 1
	for {
		if div&ret != 0 {
			break
		}
		div <<= 1
	}
	var a, b int
	for _, num := range nums {
		if num&div != 0 {
			a ^= num
		} else {
			b ^= num
		}
	}
	return []int{a, b}
}
