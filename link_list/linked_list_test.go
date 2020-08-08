package link_list

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {
	node1 := &ListNode{
		Val: 1,
	}
	node2 := &ListNode{
		Val: 2,
	}
	node3 := &ListNode{
		Val: 4,
	}
	node1.Next = node2
	node2.Next = node3

	fmt.Println(ReverseList(node1))
	fmt.Println(ReverseListRecursive(node1))

}

func TestSwapPairs(t *testing.T) {
	node1 := &ListNode{
		Val: 1,
	}
	node2 := &ListNode{
		Val: 2,
	}
	node3 := &ListNode{
		Val: 3,
	}
	node4 := &ListNode{
		Val: 4,
	}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	// fmt.Println(SwapPairs(node1))
	fmt.Println(SwapPairsRecursive(node1))
}
