//https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/
package tree

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}
