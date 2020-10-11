package link_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	res := make([]int, 0)
	var reversePrintRecv func(*ListNode)
	reversePrintRecv = func(node *ListNode) {
		if node == nil {
			return
		}
		reversePrintRecv(node.Next)
		res = append(res, node.Val)
	}
	reversePrintRecv(head)
	return res
}
