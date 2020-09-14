package dp

//剑指offer48
//https://leetcode-cn.com/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/

func LengthOfLongestSubstring(s string) int {
	//以第i个字符为结尾的最长子字符串的长度
	subStrLen := make([]int, 0)
	maxLen := 0
	for idx := range s {
		subStrLen = append(subStrLen, 1)
		if idx > 0 {
			dup := false
			i := 1
			for ; i <= subStrLen[idx-1]; i++ {
				if s[idx-i] == s[idx] {
					dup = true
					break
				}
			}
			if dup {
				subStrLen[idx] = i
			} else {
				subStrLen[idx] = subStrLen[idx-1] + 1
			}
		}

		if subStrLen[idx] > maxLen {
			maxLen = subStrLen[idx]
		}
	}
	return maxLen
}

func LengthOfLongestSubstring2(s string) int {
	//以第i个字符为结尾的最长子字符串的长度
	subStrLen := make([]int, 0)
	//start 没有重复字符的子字符串的最小索引
	maxLen, start := 0, -1
	//某个字符出现的最大索引
	charIndexMap := make(map[rune]int)
	for idx, char := range s {
		subStrLen = append(subStrLen, 1)
		preIdx, ok := charIndexMap[char]
		if ok {
			if preIdx > start {
				start = preIdx
			}
		}
		subStrLen[idx] = idx - start
		charIndexMap[char] = idx
		if maxLen < subStrLen[idx] {
			maxLen = subStrLen[idx]
		}
	}
	return maxLen
}
