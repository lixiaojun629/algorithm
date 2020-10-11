package stack

//https://leetcode-cn.com/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof/
func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0)

	for _, val := range pushed {
		stack = append(stack, val)
		for len(stack) > 0 && stack[len(stack)-1] == popped[0] {
			popped = popped[1:]
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0 && len(popped) == 0

}
