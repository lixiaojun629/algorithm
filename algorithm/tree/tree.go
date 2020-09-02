package tree

import "strconv"

type BinaryTreeNode struct {
	value  int
	left   *BinaryTreeNode
	right  *BinaryTreeNode
	parent *BinaryTreeNode
}

func (n *BinaryTreeNode) String() string {
	return strconv.Itoa(n.value)
}

func NewBinaryTreeNode(value int) *BinaryTreeNode {
	return &BinaryTreeNode{
		value: value,
	}
}

func ConnectTreeNodes(parent, left, right *BinaryTreeNode) {
	if parent == nil {
		return
	}
	if left != nil {
		parent.left = left
		left.parent = parent
	}
	if right != nil {
		parent.right = right
		right.parent = parent
	}
}
