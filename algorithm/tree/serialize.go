package tree

import (
	"strconv"
	"strings"
)

type Codec struct {
	sep  string
	null string
}

func Constructor() Codec {
	return Codec{
		sep:  ",",
		null: "null",
	}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	strs := make([]string, 0)
	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			strs = append(strs, this.null)
			return
		}

		strs = append(strs, strconv.Itoa(node.Val))
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return strings.Join(strs, this.sep)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	strs := strings.Split(data, this.sep)
	pos := 0
	var buildTree func() *TreeNode
	buildTree = func() *TreeNode {
		val := strs[pos]
		pos++
		if val == this.null {
			return nil
		}
		num, error := strconv.Atoi(val)
		if error != nil {
			panic(error)
		}
		node := &TreeNode{Val: num}
		node.Left = buildTree()
		node.Right = buildTree()
		return node
	}
	return buildTree()
}
