package sort

import (
	"fmt"
	"strconv"
	"strings"
)

func minNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i, num := range nums {
		strs[i] = strconv.Itoa(num)
	}
	qsort(strs, 0, len(nums)-1)
	fmt.Println(strs)
	return strings.Join(strs, "")
}
func qsort(strs []string, low, high int) {
	if low >= high {
		return
	}
	mid := partition4(strs, low, high)
	qsort(strs, low, mid-1)
	qsort(strs, mid+1, high)
}
func partition4(strs []string, low, high int) int {
	slow, fast := low, low
	pivot := strs[high]
	for ; fast < high; fast++ {
		if strs[fast]+pivot < pivot+strs[fast] {
			strs[slow], strs[fast] = strs[fast], strs[slow]
			slow++
		}
	}
	strs[slow], strs[high] = strs[high], strs[slow]
	return slow
}
