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