package dp

//https://leetcode-cn.com/problems/edit-distance/

//minDist[i][j] word1[0...i-1]和word2[0...j-1]的编辑距离
//word1[i-1]==word2[j-1] : minDist[i][j] = minDist[i-1][j-1]
//word1[i-1]!=word2[j-1] : minDist[i][j] = 1 + min(minDist[i-1][j],minDist[i][j-1],minDist[i-1][j-1])
func minDistance(word1 string, word2 string) int {
	len1, len2 := len(word1), len(word2)
	//minDist[i][j] word1[1..i]=>word2[1..j]
	minDist := make([][]int, len1+1)
	for idx := range minDist {
		minDist[idx] = make([]int, len2+1)
	}

	for i := 0; i < len1+1; i++ {
		minDist[i][0] = i
	}
	for j := 0; j < len2+1; j++ {
		minDist[0][j] = j
	}

	for i := 1; i < len1+1; i++ {
		for j := 1; j < len2+1; j++ {
			if word1[i-1] == word2[j-1] {
				minDist[i][j] = minDist[i-1][j-1]
			} else {
				minDist[i][j] = 1 + minThree(minDist[i][j-1], minDist[i-1][j], minDist[i-1][j-1])
			}
		}
	}
	return minDist[len1][len2]
}

func minThree(a, b, c int) int {
	if a > b {
		a = b
	}
	if a > c {
		a = c
	}
	return a
}
