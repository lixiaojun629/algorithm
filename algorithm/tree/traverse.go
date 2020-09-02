package tree

//https://leetcode-cn.com/problems/binary-tree-preorder-traversal/solution/dai-ma-sui-xiang-lu-chi-tou-qian-zhong-hou-xu-de-d/

//PreOrder traverse in pre order
func PreOrder(root *BinaryTreeNode) []*BinaryTreeNode {
	res := make([]*BinaryTreeNode, 0)
	var dfs func(*BinaryTreeNode)
	dfs = func(node *BinaryTreeNode) {
		if node == nil {
			return
		}
		res = append(res, node)
		dfs(node.left)
		dfs(node.right)
	}

	dfs(root)
	return res
}

func PostOrder(root *BinaryTreeNode) []*BinaryTreeNode {
	var res = make([]*BinaryTreeNode, 0)
	var dfs func(*BinaryTreeNode)
	dfs = func(node *BinaryTreeNode) {
		if node == nil {
			return
		}
		dfs(node.left)
		dfs(node.right)
		res = append(res, node)
	}
	dfs(root)
	return res
}

func InOrder(root *BinaryTreeNode) []*BinaryTreeNode {
	var res = make([]*BinaryTreeNode, 0)
	var dfs func(*BinaryTreeNode)
	dfs = func(node *BinaryTreeNode) {
		if node == nil {
			return
		}
		dfs(node.left)
		res = append(res, node)
		dfs(node.right)
	}
	dfs(root)
	return res
}

func PreOrderStack(root *BinaryTreeNode) []*BinaryTreeNode {
	stack := make([]*BinaryTreeNode, 0)
	res := make([]*BinaryTreeNode, 0)
	for root != nil || len(stack) > 0 {
		if root != nil {
			res = append(res, root)
			if root.right != nil {
				stack = append(stack, root.right)
			}
			root = root.left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}
	return res
}

func InOrderStack(root *BinaryTreeNode) []*BinaryTreeNode {
	res := make([]*BinaryTreeNode, 0)
	stack := make([]*BinaryTreeNode, 0)

	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, root)
			root = root.right
		}
	}
	return res
}

func PostOrderStack(root *BinaryTreeNode) []*BinaryTreeNode {
	res := make([]*BinaryTreeNode, 0)
	stack := make([]*BinaryTreeNode, 0)

	for root != nil || len(stack) > 0 {
		if root != nil {
			res = append(res, root)
			if root.left != nil {
				stack = append(stack, root.left)
			}
			root = root.right
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

//LevelOrder 层级遍历
func LevelOrder(root *TreeNode) [][]int {
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	res := make([][]int, 0)
	for len(queue) > 0 {
		n := len(queue)
		values := make([]int, n)
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			values[i] = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, values)
	}
	return res
}
