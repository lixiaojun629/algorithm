package link_list

import "math"

//https://leetcode-cn.com/problems/merge-two-sorted-lists/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	prev := dummy

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}
	if l1 == nil {
		prev.Next = l2
	} else {
		prev.Next = l1
	}
	return dummy.Next
}

//https://leetcode-cn.com/problems/merge-k-sorted-lists/

//遍历找到最小节点
//把最小节点放到合并链表末尾
//在lists中，找到该最小节点node, node = node.Next
func mergeKLists1(lists []*ListNode) *ListNode {
	minNode := &ListNode{Val: math.MaxInt64}
	allEmpty := true
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			allEmpty = false
			if minNode.Val > lists[i].Val {
				minNode = lists[i]
			}
		}
	}
	if allEmpty {
		return nil
	}
	for i := 0; i < len(lists); i++ {
		if lists[i] == minNode {
			lists[i] = minNode.Next
			break
		}
	}
	minNode.Next = mergeKLists1(lists)
	return minNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, low, high int) *ListNode {
	if low == high {
		return lists[low]
	}
	mid := low + (high-low)>>1
	l1 := merge(lists, low, mid)
	l2 := merge(lists, mid+1, high)
	return mergeTwoLists1(l1, l2)
}
