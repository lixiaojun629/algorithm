package link_list

import (
	"fmt"
	"strings"
)

//https://leetcode-cn.com/problems/reverse-linked-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) String() string {
	list := make([]string, 0)
	for n != nil {
		list = append(list, fmt.Sprintf("%d", n.Val))
		n = n.Next
	}
	return strings.Join(list, "->")
}

func ReverseList(head *ListNode) *ListNode {
	var pre *ListNode
	var curr = head

	for curr != nil {
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	return pre
}

func ReverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := ReverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
