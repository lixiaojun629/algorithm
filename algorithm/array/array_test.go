package array

import (
	"testing"
	"time"
)

func TestTrap(t *testing.T) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	t.Log(trap(height))
}

func TestFindMedianSortedArray(t *testing.T) {
	// nums1 := []int{1, 4}
	// nums2 := []int{2, 3, 5, 6}
	nums1 := []int{0, 0, 0, 0, 0}
	nums2 := []int{-1, 0, 0, 0, 0, 0, 1}
	begin := time.Now()
	t.Log(findMedianSortedArrays(nums1, nums2))
	t.Log(time.Since(begin))
}
