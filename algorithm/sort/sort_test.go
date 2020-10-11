package sort

import (
	"fmt"
	"testing"
)

func TestReversePairs(t *testing.T) {
	reversePairs([]int{7, 5, 6, 4})
}
func TestMinNumber(t *testing.T) {
	nums := []int{10, 23, 3, 43}
	fmt.Println(minNumber(nums))
}

func TestFindKthLargest(t *testing.T) {
	nums := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(FindKthLargest(nums, 2))
}
