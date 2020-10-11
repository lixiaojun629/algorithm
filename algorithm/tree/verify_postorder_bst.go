package tree

//https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/
func verifyPostorder(postorder []int) bool {
	return verify(postorder, 0, len(postorder)-1)
}

func verify(postorder []int, low, high int) bool {
	if low >= high {
		return true
	}
	root := postorder[high]
	mid := low
	for ; mid < high; mid++ {
		if postorder[mid] > root {
			break
		}
	}
	for i := mid; i < high; i++ {
		if postorder[i] < root {
			return false
		}
	}
	return verify(postorder, low, mid-1) && verify(postorder, mid, high-1)
}
