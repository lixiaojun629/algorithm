package link_list

import (
	"fmt"
	"testing"
)

func TestMergeKLink(t *testing.T) {
	l1_node1 := &ListNode{Val: 1}
	l1_node4 := &ListNode{Val: 4}
	l1_node5 := &ListNode{Val: 5}

	l2_node1 := &ListNode{Val: 1}
	l2_node3 := &ListNode{Val: 3}
	l2_node4 := &ListNode{Val: 4}

	l3_node2 := &ListNode{Val: 2}
	l3_node6 := &ListNode{Val: 6}

	l1_node1.Next = l1_node4
	l1_node4.Next = l1_node5

	l2_node1.Next = l2_node3
	l2_node3.Next = l2_node4

	l3_node2.Next = l3_node6

	head := mergeKLists([]*ListNode{l1_node1, l2_node1, l3_node2})
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
func TestAddTwoNumbers(t *testing.T) {
	// anode1 := &ListNode{Val: 9}
	// anode2 := &ListNode{Val: 9}
	// anode3 := &ListNode{Val: 9}
	// anode4 := &ListNode{Val: 9}

	// bnode1 := &ListNode{Val: 9}

	// anode1.Next = anode2
	// anode2.Next = anode3
	// anode3.Next = anode4

	addTwoNumbers1(NewList([]int{9, 9, 1}), NewList([]int{1}))

}
