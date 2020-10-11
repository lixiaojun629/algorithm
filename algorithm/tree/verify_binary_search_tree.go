package tree

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsValidBST(root *TreeNode) bool {
	var lastVal int = -1 << 31
	var check func(*TreeNode) bool
	check = func(node *TreeNode) bool {
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

	return check(root)
}
