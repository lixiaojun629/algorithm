package tree

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var lastVal int = -1 << 31

func IsValidBST(root *TreeNode) bool {
	return check(root)
}

func check(node *TreeNode) bool {
	if node == nil {
		return true
	}

	if !check(node.Left) {
		return false
	}

	if lastVal >= node.Val {
		return false
	}
	lastVal = node.Val

	if !check(node.Right) {
		return false
	}

	return true
}
