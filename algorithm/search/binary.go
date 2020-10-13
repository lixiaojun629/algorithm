package search

//Binary search
func Binary(items []int, target int) int {
	low := 0
	high := len(items) - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if items[mid] == target {
			return mid
		}

		if items[mid] > target {
			high = mid - 1
		}
		if items[mid] < target {
			low = mid + 1
		}
	}
	return -1
}

//367 https://leetcode-cn.com/problems/valid-perfect-square/
func isPerfectSquare(num int) bool {
	low, high := 0, num
	for low <= high {
		mid := low + (high-low)>>1
		v := mid * mid
		if v == num {
			return true
		} else if v > num {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}

//33 https://leetcode-cn.com/problems/search-in-rotated-sorted-array/
func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] >= nums[low] {
			if nums[mid] > target && nums[low] <= target {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if nums[mid] < target && nums[high] >= target {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

//154 offer11 https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/
//以右边界为轴元素，numbers[high]把数组一分为二，用二分查找最小元素
//mid= (low +high)/2 下取整，当nums[mid]==nums[right]时，high--不会错过最小元素
func minArray(numbers []int) int {
	low, high := 0, len(numbers)-1

	for low < high {
		mid := low + (high-low)>>1
		if numbers[mid] > numbers[high] {
			low = mid + 1
		} else if numbers[mid] < numbers[high] {
			high = mid
		} else if numbers[mid] == numbers[high] {
			high--
		}
	}
	return numbers[low]
}
