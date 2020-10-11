package link_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//https://leetcode-cn.com/problems/add-two-numbers/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	sum := l1.Val + l2.Val
	p := &ListNode{Val: sum % 10}
	p.Next = addTwoNumbers(l1.Next, l2.Next)
	if sum >= 10 {
		p.Next = addTwoNumbers(p.Next, &ListNode{Val: 1})
	}
	return p
}
