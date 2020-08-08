package tree

import (
	"fmt"
	"testing"
)

func TestRebuildBinaryTree(t *testing.T) {
	preOrder := []int{1, 2, 4, 7, 3, 5, 6, 8}
	inOrder := []int{4, 7, 2, 1, 5, 3, 6, 8}
	root := RebuildBinaryTree(preOrder, inOrder)
	fmt.Println(root)
}

func TestFindNextNode(t *testing.T) {
	node1 := NewBinaryTreeNode(1)
	node2 := NewBinaryTreeNode(2)
	node3 := NewBinaryTreeNode(3)
	node4 := NewBinaryTreeNode(4)
	node5 := NewBinaryTreeNode(5)
	node6 := NewBinaryTreeNode(6)
	node7 := NewBinaryTreeNode(7)

	ConnectTreeNodes(node1, node2, node3)
	ConnectTreeNodes(node2, node4, node5)
	ConnectTreeNodes(node3, node7, node6)
	fmt.Println(FindNextNode(node6))
	fmt.Println(FindNextNode(node4))
	fmt.Println(FindNextNode(node2))
	fmt.Println(FindNextNode(node1))
}

func TestIsValidBST(t *testing.T) {
	node := &TreeNode{
		Val: 0,
	}
	t.Log(IsValidBST(node))
}
