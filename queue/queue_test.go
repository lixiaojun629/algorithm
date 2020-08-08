package queue

import (
	"fmt"
	"testing"
)

func TestMaxSlidingWindowMax(t *testing.T) {
	fmt.Println(maxSlidingWindow([]int{1, -1}, 1))
}
