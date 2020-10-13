package tree

//https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/

//先从上到下，在每一个节点的左右子树中查找p,q节点，找到以后回溯，
//回溯过程中，如果左右子树均找到了节点，则此节点即为最近公共祖先，从此以后返回该公共祖先节点
//如果左右子树仅有一个子树找到了节点，则返回找到的节点。
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left == nil {
		return right
	} else if right == nil {
		return left
	} else {
		return root
	}
}

//https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
func lowestCommonAncestorOfSearchBinaryTree(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestorOfSearchBinaryTree(root.Right, p, q)
	}
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestorOfSearchBinaryTree(root.Left, p, q)
	}
	return root
}
