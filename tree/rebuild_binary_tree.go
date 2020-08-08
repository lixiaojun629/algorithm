package tree

//剑指offer7：输入某二叉树前序遍历和中序遍历的结果，请重建该二叉树，假设输入的前序遍历和中序便利结果中都不含有重复的数字，例如
//输入前序遍历序列 1，2，4，7，3，5，6，8和中序遍历序列4，7，2，1，5，3，8，6；则重建次二叉树并输出头节点。

func indexOf(value int, list []int) int {
	index := -1
	for idx, item := range list {
		if value == item {
			index = idx
			break
		}
	}
	return index
}

func RebuildBinaryTree(preOrder []int, inOrder []int) *BinaryTreeNode {
	lenP, lenI := len(preOrder), len(inOrder)
	if lenP != lenI {
		panic("invalid input")
	}
	if lenP == 0 {
		return nil
	}

	root := preOrder[0]
	indexOfRoot := indexOf(root, inOrder)

	leftInOrder := inOrder[:indexOfRoot]
	rightInOrder := inOrder[indexOfRoot+1:]
	leftPreOrder := preOrder[1 : len(leftInOrder)+1]
	rightPreOrder := preOrder[len(leftInOrder)+1:]

	return &BinaryTreeNode{
		value: root,
		left:  RebuildBinaryTree(leftPreOrder, leftInOrder),
		right: RebuildBinaryTree(rightPreOrder, rightInOrder),
	}
}
