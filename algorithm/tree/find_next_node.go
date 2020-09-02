package tree

//jz8, 给定一颗二叉树和其中一个节点，如何找出中序遍历的序列的下一个节点？树中的节点除了两个指向左右孩子节点的指针，还有一个指向
//父节点的指针。

//思路: 中序遍历顺序，左根右，下一个要遍历的节点只能是当前节点的根节点或右节点孩子中的最左节点。
//1、先找右孩子的最左子节点
//2、再找未遍历的父节点(当前节点是其父节点的左子节点)
func FindNextNode(current *BinaryTreeNode) *BinaryTreeNode {
	if current == nil {
		return nil
	}

	if current.right != nil {
		next := current.right
		for next.left != nil {
			next = next.left
		}
		return next
	}

	for current.parent != nil && current.parent.left != current {
		current = current.parent
	}

	return current.parent
}
