package recursive

//https://leetcode-cn.com/problems/permutations/
func permute(nums []int) [][]int {
	var res = make([][]int, 0)
	var dfs func([]int, int)
	n := len(nums)
	dfs = func(output []int, pos int) {
		if pos == n {
			tmp := make([]int, n)
			copy(tmp, output)
			res = append(res, tmp)
		}
		for i := pos; i < n; i++ {
			output[pos], output[i] = output[i], output[pos]
			dfs(output, pos+1)
			output[pos], output[i] = output[i], output[pos]
		}
	}

	output := make([]int, len(nums))
	copy(output, nums)
	dfs(output, 0)
	return res
}

//https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/
func permutation(s string) []string {
	if s == "" {
		return nil
	}
	runes := []rune(s)
	ans := make([]string, 0)

	var helper func(int)
	helper = func(index int) {
		if index == len(runes)-1 {
			ans = append(ans, string(runes))
			return
		}
		set := map[rune]bool{}
		for i := index; i < len(runes); i++ {
			if set[runes[i]] {
				continue
			}
			set[runes[i]] = true
			runes[index], runes[i] = runes[i], runes[index]
			helper(index + 1)
			runes[index], runes[i] = runes[i], runes[index]
		}
	}
	helper(0)
	return ans
}
