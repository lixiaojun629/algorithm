package link_list

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

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	dummy := &ListNode{}
	prev := dummy
	last := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + last
		if sum >= 10 {
			last = 1
		} else {
			last = 0
		}
		prev.Next = &ListNode{Val: sum % 10}
		prev = prev.Next
	}

	if last != 0 {
		prev.Next = &ListNode{Val: last}
	}
	return dummy.Next
}
