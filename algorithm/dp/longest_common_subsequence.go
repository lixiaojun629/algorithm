package dp

//dp[i][j] text1[0...i-1]和text2[0...j-1]的最长公共子序列
//text[i-1]==text[j-1]: dp[i][j] = dp[i-1][j-1]+1
//text[i-1]!=text[j-1]: dp[i][j] = max(dp[i-1][j],dp[i][j-1])
func LongestCommonSubsequence(text1 string, text2 string) int {
	l1, l2 := len(text1), len(text2)
	if l1 == 0 || l2 == 0 {
		return 0
	}
	dp := make([][]int, l1+1)
	for idx := range dp {
		dp[idx] = make([]int, l2+1)
	}

	for i := 1; i < l1+1; i++ {
		for j := 1; j < l2+1; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = maxInt(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return dp[l1][l2]
}
