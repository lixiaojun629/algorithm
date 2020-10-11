package link_list

//https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil || k <= 0 {
		return nil
	}
	slow, fast := head, head
	for i := k - 1; i > 0; i-- {
		if fast.Next != nil {
			fast = fast.Next
		} else {
			return nil
		}
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
