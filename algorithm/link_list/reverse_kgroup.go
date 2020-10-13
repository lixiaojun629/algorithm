package link_list

//https://leetcode-cn.com/problems/reverse-nodes-in-k-group/
func reverseKGroup(head *ListNode, k int) *ListNode {
	var p = head
	for i := 0; i < k; i++ {
		if p != nil {
			p = p.Next
		} else {
			return head
		}
	}

	var pre *ListNode
	var curr = head
	for i := 0; i < k; i++ {
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	head.Next = reverseKGroup(curr, k)
	return pre
}
