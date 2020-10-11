package stack

import "testing"

func TestValidateStackSequences(t *testing.T) {
	pushed := []int{1, 2, 3, 4, 5}
	poped := []int{4, 5, 3, 2, 1}
	validateStackSequences(pushed, poped)
}
