package link_list

//https://leetcode-cn.com/problems/fu-za-lian-biao-de-fu-zhi-lcof/
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	curr := head
	for curr != nil {
		cloned := *curr
		curr.Next = &cloned
		curr = cloned.Next
	}
	curr = head
	newHead := head.Next
	for curr != nil {
		cloned := curr.Next
		if cloned.Random != nil {
			cloned.Random = cloned.Random.Next
		}
		curr = cloned.Next
	}
	curr = head
	for curr != nil {
		cloned := curr.Next
		nextOrigin := cloned.Next
		if nextOrigin != nil {
			cloned.Next = nextOrigin.Next
		}
		curr.Next = nextOrigin
		curr = nextOrigin
	}
	return newHead
}

func copyRandomList1(head *Node) *Node {
	//key *Node in original link list , value *Node in cloned link list
	cache := make(map[*Node]*Node)
	curr := head
	for curr != nil {
		cloned := *curr
		cache[curr] = &cloned
		curr = curr.Next
	}
	curr = head
	for curr != nil {
		cache[curr].Next = cache[curr.Next]
		cache[curr].Random = cache[curr.Random]
		curr = curr.Next
	}
	return cache[head]
}
