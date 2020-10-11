package tree

//https://leetcode-cn.com/problems/ping-heng-er-cha-shu-lcof/
//时间复杂度O(N), 剪枝：子树不平衡时快速失败，直接返回结果
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var checkByDepth func(*TreeNode) int
	checkByDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := checkByDepth(node.Left)
		if left == -1 {
			return -1
		}
		right := checkByDepth(node.Right)
		if right == -1 {
			return -1
		}
		if left-right > 1 || right-left > 1 {
			return -1
		}
		if left > right {
			return left + 1
		}
		return right + 1
	}
	return checkByDepth(root) != -1
}
