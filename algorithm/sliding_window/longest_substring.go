package sliding_window

//剑指offer48
//https://leetcode-cn.com/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/
//时间复杂度O(n) 空间复杂度O(alpha)（字符集大小）
//滑动窗口
func lengthOfLongestSubstring(s string) int {
	cache := map[byte]bool{}
	right, ans, n := 0, 0, len(s)
	for left := range s {
		if left > 0 {
			delete(cache, s[left-1])
		}
		for right < n && !cache[s[right]] {
			cache[s[right]] = true
			right++
		}
		substrLen := right - left
		if ans < substrLen {
			ans = substrLen
		}
	}
	return ans
}

func lengthOfLongestSubstring1(s string) int {
	if len(s) == 0 {
		return 0
	}
	//以第i个字符为结尾的最长子字符串的长度
	maxLen := make([]int, len(s))
	maxLen[0] = 1
	res := 1
	for i := 1; i < len(s); i++ {
		j := i - 1
		for ; j >= i-maxLen[i-1]; j-- {
			if s[i] == s[j] {
				break
			}
		}

		maxLen[i] = i - j

		if res < maxLen[i] {
			res = maxLen[i]
		}
	}
	return res
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
