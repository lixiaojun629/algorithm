package sort

import (
	"fmt"
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	nums := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(FindKthLargest(nums, 2))
}
