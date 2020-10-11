package dp

//https://leetcode-cn.com/problems/regular-expression-matching/
//dp[i][j] str[0...i-1]与pattern[0...j-1]是否匹配
//pattern[j-1] == '*': dp[i][j] = dp[i][j] || dp[i][j-2]
//pattern[j-1] == '*' && str[i-1] == pattern[j-1]: dp[i][j] = dp[i][j] || dp[i-1][j]
//str[i-1]==pattern[j-1]: dp[i][j] = dp[i][j] || dp[i-1][j-1]
//
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	match := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		match[i] = make([]bool, n+1)
	}

	isEqual := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		if s[i-1] == p[j-1] {
			return true
		}
		return false
	}

	match[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				if j >= 2 {
					match[i][j] = match[i][j] || match[i][j-2]
					if isEqual(i, j-1) {
						match[i][j] = match[i][j] || match[i-1][j]
					}
				}

			} else {
				if isEqual(i, j) {
					match[i][j] = match[i][j] || match[i-1][j-1]
				}
			}
		}
	}
	return match[m][n]
}
