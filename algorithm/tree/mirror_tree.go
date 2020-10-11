package tree

//https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	newRoot := &TreeNode{}
	copyTree(root, newRoot)
	return newRoot
}

func copyTree(src, target *TreeNode) {
	target.Val = src.Val
	if src.Right != nil {
		target.Left = &TreeNode{}
		copyTree(src.Right, target.Left)
	}
	if src.Left != nil {
		target.Right = &TreeNode{}
		copyTree(src.Left, target.Right)
	}
}
