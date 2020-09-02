package link_list

//https://leetcode-cn.com/problems/swap-nodes-in-pairs/
func SwapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Val: -1, Next: head} //哨兵，操作完成后，找不到链表头的时候加哨兵
	pre, p1, p2 := dummy, head, head.Next

	for p1 != nil && p2 != nil {
		p1.Next = p2.Next
		p2.Next = p1
		pre.Next = p2

		pre = p1
		p1 = p1.Next
		if p1 != nil {
			p2 = p1.Next
		}
	}
	return dummy.Next
}

func SwapPairsRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = SwapPairsRecursive(next.Next)
	next.Next = head
	return next
}
