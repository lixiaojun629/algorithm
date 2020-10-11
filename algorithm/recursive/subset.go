package recursive

//https://leetcode-cn.com/problems/subsets/
func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	var dfs func(int, []int)

	dfs = func(curr int, subset []int) {
		if curr == len(nums) {
			result = append(result, append([]int(nil), subset...))
			return
		}
		subset = append(subset, nums[curr])
		dfs(curr+1, subset)
		dfs(curr+1, subset[:len(subset)-1])
	}

	dfs(0, make([]int, 0))
	return result
}
