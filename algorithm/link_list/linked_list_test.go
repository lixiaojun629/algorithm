package link_list

import (
	"fmt"
	"testing"
)

func TestCopyLink(t *testing.T) {
	node1 := &Node{Val: 7}
	node2 := &Node{Val: 13}
	node3 := &Node{Val: 11}
	node4 := &Node{Val: 10}
	node5 := &Node{Val: 1}
	node1.Next = node2
	node1.Random = nil
	node2.Next = node3
	node2.Random = node1
	node3.Next = node4
	node3.Random = node5
	node4.Next = node5
	node4.Random = node3
	node5.Random = node1
	copyRandomList(node1)

}
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
