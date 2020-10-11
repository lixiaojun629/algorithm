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

//https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/
func getLeastNumbers(arr []int, k int) []int {
	low, high := 0, len(arr)-1
	for low <= high {
		rank := partition3(arr, low, high)
		if rank == k-1 {
			return arr[:k]
		}
		if rank > k-1 {
			high = rank - 1
		} else {
			low = rank + 1
		}
	}
	return nil
}

func partition3(arr []int, low, high int) int {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i
}
