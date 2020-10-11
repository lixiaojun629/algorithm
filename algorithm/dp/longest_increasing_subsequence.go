package dp

//https://leetcode-cn.com/problems/longest-increasing-subsequence/

func LengthOfLISDP(nums []int) int {
	//dp[i] 以第i个数字结尾的最长子序列长度
	dp := make([]int, len(nums))
	max := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = maxInt(dp[j]+1, dp[i])
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

func maxInt(a, b int) int {
	if a < b {
		a = b
	}
	return a
}

func LengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	//tails[i] 长度为i+1的最长上升子序列的末尾值，tails[i]单调递增
	tails := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if nums[i] > tails[len(tails)-1] {
			tails = append(tails, nums[i])
			continue
		}
		idx := binarySearch(tails, nums[i])
		tails[idx] = nums[i]
	}
	return len(tails)
}

//二分查找数组中第一个大于等于给定值的索引，如果数组中元素都小于给定值，则返回最大索引加一
func binarySearch(list []int, target int) int {
	low, high := 0, len(list)-1
	for low <= high {
		mid := low + (high-low)>>1
		if list[mid] >= target {
			if mid == 0 || list[mid-1] < target {
				return mid
			} else {
				high = mid - 1
			}
		} else {
			low = mid + 1
		}
	}
	return -1
}
