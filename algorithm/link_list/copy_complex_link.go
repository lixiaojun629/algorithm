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

// 第一步：循环把每个节点复制一份，放在原节点后面
// 第二步：循环把复制出的节点Random指针改为指向节点的后一个节点
// 第三步：分离原链表和复制链表
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
