package sort

func FindKthLargest(nums []int, k int) int {
	low, high := 0, len(nums)-1
	k = len(nums) - k
	for low <= high {
		p := partition2(nums, low, high)
		if p == k {
			return nums[p]
		} else if p > k {
			high = p - 1
		} else if p < k {
			low = p + 1
		}
	}
	return nums[low]
}

func partition2(nums []int, low, high int) int {
	pivot := nums[high]
	slow := low
	for fast := low; fast < high; fast++ {
		if nums[fast] < pivot {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
	nums[slow], nums[high] = nums[high], nums[slow]
	return slow
}
