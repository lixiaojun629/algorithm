package queue

//https://leetcode-cn.com/problems/sliding-window-maximum/

func maxSlidingWindow(nums []int, k int) []int {
	//res保存结果
	//window保存滑动窗口里数字在nums中的索引
	res, window := make([]int, 0), make([]int, 0)

	//初始化滑动窗口，滑动窗口忽略旧的小的值
	for i := 0; i < k-1; i++ {
		window = appendWindow(window, nums, i)
	}

	for i := k - 1; i < len(nums); i++ {
		if len(window) > 0 && window[0] < i-k+1 {
			window = window[1:]
		}
		window = appendWindow(window, nums, i)
		res = append(res, nums[window[0]])
	}
	return res
}

//滑动窗口存储索引值，滑动窗口中，原来存在的值如果必新加入的值小，则没有存在价值，删除即可。
func appendWindow(window, nums []int, i int) []int {
	for len(window) > 0 && nums[i] > nums[window[len(window)-1]] {
		window = window[:len(window)-1]
	}
	window = append(window, i)
	return window
}
