package tree

//https://leetcode-cn.com/problems/check-completeness-of-a-binary-tree/
func isCompleteTree(root *TreeNode) bool {
	queue := []*TreeNode{root}
	hasNilNode := false
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node != nil {
			if hasNilNode {
				return false
			}
			queue = append(queue, node.Left, node.Right)
		} else {
			hasNilNode = true
		}
	}
	return true
}
