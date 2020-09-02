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