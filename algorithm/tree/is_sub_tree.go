package tree

//https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//递归判断，前序遍历树A，判断A树中是否包含B树
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	if isTreeAHasTreeB(A, B) {
		return true
	}

	return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

//递归判断，前序遍历树A和树B，判断A和B是否相等，A树可以有多余的叶子节点
func isTreeAHasTreeB(A, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	if A.Val != B.Val {
		return false
	}

	return isTreeAHasTreeB(A.Left, B.Left) && isTreeAHasTreeB(A.Right, B.Right)
}
